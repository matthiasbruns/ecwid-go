package orders

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/cli/internal/cmdutil"
	apiorders "github.com/matthiasbruns/ecwid-go/ecwid/orders"
)

// Cmd is the top-level orders command.
var Cmd = &cobra.Command{
	Use:   "orders",
	Short: "Manage orders",
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List orders",
	RunE: func(cmd *cobra.Command, _ []string) error {
		opts := &apiorders.SearchOptions{}

		if v, _ := cmd.Flags().GetString("payment-status"); v != "" {
			opts.PaymentStatus = v
		}
		if v, _ := cmd.Flags().GetString("fulfillment-status"); v != "" {
			opts.FulfillmentStatus = v
		}
		if v, _ := cmd.Flags().GetString("email"); v != "" {
			opts.Email = v
		}
		if v, _ := cmd.Flags().GetString("keyword"); v != "" {
			opts.Keywords = v
		}
		limit, err := cmdutil.GetNonNegativeInt(cmd, "limit")
		if err != nil {
			return err
		}
		opts.Limit = limit

		offset, err := cmdutil.GetNonNegativeInt(cmd, "offset")
		if err != nil {
			return err
		}
		opts.Offset = offset

		result, err := cmdutil.AppClient.Orders.Search(cmd.Context(), opts)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

var getCmd = &cobra.Command{
	Use:   "get <orderID>",
	Short: "Get an order by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := cmdutil.AppClient.Orders.Get(cmd.Context(), args[0])
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new order (JSON via --data or stdin)",
	RunE: func(cmd *cobra.Command, _ []string) error {
		rawData, err := cmdutil.ReadJSONInput(cmd)
		if err != nil {
			return err
		}

		var req apiorders.CreateRequest
		if err := json.Unmarshal(rawData, &req); err != nil {
			return fmt.Errorf("parse JSON: %w", err)
		}

		result, err := cmdutil.AppClient.Orders.Create(cmd.Context(), &req)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

var updateCmd = &cobra.Command{
	Use:   "update <orderID>",
	Short: "Update an order by ID (JSON via --data or stdin)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		rawData, err := cmdutil.ReadJSONInput(cmd)
		if err != nil {
			return err
		}

		var req apiorders.UpdateRequest
		if err := json.Unmarshal(rawData, &req); err != nil {
			return fmt.Errorf("parse JSON: %w", err)
		}

		result, err := cmdutil.AppClient.Orders.Update(cmd.Context(), args[0], &req)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete <orderID>",
	Short: "Delete an order by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := cmdutil.AppClient.Orders.Delete(cmd.Context(), args[0])
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

func init() {
	Cmd.AddCommand(listCmd)
	Cmd.AddCommand(getCmd)
	Cmd.AddCommand(createCmd)
	Cmd.AddCommand(updateCmd)
	Cmd.AddCommand(deleteCmd)

	// List flags.
	listCmd.Flags().String("payment-status", "", "Filter by payment status")
	listCmd.Flags().String("fulfillment-status", "", "Filter by fulfillment status")
	listCmd.Flags().String("email", "", "Filter by customer email")
	listCmd.Flags().String("keyword", "", "Filter by keyword")
	listCmd.Flags().Int("limit", 0, "Maximum number of orders to return")
	listCmd.Flags().Int("offset", 0, "Number of orders to skip")

	// Create/update data flag.
	createCmd.Flags().String("data", "", "Order JSON (reads from stdin if omitted)")
	updateCmd.Flags().String("data", "", "Order update JSON (reads from stdin if omitted)")
}
