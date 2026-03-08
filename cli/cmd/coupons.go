package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/ecwid/discounts"
)

var couponsCmd = &cobra.Command{
	Use:   "coupons",
	Short: "Manage discount coupons",
}

var couponsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List discount coupons",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		opts := &discounts.CouponSearchOptions{}

		if v, _ := cmd.Flags().GetInt("limit"); v > 0 {
			opts.Limit = v
		}
		if v, _ := cmd.Flags().GetInt("offset"); v > 0 {
			opts.Offset = v
		}
		if v, _ := cmd.Flags().GetString("code"); v != "" {
			opts.Code = v
		}

		result, err := AppClient.Discounts.SearchCoupons(cmd.Context(), opts)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

var couponsGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a coupon by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid coupon ID: %w", err)
		}
		if id <= 0 {
			return fmt.Errorf("invalid coupon ID: must be a positive integer")
		}

		result, err := AppClient.Discounts.GetCoupon(cmd.Context(), id)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

var couponsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a coupon (reads JSON from stdin or --file)",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		data, err := readInput(cmd)
		if err != nil {
			return err
		}

		var c discounts.Coupon
		if err := json.Unmarshal(data, &c); err != nil {
			return fmt.Errorf("invalid coupon JSON: %w", err)
		}

		result, err := AppClient.Discounts.CreateCoupon(cmd.Context(), &c)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

var couponsUpdateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: "Update a coupon (reads JSON from stdin or --file)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid coupon ID: %w", err)
		}
		if id <= 0 {
			return fmt.Errorf("invalid coupon ID: must be a positive integer")
		}

		data, err := readInput(cmd)
		if err != nil {
			return err
		}

		var c discounts.Coupon
		if err := json.Unmarshal(data, &c); err != nil {
			return fmt.Errorf("invalid coupon JSON: %w", err)
		}

		result, err := AppClient.Discounts.UpdateCoupon(cmd.Context(), id, &c)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

var couponsDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a coupon",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid coupon ID: %w", err)
		}
		if id <= 0 {
			return fmt.Errorf("invalid coupon ID: must be a positive integer")
		}

		result, err := AppClient.Discounts.DeleteCoupon(cmd.Context(), id)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

func init() {
	rootCmd.AddCommand(couponsCmd)
	couponsCmd.AddCommand(couponsListCmd)
	couponsCmd.AddCommand(couponsGetCmd)
	couponsCmd.AddCommand(couponsCreateCmd)
	couponsCmd.AddCommand(couponsUpdateCmd)
	couponsCmd.AddCommand(couponsDeleteCmd)

	// List flags.
	couponsListCmd.Flags().Int("limit", 0, "Maximum number of results")
	couponsListCmd.Flags().Int("offset", 0, "Offset for pagination")
	couponsListCmd.Flags().String("code", "", "Filter by coupon code")

	// Create/update flags.
	couponsCreateCmd.Flags().String("file", "", "path to JSON file (reads stdin if omitted)")
	couponsUpdateCmd.Flags().String("file", "", "path to JSON file (reads stdin if omitted)")
}
