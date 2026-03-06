package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/config"
	"github.com/matthiasbruns/ecwid-go/ecwid"
)

var (
	cfgFile  string
	storeID  string
	token    string
	output   string
	logLevel string

	// AppClient is the initialized Ecwid API client, available to subcommands.
	AppClient *ecwid.Client
)

var rootCmd = &cobra.Command{
	Use:   "ecwid",
	Short: "Ecwid API client CLI",
	Long:  "Command-line interface for the Ecwid REST API.\nManage products, orders, customers, and more from your terminal.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Skip client init for version command.
		if cmd.Name() == "version" {
			return nil
		}

		overrides := config.Config{
			StoreID: storeID,
			Token:   token,
			Output:  output,
		}

		cfg, err := config.Load(cfgFile, overrides)
		if err != nil {
			return fmt.Errorf("load config: %w", err)
		}

		// Apply log level from config.
		setLogLevel(cfg.LogLevel)

		if err := cfg.Validate(); err != nil {
			return err
		}

		AppClient = ecwid.NewClient(cfg, ecwid.WithLogger(slog.Default()))
		return nil
	},
}

// Execute runs the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default: ~/.ecwid.yaml)")
	rootCmd.PersistentFlags().StringVar(&storeID, "store-id", "", "Ecwid store ID (env: ECWID_STORE_ID)")
	rootCmd.PersistentFlags().StringVar(&token, "token", "", "API access token (env: ECWID_TOKEN)")
	rootCmd.PersistentFlags().StringVar(&output, "output", "", "output format: json|table (default: json)")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "", "log level: debug|info|warn|error (default: info)")
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
