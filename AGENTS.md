# AGENTS.md — ecwid-go

Instructions for AI agents working on this codebase. See also https://agents.md/.

## Project Overview

Go API client and CLI for the [Ecwid REST API](https://docs.ecwid.com/api-reference).

**Three separate Go modules** in one repo, connected by `go.work`:

| Module | Path | Purpose | Dependencies |
|--------|------|---------|--------------|
| `config` | `config/` | Config loading (file + env + flags) | stdlib only |
| `ecwid` | `ecwid/` | API client library | stdlib + config |
| `cli` | `cli/` | Cobra CLI wrapping the client | config + ecwid + cobra |

Each module has its own `go.mod`. Local inter-module deps use `replace` directives.

## Architecture Principles

### Stateless API Client

The API client has **zero internal state**. No stored tokens, no singletons, no package-level vars.

- Every request starts from scratch.
- Credentials (`StoreID`, `Token`) live in `config.Config`, passed into `NewClient()`.
- The `Client` is safe for concurrent use — it holds only immutable config + an `*http.Client`.

```go
// Good: credentials flow through explicitly
client := ecwid.NewClient(config.Config{
    StoreID: "12345",
    Token:   "secret_abc",
})
products, err := client.Products.Search(ctx, ecwid.SearchProductsRequest{Keyword: "shirt"})

// Bad: package-level state
ecwid.SetToken("secret_abc") // ← NEVER do this
```

### Module Separation

- `config/` is its own module so CLI dependencies (cobra) don't pollute it.
- `ecwid/` imports `config/` but nothing else external — stdlib only.
- `cli/` imports both `config/` and `ecwid/`, plus cobra.
- Users who only need the API client import `ecwid/` and get zero transitive deps beyond `config/`.

### Package Layout

```
ecwid-go/
├── go.work             # Workspace: ./config, ./ecwid, ./cli
├── config/             # Config module (stdlib only)
│   ├── go.mod
│   ├── config.go
│   └── config_test.go
├── ecwid/              # API client module (stdlib + config)
│   ├── go.mod
│   ├── client.go       # Client, HTTP plumbing (get/post/put/delete)
│   ├── errors.go       # APIError, RateLimitError
│   ├── retry.go        # Configurable retry transport for 429s
│   ├── services.go     # Stub service structs
│   ├── doc.go          # Package documentation
│   ├── client_test.go  # Core HTTP tests
│   ├── testdata/       # JSON fixtures
│   ├── AGENTS.md       # Package-level agent instructions
│   ├── products.go     # ProductService (future)
│   ├── orders.go       # OrderService (future)
│   └── ...             # One file per domain
├── cli/                # CLI module (config + ecwid + cobra)
│   ├── go.mod
│   ├── main.go         # Entry point, slog JSON handler
│   ├── cmd/
│   │   ├── root.go     # Root command, global flags, config loading
│   │   ├── version.go  # `ecwid version`
│   │   └── ...         # One file per domain command
│   └── AGENTS.md       # CLI agent instructions
├── e2e/                # E2E tests (future, gated behind ECWID_E2E=1)
├── Taskfile.yml        # Task runner (lint, test, e2e, build, all)
├── .pre-commit-config.yaml
├── .golangci.yml
├── .github/workflows/ci.yml
├── LICENSE
├── README.md
└── AGENTS.md           # This file
```

## Coding Standards

### Go Idioms

Follow [Effective Go](https://go.dev/doc/effective_go) strictly:

- **Naming**: `MixedCaps`, not underscores. Exported names are capitalized.
- **Errors**: Return `error` as last return value. Wrap with `fmt.Errorf("...: %w", err)`.
- **Interfaces**: Accept interfaces, return structs. Keep interfaces small.
- **Context**: Every API method takes `context.Context` as first parameter.
- **Zero values**: Design structs so zero values are useful.
- **Deferred closes**: Use `defer func() { _ = r.Close() }()` pattern (errcheck-safe).

### Error Handling

```go
// All API errors are typed
var apiErr *ecwid.APIError
if errors.As(err, &apiErr) && apiErr.StatusCode == 404 {
    // handle not found
}

// Rate limit errors include RetryAfter
var rlErr *ecwid.RateLimitError
if errors.As(err, &rlErr) {
    time.Sleep(rlErr.RetryAfter)
}
```

### Retry Transport

Configurable via `config.Config.MaxRetries`:
- `0` (default): no retries, caller handles 429s manually.
- `> 0`: wraps `http.Client` transport, respects `Retry-After` header, respects context cancellation.

### Logging

- Use `log/slog` with JSON handler by default.
- **NEVER** log credentials, tokens, or API keys. Not even at debug level.
- Log request method + path + status code + duration. Redact the `Authorization` header.

### Testing

- Every service method needs unit tests with `net/http/httptest`.
- Mock the Ecwid API responses using recorded JSON fixtures in `testdata/`.
- Test both success and error paths (400, 404, 429, 500).
- Test query parameter encoding for search/filter endpoints.
- Use stdlib `testing` only — no testify, no gomock.
- E2E tests are separate in `e2e/` — gated behind `ECWID_E2E=1`.

### Dependencies

- **config module**: stdlib ONLY. Zero external dependencies.
- **ecwid module**: stdlib + config. Zero other external dependencies.
- **cli module**: config + ecwid + `github.com/spf13/cobra`.
- **Tests**: stdlib `testing` package only.

## Git Workflow

### Commits

Use [Conventional Commits](https://www.conventionalcommits.org/):

```
feat(products): add search products endpoint
fix(orders): handle 429 rate limit response
test(categories): add unit tests for create category
docs: update README with CLI usage
chore: add golangci-lint config
```

Scopes: `products`, `orders`, `categories`, `customers`, `carts`, `subscriptions`,
`discounts`, `profile`, `reviews`, `staff`, `domains`, `dictionaries`, `cli`, `config`, `e2e`

### Branches

Use conventional branch names:

```
feat/products-search
fix/orders-rate-limit
test/categories-unit-tests
docs/readme-cli-usage
chore/pre-commit-setup
```

### Pull Requests

- `main` is protected — all changes go through PRs.
- Every PR needs reviews (Copilot + human).
- Address ALL review comments before merging.
- Squash merge preferred.

### Pre-commit Hooks

Mandatory before push:
1. `golangci-lint run` (all modules)
2. `go test ./config/... ./ecwid/...` with `-race`

Configured via `.pre-commit-config.yaml`. Install with `pre-commit install`.

## Rate Limiting

Ecwid enforces **600 req/min per token**. The client:
- Parses `Retry-After` header on 429 responses.
- Returns `*RateLimitError` so callers can handle it.
- Optionally auto-retries if `MaxRetries > 0` in config.

## API Coverage (planned — stub services only in bootstrap)

Base URL: `https://app.ecwid.com/api/v3/{storeId}`

All services below are **stubbed** (empty structs wired to `Client`). Actual endpoint methods will be implemented in follow-up issues.

| Domain | Service | Planned Endpoints |
|--------|---------|-----------|
| Store Profile | `ProfileService` | `/profile`, staffScopes, order_statuses, extrafields, logos, shipping/payment options |
| Orders | `OrderService` | `/orders`, `/orders/{id}`, last, deleted, extra fields, invoices, calculate |
| Abandoned Carts | `CartService` | `/carts`, `/carts/{id}`, place |
| Subscriptions | `SubscriptionService` | `/subscriptions`, `/subscriptions/{id}` |
| Products | `ProductService` | `/products`, variations, inventory, images, gallery, files, sort, filters, brands, classes, swatches |
| Product Reviews | `ReviewService` | `/reviews`, filters_data, deleted, mass_update |
| Categories | `CategoryService` | `/categories`, byPath, sort, images, assign/unassign products |
| Customers | `CustomerService` | `/customers`, contacts, extra fields, customer groups, deleted |
| Promotions | `PromotionService` | `/promotions`, `/promotions/{id}` |
| Discount Coupons | `CouponService` | `/discount_coupons`, `/discount_coupons/{id}`, deleted |
| Domains | `DomainService` | `/domains`, search, purchase, verification, reset password |
| Dictionaries | `DictionaryService` | `/countries`, `/currencies`, `/currencyByCountry`, `/states`, `/taxClasses` |
| Staff | `StaffService` | `/staff`, `/staff/{id}` |
| Reports | `ReportService` | `/reports/{type}`, `/latest-stats` |

## CLI Design (planned — only `version` implemented)

Currently implemented: `ecwid version`

Planned structure for future domain commands:

```
ecwid [command] [subcommand] [flags]

Global flags (implemented):
  --store-id     Ecwid store ID (env: ECWID_STORE_ID)
  --token        API token (env: ECWID_TOKEN)
  --config       Config file path (default: ~/.ecwid.yaml)
  --output       Output format: json|table (default: json)
  --log-level    Log level: debug|info|warn|error (default: info)

Planned commands:
  ecwid products list --keyword "shirt" --limit 10
  ecwid orders get 12345
  ecwid profile get
  ecwid categories list --parent 0
```

Config file (`~/.ecwid.yaml`):
```yaml
store_id: "12345"
token: "secret_abc"
output: json
log_level: info
max_retries: 3
```

## Security

- **NEVER** log tokens, API keys, or credentials at any log level.
- **NEVER** include credentials in error messages.
- Use `config.RedactedToken()` for any diagnostic output.
- Config file should have 0600 permissions.
