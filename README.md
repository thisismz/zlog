# zlog

The logs can be saved to the related files inside the log files. It can be done easily and quickly.

### Installing

 `go get -u github.com/thisismz/zlog@latest`

## ⚡️ Quickstart

```go

package main

 
import  "github.com/thisismz/zlog"

func main() {
	var config = zlog.Config{
		Level:         "debug",
		Prefix:        "zlog",
		Format:        "json",
		Director:      "log",
		EncodeLevel:   "CapitalLevelEncoder",
		StacktraceKey: "stacktrace",
		MaxAge:        7,
		ShowLine:      true,
		LogInConsole:  true,
		SaveInfile: false,
	}
	z := zlog.New(config).Log()
	z.Info("Hello World")
	err := errors.New("this is your error")
	z.Error("Error", zap.Error(err))
}

```

### Default Config

```go

	var config = zlog.Config{
		Level:         "debug",
		Prefix:        "zlog",
		Format:        "json",
		Director:      "log",
		EncodeLevel:   "CapitalLevelEncoder",
		StacktraceKey: "stacktrace",
		MaxAge:        7,
		ShowLine:      true,
		LogInConsole:  true,
	}
```
