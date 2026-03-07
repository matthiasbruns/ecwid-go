# cli/ — Ecwid CLI

## Overview

Cobra-based CLI wrapping the `ecwid` client library. Loads credentials from env, config file, or flags.

## Structure

```
cli/
├── main.go          # Entry point, slog setup, version injection
├── cmd/
│   ├── root.go      # Root command, global flags, config loading, client init
│   ├── version.go   # `ecwid version`
│   ├── products.go  # `ecwid products [list|get|create|update|delete]`
│   ├── orders.go    # `ecwid orders [list|get|create|update|delete]`
│   └── ...          # One file per domain
└── AGENTS.md        # This file
```

## Adding a New Command (pattern for future implementations)

Currently only `version` is implemented. When adding domain commands:

1. Create `cmd/<domain>.go`.
2. Define the parent command and subcommands.
3. Use `AppClient.<Service>.<Method>()` for API calls (client is initialized in `PersistentPreRunE`).
4. Add an output helper (respects `--output` flag) — not yet implemented.

```go
// Example pattern (planned, not yet in codebase):
var productsCmd = &cobra.Command{
    Use:   "products",
    Short: "Manage products",
}

var productsListCmd = &cobra.Command{
    Use:   "list",
    Short: "List products",
    RunE: func(cmd *cobra.Command, args []string) error {
        resp, err := AppClient.Products.Search(cmd.Context(), ecwid.SearchProductsRequest{})
        if err != nil {
            return err
        }
        return outputResult(cmd, resp) // outputResult TBD
    },
}

func init() {
    rootCmd.AddCommand(productsCmd)
    productsCmd.AddCommand(productsListCmd)
}
```

## Rules

- Never log tokens or credentials (not even at debug level).
- All commands use `RunE` (return errors, don't os.Exit).
- JSON output is the default; table output is optional.
- Version is injected at build time via `-ldflags "-X main.version=..."`.
