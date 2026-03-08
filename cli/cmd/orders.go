package cmd

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/ecwid/orders"
)

var ordersCmd = &cobra.Command{
	Use:   "orders",
	Short: "Manage orders",
}

var ordersListCmd = &cobra.Command{
	Use:   "list",
	Short: "List orders",
	RunE: func(cmd *cobra.Command, _ []string) error {
		opts := &orders.SearchOptions{}

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
		if v, _ := cmd.Flags().GetInt("limit"); v > 0 {
			opts.Limit = v
		}
		if v, _ := cmd.Flags().GetInt("offset"); v > 0 {
			opts.Offset = v
		}

		result, err := AppClient.Orders.Search(cmd.Context(), opts)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

var ordersGetCmd = &cobra.Command{
	Use:   "get <orderID>",
	Short: "Get an order by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := AppClient.Orders.Get(cmd.Context(), args[0])
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

var ordersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new order (JSON via --data or stdin)",
	RunE: func(cmd *cobra.Command, _ []string) error {
		rawData, err := readJSONInput(cmd)
		if err != nil {
			return err
		}

		var req orders.CreateRequest
		if err := json.Unmarshal(rawData, &req); err != nil {
			return fmt.Errorf("parse JSON: %w", err)
		}

		result, err := AppClient.Orders.Create(cmd.Context(), &req)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

var ordersUpdateCmd = &cobra.Command{
	Use:   "update <orderID>",
	Short: "Update an order by ID (JSON via --data or stdin)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		rawData, err := readJSONInput(cmd)
		if err != nil {
			return err
		}

		var req orders.UpdateRequest
		if err := json.Unmarshal(rawData, &req); err != nil {
			return fmt.Errorf("parse JSON: %w", err)
		}

		result, err := AppClient.Orders.Update(cmd.Context(), args[0], &req)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

var ordersDeleteCmd = &cobra.Command{
	Use:   "delete <orderID>",
	Short: "Delete an order by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := AppClient.Orders.Delete(cmd.Context(), args[0])
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

// readJSONInput reads JSON from the --data flag, or stdin if the flag is empty.
func readJSONInput(cmd *cobra.Command) ([]byte, error) {
	if data, _ := cmd.Flags().GetString("data"); data != "" {
		return []byte(data), nil
	}
	raw, err := io.ReadAll(cmd.InOrStdin())
	if err != nil {
		return nil, fmt.Errorf("read input: %w", err)
	}
	if len(raw) == 0 {
		return nil, fmt.Errorf("no input: use --data flag or pipe JSON to stdin")
	}
	return raw, nil
}

func init() {
	rootCmd.AddCommand(ordersCmd)
	ordersCmd.AddCommand(ordersListCmd)
	ordersCmd.AddCommand(ordersGetCmd)
	ordersCmd.AddCommand(ordersCreateCmd)
	ordersCmd.AddCommand(ordersUpdateCmd)
	ordersCmd.AddCommand(ordersDeleteCmd)

	// List flags.
	ordersListCmd.Flags().String("payment-status", "", "Filter by payment status")
	ordersListCmd.Flags().String("fulfillment-status", "", "Filter by fulfillment status")
	ordersListCmd.Flags().String("email", "", "Filter by customer email")
	ordersListCmd.Flags().String("keyword", "", "Filter by keyword")
	ordersListCmd.Flags().Int("limit", 0, "Maximum number of orders to return")
	ordersListCmd.Flags().Int("offset", 0, "Number of orders to skip")

	// Create/update data flag.
	ordersCreateCmd.Flags().String("data", "", "Order JSON (reads from stdin if omitted)")
	ordersUpdateCmd.Flags().String("data", "", "Order update JSON (reads from stdin if omitted)")
}
