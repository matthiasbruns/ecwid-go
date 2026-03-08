package promotions

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/cli/internal/cmdutil"
	"github.com/matthiasbruns/ecwid-go/ecwid/discounts"
)

// Cmd is the top-level promotions command.
var Cmd = &cobra.Command{
	Use:   "promotions",
	Short: "Manage promotions",
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List promotions",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		result, err := cmdutil.AppClient.Discounts.SearchPromotions(cmd.Context())
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a promotion (reads JSON from stdin or --file)",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		data, err := cmdutil.ReadInput(cmd)
		if err != nil {
			return err
		}

		var p discounts.Promotion
		if err := json.Unmarshal(data, &p); err != nil {
			return fmt.Errorf("invalid promotion JSON: %w", err)
		}

		result, err := cmdutil.AppClient.Discounts.CreatePromotion(cmd.Context(), &p)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

var updateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: "Update a promotion (reads JSON from stdin or --file)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid promotion ID: %w", err)
		}
		if id <= 0 {
			return fmt.Errorf("invalid promotion ID: must be a positive integer")
		}

		data, err := cmdutil.ReadInput(cmd)
		if err != nil {
			return err
		}

		var p discounts.Promotion
		if err := json.Unmarshal(data, &p); err != nil {
			return fmt.Errorf("invalid promotion JSON: %w", err)
		}

		if p.ID != 0 && p.ID != id {
			return fmt.Errorf("promotion JSON id %d does not match argument %d", p.ID, id)
		}
		p.ID = id

		result, err := cmdutil.AppClient.Discounts.UpdatePromotion(cmd.Context(), id, &p)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a promotion",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid promotion ID: %w", err)
		}
		if id <= 0 {
			return fmt.Errorf("invalid promotion ID: must be a positive integer")
		}

		result, err := cmdutil.AppClient.Discounts.DeletePromotion(cmd.Context(), id)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

func init() {
	Cmd.AddCommand(listCmd, createCmd, updateCmd, deleteCmd)

	// Create/update flags.
	createCmd.Flags().String("file", "", "path to JSON file (reads stdin if omitted)")
	updateCmd.Flags().String("file", "", "path to JSON file (reads stdin if omitted)")
}
