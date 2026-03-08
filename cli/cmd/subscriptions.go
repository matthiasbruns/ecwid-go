package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/ecwid/subscriptions"
)

var subscriptionsCmd = &cobra.Command{
	Use:   "subscriptions",
	Short: "Manage recurring subscriptions",
}

var subscriptionsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List subscriptions",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		opts := &subscriptions.SearchOptions{}

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

		result, err := AppClient.Subscriptions.Search(cmd.Context(), opts)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

var subscriptionsGetCmd = &cobra.Command{
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

		result, err := AppClient.Subscriptions.Get(cmd.Context(), id)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

var subscriptionsUpdateCmd = &cobra.Command{
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

		data, err := readInput(cmd)
		if err != nil {
			return err
		}

		var req subscriptions.UpdateRequest
		if err := json.Unmarshal(data, &req); err != nil {
			return fmt.Errorf("invalid subscription JSON: %w", err)
		}

		result, err := AppClient.Subscriptions.Update(cmd.Context(), id, &req)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

func init() {
	rootCmd.AddCommand(subscriptionsCmd)
	subscriptionsCmd.AddCommand(subscriptionsListCmd)
	subscriptionsCmd.AddCommand(subscriptionsGetCmd)
	subscriptionsCmd.AddCommand(subscriptionsUpdateCmd)

	// List flags.
	subscriptionsListCmd.Flags().Int("limit", 0, "Maximum number of results")
	subscriptionsListCmd.Flags().Int("offset", 0, "Offset for pagination")
	subscriptionsListCmd.Flags().String("status", "", "Filter by status")
	subscriptionsListCmd.Flags().Int64("customer-id", 0, "Filter by customer ID")
	subscriptionsListCmd.Flags().Int64("product-id", 0, "Filter by product ID")

	// Update flags.
	subscriptionsUpdateCmd.Flags().String("file", "", "path to JSON file (reads stdin if omitted)")
}
