package zlog

import (
	"github.com/thisismz/zlog/core"
	"github.com/thisismz/zlog/global"

	"go.uber.org/zap"
)

type zlog struct {
}

type Config struct {
	Level         string
	Prefix        string
	Format        string
	Director      string
	EncodeLevel   string
	StacktraceKey string
	MaxAge        int
	ShowLine      string
	LogInConsole  string
	SaveInFile    string
}

func (z *zlog) Log() *zap.Logger {
	return core.Zap()
}

func New(config ...Config) *zlog {
	if len(config) > 0 {
		global.CONFIG.Director = config[0].Director
		global.CONFIG.EncodeLevel = config[0].EncodeLevel
		global.CONFIG.Format = config[0].Format
		global.CONFIG.Level = config[0].Level
		global.CONFIG.MaxAge = config[0].MaxAge
		global.CONFIG.Prefix = config[0].Prefix
		global.CONFIG.StacktraceKey = config[0].StacktraceKey
		if config[0].LogInConsole == "true" {
			global.CONFIG.LogInConsole = true
		} else {
			global.CONFIG.LogInConsole = false
		}
		if config[0].ShowLine == "true" {
			global.CONFIG.ShowLine = true
		} else {
			global.CONFIG.ShowLine = false
		}
		if config[0].SaveInFile == "true" {
			global.CONFIG.SaveInFile = true
		} else {
			global.CONFIG.SaveInFile = false
		}
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
	if global.CONFIG.ShowLine {
		global.CONFIG.ShowLine = true
	}
	if global.CONFIG.LogInConsole {
		global.CONFIG.LogInConsole = true
	}
	if global.CONFIG.SaveInFile {
		global.CONFIG.SaveInFile = true
	}
	return &zlog{}
}
