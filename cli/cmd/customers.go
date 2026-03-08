package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/ecwid/customers"
)

var customersCmd = &cobra.Command{
	Use:   "customers",
	Short: "Manage Ecwid customers",
}

// customers list

var customersListCmd = &cobra.Command{
	Use:   "list",
	Short: "Search / list customers",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		keyword, _ := cmd.Flags().GetString("keyword")
		email, _ := cmd.Flags().GetString("email")
		limit, _ := cmd.Flags().GetInt("limit")
		offset, _ := cmd.Flags().GetInt("offset")

		opts := &customers.SearchOptions{
			Keyword: keyword,
			Email:   email,
			Limit:   limit,
			Offset:  offset,
		}

		result, err := AppClient.Customers.Search(cmd.Context(), opts)
		if err != nil {
			return err
		}

		return outputResult(cmd, result.Items)
	},
}

// customers get

var customersGetCmd = &cobra.Command{
	Use:   "get <customerID>",
	Short: "Get a customer by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid customer ID %q: %w", args[0], err)
		}
		if id <= 0 {
			return fmt.Errorf("invalid customer ID %q: must be a positive integer", args[0])
		}

		result, err := AppClient.Customers.Get(cmd.Context(), id)
		if err != nil {
			return err
		}

		return outputResult(cmd, result)
	},
}

// customers create

var customersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new customer (reads JSON from stdin)",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		data, err := io.ReadAll(cmd.InOrStdin())
		if err != nil {
			return fmt.Errorf("read stdin: %w", err)
		}
		if len(data) == 0 {
			return fmt.Errorf("no input provided: pipe JSON to stdin")
		}

		var cust customers.Customer
		if err := json.Unmarshal(data, &cust); err != nil {
			return fmt.Errorf("parse customer JSON: %w", err)
		}

		result, err := AppClient.Customers.Create(cmd.Context(), &cust)
		if err != nil {
			return err
		}

		return outputResult(cmd, result)
	},
}

// customers update

var customersUpdateCmd = &cobra.Command{
	Use:   "update <customerID>",
	Short: "Update a customer by ID (reads JSON from stdin)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid customer ID %q: %w", args[0], err)
		}
		if id <= 0 {
			return fmt.Errorf("invalid customer ID %q: must be a positive integer", args[0])
		}

		data, err := io.ReadAll(cmd.InOrStdin())
		if err != nil {
			return fmt.Errorf("read stdin: %w", err)
		}
		if len(data) == 0 {
			return fmt.Errorf("no input provided: pipe JSON to stdin")
		}

		var cust customers.Customer
		if err := json.Unmarshal(data, &cust); err != nil {
			return fmt.Errorf("parse customer JSON: %w", err)
		}

		result, err := AppClient.Customers.Update(cmd.Context(), id, &cust)
		if err != nil {
			return err
		}

		return outputResult(cmd, result)
	},
}

// customers delete

var customersDeleteCmd = &cobra.Command{
	Use:   "delete <customerID>",
	Short: "Delete a customer by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid customer ID %q: %w", args[0], err)
		}
		if id <= 0 {
			return fmt.Errorf("invalid customer ID %q: must be a positive integer", args[0])
		}

		result, err := AppClient.Customers.Delete(cmd.Context(), id)
		if err != nil {
			return err
		}

		return outputResult(cmd, result)
	},
}

func init() {
	customersListCmd.Flags().String("keyword", "", "Search keyword (name, email, etc.)")
	customersListCmd.Flags().String("email", "", "Filter by email address")
	customersListCmd.Flags().Int("limit", 0, "Maximum number of results")
	customersListCmd.Flags().Int("offset", 0, "Offset for pagination")

	customersCmd.AddCommand(customersListCmd)
	customersCmd.AddCommand(customersGetCmd)
	customersCmd.AddCommand(customersCreateCmd)
	customersCmd.AddCommand(customersUpdateCmd)
	customersCmd.AddCommand(customersDeleteCmd)

	rootCmd.AddCommand(customersCmd)
}
