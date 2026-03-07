// Package cfg provides viper-based configuration loading for the CLI.
package cfg

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/matthiasbruns/ecwid-go/config"
)

// Load reads configuration from file, environment, and flags using viper.
// Precedence (highest to lowest): flags > env > config file > defaults.
func Load(configPath string, cmd *cobra.Command) (config.Config, error) {
	v := viper.New()

	// Config file
	if configPath != "" {
		v.SetConfigFile(configPath)
	} else {
		home, err := os.UserHomeDir()
		if err == nil {
			v.SetConfigFile(filepath.Join(home, ".ecwid.yaml"))
		}
	}

	// Environment
	v.SetEnvPrefix("ECWID")
	v.AutomaticEnv()

	// Map flag names (kebab-case) to config keys (snake_case).
	v.RegisterAlias("store_id", "store_id")
	v.RegisterAlias("base_url", "base_url")
	v.RegisterAlias("log_level", "log_level")
	v.RegisterAlias("max_retries", "max_retries")

	// Bind CLI flags so they participate in viper precedence.
	if cmd != nil {
		if f := cmd.Flags().Lookup("store-id"); f != nil {
			_ = v.BindPFlag("store_id", f)
		}
		if f := cmd.Flags().Lookup("token"); f != nil {
			_ = v.BindPFlag("token", f)
		}
		if f := cmd.Flags().Lookup("output"); f != nil {
			_ = v.BindPFlag("output", f)
		}
		if f := cmd.Flags().Lookup("log-level"); f != nil {
			_ = v.BindPFlag("log_level", f)
		}
		if f := cmd.Flags().Lookup("base-url"); f != nil {
			_ = v.BindPFlag("base_url", f)
		}
	}

	// Read config file (ignore not-found).
	if err := v.ReadInConfig(); err != nil {
		var notFound viper.ConfigFileNotFoundError
		if !errors.As(err, &notFound) && !os.IsNotExist(err) {
			// Only fail on real errors, not missing file.
			if _, statErr := os.Stat(v.ConfigFileUsed()); os.IsNotExist(statErr) {
				// File doesn't exist — that's fine.
			} else {
				return config.Config{}, fmt.Errorf("read config: %w", err)
			}
		}
	}

	var cfg config.Config
	if err := v.Unmarshal(&cfg); err != nil {
		return config.Config{}, fmt.Errorf("unmarshal config: %w", err)
	}

	cfg = cfg.WithDefaults()

	return cfg, nil
}
