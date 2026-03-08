package profile

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/cli/internal/cmdutil"
	apiprofile "github.com/matthiasbruns/ecwid-go/ecwid/profile"
)

// Cmd is the top-level profile command.
var Cmd = &cobra.Command{
	Use:   "profile",
	Short: "Manage store profile",
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get store profile",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		result, err := cmdutil.AppClient.Profile.Get(cmd.Context())
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update store profile (reads JSON from stdin or --file)",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		data, err := cmdutil.ReadInput(cmd)
		if err != nil {
			return err
		}

		var req apiprofile.UpdateRequest
		if err := json.Unmarshal(data, &req); err != nil {
			return fmt.Errorf("invalid profile JSON: %w", err)
		}

		result, err := cmdutil.AppClient.Profile.Update(cmd.Context(), &req)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

func init() {
	Cmd.AddCommand(getCmd, updateCmd)
	updateCmd.Flags().String("file", "", "path to JSON file (reads stdin if omitted)")
}
