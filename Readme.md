# zlog

`zlog` is a logging package for Go built on top of `zap`, designed for simple and structured logging. It provides multiple logging levels like Debug, Info, Warn, Error, and Fatal, with flexibility in configuring console output and log file directory.

## Installation

To use `zlog`, add it to your Go project by running:

```bash
go get github.com/thisismz/zlog@latest
```

## Configuration

You can configure `zlog` to control console output and specify a directory for log files:

```go
// Disable console output
log.LogInConsole = false

// Set directory path for log files
log.Director = "./log"
```

## Usage

Each logging function provides an easy way to log messages with various severity levels.

### Log Levels

- **Debug**: Logs debug information

  ```go
  log.Debug("Debug message", zap.String("key", "value"))
  ```
- **Info**: Logs informational messages

  ```go
  log.Info("Info message", zap.String("key", "value"))
  ```
- **Warn**: Logs warnings

  ```go
  log.Warn("Warning message", zap.String("key", "value"))
  ```
- **Error**: Logs errors

  ```go
  log.Error("Error occurred", err)
  ```
- **Fatal**: Logs fatal errors and terminates the application

  ```go
  log.Fatal("Fatal error", zap.String("key", "value"))
  ```

Each function takes a message string and optional fields (key-value pairs) for structured logging.

## Example

```go
package main

import (
    "log"
    "go.uber.org/zap"
)

func main() {
    // Configure logging
    log.LogInConsole = false
    log.Director = "./log"

    // Log messages
    log.Debug("Starting the application", zap.String("status", "debugging"))
    log.Info("Application running", zap.String("status", "info"))
    log.Warn("Potential issue detected", zap.String("status", "warning"))
    log.Error("An error occurred", err)
    log.Fatal("Critical error, shutting down", zap.String("status", "fatal"))
}
```

## Benchmarking

The package includes a benchmark test (`log_benchmark_test.go`) to evaluate logging performance. Run it with:

```bash
go test -bench .
```
