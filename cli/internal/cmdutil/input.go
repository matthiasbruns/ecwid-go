package cmdutil

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// ReadInput reads JSON input from --file flag or stdin.
func ReadInput(cmd *cobra.Command) ([]byte, error) {
	file, _ := cmd.Flags().GetString("file")
	if file != "" {
		data, err := os.ReadFile(file)
		if err != nil {
			return nil, fmt.Errorf("read file %q: %w", file, err)
		}
		return data, nil
	}
	data, err := io.ReadAll(cmd.InOrStdin())
	if err != nil {
		return nil, fmt.Errorf("read stdin: %w", err)
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("no input provided: specify --file or provide JSON on stdin")
	}
	return data, nil
}

// ReadJSONInput reads JSON from the --data flag, or stdin if the flag is empty.
func ReadJSONInput(cmd *cobra.Command) ([]byte, error) {
	if data, _ := cmd.Flags().GetString("data"); data != "" {
		return []byte(data), nil
	}
	raw, err := io.ReadAll(cmd.InOrStdin())
	if err != nil {
		return nil, fmt.Errorf("read input: %w", err)
	}
	if len(raw) == 0 {
		return nil, fmt.Errorf("no input: use --data flag or pipe JSON to stdin")
	}
	return raw, nil
}
