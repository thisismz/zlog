package zlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Level         = "debug"
	Prefix        = "mobinio"
	Format        = "console"
	Director      = "./logs"
	EncodeLevel   = "uppercase"
	StacktraceKey = "trace"
	ShowLine      = false
	LogInConsole  = true
	RetentionDay  = 30
	SaveInFile    = false
	SentryDns     = ""
)
var logger = newlog()

// Debug starts a new message with debug level.
//
// You must call Msg on the returned event in order to send the event.
func Debug(msg string, fields ...zapcore.Field) {
	logger.Debug(msg, fields...)
}

// Info starts a new message with info level.
//
// You must call Msg on the returned event in order to send the event.
func Info(msg string, fields ...zapcore.Field) {
	logger.Info(msg, fields...)
}

// Warn starts a new message with warn level.
//
// You must call Msg on the returned event in order to send the event.
func Warn(msg string, fields ...zapcore.Field) {
	logger.Warn(msg, fields...)
}

// Error starts a new message with error level.
//
// You must call Msg on the returned event in order to send the event.
func Error(msg string, err error) {
	logger.Error(msg, zap.Error(err))
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
//
// You must call Msg on the returned event in order to send the event.
func Fatal(msg string, fields ...zapcore.Field) {
	logger.Fatal(msg, fields...)
}

// Panic starts a new message with panic level. The message is also sent
// to the panic function.
//
// You must call Msg on the returned event in order to send the event.
func Panic(msg string, fields ...zapcore.Field) {
	logger.Panic(msg, fields...)
}

// WithLevel starts a new message with level.

// Log starts a new message with no level. Setting zerolog.GlobalLevel to
// zerolog.Disabled will still disable events produced by this method.
//
// You must call Msg on the returned event in order to send the event.
func Log() *zap.Logger {
	return logger
}
