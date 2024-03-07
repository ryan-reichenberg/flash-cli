package internal

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/montanaflynn/stats"
)

type HttpRequest struct {
	RequestUrl string
	Verb       string
	Headers    []string
	Body       string
	Times      int
	Threads    int
	Verbose    bool
}

type HttpResponse struct {
	statusCode    int
	executionTime time.Duration
}

func Execute(httpRequest HttpRequest) {

	channel := make(chan HttpResponse, httpRequest.Times)
	limiter := make(chan struct{}, httpRequest.Threads)
	defer close(channel)
	defer close(limiter)

	for i := 0; i < httpRequest.Times; i++ {
		limiter <- struct{}{}
		go func() {
			defer func() { <-limiter }()
			var req *http.Request
			var err error

			if httpRequest.Body != "" {
				req, err = http.NewRequest(httpRequest.Verb, httpRequest.RequestUrl, bytes.NewBuffer([]byte(httpRequest.Body)))
			} else {
				req, err = http.NewRequest(httpRequest.Verb, httpRequest.RequestUrl, nil)
			}

			if err != nil {
				fmt.Printf("error to create http request: %s\n", err)
				os.Exit(1)
			}
			for _, element := range httpRequest.Headers {
				if name, value, found := strings.Cut(element, ":"); found {
					req.Header.Add(strings.TrimSpace(name), strings.TrimSpace(value))
				} else if !found && httpRequest.Verbose {
					fmt.Printf("Incorrect header format [%s], skipping...", element)
				}
			}
			startTime := time.Now()
			resp, err := http.DefaultClient.Do(req)
			duration := time.Since(startTime)

			if httpRequest.Verbose {
				fmt.Printf("[%d] - %fs\n", resp.StatusCode, duration.Seconds())
			}
			if err != nil {
				fmt.Printf("Failed to send request: %s\n", err)
				os.Exit(1)
			}
			channel <- HttpResponse{
				statusCode:    resp.StatusCode,
				executionTime: duration,
			}
		}()
	}

	for i := 0; i < cap(limiter); i++ {
		limiter <- struct{}{}
	}

	PrintMetrics(channel, httpRequest.Times)
}

func PrintMetrics(requests <-chan HttpResponse, requestCount int) {
	timings := []float64{}
	successCount := 0
	failCount := 0
	for i := 0; i < requestCount; i++ {
		request := <-requests
		if IsSuccessful(request.statusCode) {
			successCount++
		} else if isFail(request.statusCode) {
			failCount++
		}
		timings = append(timings, request.executionTime.Seconds())
	}
	average, _ := stats.Mean(timings)
	median, _ := stats.Median(timings)
	p90, _ := stats.Percentile(timings, 90)
	p99, _ := stats.Percentile(timings, 99)
	fmt.Println("----- Timings Summary -----")
	fmt.Printf("Successful Requests: %d, Failed Requests: %d\n", successCount, failCount)
	fmt.Printf("Average: %fs\nMedian: %fs\np90: %fms\np99: %fs\n", average, median, p90, p99)
}
