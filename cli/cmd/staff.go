package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/ecwid/staff"
)

var staffCmd = &cobra.Command{
	Use:   "staff",
	Short: "Manage staff accounts",
}

var staffListCmd = &cobra.Command{
	Use:   "list",
	Short: "List staff accounts",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		result, err := AppClient.Staff.List(cmd.Context())
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

var staffGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a staff account by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := AppClient.Staff.Get(cmd.Context(), args[0])
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

var staffCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Invite a staff member (reads JSON from stdin or --file)",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		data, err := readInput(cmd)
		if err != nil {
			return err
		}

		var req staff.CreateRequest
		if err := json.Unmarshal(data, &req); err != nil {
			return fmt.Errorf("invalid staff JSON: %w", err)
		}

		result, err := AppClient.Staff.Create(cmd.Context(), &req)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

var staffUpdateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: "Update a staff account (reads JSON from stdin or --file)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := readInput(cmd)
		if err != nil {
			return err
		}

		var req staff.UpdateRequest
		if err := json.Unmarshal(data, &req); err != nil {
			return fmt.Errorf("invalid staff JSON: %w", err)
		}

		result, err := AppClient.Staff.Update(cmd.Context(), args[0], &req)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

var staffDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a staff account",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := AppClient.Staff.Delete(cmd.Context(), args[0])
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

func init() {
	rootCmd.AddCommand(staffCmd)
	staffCmd.AddCommand(staffListCmd)
	staffCmd.AddCommand(staffGetCmd)
	staffCmd.AddCommand(staffCreateCmd)
	staffCmd.AddCommand(staffUpdateCmd)
	staffCmd.AddCommand(staffDeleteCmd)

	// Create/update flags.
	staffCreateCmd.Flags().String("file", "", "path to JSON file (reads stdin if omitted)")
	staffUpdateCmd.Flags().String("file", "", "path to JSON file (reads stdin if omitted)")
}
