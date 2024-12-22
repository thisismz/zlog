package zlog

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"testing"
	"time"

	"go.uber.org/zap"
)

func BenchmarkNewLogger(b *testing.B) {
	// Running the benchmark for creating a logger with a custom configuration

	logger := newlog()
	b.Run("CustomConfig", func(b *testing.B) {

		for i := 0; i < b.N; i++ {

			logger.Info("benchmark", zap.Any("I", i))
		}
	})
}

func BenchmarkSlog(b *testing.B) {
	// Setup logger with discard handler
	logger := slog.New(&discardHandler{})

	// Different log levels for benchmark
	levels := []struct {
		name  string
		level slog.Level
	}{
		{"Info", slog.LevelInfo},
		{"Debug", slog.LevelDebug},
		{"Warn", slog.LevelWarn},
		{"Error", slog.LevelError},
	}

	// Different payload sizes for benchmark
	payloads := []struct {
		name    string
		message string
		attrs   []slog.Attr
	}{
		{
			name:    "Small",
			message: "test message",
			attrs: []slog.Attr{
				slog.String("key1", "value1"),
			},
		},
		{
			name:    "Medium",
			message: "test message with more context",
			attrs: []slog.Attr{
				slog.String("key1", "value1"),
				slog.Int("key2", 123),
				slog.String("key3", "value3"),
				slog.Float64("key4", 123.456),
			},
		},
		{
			name:    "Large",
			message: "test message with extensive context",
			attrs: []slog.Attr{
				slog.String("key1", "value1"),
				slog.Int("key2", 123),
				slog.String("key3", "value3"),
				slog.Float64("key4", 123.456),
				slog.Time("key5", time.Now()),
				slog.String("key6", "value6"),
				slog.Int("key7", 789),
				slog.Bool("key8", true),
			},
		},
	}

	// Run benchmarks for different combinations
	for _, level := range levels {
		for _, payload := range payloads {
			name := fmt.Sprintf("Slog/%s/%s", level.name, payload.name)
			b.Run(name, func(b *testing.B) {
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					logger.LogAttrs(context.Background(), level.level, payload.message, payload.attrs...)
				}
			})
		}
	}
}

// Benchmark with JSON handler
func BenchmarkSlogJSON(b *testing.B) {
	// Setup JSON logger with discard writer
	jsonLogger := slog.New(slog.NewJSONHandler(io.Discard, nil))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		jsonLogger.Info("test message",
			"string", "value",
			"int", 123,
			"float", 123.456,
			"bool", true,
		)
	}
}

// Benchmark with Text handler
func BenchmarkSlogText(b *testing.B) {
	// Setup Text logger with discard writer
	textLogger := slog.New(slog.NewTextHandler(io.Discard, nil))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		textLogger.Info("test message",
			"string", "value",
			"int", 123,
			"float", 123.456,
			"bool", true,
		)
	}
}

// discardHandler implements slog.Handler interface and discards all output
type discardHandler struct{}

func (h *discardHandler) Handle(_ context.Context, _ slog.Record) error {
	return nil
}

func (h *discardHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return h
}

func (h *discardHandler) WithGroup(_ string) slog.Handler {
	return h
}

func (h *discardHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return true
}
