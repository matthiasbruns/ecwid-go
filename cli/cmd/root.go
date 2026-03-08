package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/cli/cmd/carts"
	"github.com/matthiasbruns/ecwid-go/cli/cmd/categories"
	"github.com/matthiasbruns/ecwid-go/cli/cmd/coupons"
	"github.com/matthiasbruns/ecwid-go/cli/cmd/customers"
	"github.com/matthiasbruns/ecwid-go/cli/cmd/dictionaries"
	"github.com/matthiasbruns/ecwid-go/cli/cmd/domains"
	"github.com/matthiasbruns/ecwid-go/cli/cmd/orders"
	"github.com/matthiasbruns/ecwid-go/cli/cmd/products"
	"github.com/matthiasbruns/ecwid-go/cli/cmd/profile"
	"github.com/matthiasbruns/ecwid-go/cli/cmd/promotions"
	"github.com/matthiasbruns/ecwid-go/cli/cmd/reports"
	"github.com/matthiasbruns/ecwid-go/cli/cmd/reviews"
	"github.com/matthiasbruns/ecwid-go/cli/cmd/staff"
	"github.com/matthiasbruns/ecwid-go/cli/cmd/subscriptions"
	"github.com/matthiasbruns/ecwid-go/cli/internal/cfg"
	"github.com/matthiasbruns/ecwid-go/cli/internal/cmdutil"
	"github.com/matthiasbruns/ecwid-go/ecwid"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "ecwid",
	Short: "Ecwid API client CLI",
	Long:  "Command-line interface for the Ecwid REST API.\nManage products, orders, customers, and more from your terminal.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Skip client init for version command.
		if cmd.Name() == "version" {
			return nil
		}

		loadedCfg, err := cfg.Load(cfgFile, cmd)
		if err != nil {
			return fmt.Errorf("load config: %w", err)
		}

		// Apply log level from config.
		setLogLevel(loadedCfg.LogLevel)

		if err := loadedCfg.Validate(); err != nil {
			return err
		}

		cmdutil.AppClient = ecwid.NewClient(loadedCfg, ecwid.WithLogger(slog.Default()))
		return nil
	},
}

// Execute runs the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default: ~/.ecwid.yaml)")
	rootCmd.PersistentFlags().String("store-id", "", "Ecwid store ID (env: ECWID_STORE_ID)")
	rootCmd.PersistentFlags().String("token", "", "API access token (env: ECWID_TOKEN)")
	rootCmd.PersistentFlags().String("output", "", "output format: json|table (default: json)")
	rootCmd.PersistentFlags().String("log-level", "", "log level: debug|info|warn|error (default: info)")

	// Register domain commands.
	rootCmd.AddCommand(
		carts.Cmd,
		categories.Cmd,
		coupons.Cmd,
		customers.Cmd,
		dictionaries.Cmd,
		domains.Cmd,
		orders.Cmd,
		products.Cmd,
		profile.Cmd,
		promotions.Cmd,
		reports.Cmd,
		reviews.Cmd,
		staff.Cmd,
		subscriptions.Cmd,
	)
}

func setLogLevel(level string) {
	var l slog.Level
	switch level {
	case "debug":
		l = slog.LevelDebug
	case "warn":
		l = slog.LevelWarn
	case "error":
		l = slog.LevelError
	default:
		l = slog.LevelInfo
	}

	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: l}))
	slog.SetDefault(logger)
}
