# zlog

The logs can be saved to the related files inside the log files. It can be done easily and quickly.

### Installing

 `go get -u github.com/thisismz/zlog`

##  ⚡️ Quickstart

```go

package main

 
import  "github.com/thisismz/zlog"

func  main() {
	z := zlog.New()
	z.Log().Info("Hello World")
	err := errors.New("this is your error")
	z.Log().Error("Error", zap.Error(err))
}

```
###  Default Config

```go

var ConfigDefault = Config{
	Level =  "info"
	Format =  "console"
	Prefix =  "[zlog]"
	Director =  "log"
	EncodeLevel =  "CapitalLevelEncoder"
	StacktraceKey =  "stacktrace"
)
```
