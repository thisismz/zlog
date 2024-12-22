package zlog

import (
	"log/slog"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newlog() *zap.Logger {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	if SaveInFile {
		if ok, _ := pathExists(Director); !ok {
			log.Info("created directory" + Director)
			_ = os.Mkdir(Director, os.ModePerm)
		}
	}

	levels := levels()
	length := len(levels)
	cores := make([]zapcore.Core, 0, length)

	for i := 0; i < length; i++ {
		core := newZapCore(levels[i])
		cores = append(cores, core)
	}
	logger := zap.New(zapcore.NewTee(cores...))
	if ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	if SentryDns != "" {
		logger = modifyToSentryLogger(logger)
	}
	return logger
}
