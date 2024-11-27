package logger

import (
	"main/internal/lib/logger/prettylog"
	"log/slog"
	"os"
)

func NewLogger(level string) *slog.Logger {
	opts := prettylog.PrettyHandlerOptions{}
	
	if level == "debug" {
		opts = prettylog.PrettyHandlerOptions{
			SlogOpts: &slog.HandlerOptions{
				Level: slog.LevelDebug,
			},
		}
	}

	if level == "info" {
		opts = prettylog.PrettyHandlerOptions{
			SlogOpts: &slog.HandlerOptions{
				Level: slog.LevelInfo,
			},
		}
	}

	handler := opts.NewPrettyHandler(os.Stdout)
	log := slog.New(handler)
	slog.SetDefault(log)

	return log
}
