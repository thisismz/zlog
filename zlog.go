package zlog

import (
	"observ/zlog/config"
	"observ/zlog/core"
	"observ/zlog/global"

	"go.uber.org/zap"
)


var Config = config.Zap{}

type zlog struct{}

func (z *zlog) Log() *zap.Logger {
	// Set default values
	return core.Zap()
}

func New(config ...config.Zap) *zlog {
	if len(config) > 0 {
		global.CONFIG = config[0]
	}
	if global.CONFIG.Level == "" {
		global.CONFIG.Level = "debug"
	}
	if global.CONFIG.Prefix == "" {
		global.CONFIG.Prefix = "zlog"
	}
	if global.CONFIG.Format == "" {
		global.CONFIG.Format = "json"
	}
	if global.CONFIG.Director == "" {
		global.CONFIG.Director = "logs"
	}
	if global.CONFIG.EncodeLevel == "" {
		global.CONFIG.EncodeLevel = "LowercaseLevelEncoder"
	}
	if global.CONFIG.StacktraceKey == "" {
		global.CONFIG.StacktraceKey = "stacktrace"
	}
	if global.CONFIG.MaxAge == 0 {
		global.CONFIG.MaxAge = 7
	}
	if global.CONFIG.ShowLine == true {
		global.CONFIG.ShowLine = true
	}
	if global.CONFIG.LogInConsole == true {
		global.CONFIG.LogInConsole = true
	}
	return &zlog{}
}
