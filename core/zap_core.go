package core

import (
	"fmt"
	"observ/zlog/global"
	"observ/zlog/internal"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Zap() (logger *zap.Logger) {
	if ok, _ := internal.PathExists(global.CONFIG.Director); !ok {
		fmt.Printf("create %v directory\n", global.CONFIG.Director)
		_ = os.Mkdir(global.CONFIG.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.CONFIG.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
