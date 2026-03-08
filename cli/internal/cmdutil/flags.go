package cmdutil

import (
	"fmt"

	"github.com/spf13/cobra"
)

// GetNonNegativeInt returns the int value for a flag, returning an error if negative.
func GetNonNegativeInt(cmd *cobra.Command, name string) (int, error) {
	v, err := cmd.Flags().GetInt(name)
	if err != nil {
		return 0, err
	}
	if v < 0 {
		return 0, fmt.Errorf("--%s must be zero or greater", name)
	}
	return v, nil
}

// GetNonNegativeInt64 returns the int64 value for a flag, returning an error if negative.
func GetNonNegativeInt64(cmd *cobra.Command, name string) (int64, error) {
	v, err := cmd.Flags().GetInt64(name)
	if err != nil {
		return 0, err
	}
	if v < 0 {
		return 0, fmt.Errorf("--%s must be zero or greater", name)
	}
	return v, nil
}

// GetPositiveInt64IfChanged returns the int64 value for a flag only if explicitly set.
// Returns an error if the value is not positive when provided.
func GetPositiveInt64IfChanged(cmd *cobra.Command, name string) (value int64, changed bool, err error) {
	if !cmd.Flags().Changed(name) {
		return 0, false, nil
	}
	v, err := cmd.Flags().GetInt64(name)
	if err != nil {
		return 0, true, err
	}
	if v <= 0 {
		return 0, true, fmt.Errorf("--%s must be a positive integer", name)
	}
	return v, true, nil
}
