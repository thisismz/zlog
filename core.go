package zlog

import (
	"io"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapCore struct {
	level zapcore.Level
	zapcore.Core
}

func newZapCore(level zapcore.Level) *zapCore {
	entity := &zapCore{level: level}
	syncer := entity.WriteSyncer()
	levelEnabler := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l == level
	})
	entity.Core = zapcore.NewCore(encoder(), syncer, levelEnabler)
	return entity
}

func (z *zapCore) WriteSyncer(formats ...string) zapcore.WriteSyncer {
	cutter := newCutter(
		Director,
		z.level.String(),
		RetentionDay,
		cutterWithLayout(time.DateOnly),
		cutterWithFormats(formats...),
	)
	if LogInConsole {
		multiSyncer := zapcore.NewMultiWriteSyncer(os.Stdout)
		return zapcore.AddSync(multiSyncer)
	} else if SaveInFile {
		return zapcore.AddSync(cutter)
	}
	// temp io.Writer Discard for calls succeed without doing anything
	return zapcore.AddSync(io.Discard)
}

func (z *zapCore) Enabled(level zapcore.Level) bool {
	return z.level == level
}

func (z *zapCore) With(fields []zapcore.Field) zapcore.Core {
	return z.Core.With(fields)
}

func (z *zapCore) Check(entry zapcore.Entry, check *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if z.Enabled(entry.Level) {
		return check.AddCore(entry, z)
	}
	return check
}

func (z *zapCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	for i := 0; i < len(fields); i++ {
		if fields[i].Key == "business" || fields[i].Key == "folder" || fields[i].Key == "directory" {
			syncer := z.WriteSyncer(fields[i].String)
			z.Core = zapcore.NewCore(encoder(), syncer, z.level)
		}
	}
	return z.Core.Write(entry, fields)
}

func (z *zapCore) Sync() error {
	return z.Core.Sync()
}
