// Package cmd wires the ecwid-mock cobra command tree.
package cmd

import (
	"github.com/spf13/cobra"
)

var appVersion = "dev"

// SetVersion sets the application version, injected at build time and reported
// by --version. Called from main.
func SetVersion(v string) {
	appVersion = v
	rootCmd.Version = v
}

var rootCmd = &cobra.Command{
	Use:     "ecwid-mock",
	Short:   "Local mock of an Ecwid store for embedded-app development",
	Long:    "ecwid-mock runs a local stand-in for an Ecwid store: it serves the app in an\niframe, simulates the REST API, and lets you trigger webhooks — none of which\nEcwid provides tooling for.",
	Version: appVersion,
}

// Execute runs the root command.
func Execute() error {
	return rootCmd.Execute()
}
