package staff

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/cli/internal/cmdutil"
	apistaff "github.com/matthiasbruns/ecwid-go/ecwid/staff"
)

// Cmd is the top-level staff command.
var Cmd = &cobra.Command{
	Use:   "staff",
	Short: "Manage staff accounts",
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List staff accounts",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		result, err := cmdutil.AppClient.Staff.List(cmd.Context())
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

var getCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a staff account by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := cmdutil.AppClient.Staff.Get(cmd.Context(), args[0])
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Invite a staff member (reads JSON from stdin or --file)",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		data, err := cmdutil.ReadInput(cmd)
		if err != nil {
			return err
		}

		var req apistaff.CreateRequest
		if err := json.Unmarshal(data, &req); err != nil {
			return fmt.Errorf("invalid staff JSON: %w", err)
		}

		result, err := cmdutil.AppClient.Staff.Create(cmd.Context(), &req)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

var updateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: "Update a staff account (reads JSON from stdin or --file)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := cmdutil.ReadInput(cmd)
		if err != nil {
			return err
		}

		var req apistaff.UpdateRequest
		if err := json.Unmarshal(data, &req); err != nil {
			return fmt.Errorf("invalid staff JSON: %w", err)
		}

		result, err := cmdutil.AppClient.Staff.Update(cmd.Context(), args[0], &req)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a staff account",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := cmdutil.AppClient.Staff.Delete(cmd.Context(), args[0])
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

func init() {
	Cmd.AddCommand(listCmd, getCmd, createCmd, updateCmd, deleteCmd)

	// Create/update flags.
	createCmd.Flags().String("file", "", "path to JSON file (reads stdin if omitted)")
	updateCmd.Flags().String("file", "", "path to JSON file (reads stdin if omitted)")
}
