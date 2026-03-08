package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/ecwid/profile"
)

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Manage store profile",
}

var profileGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get store profile",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		result, err := AppClient.Profile.Get(cmd.Context())
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

var profileUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update store profile (reads JSON from stdin or --file)",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		data, err := readInput(cmd)
		if err != nil {
			return err
		}

		var req profile.UpdateRequest
		if err := json.Unmarshal(data, &req); err != nil {
			return fmt.Errorf("invalid profile JSON: %w", err)
		}

		result, err := AppClient.Profile.Update(cmd.Context(), &req)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)
	profileCmd.AddCommand(profileGetCmd)
	profileCmd.AddCommand(profileUpdateCmd)

	// Update flags.
	profileUpdateCmd.Flags().String("file", "", "path to JSON file (reads stdin if omitted)")
}
