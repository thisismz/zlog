package zlog

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Zlog struct {
	// App config
	config Config
}

type Config struct {
	Level         string
	Format        string
	Prefix        string
	Director      string
	ShowLine      *bool
	EncodeLevel   string
	StacktraceKey string
	LogInConsole  *bool
}

// Default Config values
const (
	Level         = "info"
	Format        = "console"
	Prefix        = "[zlog]"
	Director      = "log"
	EncodeLevel   = "CapitalLevelEncoder"
	StacktraceKey = "stacktrace"
)

var (
	ShowLine     = true
	LogInConsole = true
)

// NewZlog configures Zlog
func New(config ...Config) *Zlog {
	zlog := &Zlog{
		config: Config{},
	}
	// Override config if provided
	if len(config) > 0 {
		zlog.config = config[0]
	}
	if zlog.config.Level == "" {
		zlog.config.Level = Level
	}
	if zlog.config.Format == "" {
		zlog.config.Format = Format
	}
	if zlog.config.Prefix == "" {
		zlog.config.Prefix = Prefix
	}
	if zlog.config.Director == "" {
		zlog.config.Director = Director
	}
	if zlog.config.ShowLine == nil {
		zlog.config.ShowLine = &ShowLine
	}
	if zlog.config.EncodeLevel == "" {
		zlog.config.EncodeLevel = EncodeLevel
	}
	if zlog.config.StacktraceKey == "" {
		zlog.config.StacktraceKey = StacktraceKey
	}
	if zlog.config.LogInConsole == nil {
		zlog.config.LogInConsole = &LogInConsole
	}
	return zlog
}
func (zlog *Zlog) Log() *zap.Logger {
	if ok, _ := pathExists(zlog.config.Director); !ok { // Determine if there is a Director folder
		fmt.Printf("create %v directory\n", zlog.config.Director)
		_ = os.Mkdir(zlog.config.Director, os.ModePerm)
	}
	// Debug level
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	// Log level
	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	// Warning level
	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})
	// Error level
	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	cores := [...]zapcore.Core{
		zlog.getEncoderCore(fmt.Sprintf("./%s/zlog_debug.log", zlog.config.Director), debugPriority),
		zlog.getEncoderCore(fmt.Sprintf("./%s/zlog_info.log", zlog.config.Director), infoPriority),
		zlog.getEncoderCore(fmt.Sprintf("./%s/zlog_warn.log", zlog.config.Director), warnPriority),
		zlog.getEncoderCore(fmt.Sprintf("./%s/zlog_error.log", zlog.config.Director), errorPriority),
	}
	logger := zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())

	if *zlog.config.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

// getEncoderConfig get zapcore. EncoderConfig
func (zlog *Zlog) getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  zlog.config.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zlog.customTimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // Print the file name and the line which the error occurred
	}
	switch {
	case zlog.config.EncodeLevel == "LowercaseLevelEncoder": // Lowercase encoder (default)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case zlog.config.EncodeLevel == "LowercaseColorLevelEncoder": // Lowercase encoder with color
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case zlog.config.EncodeLevel == "CapitalLevelEncoder": // Uppercase encoder
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case zlog.config.EncodeLevel == "CapitalColorLevelEncoder": // Uppercase encoder with color
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder get zapcore. Encoder
func (zlog *Zlog) getEncoder() zapcore.Encoder {
	if zlog.config.Format == "json" {
		return zapcore.NewJSONEncoder(zlog.getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(zlog.getEncoderConfig())
}

// getEncoderCore gets Encoder's zapcore.Core
func (zlog *Zlog) getEncoderCore(fileName string, level zapcore.LevelEnabler) (core zapcore.Core) {
	writer := GetWriteSyncer(fileName, *zlog.config.LogInConsole) // Use file-rotatelogs for log splitting
	return zapcore.NewCore(zlog.getEncoder(), writer, level)
}

// Customize the log output time format
func (zlog *Zlog) customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 - 15:04:05.000"))
}
