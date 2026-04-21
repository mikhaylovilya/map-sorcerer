package logger

import (
	"log/slog"
	"mikhaylovilya/map-sorcer/internal/config"
	"os"
)

func NewLogger(env string) *slog.Logger {
	switch env {
	case config.EnvDebug:
		return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case config.EnvProd:
		return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
}
