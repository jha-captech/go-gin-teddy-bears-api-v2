package config

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

func LoggerInit() {
	w := os.Stderr
	// set global logger with custom options
	slog.SetDefault(slog.New(
		tint.NewHandler(w, &tint.Options{
			Level: slog.LevelDebug,
		}),
	))
}
