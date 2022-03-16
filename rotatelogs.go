package zlog

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap/zapcore"
)

func GetWriteSyncer(file string, logInConsole bool) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file, // The location of the log file
		MaxSize:    10,   // Maximum size of the log file (in MB) before the cut is made
		MaxBackups: 200,  // The maximum number of old files to keep
		MaxAge:     30,   // The maximum number of days to retain old files
		Compress:   true, // Whether to compress/archive old files
	}

	if logInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return zapcore.AddSync(lumberJackLogger)
}
