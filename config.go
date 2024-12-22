package zlog

import (
	"time"

	"go.uber.org/zap/zapcore"
)

var (
	Level         = "debug"
	Prefix        = "mobinio"
	Format        = "console"
	Director      = ""
	EncodeLevel   = "uppercase"
	StacktraceKey = "trace"
	ShowLine      = false
	LogInConsole  = true
	RetentionDay  = 30
	SentryDns     = ""
)

func levels() []zapcore.Level {
	levels := make([]zapcore.Level, 0, 7)
	level, err := zapcore.ParseLevel(Level)
	if err != nil {
		level = zapcore.DebugLevel
	}
	for ; level <= zapcore.FatalLevel; level++ {
		levels = append(levels, level)
	}
	return levels
}

func encoder() zapcore.Encoder {
	config := zapcore.EncoderConfig{
		TimeKey:       "time",
		NameKey:       "name",
		LevelKey:      "level",
		CallerKey:     "caller",
		MessageKey:    "message",
		StacktraceKey: StacktraceKey,
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(Prefix + t.Format("2006-01-02 15:04:05.000"))
		},
		EncodeLevel:    LevelEncoder(),
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}
	if Format == "json" {
		return zapcore.NewJSONEncoder(config)
	}
	return zapcore.NewConsoleEncoder(config)

}

func LevelEncoder() zapcore.LevelEncoder {
	switch {
	case EncodeLevel == "LowercaseLevelEncoder":
		return zapcore.LowercaseLevelEncoder
	case EncodeLevel == "LowercaseColorLevelEncoder":
		return zapcore.LowercaseColorLevelEncoder
	case EncodeLevel == "CapitalLevelEncoder":
		return zapcore.CapitalLevelEncoder
	case EncodeLevel == "CapitalColorLevelEncoder":
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

// defaultConfig returns a Config instance with default values
