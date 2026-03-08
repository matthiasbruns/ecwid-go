package products

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/cli/internal/cmdutil"
	apiproducts "github.com/matthiasbruns/ecwid-go/ecwid/products"
)

// Cmd is the top-level products command.
var Cmd = &cobra.Command{
	Use:   "products",
	Short: "Manage products",
	Long:  "List, get, create, update, and delete products in your Ecwid store.",
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Search products",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		opts := &apiproducts.SearchOptions{}

		if v, _ := cmd.Flags().GetString("keyword"); v != "" {
			opts.Keyword = v
		}
		if v, _ := cmd.Flags().GetInt64("category"); v > 0 {
			opts.Category = v
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
		if cmd.Flags().Changed("enabled") {
			v, _ := cmd.Flags().GetBool("enabled")
			opts.Enabled = &v
		}
		if cmd.Flags().Changed("in-stock") {
			v, _ := cmd.Flags().GetBool("in-stock")
			opts.InStock = &v
		}
		if v, _ := cmd.Flags().GetString("sku"); v != "" {
			opts.SKU = v
		}
		if v, _ := cmd.Flags().GetString("sort-by"); v != "" {
			opts.SortBy = v
		}

		resp, err := cmdutil.AppClient.Products.Search(cmd.Context(), opts)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, resp)
	},
}

var getCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a product by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid product ID: %w", err)
		}
		if id <= 0 {
			return fmt.Errorf("invalid product ID: must be a positive integer")
		}

		resp, err := cmdutil.AppClient.Products.Get(cmd.Context(), id)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, resp)
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a product (reads JSON from stdin or --file)",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		data, err := cmdutil.ReadInput(cmd)
		if err != nil {
			return err
		}

		var p apiproducts.Product
		if err := json.Unmarshal(data, &p); err != nil {
			return fmt.Errorf("invalid product JSON: %w", err)
		}

		resp, err := cmdutil.AppClient.Products.Create(cmd.Context(), &p)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, resp)
	},
}

var updateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: "Update a product (reads JSON from stdin or --file)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid product ID: %w", err)
		}
		if id <= 0 {
			return fmt.Errorf("invalid product ID: must be a positive integer")
		}

		data, err := cmdutil.ReadInput(cmd)
		if err != nil {
			return err
		}

		var p apiproducts.Product
		if err := json.Unmarshal(data, &p); err != nil {
			return fmt.Errorf("invalid product JSON: %w", err)
		}

		if p.ID != 0 && p.ID != id {
			return fmt.Errorf("product JSON id %d does not match argument %d", p.ID, id)
		}
		p.ID = id

		resp, err := cmdutil.AppClient.Products.Update(cmd.Context(), id, &p)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, resp)
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a product",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid product ID: %w", err)
		}
		if id <= 0 {
			return fmt.Errorf("invalid product ID: must be a positive integer")
		}

		resp, err := cmdutil.AppClient.Products.Delete(cmd.Context(), id)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, resp)
	},
}

func init() {
	Cmd.AddCommand(listCmd)
	Cmd.AddCommand(getCmd)
	Cmd.AddCommand(createCmd)
	Cmd.AddCommand(updateCmd)
	Cmd.AddCommand(deleteCmd)

	// List flags.
	listCmd.Flags().String("keyword", "", "search keyword")
	listCmd.Flags().Int64("category", 0, "filter by category ID")
	listCmd.Flags().Int("limit", 0, "maximum number of results")
	listCmd.Flags().Int("offset", 0, "offset for pagination")
	listCmd.Flags().Bool("enabled", false, "filter by enabled status")
	listCmd.Flags().Bool("in-stock", false, "filter by in-stock status")
	listCmd.Flags().String("sku", "", "filter by SKU")
	listCmd.Flags().String("sort-by", "", "sort order (e.g., NAME_ASC, PRICE_DESC)")

	// Create/update flags.
	createCmd.Flags().String("file", "", "path to JSON file (reads stdin if omitted)")
	updateCmd.Flags().String("file", "", "path to JSON file (reads stdin if omitted)")
}
