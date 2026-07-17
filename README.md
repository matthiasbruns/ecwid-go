# ecwid-go

Go client library and CLI for the [Ecwid REST API](https://docs.ecwid.com/api-reference).

[![Go Reference](https://pkg.go.dev/badge/github.com/matthiasbruns/ecwid-go/ecwid.svg)](https://pkg.go.dev/github.com/matthiasbruns/ecwid-go/ecwid)
[![CI](https://github.com/matthiasbruns/ecwid-go/actions/workflows/ci.yml/badge.svg)](https://github.com/matthiasbruns/ecwid-go/actions/workflows/ci.yml)

## Features

- **Full Ecwid API coverage** — Products, Orders, Customers, Categories, Carts, Subscriptions, Promotions, Coupons, Reviews, Store Profile, Staff, Domains, Instant Site, Dictionaries, Reports
- **Webhooks** — Typed events, constant-time signature verification that fails closed, and an `http.Handler` with replay defenses
- **Stdlib only** — Zero external dependencies in config and client modules
- **Stateless** — No internal state; credentials passed explicitly per client
- **Multi-module** — Clean separation: `config/`, `ecwid/`, `cli/` with independent `go.mod`s
- **Configurable retry** — Optional auto-retry on 429 with `Retry-After` support
- **CLI included** — Cobra-based CLI for terminal access
- **Structured logging** — `slog` with JSON output, credentials never logged
- **Fully tested** — Unit tests per endpoint + E2E tests against real stores

## Project Structure

```
ecwid-go/
├── config/     # Config loading (file + env + flags) — stdlib only
├── ecwid/      # API client library — stdlib + config
├── cli/        # Cobra CLI — config + ecwid + cobra
├── e2e/        # E2E tests (future)
└── go.work     # Go workspace
```

Three separate Go modules. Users importing only the client library get zero transitive dependencies beyond `config/`.

## Installation

### Library

```bash
go get github.com/matthiasbruns/ecwid-go/ecwid
```

### CLI

```bash
go install github.com/matthiasbruns/ecwid-go/cli@latest
```

## Quick Start

### Library Usage (planned — endpoints not yet implemented)

> **Note:** The API service methods shown below are planned. The current bootstrap includes the client core, error types, and stub services. Endpoint implementations will follow in separate issues.

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/matthiasbruns/ecwid-go/config"
	"github.com/matthiasbruns/ecwid-go/ecwid"
)

func main() {
	cfg := config.Config{
		StoreID:    "12345",
		Token:      "secret_abc123",
		MaxRetries: 3, // Auto-retry on 429
	}

	client := ecwid.NewClient(cfg)

	// Search products (planned)
	resp, err := client.Products.Search(context.Background(), ecwid.SearchProductsRequest{
		Keyword: "shirt",
		Limit:   10,
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range resp.Items {
		fmt.Printf("%d: %s ($%.2f)\n", p.ID, p.Name, p.Price)
	}
}
```

### CLI Usage (planned — only `version` is implemented)

```bash
# Currently available
ecwid version

# Set credentials via environment
export ECWID_STORE_ID=12345
export ECWID_TOKEN=secret_abc123

# Planned commands (not yet implemented):
ecwid products list --keyword "shirt" --limit 10
ecwid products get 456789
ecwid orders list --status PAID
ecwid orders get 78901
ecwid profile get
ecwid categories list --parent 0
```

### Config File

Create `~/.ecwid.yaml`:

```yaml
store_id: "12345"
token: "secret_abc123"
output: json       # json | table
log_level: info    # debug | info | warn | error
max_retries: 3     # 0 = no retry
```

> ⚠️ Set file permissions to `0600` — the file contains your API token.

### Config Precedence

1. CLI flags (`--store-id`, `--token`)
2. Environment variables (`ECWID_STORE_ID`, `ECWID_TOKEN`)
3. Config file (`~/.ecwid.yaml`)
4. Defaults

## Error Handling

All API errors are typed for precise handling:

```go
import "errors"

resp, err := client.Products.Get(ctx, 999999)
if err != nil {
    var apiErr *ecwid.APIError
    if errors.As(err, &apiErr) {
        switch apiErr.StatusCode {
        case 404:
            fmt.Println("Product not found")
        case 429:
            // Only if MaxRetries=0 (no auto-retry)
            fmt.Println("Rate limited")
        default:
            fmt.Printf("API error %d: %s\n", apiErr.StatusCode, apiErr.Message)
        }
    }
}
```

## Rate Limiting

Ecwid allows **600 requests/minute per token**.

- `MaxRetries: 0` (default): 429 responses surface as `*ecwid.RateLimitError` with `RetryAfter`.
- `MaxRetries: N`: Client auto-retries up to N times, respecting `Retry-After` and context cancellation.

## Webhooks

`ecwid/webhooks` provides typed events, signature verification, and an `http.Handler`:

```go
h, err := webhooks.NewHandler(clientSecret, func(ctx context.Context, e webhooks.Event) {
	if e.EventType != webhooks.EventOrderCreated {
		return
	}
	// Order events carry the *internal* ID in EntityID, which the order
	// endpoints don't accept — the usable one is in the payload.
	d, err := e.OrderData()
	if err != nil {
		return
	}
	// Re-fetch rather than trusting the rest of e.Data — see the caveat below.
	order, err := client.Orders.Get(ctx, d.OrderID)
	// ...
}, &webhooks.Options{MaxAge: 5 * time.Minute})
if err != nil {
	return err
}
mux.Handle("POST /webhooks/ecwid", h)
```

The handler responds before running your callback (Ecwid allows 10s), fails closed on a
missing signature, and defaults to a 200 — Ecwid counts 203, 208 and every 3xx as a failed
delivery, so an endpoint that redirects HTTP to HTTPS silently never receives a webhook.

> **The webhook signature does not cover the request body.** It is an HMAC over only
> `"<eventCreated>.<eventId>"`, keyed with the app's `client_secret`. Everything else —
> `storeId`, `eventType`, `entityId`, `data` — is unauthenticated, and a captured webhook
> stays replayable forever. Dedupe on `EventID` (`Options.Deduper`), bound staleness with
> `Options.MaxAge`, and re-fetch the entity by `EntityID` instead of trusting `Data`.
> See the [package docs](https://pkg.go.dev/github.com/matthiasbruns/ecwid-go/ecwid/webhooks).

## API Coverage

| Domain | Service | Status |
|--------|---------|--------|
| Store Profile | `ProfileService` | 🔲 |
| Orders | `OrderService` | 🔲 |
| Abandoned Carts | `CartService` | 🔲 |
| Subscriptions | `SubscriptionService` | 🔲 |
| Products | `ProductService` | 🔲 |
| Product Reviews | `ReviewService` | 🔲 |
| Categories | `CategoryService` | 🔲 |
| Customers | `CustomerService` | 🔲 |
| Promotions | `PromotionService` | 🔲 |
| Discount Coupons | `CouponService` | 🔲 |
| Domains | `DomainService` | 🔲 |
| Instant Site | `InstantSiteService` | ✅ |
| Dictionaries | `DictionaryService` | 🔲 |
| Staff | `StaffService` | 🔲 |
| Reports | `ReportService` | 🔲 |
| Webhooks | `webhooks` package | ✅ |
| CLI | Cobra commands | 🔲 |
| E2E Tests | Real store tests | 🔲 |

## Development

### Prerequisites

- Go 1.26+
- [Task](https://taskfile.dev/) (`go install github.com/go-task/task/v3/cmd/task@latest`)
- [golangci-lint](https://golangci-lint.run/) v2+
- [pre-commit](https://pre-commit.com/)

### Setup

```bash
git clone https://github.com/matthiasbruns/ecwid-go.git
cd ecwid-go
pre-commit install
```

### Commands

```bash
task lint       # Run golangci-lint across all modules
task test       # Run unit tests with -race
task e2e        # Run E2E tests (requires ECWID_STORE_ID + ECWID_TOKEN)
task build      # Build CLI binary to ./bin/ecwid
task all        # lint + test + build
```

### E2E Tests

E2E tests run against a real Ecwid store, gated behind `ECWID_E2E=1`:

```bash
export ECWID_E2E=1
export ECWID_STORE_ID=12345
export ECWID_TOKEN=secret_abc123
task e2e
```

## Contributing

1. Fork & clone
2. Create a branch: `feat/your-feature`
3. Use [conventional commits](https://www.conventionalcommits.org/)
4. Ensure `task all` passes
5. Open a PR — all PRs require review

## License

MIT — see [LICENSE](LICENSE).
