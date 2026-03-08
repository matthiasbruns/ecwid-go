package reviews

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/cli/internal/cmdutil"
	apireviews "github.com/matthiasbruns/ecwid-go/ecwid/reviews"
)

// Cmd is the top-level reviews command.
var Cmd = &cobra.Command{
	Use:   "reviews",
	Short: "Manage product reviews",
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List product reviews",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		opts := &apireviews.SearchOptions{}

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

		if v, _ := cmd.Flags().GetString("status"); v != "" {
			opts.Status = v
		}

		productID, err := cmdutil.GetNonNegativeInt64(cmd, "product-id")
		if err != nil {
			return err
		}
		opts.ProductID = productID

		result, err := cmdutil.AppClient.Reviews.Search(cmd.Context(), opts)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

var updateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: "Update a review status",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid review ID: %w", err)
		}
		if id <= 0 {
			return fmt.Errorf("invalid review ID: must be a positive integer")
		}

		status, _ := cmd.Flags().GetString("status")
		if status == "" {
			return fmt.Errorf("--status flag is required")
		}

		result, err := cmdutil.AppClient.Reviews.UpdateStatus(cmd.Context(), id, status)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a review",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid review ID: %w", err)
		}
		if id <= 0 {
			return fmt.Errorf("invalid review ID: must be a positive integer")
		}

		result, err := cmdutil.AppClient.Reviews.Delete(cmd.Context(), id)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

func init() {
	Cmd.AddCommand(listCmd, updateCmd, deleteCmd)

	// List flags.
	listCmd.Flags().Int("limit", 0, "Maximum number of results")
	listCmd.Flags().Int("offset", 0, "Offset for pagination")
	listCmd.Flags().String("status", "", "Filter by status")
	listCmd.Flags().Int64("product-id", 0, "Filter by product ID")

	// Update flags.
	updateCmd.Flags().String("status", "", "New review status (required)")
}
