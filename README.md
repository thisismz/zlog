# zlog

The logs can be saved to the related files inside the log files. It can be done easily and quickly.

### Installing

 `go get -u github.com/thisismz/zlog@latest`

## ⚡️ Quickstart

```go

package main

import (
	"errors"
	"time"

	"github.com/TheZeroSlave/zapsentry"
	"github.com/getsentry/sentry-go"
	"github.com/thisismz/zlog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	var config = zlog.Config{
		Level:         "debug",
		Prefix:        "zlog",
		Format:        "json",
		Director:      "log",
		EncodeLevel:   "CapitalLevelEncoder",
		StacktraceKey: "stacktrace",
		MaxAge:        7,
		ShowLine:      "true",
		LogInConsole:  "true",
		SaveInFile:    "false",
	}
	sentryClient, err := sentry.NewClient(sentry.ClientOptions{
		Dsn: "Sentry DSN",
	})
	if err != nil {
		// Handle the error here
		panic(err)
	}
	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	defer sentryClient.Flush(2 * time.Second)
	// Create a New zap logger
	zapLog := zlog.New(config).Log()
	// Modify zap logger to sentry logger
	z := modifyToSentryLogger(zapLog, sentryClient)
	z.Info("Hello World")
	err = errors.New("this is new error")
	z.Error("Error", zap.Error(err))
}
func modifyToSentryLogger(log *zap.Logger, client *sentry.Client) *zap.Logger {
	cfg := zapsentry.Configuration{
		Level:             zapcore.ErrorLevel, //when to send message to sentry
		EnableBreadcrumbs: true,               // enable sending breadcrumbs to Sentry
		BreadcrumbLevel:   zapcore.InfoLevel,  // at what level should we sent breadcrumbs to sentry
		Tags: map[string]string{
			"component": "system",
		},
	}
	core, err := zapsentry.NewCore(cfg, zapsentry.NewSentryClientFromClient(client))

	//in case of err it will return noop core. so we can safely attach it
	if err != nil {
		log.Warn("failed to init zap", zap.Error(err))
	}

	log = zapsentry.AttachCoreToLogger(core, log)

	// to use breadcrumbs feature - create new scope explicitly
	// and attach after attaching the core
	return log.With(zapsentry.NewScope())
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
		ShowLine:      "true",
		LogInConsole:  "true",
		SaveInFile:    "false",
	}
```
