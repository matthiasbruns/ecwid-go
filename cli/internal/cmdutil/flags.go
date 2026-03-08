package cmdutil

import (
	"fmt"

	"github.com/spf13/cobra"
)

// GetNonNegativeInt returns the int value for a flag, returning an error if negative.
func GetNonNegativeInt(cmd *cobra.Command, name string) (int, error) {
	v, _ := cmd.Flags().GetInt(name)
	if v < 0 {
		return 0, fmt.Errorf("--%s must be zero or greater", name)
	}
	return v, nil
}

// GetNonNegativeInt64 returns the int64 value for a flag, returning an error if negative.
func GetNonNegativeInt64(cmd *cobra.Command, name string) (int64, error) {
	v, _ := cmd.Flags().GetInt64(name)
	if v < 0 {
		return 0, fmt.Errorf("--%s must be zero or greater", name)
	}
	return v, nil
}
