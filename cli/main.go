package main

import (
	"log/slog"
	"os"

	"github.com/matthiasbruns/ecwid-go/cli/cmd"
)

// version is set at build time via -ldflags.
var version = "dev"

func main() {
	// Default: JSON structured logging.
	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	cmd.SetVersion(version)

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
