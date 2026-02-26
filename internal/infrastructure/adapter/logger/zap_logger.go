package logger

import (
	"context"
	"go-subscription-service/internal/application/port"
	"log/slog"
	"os"
)

type SlogLogger struct {
	logger *slog.Logger
}

func NewSlogLogger(level slog.Level) *SlogLogger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})
	return &SlogLogger{
		logger: slog.New(handler),
	}
}

func (l *SlogLogger) Info(ctx context.Context, msg string, fields ...port.Field) {
	l.logger.InfoContext(ctx, msg, convert(fields)...)
}

func (l *SlogLogger) Error(ctx context.Context, msg string, fields ...port.Field) {
	l.logger.ErrorContext(ctx, msg, convert(fields)...)
}

func (l *SlogLogger) Debug(ctx context.Context, msg string, fields ...port.Field) {
	l.logger.DebugContext(ctx, msg, convert(fields)...)
}

func convert(fields []port.Field) []any {
	args := make([]any, 0, len(fields)*2)
	for _, f := range fields {
		args = append(args, f.Key, f.Value)
	}
	return args
}
