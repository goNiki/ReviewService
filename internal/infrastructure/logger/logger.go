package logger

import (
	"context"
	"log/slog"
	"os"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/goNiki/ReviewService/internal/infrastructure/logger/sl"
)

type Log struct {
	Log *slog.Logger
}

func InitLogger(env string) *Log {
	var log *slog.Logger
	switch env {
	case "local":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))

	default:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))
	}

	return &Log{
		Log: log,
	}
}

func (l *Log) Error(ctx context.Context, operation string, err error) {
	l.Log.Error("operation Error",
		slog.String("operation", operation),
		slog.String("request_id", middleware.GetReqID(ctx)),
		sl.Error(err))
}
