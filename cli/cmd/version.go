package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var appVersion = "dev"

// SetVersion sets the application version (called from main).
func SetVersion(v string) {
	appVersion = v
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	RunE: func(cmd *cobra.Command, _ []string) error {
		_, err := fmt.Fprintf(cmd.OutOrStdout(), "ecwid-cli %s\n", appVersion)
		return err
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
