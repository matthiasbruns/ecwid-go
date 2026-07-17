// Command ecwid-mock runs a local stand-in for an Ecwid store so a Go developer
// can exercise an embedded app without a real store.
package main

import (
	"log/slog"
	"os"

	"github.com/matthiasbruns/ecwid-go/mock/cmd"
)

// version is set at build time via -ldflags "-X main.version=...".
var version = "dev"

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	cmd.SetVersion(version)

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
