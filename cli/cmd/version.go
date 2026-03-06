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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("ecwid-cli %s\n", appVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
