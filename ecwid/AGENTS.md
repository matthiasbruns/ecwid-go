# ecwid/ — API Client Package

## Overview

Stateless Go client for the Ecwid REST API. Stdlib only — zero external dependencies.

## Key Types

- `Client` — Main entry point. Created via `NewClient(cfg, ...opts)`. Holds immutable config + `*http.Client`.
- `*Service` — One per API domain (e.g., `ProductService`, `OrderService`). Accessed via `client.Products`, `client.Orders`, etc.
- `APIError` — Typed error for all non-2xx responses. Includes `StatusCode`, `Code`, `Message`.
- `RateLimitError` — Wraps `APIError` for 429 responses. Includes `RetryAfter` duration.

## Adding a New Endpoint

1. Add request/response types in the domain file (e.g., `products.go`).
2. Implement the method on the service struct.
3. Use `client.get()`, `client.post()`, `client.put()`, `client.delete()` for HTTP calls.
4. Add unit tests in `*_test.go` using `httptest.NewServer`.
5. Add JSON fixtures in `testdata/` for mock responses.

```go
// Example: Add "Get Product" to ProductService
func (s *ProductService) Get(ctx context.Context, productID int) (*Product, error) {
    path := fmt.Sprintf("/products/%d", productID)
    var product Product
    if err := s.client.get(ctx, path, nil, &product); err != nil {
        return nil, err
    }
    return &product, nil
}
```

## Testing Pattern

```go
func TestProductService_Get(t *testing.T) {
    c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
        // Assert request
        // Return fixture
        w.Write(loadFixture(t, "product_get.json"))
    })

    product, err := c.Products.Get(context.Background(), 123)
    // assertions...
}
```

## Rules

- Every method takes `context.Context` as first parameter.
- Return `(*Type, error)` — never return both nil.
- Use `fmt.Errorf("...: %w", err)` for error wrapping.
- No package-level state. No `init()`.
- JSON struct tags use `snake_case` matching the API.
- Optional query params use pointer types or zero-value omission.
