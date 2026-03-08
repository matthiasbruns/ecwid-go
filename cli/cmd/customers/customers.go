package customers

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/cli/internal/cmdutil"
	apicustomers "github.com/matthiasbruns/ecwid-go/ecwid/customers"
)

// Cmd is the top-level customers command.
var Cmd = &cobra.Command{
	Use:   "customers",
	Short: "Manage Ecwid customers",
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Search / list customers",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		keyword, _ := cmd.Flags().GetString("keyword")
		email, _ := cmd.Flags().GetString("email")
		limit, _ := cmd.Flags().GetInt("limit")
		offset, _ := cmd.Flags().GetInt("offset")

		opts := &apicustomers.SearchOptions{
			Keyword: keyword,
			Email:   email,
			Limit:   limit,
			Offset:  offset,
		}

		result, err := cmdutil.AppClient.Customers.Search(cmd.Context(), opts)
		if err != nil {
			return err
		}

		return cmdutil.OutputResult(cmd, result.Items)
	},
}

var getCmd = &cobra.Command{
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

		result, err := cmdutil.AppClient.Customers.Get(cmd.Context(), id)
		if err != nil {
			return err
		}

		return cmdutil.OutputResult(cmd, result)
	},
}

var createCmd = &cobra.Command{
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

		var cust apicustomers.Customer
		if err := json.Unmarshal(data, &cust); err != nil {
			return fmt.Errorf("parse customer JSON: %w", err)
		}

		result, err := cmdutil.AppClient.Customers.Create(cmd.Context(), &cust)
		if err != nil {
			return err
		}

		return cmdutil.OutputResult(cmd, result)
	},
}

var updateCmd = &cobra.Command{
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

		var cust apicustomers.Customer
		if err := json.Unmarshal(data, &cust); err != nil {
			return fmt.Errorf("parse customer JSON: %w", err)
		}

		if cust.ID != 0 && cust.ID != id {
			return fmt.Errorf("customer JSON id %d does not match argument %d", cust.ID, id)
		}

		result, err := cmdutil.AppClient.Customers.Update(cmd.Context(), id, &cust)
		if err != nil {
			return err
		}

		return cmdutil.OutputResult(cmd, result)
	},
}

var deleteCmd = &cobra.Command{
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

		result, err := cmdutil.AppClient.Customers.Delete(cmd.Context(), id)
		if err != nil {
			return err
		}

		return cmdutil.OutputResult(cmd, result)
	},
}

func init() {
	listCmd.Flags().String("keyword", "", "Search keyword (name, email, etc.)")
	listCmd.Flags().String("email", "", "Filter by email address")
	listCmd.Flags().Int("limit", 0, "Maximum number of results")
	listCmd.Flags().Int("offset", 0, "Offset for pagination")

	Cmd.AddCommand(listCmd, getCmd, createCmd, updateCmd, deleteCmd)
}
