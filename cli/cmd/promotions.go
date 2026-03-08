package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/ecwid/discounts"
)

var promotionsCmd = &cobra.Command{
	Use:   "promotions",
	Short: "Manage promotions",
}

var promotionsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List promotions",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		result, err := AppClient.Discounts.SearchPromotions(cmd.Context())
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

var promotionsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a promotion (reads JSON from stdin or --file)",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		data, err := readInput(cmd)
		if err != nil {
			return err
		}

		var p discounts.Promotion
		if err := json.Unmarshal(data, &p); err != nil {
			return fmt.Errorf("invalid promotion JSON: %w", err)
		}

		result, err := AppClient.Discounts.CreatePromotion(cmd.Context(), &p)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

var promotionsUpdateCmd = &cobra.Command{
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

		data, err := readInput(cmd)
		if err != nil {
			return err
		}

		var p discounts.Promotion
		if err := json.Unmarshal(data, &p); err != nil {
			return fmt.Errorf("invalid promotion JSON: %w", err)
		}

		result, err := AppClient.Discounts.UpdatePromotion(cmd.Context(), id, &p)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

var promotionsDeleteCmd = &cobra.Command{
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

		result, err := AppClient.Discounts.DeletePromotion(cmd.Context(), id)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

func init() {
	rootCmd.AddCommand(promotionsCmd)
	promotionsCmd.AddCommand(promotionsListCmd)
	promotionsCmd.AddCommand(promotionsCreateCmd)
	promotionsCmd.AddCommand(promotionsUpdateCmd)
	promotionsCmd.AddCommand(promotionsDeleteCmd)

	// Create/update flags.
	promotionsCreateCmd.Flags().String("file", "", "path to JSON file (reads stdin if omitted)")
	promotionsUpdateCmd.Flags().String("file", "", "path to JSON file (reads stdin if omitted)")
}
