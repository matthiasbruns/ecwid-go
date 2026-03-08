package cmd

import (
	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/ecwid/carts"
)

var cartsCmd = &cobra.Command{
	Use:   "carts",
	Short: "Manage abandoned carts",
}

var cartsListCmd = &cobra.Command{
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

		opts := &carts.SearchOptions{
			CustomerID: customerID,
			Limit:      limit,
			Offset:     offset,
		}

		result, err := AppClient.Carts.Search(cmd.Context(), opts)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

var cartsGetCmd = &cobra.Command{
	Use:   "get <cartId>",
	Short: "Get an abandoned cart by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := AppClient.Carts.Get(cmd.Context(), args[0])
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

var cartsUpdateCmd = &cobra.Command{
	Use:   "update <cartId>",
	Short: "Update an abandoned cart",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		req := &carts.UpdateRequest{}
		if cmd.Flags().Changed("hidden") {
			hidden, err := cmd.Flags().GetBool("hidden")
			if err != nil {
				return err
			}
			req.Hidden = &hidden
		}

		result, err := AppClient.Carts.Update(cmd.Context(), args[0], req)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

var cartsPlaceCmd = &cobra.Command{
	Use:   "place <cartId>",
	Short: "Convert an abandoned cart into an order",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := AppClient.Carts.Place(cmd.Context(), args[0])
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

func init() {
	cartsListCmd.Flags().Int64("customer-id", 0, "Filter by customer ID")
	cartsListCmd.Flags().Int("limit", 0, "Maximum number of carts to return (0 = API default)")
	cartsListCmd.Flags().Int("offset", 0, "Number of carts to skip for pagination (0 = start from beginning)")

	cartsUpdateCmd.Flags().Bool("hidden", false, "Mark cart as hidden")

	cartsCmd.AddCommand(cartsListCmd, cartsGetCmd, cartsUpdateCmd, cartsPlaceCmd)
	rootCmd.AddCommand(cartsCmd)
}
