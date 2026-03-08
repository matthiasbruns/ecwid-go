package subscriptions

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/cli/internal/cmdutil"
	apisubs "github.com/matthiasbruns/ecwid-go/ecwid/subscriptions"
)

// Cmd is the top-level subscriptions command.
var Cmd = &cobra.Command{
	Use:   "subscriptions",
	Short: "Manage recurring subscriptions",
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List subscriptions",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		opts := &apisubs.SearchOptions{}

		if v, _ := cmd.Flags().GetInt("limit"); v > 0 {
			opts.Limit = v
		}
		if v, _ := cmd.Flags().GetInt("offset"); v > 0 {
			opts.Offset = v
		}
		if v, _ := cmd.Flags().GetString("status"); v != "" {
			opts.Status = v
		}
		if v, _ := cmd.Flags().GetInt64("customer-id"); v > 0 {
			opts.CustomerID = v
		}
		if v, _ := cmd.Flags().GetInt64("product-id"); v > 0 {
			opts.ProductID = v
		}

		result, err := cmdutil.AppClient.Subscriptions.Search(cmd.Context(), opts)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

var getCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a subscription by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid subscription ID: %w", err)
		}
		if id <= 0 {
			return fmt.Errorf("invalid subscription ID: must be a positive integer")
		}

		result, err := cmdutil.AppClient.Subscriptions.Get(cmd.Context(), id)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

var updateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: "Update a subscription (reads JSON from stdin or --file)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid subscription ID: %w", err)
		}
		if id <= 0 {
			return fmt.Errorf("invalid subscription ID: must be a positive integer")
		}

		data, err := cmdutil.ReadInput(cmd)
		if err != nil {
			return err
		}

		var req apisubs.UpdateRequest
		if err := json.Unmarshal(data, &req); err != nil {
			return fmt.Errorf("invalid subscription JSON: %w", err)
		}

		result, err := cmdutil.AppClient.Subscriptions.Update(cmd.Context(), id, &req)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

func init() {
	Cmd.AddCommand(listCmd, getCmd, updateCmd)

	// List flags.
	listCmd.Flags().Int("limit", 0, "Maximum number of results")
	listCmd.Flags().Int("offset", 0, "Offset for pagination")
	listCmd.Flags().String("status", "", "Filter by status")
	listCmd.Flags().Int64("customer-id", 0, "Filter by customer ID")
	listCmd.Flags().Int64("product-id", 0, "Filter by product ID")

	// Update flags.
	updateCmd.Flags().String("file", "", "path to JSON file (reads stdin if omitted)")
}
