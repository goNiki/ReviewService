package logger

import (
	"log/slog"
	"os"
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
