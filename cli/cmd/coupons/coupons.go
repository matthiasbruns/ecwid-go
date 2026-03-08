package coupons

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/cli/internal/cmdutil"
	"github.com/matthiasbruns/ecwid-go/ecwid/discounts"
)

// Cmd is the top-level coupons command.
var Cmd = &cobra.Command{
	Use:   "coupons",
	Short: "Manage discount coupons",
}

func parsePositiveID(arg, label string) (int64, error) {
	id, err := strconv.ParseInt(arg, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid %s ID: %w", label, err)
	}
	if id <= 0 {
		return 0, fmt.Errorf("invalid %s ID: must be a positive integer", label)
	}
	return id, nil
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List discount coupons",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		opts := &discounts.CouponSearchOptions{}

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

		if v, _ := cmd.Flags().GetString("code"); v != "" {
			opts.Code = v
		}

		result, err := cmdutil.AppClient.Discounts.SearchCoupons(cmd.Context(), opts)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

var getCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a coupon by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parsePositiveID(args[0], "coupon")
		if err != nil {
			return err
		}

		result, err := cmdutil.AppClient.Discounts.GetCoupon(cmd.Context(), id)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a coupon (reads JSON from stdin or --file)",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		data, err := cmdutil.ReadInput(cmd)
		if err != nil {
			return err
		}

		var c discounts.Coupon
		if err := json.Unmarshal(data, &c); err != nil {
			return fmt.Errorf("invalid coupon JSON: %w", err)
		}

		result, err := cmdutil.AppClient.Discounts.CreateCoupon(cmd.Context(), &c)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

var updateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: "Update a coupon (reads JSON from stdin or --file)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parsePositiveID(args[0], "coupon")
		if err != nil {
			return err
		}

		data, err := cmdutil.ReadInput(cmd)
		if err != nil {
			return err
		}

		var c discounts.Coupon
		if err := json.Unmarshal(data, &c); err != nil {
			return fmt.Errorf("invalid coupon JSON: %w", err)
		}

		if c.ID != 0 && c.ID != id {
			return fmt.Errorf("coupon JSON id %d does not match argument %d", c.ID, id)
		}
		c.ID = id

		result, err := cmdutil.AppClient.Discounts.UpdateCoupon(cmd.Context(), id, &c)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a coupon",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parsePositiveID(args[0], "coupon")
		if err != nil {
			return err
		}

		result, err := cmdutil.AppClient.Discounts.DeleteCoupon(cmd.Context(), id)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

func init() {
	Cmd.AddCommand(listCmd, getCmd, createCmd, updateCmd, deleteCmd)

	// List flags.
	listCmd.Flags().Int("limit", 0, "Maximum number of results")
	listCmd.Flags().Int("offset", 0, "Offset for pagination")
	listCmd.Flags().String("code", "", "Filter by coupon code")

	// Create/update flags.
	createCmd.Flags().String("file", "", "path to JSON file (reads stdin if omitted)")
	updateCmd.Flags().String("file", "", "path to JSON file (reads stdin if omitted)")
}
