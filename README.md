# ecwid-go

Go client library and CLI for the [Ecwid REST API](https://docs.ecwid.com/api-reference).

[![Go Reference](https://pkg.go.dev/badge/github.com/matthiasbruns/ecwid-go/ecwid.svg)](https://pkg.go.dev/github.com/matthiasbruns/ecwid-go/ecwid)
[![CI](https://github.com/matthiasbruns/ecwid-go/actions/workflows/ci.yml/badge.svg)](https://github.com/matthiasbruns/ecwid-go/actions/workflows/ci.yml)

## Features

- **Full Ecwid API coverage** — Products, Orders, Customers, Categories, Carts, Subscriptions, Promotions, Coupons, Reviews, Store Profile, Staff, Domains, Dictionaries, Reports
- **Stdlib only** — Zero external dependencies in the client library
- **Stateless** — No internal state; credentials passed explicitly per client
- **CLI included** — Cobra-based CLI for quick terminal access
- **Structured logging** — `slog` with JSON output, credentials never logged
- **Fully tested** — Unit tests per endpoint + E2E tests against real stores

## Installation

### Library

```bash
go get github.com/matthiasbruns/ecwid-go/ecwid
```

### CLI

```bash
go install github.com/matthiasbruns/ecwid-go/cmd/ecwid@latest
```

## Quick Start

### Library Usage

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/matthiasbruns/ecwid-go/ecwid"
)

func main() {
	client := ecwid.NewClient(ecwid.Config{
		StoreID: "12345",
		Token:   "secret_abc123",
	})

	// Search products
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

	// Get a single order
	order, err := client.Orders.Get(context.Background(), 78901)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Order #%s: %s\n", order.VendorNumber, order.PaymentStatus)
}
```

### CLI Usage

```bash
# Set credentials via environment
export ECWID_STORE_ID=12345
export ECWID_TOKEN=secret_abc123

# Or via config file (~/.ecwid.yaml)
ecwid config init

# Products
ecwid products list --keyword "shirt" --limit 10
ecwid products get 456789

# Orders
ecwid orders list --status PAID
ecwid orders get 78901

# Store profile
ecwid profile get

# Categories
ecwid categories list --parent 0
```

### Config File

Create `~/.ecwid.yaml`:

```yaml
store_id: "12345"
token: "secret_abc123"
output: json       # json | table
log_level: info    # debug | info | warn | error
```

> ⚠️ Set file permissions to `0600` — the file contains your API token.

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
            fmt.Printf("Rate limited, retry after %s\n", apiErr.RetryAfter)
        default:
            fmt.Printf("API error %d: %s\n", apiErr.StatusCode, apiErr.Message)
        }
    }
}
```

## Rate Limiting

Ecwid allows **600 requests/minute per token**. The client surfaces `429` responses as `*ecwid.RateLimitError` with the `Retry-After` value. It does **not** auto-retry — your code controls the backoff strategy.

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
| Dictionaries | `DictionaryService` | 🔲 |
| Staff | `StaffService` | 🔲 |
| Reports | `ReportService` | 🔲 |
| CLI | Cobra commands | 🔲 |
| E2E Tests | Real store tests | 🔲 |

## Development

### Prerequisites

- Go 1.26+
- [pre-commit](https://pre-commit.com/)
- [golangci-lint](https://golangci-lint.run/)

### Setup

```bash
git clone https://github.com/matthiasbruns/ecwid-go.git
cd ecwid-go
go mod download
pre-commit install
```

### Commands

```bash
make lint       # Run golangci-lint
make test       # Run unit tests
make e2e        # Run E2E tests (requires ECWID_STORE_ID + ECWID_TOKEN)
make build      # Build CLI binary
make all        # lint + test + build
```

### E2E Tests

E2E tests run against a real Ecwid store. They're gated behind `ECWID_E2E=1`:

```bash
export ECWID_E2E=1
export ECWID_STORE_ID=12345
export ECWID_TOKEN=secret_abc123
make e2e
```

## Contributing

1. Fork & clone
2. Create a branch: `feat/your-feature`
3. Use [conventional commits](https://www.conventionalcommits.org/)
4. Ensure `make all` passes
5. Open a PR — all PRs require review

## License

MIT — see [LICENSE](LICENSE).
