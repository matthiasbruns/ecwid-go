package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// testStoreID is the fake store ID used by all integration tests.
const testStoreID = "12345"

// newMockServer creates a test HTTP server that routes requests matching the Ecwid API.
// The storeID prefix (e.g. "/12345") is stripped before matching routes.
func newMockServer(t *testing.T) *httptest.Server {
	t.Helper()

	mux := http.NewServeMux()
	prefix := "/" + testStoreID

	// Helper to register routes with the store ID prefix.
	handle := func(pattern string, h http.HandlerFunc) {
		mux.HandleFunc(prefix+pattern, h)
	}

	// ── Profile ──────────────────────────────────────────────────────────
	handle("/profile", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, map[string]any{
				"generalInfo": map[string]any{"storeId": 12345},
				"settings":    map[string]any{"storeName": "Test Store"},
			})
		case http.MethodPut:
			writeJSON(w, map[string]any{"updateCount": 1})
		}
	})

	// ── Dictionaries ─────────────────────────────────────────────────────
	handle("/countries", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, []map[string]any{
			{"code": "US", "name": "United States"},
			{"code": "DE", "name": "Germany"},
		})
	})
	handle("/currencies", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, []map[string]any{
			{"code": "USD", "name": "US Dollar"},
			{"code": "EUR", "name": "Euro"},
		})
	})

	// ── Categories ───────────────────────────────────────────────────────
	handle("/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, map[string]any{
				"total": 1, "count": 1, "offset": 0, "limit": 100,
				"items": []map[string]any{
					{"id": 1001, "name": "Test Category", "enabled": true},
				},
			})
		case http.MethodPost:
			writeJSON(w, map[string]any{"id": 1001})
		}
	})
	handle("/categories/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, map[string]any{"id": 1001, "name": "Test Category", "enabled": true})
		case http.MethodPut:
			writeJSON(w, map[string]any{"updateCount": 1})
		case http.MethodDelete:
			writeJSON(w, map[string]any{"deleteCount": 1})
		}
	})

	// ── Products ─────────────────────────────────────────────────────────
	handle("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, map[string]any{
				"total": 1, "count": 1, "offset": 0, "limit": 100,
				"items": []map[string]any{
					{"id": 2001, "name": "Test Product", "price": 19.99, "sku": "TEST-SKU"},
				},
			})
		case http.MethodPost:
			writeJSON(w, map[string]any{"id": 2001})
		}
	})
	handle("/products/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, map[string]any{"id": 2001, "name": "Test Product", "price": 19.99, "sku": "TEST-SKU"})
		case http.MethodPut:
			writeJSON(w, map[string]any{"updateCount": 1})
		case http.MethodDelete:
			writeJSON(w, map[string]any{"deleteCount": 1})
		}
	})

	// ── Orders ───────────────────────────────────────────────────────────
	handle("/orders", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, map[string]any{
				"total": 1, "count": 1, "offset": 0, "limit": 100,
				"items": []map[string]any{
					{"id": "ORD-001", "total": 42.00, "email": "test@example.com", "paymentStatus": "PAID"},
				},
			})
		case http.MethodPost:
			writeJSON(w, map[string]any{"id": 3001, "orderid": "ORD-001"})
		}
	})
	handle("/orders/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, map[string]any{"id": "ORD-001", "total": 42.00, "email": "test@example.com", "paymentStatus": "PAID"})
		case http.MethodPut:
			writeJSON(w, map[string]any{"updateCount": 1})
		case http.MethodDelete:
			writeJSON(w, map[string]any{"deleteCount": 1})
		}
	})

	// ── Customers ────────────────────────────────────────────────────────
	handle("/customers", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, map[string]any{
				"total": 1, "count": 1, "offset": 0, "limit": 100,
				"items": []map[string]any{
					{"id": 4001, "email": "customer@example.com", "name": "Test Customer"},
				},
			})
		case http.MethodPost:
			writeJSON(w, map[string]any{"id": 4001})
		}
	})
	handle("/customers/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, map[string]any{"id": 4001, "email": "customer@example.com", "name": "Test Customer"})
		case http.MethodPut:
			writeJSON(w, map[string]any{"updateCount": 1})
		case http.MethodDelete:
			writeJSON(w, map[string]any{"deleteCount": 1})
		}
	})

	// ── Promotions ───────────────────────────────────────────────────────
	handle("/promotions", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, map[string]any{
				"total": 1, "count": 1, "offset": 0, "limit": 100,
				"items": []map[string]any{
					{"id": 5001, "name": "Test Promo"},
				},
			})
		case http.MethodPost:
			writeJSON(w, map[string]any{"id": 5001})
		}
	})
	handle("/promotions/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			writeJSON(w, map[string]any{"updateCount": 1})
		case http.MethodDelete:
			writeJSON(w, map[string]any{"deleteCount": 1})
		}
	})

	// ── Coupons (discount_coupons) ───────────────────────────────────────
	handle("/discount_coupons", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, map[string]any{
				"total": 1, "count": 1, "offset": 0, "limit": 100,
				"items": []map[string]any{
					{"id": 6001, "code": "TEST10", "name": "Test Coupon"},
				},
			})
		case http.MethodPost:
			writeJSON(w, map[string]any{"id": 6001})
		}
	})
	handle("/discount_coupons/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, map[string]any{"id": 6001, "code": "TEST10", "name": "Test Coupon"})
		case http.MethodPut:
			writeJSON(w, map[string]any{"updateCount": 1})
		case http.MethodDelete:
			writeJSON(w, map[string]any{"deleteCount": 1})
		}
	})

	// ── Reviews ──────────────────────────────────────────────────────────
	handle("/reviews", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, map[string]any{
			"total": 1, "count": 1, "offset": 0, "limit": 100,
			"items": []map[string]any{
				{"id": 7001, "productId": 2001, "status": "APPROVED", "reviewText": "Great!"},
			},
		})
	})
	handle("/reviews/mass_update", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, map[string]any{"updateCount": 1})
	})
	handle("/reviews/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			writeJSON(w, map[string]any{"updateCount": 1})
		case http.MethodDelete:
			writeJSON(w, map[string]any{"deleteCount": 1})
		}
	})

	// ── Staff ────────────────────────────────────────────────────────────
	handle("/staff", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, map[string]any{
				"staffList": []map[string]any{
					{"id": "staff-001", "email": "staff@example.com", "name": "Test Staff"},
				},
			})
		case http.MethodPost:
			writeJSON(w, map[string]any{"success": true})
		}
	})
	handle("/staff/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, map[string]any{"id": "staff-001", "email": "staff@example.com", "name": "Test Staff"})
		case http.MethodPut:
			writeJSON(w, map[string]any{"updateCount": 1})
		case http.MethodDelete:
			writeJSON(w, map[string]any{"deleteCount": 1})
		}
	})

	// ── Domains ──────────────────────────────────────────────────────────
	handle("/domains", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, map[string]any{
			"instantSiteDomain": map[string]any{
				"ecwidSubdomain": "test.ecwid.com",
			},
		})
	})
	handle("/domains/purchase", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, map[string]any{"id": 1, "name": "mydomain.com", "status": "ACTIVE"})
	})

	// ── Reports ──────────────────────────────────────────────────────────
	handle("/reports/", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, map[string]any{
			"reportType": "allOrders",
			"reportData": []map[string]any{},
		})
	})
	handle("/latest-stats", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, map[string]any{
			"productsUpdated": "2026-01-01 00:00:00",
			"ordersUpdated":   "2026-01-01 00:00:00",
			"profileUpdated":  "2026-01-01 00:00:00",
		})
	})

	// ── Carts ────────────────────────────────────────────────────────────
	handle("/carts", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, map[string]any{
			"total": 1, "count": 1, "offset": 0, "limit": 100,
			"items": []map[string]any{
				{"cartId": "cart-001", "hidden": false},
			},
		})
	})
	handle("/carts/", func(w http.ResponseWriter, r *http.Request) {
		// Place endpoint: POST /carts/{id}/place
		if strings.HasSuffix(r.URL.Path, "/place") && r.Method == http.MethodPost {
			writeJSON(w, map[string]any{"id": 9001, "orderid": "ORD-CART"})
			return
		}
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, map[string]any{"cartId": "cart-001", "hidden": false})
		case http.MethodPut:
			writeJSON(w, map[string]any{"updateCount": 1})
		}
	})

	// ── Subscriptions ────────────────────────────────────────────────────
	handle("/subscriptions", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, map[string]any{
			"total": 1, "count": 1, "offset": 0, "limit": 100,
			"items": []map[string]any{
				{"subscriptionId": 8001, "customerId": 100, "status": "ACTIVE"},
			},
		})
	})
	handle("/subscriptions/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, map[string]any{"subscriptionId": 8001, "customerId": 100, "status": "ACTIVE"})
		case http.MethodPut:
			writeJSON(w, map[string]any{"updateCount": 1})
		}
	})

	// ── Error simulation endpoints ───────────────────────────────────────
	// These use a different store prefix so we can test error codes
	// by pointing ECWID_BASE_URL at srv.URL + "/error-NNN"
	for _, errCase := range []struct {
		prefix string
		code   int
		body   map[string]any
	}{
		{"/error-401/" + testStoreID, 401, map[string]any{"errorCode": "UNAUTHORIZED", "errorMessage": "Invalid token"}},
		{"/error-404/" + testStoreID, 404, map[string]any{"errorCode": "NOT_FOUND", "errorMessage": "Not found"}},
		{"/error-429/" + testStoreID, 429, map[string]any{"errorCode": "RATE_LIMIT", "errorMessage": "Rate limit exceeded"}},
	} {
		errCase := errCase
		mux.HandleFunc(errCase.prefix+"/", func(w http.ResponseWriter, _ *http.Request) {
			if errCase.code == 429 {
				w.Header().Set("Retry-After", "5")
			}
			w.WriteHeader(errCase.code)
			writeJSON(w, errCase.body)
		})
	}

	// Catch-all for debugging.
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		writeJSON(w, map[string]any{
			"errorCode":    "ROUTE_NOT_FOUND",
			"errorMessage": "mock server: no handler for " + r.Method + " " + r.URL.Path,
		})
	})

	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	return srv
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}
