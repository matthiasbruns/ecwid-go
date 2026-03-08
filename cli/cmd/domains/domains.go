package domains

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/cli/internal/cmdutil"
	apidomains "github.com/matthiasbruns/ecwid-go/ecwid/domains"
)

// Cmd is the top-level domains command.
var Cmd = &cobra.Command{
	Use:   "domains",
	Short: "Manage store domains",
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get store domain settings",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		result, err := cmdutil.AppClient.Domains.Get(cmd.Context())
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

var purchaseCmd = &cobra.Command{
	Use:   "purchase",
	Short: "Purchase a domain (reads JSON from stdin or --file)",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, _ []string) error {
		data, err := cmdutil.ReadInput(cmd)
		if err != nil {
			return err
		}

		var req apidomains.PurchaseRequest
		if err := json.Unmarshal(data, &req); err != nil {
			return fmt.Errorf("invalid domain purchase JSON: %w", err)
		}

		result, err := cmdutil.AppClient.Domains.Purchase(cmd.Context(), &req)
		if err != nil {
			return err
		}
		return cmdutil.OutputResult(cmd, result)
	},
}

func init() {
	Cmd.AddCommand(getCmd, purchaseCmd)
	purchaseCmd.Flags().String("file", "", "path to JSON file (reads stdin if omitted)")
}
