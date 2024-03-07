package cmd

import (
	"github.com/ryan-reichenberg/flash-cli/internal"
	"github.com/spf13/cobra"
)

var (
	requestUrl string
	verb       string
	headers    []string
	body       string
	times      int
	threads    int
	verbose    bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "flash",
	Short:   "A tool for measuring endpoint performance",
	Version: "1.0.0",
	Long: `Flash is a CLI tool that allows you to measure response times against an http endpoint.
	This application provides detailed timing metrics against the specified endpoint`,
	Run: func(cmd *cobra.Command, args []string) {
		request := internal.HttpRequest{
			RequestUrl: requestUrl,
			Verb:       verb,
			Headers:    headers,
			Body:       body,
			Times:      times,
			Threads:    threads,
			Verbose:    verbose,
		}

		internal.Execute(request)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().StringVarP(&verb, "verb", "v", "GET", "The HTTP verb")
	rootCmd.Flags().StringVarP(&requestUrl, "url", "u", "", "The request url")
	rootCmd.Flags().StringVarP(&body, "body", "b", "", "The request body")
	rootCmd.Flags().StringArrayVarP(&headers, "headers", "H", []string{}, "The request headers")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "V", false, "Verbose mode")
	rootCmd.Flags().IntVarP(&times, "times", "t", 1, "Number of times to run request")
	rootCmd.Flags().IntVarP(&threads, "threads", "T", 10, "Number of threads to run")
}
