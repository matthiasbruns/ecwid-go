# AGENTS.md — ecwid-go

Go API client and CLI for the [Ecwid REST API](https://docs.ecwid.com/api-reference).

## Modules

Three separate Go modules connected by `go.work`:

| Module | Path | Purpose | Dependencies |
|--------|------|---------|--------------|
| `config` | `config/` | Config loading (file + env + flags) | stdlib only |
| `ecwid` | `ecwid/` | Stateless API client library | stdlib + config |
| `cli` | `cli/` | Cobra CLI wrapping the client | config + ecwid + cobra |

Each module has its own `AGENTS.md` with package-specific instructions.

## ⚠️ Exported API — CRUCIAL RULE

**Every exported symbol is a compatibility contract.** Once published, changing or removing it is a breaking change. Keep the public API surface as small as possible.

### Package structure

```
ecwid/
  client.go                  # Client, NewClient, Options (public)
  errors.go                  # APIError, RateLimitError (type aliases, public)
  internal/
    api/                     # HTTP transport, retry, errors (private)
  dictionaries/              # Public domain package
    types.go                 # Exported types (Country, Currency, etc.)
    service.go               # Service struct + methods
  reports/                   # Public domain package
    types.go
    service.go
```

### Rules

- **Public packages** (`ecwid/dictionaries/`, `ecwid/reports/`, etc.) export only what users need: Service type, request options, response types.
- **Internal packages** (`ecwid/internal/api/`) hold all implementation details: HTTP client, retry transport, error parsing, request building. Users of the lib cannot import these.
- **New domain services** follow the same pattern: `ecwid/<domain>/service.go` + `ecwid/<domain>/types.go`, with the Service accepting an `api.Requester` interface.
- **Before exporting anything**, ask: "Does the consumer need this?" If not, keep it internal or unexported.
- **Type aliases** in root `ecwid/` re-export internal types that users need (e.g., `APIError`, `RateLimitError`) without exposing the internal package.

## Key Principles

- **Stateless client** — no package-level state, credentials passed via `config.Config` into `NewClient()`
- **Stdlib only** for config + client — CLI adds cobra as the only external dep
- **Every method takes `context.Context`** as first parameter
- **Typed errors** — `*ecwid.APIError` (all non-2xx) and `*ecwid.RateLimitError` (429 + Retry-After)
- **Never log credentials** — use `config.RedactedToken()` for diagnostics
- **Domain services are sub-packages** — `ecwid/dictionaries`, `ecwid/reports`, etc.
- **Shared types** (`UpdateResult`, `CreateResult`, `DeleteResult`) go in a shared internal or the domain package that introduces them

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
