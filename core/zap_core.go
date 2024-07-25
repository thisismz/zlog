package core

import (
	"fmt"
	"os"

	"github.com/thisismz/zlog/configure"
	"github.com/thisismz/zlog/internal"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Zap() (logger *zap.Logger) {
	if configure.CONF.SaveInFile == true {
		if ok, _ := internal.PathExists(configure.CONF.Director); !ok {
			fmt.Printf("create %v directory\n", configure.CONF.Director)
			_ = os.Mkdir(configure.CONF.Director, os.ModePerm)
		}
	}
	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if configure.CONF.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
