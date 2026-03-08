package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/ecwid/products"
)

var productsCmd = &cobra.Command{
	Use:   "products",
	Short: "Manage products",
	Long:  "List, get, create, update, and delete products in your Ecwid store.",
}

var productsListCmd = &cobra.Command{
	Use:   "list",
	Short: "Search products",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		opts := &products.SearchOptions{}

		if v, _ := cmd.Flags().GetString("keyword"); v != "" {
			opts.Keyword = v
		}
		if v, _ := cmd.Flags().GetInt64("category"); v > 0 {
			opts.Category = v
		}
		if v, _ := cmd.Flags().GetInt("limit"); v > 0 {
			opts.Limit = v
		}
		if v, _ := cmd.Flags().GetInt("offset"); v > 0 {
			opts.Offset = v
		}
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

		resp, err := AppClient.Products.Search(cmd.Context(), opts)
		if err != nil {
			return err
		}
		return outputResult(cmd, resp)
	},
}

var productsGetCmd = &cobra.Command{
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

		resp, err := AppClient.Products.Get(cmd.Context(), id)
		if err != nil {
			return err
		}
		return outputResult(cmd, resp)
	},
}

var productsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a product (reads JSON from stdin or --file)",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		data, err := readInput(cmd)
		if err != nil {
			return err
		}

		var p products.Product
		if err := json.Unmarshal(data, &p); err != nil {
			return fmt.Errorf("invalid product JSON: %w", err)
		}

		resp, err := AppClient.Products.Create(cmd.Context(), &p)
		if err != nil {
			return err
		}
		return outputResult(cmd, resp)
	},
}

var productsUpdateCmd = &cobra.Command{
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

		data, err := readInput(cmd)
		if err != nil {
			return err
		}

		var p products.Product
		if err := json.Unmarshal(data, &p); err != nil {
			return fmt.Errorf("invalid product JSON: %w", err)
		}

		resp, err := AppClient.Products.Update(cmd.Context(), id, &p)
		if err != nil {
			return err
		}
		return outputResult(cmd, resp)
	},
}

var productsDeleteCmd = &cobra.Command{
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

		resp, err := AppClient.Products.Delete(cmd.Context(), id)
		if err != nil {
			return err
		}
		return outputResult(cmd, resp)
	},
}

// readInput reads JSON input from --file flag or stdin.
func readInput(cmd *cobra.Command) ([]byte, error) {
	file, _ := cmd.Flags().GetString("file")
	if file != "" {
		data, err := os.ReadFile(file)
		if err != nil {
			return nil, fmt.Errorf("read file %q: %w", file, err)
		}
		return data, nil
	}
	data, err := io.ReadAll(cmd.InOrStdin())
	if err != nil {
		return nil, fmt.Errorf("read stdin: %w", err)
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("no input provided: specify --file or provide JSON on stdin")
	}
	return data, nil
}

func init() {
	rootCmd.AddCommand(productsCmd)
	productsCmd.AddCommand(productsListCmd)
	productsCmd.AddCommand(productsGetCmd)
	productsCmd.AddCommand(productsCreateCmd)
	productsCmd.AddCommand(productsUpdateCmd)
	productsCmd.AddCommand(productsDeleteCmd)

	// List flags.
	productsListCmd.Flags().String("keyword", "", "search keyword")
	productsListCmd.Flags().Int64("category", 0, "filter by category ID")
	productsListCmd.Flags().Int("limit", 0, "maximum number of results")
	productsListCmd.Flags().Int("offset", 0, "offset for pagination")
	productsListCmd.Flags().Bool("enabled", false, "filter by enabled status")
	productsListCmd.Flags().Bool("in-stock", false, "filter by in-stock status")
	productsListCmd.Flags().String("sku", "", "filter by SKU")
	productsListCmd.Flags().String("sort-by", "", "sort order (e.g., NAME_ASC, PRICE_DESC)")

	// Create/update flags.
	productsCreateCmd.Flags().String("file", "", "path to JSON file (reads stdin if omitted)")
	productsUpdateCmd.Flags().String("file", "", "path to JSON file (reads stdin if omitted)")
}
