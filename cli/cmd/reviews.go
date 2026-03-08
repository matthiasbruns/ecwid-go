package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/ecwid/reviews"
)

var reviewsCmd = &cobra.Command{
	Use:   "reviews",
	Short: "Manage product reviews",
}

var reviewsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List product reviews",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		opts := &reviews.SearchOptions{}

		if v, _ := cmd.Flags().GetInt("limit"); v > 0 {
			opts.Limit = v
		}
		if v, _ := cmd.Flags().GetInt("offset"); v > 0 {
			opts.Offset = v
		}
		if v, _ := cmd.Flags().GetString("status"); v != "" {
			opts.Status = v
		}
		if v, _ := cmd.Flags().GetInt64("product-id"); v > 0 {
			opts.ProductID = v
		}

		result, err := AppClient.Reviews.Search(cmd.Context(), opts)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

var reviewsUpdateCmd = &cobra.Command{
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

		result, err := AppClient.Reviews.UpdateStatus(cmd.Context(), id, status)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

var reviewsDeleteCmd = &cobra.Command{
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

		result, err := AppClient.Reviews.Delete(cmd.Context(), id)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

func init() {
	rootCmd.AddCommand(reviewsCmd)
	reviewsCmd.AddCommand(reviewsListCmd)
	reviewsCmd.AddCommand(reviewsUpdateCmd)
	reviewsCmd.AddCommand(reviewsDeleteCmd)

	// List flags.
	reviewsListCmd.Flags().Int("limit", 0, "Maximum number of results")
	reviewsListCmd.Flags().Int("offset", 0, "Offset for pagination")
	reviewsListCmd.Flags().String("status", "", "Filter by status")
	reviewsListCmd.Flags().Int64("product-id", 0, "Filter by product ID")

	// Update flags.
	reviewsUpdateCmd.Flags().String("status", "", "New review status (required)")
}
