package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/ecwid/domains"
)

var domainsCmd = &cobra.Command{
	Use:   "domains",
	Short: "Manage store domains",
}

var domainsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get store domain settings",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		result, err := AppClient.Domains.Get(cmd.Context())
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

var domainsPurchaseCmd = &cobra.Command{
	Use:   "purchase",
	Short: "Purchase a domain (reads JSON from stdin or --file)",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		data, err := readInput(cmd)
		if err != nil {
			return err
		}

		var req domains.PurchaseRequest
		if err := json.Unmarshal(data, &req); err != nil {
			return fmt.Errorf("invalid domain purchase JSON: %w", err)
		}

		result, err := AppClient.Domains.Purchase(cmd.Context(), &req)
		if err != nil {
			return err
		}
		return outputResult(cmd, result)
	},
}

func init() {
	rootCmd.AddCommand(domainsCmd)
	domainsCmd.AddCommand(domainsGetCmd)
	domainsCmd.AddCommand(domainsPurchaseCmd)

	// Purchase flags.
	domainsPurchaseCmd.Flags().String("file", "", "path to JSON file (reads stdin if omitted)")
}
