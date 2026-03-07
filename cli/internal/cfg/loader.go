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

// flagBindings maps CLI flag names (kebab-case) to config keys (snake_case).
var flagBindings = map[string]string{
	"store-id":    "store_id",
	"token":       "token",
	"output":      "output",
	"log-level":   "log_level",
	"base-url":    "base_url",
	"max-retries": "max_retries",
}

// Load reads configuration from file, environment, and flags using viper.
// Precedence (highest to lowest): flags > env > config file > defaults.
func Load(configPath string, cmd *cobra.Command) (config.Config, error) {
	v := viper.New()

	// Config file.
	if configPath != "" {
		v.SetConfigFile(configPath)
	} else {
		home, err := os.UserHomeDir()
		if err == nil {
			v.SetConfigFile(filepath.Join(home, ".ecwid.yaml"))
		}
	}

	// Environment.
	v.SetEnvPrefix("ECWID")
	v.AutomaticEnv()

	// Bind CLI flags so they participate in viper precedence.
	if cmd != nil {
		for flagName, cfgKey := range flagBindings {
			if f := cmd.Flags().Lookup(flagName); f != nil {
				if err := v.BindPFlag(cfgKey, f); err != nil {
					return config.Config{}, fmt.Errorf("bind flag %q: %w", flagName, err)
				}
			}
		}
	}

	// Read config file (ignore not-found).
	if err := v.ReadInConfig(); err != nil {
		var notFound viper.ConfigFileNotFoundError
		if !errors.As(err, &notFound) && !os.IsNotExist(err) {
			if _, statErr := os.Stat(v.ConfigFileUsed()); !os.IsNotExist(statErr) {
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
