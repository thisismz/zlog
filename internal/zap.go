package internal

import (
	"fmt"
	"os"

	"github.com/thisismz/zlog/global"

	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Zap = new(_zap)

type _zap struct{}

func (z *_zap) GetEncoder() zapcore.Encoder {
	if global.CONFIG.Format == "json" {
		return zapcore.NewJSONEncoder(z.GetEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(z.GetEncoderConfig())
}

func (z *_zap) GetEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.CONFIG.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    global.CONFIG.ZapEncodeLevel(),
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func (z *_zap) GetEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	if global.CONFIG.SaveInFile == true {
		writer, err := FileRotatelogs.GetWriteSyncer(l.String())
		if err != nil {
			fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
			return nil
		}
		return zapcore.NewCore(z.GetEncoder(), writer, level)
	}
	writer := zapcore.AddSync(os.Stdout)
	return zapcore.NewCore(z.GetEncoder(), writer, level)
}

func (z *_zap) CustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(global.CONFIG.Prefix + t.Format("2006/01/02 - 15:04:05.000"))
}

func (z *_zap) GetZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := global.CONFIG.TransportLevel(); level <= zapcore.FatalLevel; level++ {
		cores = append(cores, z.GetEncoderCore(level, z.GetLevelPriority(level)))
	}
	return cores
}

func (z *_zap) GetLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool {
			return level == zap.DebugLevel
		}
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool {
			return level == zap.InfoLevel
		}
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool {
			return level == zap.WarnLevel
		}
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool {
			return level == zap.ErrorLevel
		}
	case zapcore.DPanicLevel:
		return func(level zapcore.Level) bool {
			return level == zap.DPanicLevel
		}
	case zapcore.PanicLevel:
		return func(level zapcore.Level) bool {
			return level == zap.PanicLevel
		}
	case zapcore.FatalLevel:
		return func(level zapcore.Level) bool {
			return level == zap.FatalLevel
		}
	default:
		return func(level zapcore.Level) bool {
			return level == zap.DebugLevel
		}
	}
}
