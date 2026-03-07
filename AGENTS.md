# AGENTS.md — ecwid-go

Go API client and CLI for the [Ecwid REST API](https://docs.ecwid.com/api-reference). See also [agents.md](https://agents.md/).

## Modules

Three separate Go modules connected by `go.work`:

| Module | Path | Purpose | Dependencies |
|--------|------|---------|--------------|
| `config` | `config/` | Config loading (file + env + flags) | stdlib only |
| `ecwid` | `ecwid/` | Stateless API client library | stdlib + config |
| `cli` | `cli/` | Cobra CLI wrapping the client | config + ecwid + cobra |

Each module has its own `AGENTS.md` with package-specific instructions.

## Key Principles

- **Stateless client** — no package-level state, credentials passed via `config.Config` into `NewClient()`
- **Stdlib only** for config + client — CLI adds cobra as the only external dep
- **Every method takes `context.Context`** as first parameter
- **Typed errors** — `*APIError` (all non-2xx) and `*RateLimitError` (429 + Retry-After)
- **Never log credentials** — use `config.RedactedToken()` for diagnostics

## Git Workflow

- **Conventional commits**: `feat(products):`, `fix(orders):`, `test(categories):`, `chore:`, `docs:`
- **Conventional branches**: `feat/products-search`, `fix/orders-rate-limit`
- **PRs only** — `main` is protected, squash merge preferred
- **Reviews required** — address all Copilot + CodeRabbit comments before merge

## Commands

```bash
task lint        # Lint all modules
task lint:fix    # Lint + auto-fix
task test        # Test all modules (-race)
task tidy        # go mod tidy all modules
task build       # Build CLI to ./bin/ecwid
task all         # lint + test + build
task e2e         # E2E tests (requires ECWID_E2E=1)

# Per-module:
task config:test
task ecwid:lint
task cli:tidy
```

## Security

- **NEVER** log tokens, API keys, or credentials at any log level
- **NEVER** include credentials in error messages
- Config files should have `0600` permissions
