package cmd

import (
	"fmt"
	"io"
	"log/slog"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/matthiasbruns/ecwid-go/mock/internal/config"
	"github.com/matthiasbruns/ecwid-go/mock/internal/server"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the mock server",
	Long:  "Start the mock HTTP server: the admin shell, the simulated Ecwid REST API,\nand the mock's control API.",
	RunE:  runServe,
	// A config/validation error from RunE is not a usage error, so don't bury
	// the message under a flags dump. Cobra still prints flag-parse errors.
	SilenceUsage: true,
}

func init() {
	f := serveCmd.Flags()
	f.String("app-url", "", "URL of the app to iframe, e.g. http://localhost:3000 (env: ECWID_MOCK_APP_URL) (required)")
	f.String("client-id", config.DefaultClientID, "app client_id; also EcwidApp.init({app_id}) (env: ECWID_MOCK_CLIENT_ID)")
	f.String("client-secret", "", "signs webhooks and encrypts payloads; must be >=16 bytes; generated if unset (env: ECWID_MOCK_CLIENT_SECRET)")
	f.String("store-id", config.DefaultStoreID, "store ID placed in the auth payload (env: ECWID_MOCK_STORE_ID)")
	f.String("auth-mode", config.DefaultAuthMode, "iframe auth mode: default (hex fragment) | enhanced (AES query) (env: ECWID_MOCK_AUTH_MODE)")
	f.String("webhook-url", "", "where triggered webhooks POST (env: ECWID_MOCK_WEBHOOK_URL)")
	f.Int("port", config.DefaultPort, "listen port (env: ECWID_MOCK_PORT)")
	f.String("proxy-store", "", "store ID to forward unimplemented REST calls to (env: ECWID_MOCK_PROXY_STORE)")
	f.String("proxy-token", "", "access token for the proxy store (env: ECWID_MOCK_PROXY_TOKEN)")

	rootCmd.AddCommand(serveCmd)
}

func runServe(cmd *cobra.Command, _ []string) error {
	cfg, err := config.Load(cmd)
	if err != nil {
		return err
	}
	if err := cfg.Validate(); err != nil {
		return err
	}

	if err := printBanner(cmd.OutOrStdout(), &cfg); err != nil {
		return err
	}

	ctx, stop := signal.NotifyContext(cmd.Context(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := server.New(cfg, slog.Default())
	return srv.Run(ctx)
}

// printBanner writes the human-facing startup summary. It prints the generated
// client_secret only when the mock generated it, so the developer can configure
// their app to match — a user-supplied secret is never printed.
func printBanner(w io.Writer, cfg *config.Config) error {
	banner := "ecwid-mock\n" +
		fmt.Sprintf("  admin:      http://localhost:%d/\n", cfg.Port) +
		fmt.Sprintf("  app URL:    %s\n", cfg.AppURL) +
		fmt.Sprintf("  store ID:   %s\n", cfg.StoreID) +
		fmt.Sprintf("  auth mode:  %s\n", cfg.AuthMode) +
		fmt.Sprintf("  client_id:  %s\n", cfg.ClientID)
	if cfg.SecretGenerated {
		banner += fmt.Sprintf("  client_secret (generated): %s\n", cfg.ClientSecret) +
			"  ^ configure your app with this secret; override with --client-secret\n"
	}
	_, err := io.WriteString(w, banner)
	return err
}
