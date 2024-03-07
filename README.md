# Flash CLI ⚡️

Flash is a blazingly fast http client designed for benchmarking API endpoints. Flash provides detailed metrics over a number of http requests and displays them in an easy to read format.

Flash aims to provide an easy mechanism to quickly benchmark API endpoint performance.

## Features

- Blazing fast
- Cross platform
- Intutive syntax
- Accurate and detailed timing metrics

## Installation

```
go install github.com/ryan-reichenberg/flash-cli
```

## Usage

```
Flash is a CLI tool that allows you to measure response times against an http endpoint.
	This application provides detailed timing metrics against the specified endpoint

Usage:
  flash [flags]

Flags:
  -b, --body string           The request body
  -H, --headers stringArray   The request headers
  -h, --help                  help for flash
  -T, --threads int           Number of threads to run (default 10)
  -t, --times int             Number of times to run request (default 1)
  -u, --url string            The request url
  -v, --verb string           The HTTP verb (default "GET")
  -V, --verbose               Verbose mode
      --version               version for flash
```

## Example

In this example, we run flash against the [https://jsonplaceholder.typicode.com](JSONPlaceholder API) and we specify that we want to gather metrics over 10 requests.

```
~ ❯ flash --url https://jsonplaceholder.typicode.com/posts/1 --times 10                                         
----- Timings Summary -----
Successful Requests: 10, Failed Requests: 0
Average: 0.385360s
Median: 0.383695s
p90: 0.389663ms
p99: 0.392193s
```

We can infer here, that the JSONPlaceholder API returned data within 383ms on average and within the worst case (p99) it returned data within 392ms

## Disclaimer

This tool should be used with care and should be used to benchmark APIs you own. It should be used responsibly and not with the intent of harm.
