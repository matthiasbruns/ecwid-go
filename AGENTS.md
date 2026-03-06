# AGENTS.md — ecwid-go

Instructions for AI agents working on this codebase.

## Project Overview

Go API client and CLI for the [Ecwid REST API](https://docs.ecwid.com/api-reference).

- **API Client Library** (`ecwid/`): Stateless, stdlib-only, importable by other Go projects.
- **CLI** (`cmd/ecwid/`): Cobra-based CLI wrapping the API client. Loads credentials from env or config file.

## Architecture Principles

### Stateless API Client

The API client has **zero internal state**. No stored tokens, no singletons, no package-level vars.

- Every request starts from scratch.
- Credentials (`StoreID`, `Token`) are passed via a `Credentials` struct into each service method or through a `Client` that holds config but no mutable state.
- The `Client` is safe for concurrent use — it holds only immutable config + an `*http.Client`.

```go
// Good: credentials flow through explicitly
client := ecwid.NewClient(ecwid.Config{
    StoreID: "12345",
    Token:   "secret_abc",
})
products, err := client.Products.Search(ctx, ecwid.SearchProductsRequest{Keyword: "shirt"})

// Bad: package-level state
ecwid.SetToken("secret_abc") // ← NEVER do this
```

### Package Layout

```
ecwid-go/
├── ecwid/              # Public API client library (importable)
│   ├── client.go       # Client, Config, HTTP plumbing
│   ├── errors.go       # API error types
│   ├── products.go     # ProductService
│   ├── products_test.go
│   ├── orders.go       # OrderService
│   ├── orders_test.go
│   ├── categories.go   # CategoryService
│   ├── customers.go    # CustomerService
│   ├── carts.go        # CartService (abandoned carts)
│   ├── subscriptions.go # SubscriptionService
│   ├── discounts.go    # PromotionService + CouponService
│   ├── profile.go      # StoreProfileService
│   ├── reviews.go      # ReviewService
│   ├── staff.go        # StaffService
│   ├── domains.go      # DomainService
│   ├── dictionaries.go # DictionaryService (countries, currencies, etc.)
│   └── README.md       # Package-level agent instructions
├── cmd/
│   └── ecwid/          # CLI binary
│       ├── main.go
│       ├── root.go     # Root cobra command, config loading
│       ├── products.go # `ecwid products list`, `ecwid products get`, etc.
│       ├── orders.go
│       └── README.md   # CLI agent instructions
├── internal/
│   └── config/         # Config file loading (YAML/TOML)
│       └── config.go
├── e2e/                # End-to-end tests against real Ecwid store
│   ├── e2e_test.go
│   └── README.md
├── .pre-commit-config.yaml
├── .golangci.yml
├── go.mod
├── go.sum
├── Makefile
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

### Error Handling

```go
// All API errors are typed
type APIError struct {
    StatusCode int    `json:"-"`
    Code       string `json:"errorCode"`
    Message    string `json:"errorMessage"`
}

func (e *APIError) Error() string {
    return fmt.Sprintf("ecwid: %d %s: %s", e.StatusCode, e.Code, e.Message)
}

// Callers can type-assert
var apiErr *ecwid.APIError
if errors.As(err, &apiErr) && apiErr.StatusCode == 404 {
    // handle not found
}
```

### Logging

- Use `log/slog` with JSON handler by default.
- **NEVER** log credentials, tokens, or API keys. Not even at debug level.
- Log request method + path + status code + duration. Redact the `Authorization` header.

```go
slog.Info("ecwid request",
    "method", req.Method,
    "path", req.URL.Path,
    "status", resp.StatusCode,
    "duration", time.Since(start),
)
```

### Testing

- Every service method needs unit tests with `net/http/httptest`.
- Mock the Ecwid API responses using recorded JSON fixtures in `testdata/`.
- Test both success and error paths (400, 404, 429, 500).
- Test query parameter encoding for search/filter endpoints.
- E2E tests are separate in `e2e/` — they hit a real store and are gated behind `ECWID_E2E=1`.

```go
func TestProductService_Search(t *testing.T) {
    srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Assert request
        assert(t, r.Method, http.MethodGet)
        assert(t, r.URL.Path, "/api/v3/12345/products")
        assert(t, r.URL.Query().Get("keyword"), "shirt")
        assert(t, r.Header.Get("Authorization"), "Bearer test-token")

        // Return fixture
        w.Header().Set("Content-Type", "application/json")
        w.Write(loadFixture(t, "products_search.json"))
    }))
    defer srv.Close()

    client := ecwid.NewClient(ecwid.Config{
        StoreID: "12345",
        Token:   "test-token",
        BaseURL: srv.URL, // Override for testing
    })

    resp, err := client.Products.Search(context.Background(), ecwid.SearchProductsRequest{
        Keyword: "shirt",
    })
    // assertions...
}
```

### Dependencies

- **API client (`ecwid/`)**: Go stdlib ONLY. Zero external dependencies.
- **CLI (`cmd/ecwid/`)**: `github.com/spf13/cobra` + `github.com/spf13/viper` for config.
- **Tests**: stdlib `testing` package only. No testify, no gomock.

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
`discounts`, `profile`, `reviews`, `staff`, `domains`, `dictionaries`, `cli`, `e2e`

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
1. `golangci-lint run`
2. `go test ./...`

Configured via `.pre-commit-config.yaml`. Install with `pre-commit install`.

## Rate Limiting

Ecwid enforces **600 req/min per token**. The client should:
- Parse `Retry-After` header on 429 responses.
- Return a typed `RateLimitError` so callers can handle it.
- NOT auto-retry — let the caller decide.

## API Coverage

Base URL: `https://app.ecwid.com/api/v3/{storeId}`

### Domain → Service mapping

| Domain | Service | Endpoints |
|--------|---------|-----------|
| Store Profile | `ProfileService` | `/profile`, `/profile/staffScopes`, `/profile/order_statuses`, `/profile/extrafields`, logos, shipping/payment options |
| Orders | `OrderService` | `/orders`, `/orders/{id}`, `/orders/last`, `/orders/deleted`, extra fields, invoices, calculate |
| Abandoned Carts | `CartService` | `/carts`, `/carts/{id}`, `/carts/{id}/place` |
| Subscriptions | `SubscriptionService` | `/subscriptions`, `/subscriptions/{id}` |
| Products | `ProductService` | `/products`, variations, inventory, images, gallery, files, sort, filters |
| Product Reviews | `ReviewService` | `/reviews`, `/reviews/filters_data`, `/reviews/deleted`, mass update |
| Categories | `CategoryService` | `/categories`, `/categoriesByPath`, sort, images, assign/unassign products |
| Customers | `CustomerService` | `/customers`, contacts, extra fields, customer groups, deleted |
| Promotions | `PromotionService` | `/promotions`, `/promotions/{id}` |
| Discount Coupons | `CouponService` | `/discount_coupons`, `/discount_coupons/{id}`, `/coupons/deleted` |
| Domains | `DomainService` | `/domains`, search, purchase, verification email, reset password |
| Dictionaries | `DictionaryService` | `/countries`, `/currencies`, `/currencyByCountry`, `/states`, `/taxClasses` |
| Staff | `StaffService` | `/staff`, `/staff/{id}` |
| Reports | `ReportService` | `/reports/{type}`, `/latest-stats` |

## CLI Design

```
ecwid [command] [subcommand] [flags]

Global flags:
  --store-id     Ecwid store ID (env: ECWID_STORE_ID)
  --token        API token (env: ECWID_TOKEN)
  --config       Config file path (default: ~/.ecwid.yaml)
  --output       Output format: json|table (default: json)
  --log-level    Log level: debug|info|warn|error (default: info)

Examples:
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
```

## Security

- **NEVER** log tokens, API keys, or credentials at any log level.
- **NEVER** include credentials in error messages.
- Config file should warn about permissions (0600 recommended).
- Redact `Authorization` header in any debug output.
