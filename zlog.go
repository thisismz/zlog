package zlog

import (
	"github.com/thisismz/zlog/configure"
	"github.com/thisismz/zlog/core"
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
		configure.CONF.Director = config[0].Director
		configure.CONF.EncodeLevel = config[0].EncodeLevel
		configure.CONF.Format = config[0].Format
		configure.CONF.Level = config[0].Level
		configure.CONF.MaxAge = config[0].MaxAge
		configure.CONF.Prefix = config[0].Prefix
		configure.CONF.StacktraceKey = config[0].StacktraceKey
		if config[0].LogInConsole == "true" {
			configure.CONF.LogInConsole = true
		} else {
			configure.CONF.LogInConsole = false
		}
		if config[0].ShowLine == "true" {
			configure.CONF.ShowLine = true
		} else {
			configure.CONF.ShowLine = false
		}
		if config[0].SaveInFile == "true" {
			configure.CONF.SaveInFile = true
		} else {
			configure.CONF.SaveInFile = false
		}
	}
	if configure.CONF.Level == "" {
		configure.CONF.Level = "debug"
	}
	if configure.CONF.Prefix == "" {
		configure.CONF.Prefix = "zlog"
	}
	if configure.CONF.Format == "" {
		configure.CONF.Format = "json"
	}
	if configure.CONF.Director == "" {
		configure.CONF.Director = "logs"
	}
	if configure.CONF.EncodeLevel == "" {
		configure.CONF.EncodeLevel = "LowercaseLevelEncoder"
	}
	if configure.CONF.StacktraceKey == "" {
		configure.CONF.StacktraceKey = "stacktrace"
	}
	if configure.CONF.MaxAge == 0 {
		configure.CONF.MaxAge = 7
	}
	if configure.CONF.ShowLine {
		configure.CONF.ShowLine = true
	}
	if configure.CONF.LogInConsole {
		configure.CONF.LogInConsole = true
	}
	if configure.CONF.SaveInFile {
		configure.CONF.SaveInFile = true
	}
	return &zlog{}
}
