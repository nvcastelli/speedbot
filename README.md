# speedbot
Here we will be creating a network tool that tests connectivity over long periods of time. This will primarily be written in Golang.

## Dependancies
* https://github.com/showwin/speedtest-go - pure Go API to Test Internet Speed using speedtest.net
## Usage
	go run main.go 1h 30m 30s

### Planned Features

- Visual graph representation of data using go-echarts library
- More complete value structure for metric reporting (e.g. period metrics for low/high performance)
- Fleshed out interface (Undecided on CLI tool verses windowed interface)
