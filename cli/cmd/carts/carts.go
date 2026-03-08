package carts

import (
	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/cli/internal/cmdutil"
	apicarts "github.com/matthiasbruns/ecwid-go/ecwid/carts"
)

// Cmd is the top-level carts command.
var Cmd = &cobra.Command{
	Use:   "carts",
	Short: "Manage abandoned carts",
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List abandoned carts",
	RunE: func(cmd *cobra.Command, _ []string) error {
		customerID, err := cmd.Flags().GetInt64("customer-id")
		if err != nil {
			return err
		}
		limit, err := cmd.Flags().GetInt("limit")
		if err != nil {
			return err
		}
		offset, err := cmd.Flags().GetInt("offset")
		if err != nil {
			return err
		}

		opts := &apicarts.SearchOptions{
			CustomerID: customerID,
			Limit:      limit,
			Offset:     offset,
		}

		result, err := cmdutil.AppClient.Carts.Search(cmd.Context(), opts)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

var getCmd = &cobra.Command{
	Use:   "get <cartId>",
	Short: "Get an abandoned cart by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := cmdutil.AppClient.Carts.Get(cmd.Context(), args[0])
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

var updateCmd = &cobra.Command{
	Use:   "update <cartId>",
	Short: "Update an abandoned cart",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		req := &apicarts.UpdateRequest{}
		if cmd.Flags().Changed("hidden") {
			hidden, err := cmd.Flags().GetBool("hidden")
			if err != nil {
				return err
			}
			req.Hidden = &hidden
		}

		result, err := cmdutil.AppClient.Carts.Update(cmd.Context(), args[0], req)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

var placeCmd = &cobra.Command{
	Use:   "place <cartId>",
	Short: "Convert an abandoned cart into an order",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := cmdutil.AppClient.Carts.Place(cmd.Context(), args[0])
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

func init() {
	listCmd.Flags().Int64("customer-id", 0, "Filter by customer ID")
	listCmd.Flags().Int("limit", 0, "Maximum number of carts to return (0 = API default)")
	listCmd.Flags().Int("offset", 0, "Number of carts to skip for pagination (0 = start from beginning)")

	updateCmd.Flags().Bool("hidden", false, "Mark cart as hidden")

	Cmd.AddCommand(listCmd, getCmd, updateCmd, placeCmd)
}
