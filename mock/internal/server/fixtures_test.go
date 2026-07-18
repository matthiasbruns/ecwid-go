package server

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/matthiasbruns/ecwid-go/ecwid/customers"
	"github.com/matthiasbruns/ecwid-go/ecwid/profile"
	"github.com/matthiasbruns/ecwid-go/mock/internal/config"
)

// newFixtureServer builds a server with a known access token and store, seeded
// with the default fixtures, for the simulated REST tests. scopes narrows the
// granted scopes; nil grants all (the default).
func newFixtureServer(scopes []string) *Server {
	return New(config.Config{
		Port:        0,
		AccessToken: testToken,
		StoreID:     testStoreID,
		Scopes:      scopes,
	}, discardLogger())
}

// fixtureReq builds an authorized request against a simulated REST route.
func fixtureReq(method, target string, body io.Reader) *http.Request {
	req := httptest.NewRequest(method, target, body)
	req.Header.Set("Authorization", "Bearer "+testToken)
	return req
}

// restPath builds an /api/v3/{storeId}/... path for the test store.
func restPath(suffix string) string {
	return "/api/v3/" + testStoreID + "/" + suffix
}

// ── Profile ──────────────────────────────────────────────────────────────

// TestFixtures_Profile covers GET /profile: the seeded default is returned with
// the store ID wired through, and the read_store_profile scope is enforced.
func TestFixtures_Profile(t *testing.T) {
	t.Run("returns seeded profile with store id", func(t *testing.T) {
		srv := newFixtureServer(nil)
		rec := do(srv, fixtureReq(http.MethodGet, restPath("profile"), nil))
		if rec.Code != http.StatusOK {
			t.Fatalf("GET profile status = %d, want 200", rec.Code)
		}
		var p profile.Profile
		if err := json.NewDecoder(rec.Body).Decode(&p); err != nil {
			t.Fatalf("decode profile: %v", err)
		}
		if p.GeneralInfo == nil || p.GeneralInfo.StoreID != 1003 {
			t.Errorf("profile generalInfo.storeId = %+v, want 1003", p.GeneralInfo)
		}
		if p.Settings == nil || p.Settings.StoreName == "" {
			t.Errorf("profile settings.storeName is empty, want a seeded name")
		}
	})

	t.Run("scope denied is 403", func(t *testing.T) {
		srv := newFixtureServer([]string{scopeReadCustomers}) // profile scope withheld
		rec := do(srv, fixtureReq(http.MethodGet, restPath("profile"), nil))
		if rec.Code != http.StatusForbidden {
			t.Fatalf("GET profile without scope status = %d, want 403", rec.Code)
		}
		assertErrorBody(t, rec.Body.Bytes(), scopeReadStoreProfile)
	})

	t.Run("wrong bearer is 401", func(t *testing.T) {
		srv := newFixtureServer(nil)
		req := httptest.NewRequest(http.MethodGet, restPath("profile"), http.NoBody)
		req.Header.Set("Authorization", "Bearer not-the-token")
		rec := do(srv, req)
		if rec.Code != http.StatusUnauthorized {
			t.Errorf("GET profile wrong bearer status = %d, want 401", rec.Code)
		}
	})
}

// ── Customers: list / paging / email filter ──────────────────────────────

// TestFixtures_CustomersSearch covers GET /customers: the paged envelope, offset
// and limit windowing, the email filter, and scope enforcement.
func TestFixtures_CustomersSearch(t *testing.T) {
	tests := []struct {
		name       string
		query      string
		wantStatus int
		wantTotal  int
		wantCount  int
		wantOffset int
		wantLimit  int
		wantEmails []string
	}{
		{
			name:       "unpaged returns all in default window",
			query:      "",
			wantStatus: http.StatusOK,
			wantTotal:  3,
			wantCount:  3,
			wantOffset: 0,
			wantLimit:  defaultCustomersLimit,
			wantEmails: []string{"ada@example.com", "grace@example.com", "alan@example.com"},
		},
		{
			name:       "first page of two",
			query:      "?offset=0&limit=2",
			wantStatus: http.StatusOK,
			wantTotal:  3,
			wantCount:  2,
			wantOffset: 0,
			wantLimit:  2,
			wantEmails: []string{"ada@example.com", "grace@example.com"},
		},
		{
			name:       "second page carries the remainder",
			query:      "?offset=2&limit=2",
			wantStatus: http.StatusOK,
			wantTotal:  3,
			wantCount:  1,
			wantOffset: 2,
			wantLimit:  2,
			wantEmails: []string{"alan@example.com"},
		},
		{
			name:       "offset past the end is an empty page",
			query:      "?offset=99&limit=2",
			wantStatus: http.StatusOK,
			wantTotal:  3,
			wantCount:  0,
			wantOffset: 99,
			wantLimit:  2,
			wantEmails: []string{},
		},
		{
			name:       "email filter finds the one match",
			query:      "?email=grace@example.com",
			wantStatus: http.StatusOK,
			wantTotal:  1,
			wantCount:  1,
			wantOffset: 0,
			wantLimit:  defaultCustomersLimit,
			wantEmails: []string{"grace@example.com"},
		},
		{
			name:       "email filter is case-insensitive",
			query:      "?email=ADA@example.com",
			wantStatus: http.StatusOK,
			wantTotal:  1,
			wantCount:  1,
			wantOffset: 0,
			wantLimit:  defaultCustomersLimit,
			wantEmails: []string{"ada@example.com"},
		},
		{
			name:       "email filter with no match is empty",
			query:      "?email=nobody@example.com",
			wantStatus: http.StatusOK,
			wantTotal:  0,
			wantCount:  0,
			wantOffset: 0,
			wantLimit:  defaultCustomersLimit,
			wantEmails: []string{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			srv := newFixtureServer(nil)
			rec := do(srv, fixtureReq(http.MethodGet, restPath("customers")+tc.query, nil))
			if rec.Code != tc.wantStatus {
				t.Fatalf("status = %d, want %d", rec.Code, tc.wantStatus)
			}
			var res customers.SearchResult
			if err := json.NewDecoder(rec.Body).Decode(&res); err != nil {
				t.Fatalf("decode search result: %v", err)
			}
			if res.Total != tc.wantTotal {
				t.Errorf("total = %d, want %d", res.Total, tc.wantTotal)
			}
			if res.Count != tc.wantCount {
				t.Errorf("count = %d, want %d", res.Count, tc.wantCount)
			}
			if res.Offset != tc.wantOffset {
				t.Errorf("offset = %d, want %d", res.Offset, tc.wantOffset)
			}
			if res.Limit != tc.wantLimit {
				t.Errorf("limit = %d, want %d", res.Limit, tc.wantLimit)
			}
			if res.Count != len(res.Items) {
				t.Errorf("count = %d but items has %d entries", res.Count, len(res.Items))
			}
			gotEmails := make([]string, len(res.Items))
			for i, c := range res.Items {
				gotEmails[i] = c.Email
			}
			if strings.Join(gotEmails, ",") != strings.Join(tc.wantEmails, ",") {
				t.Errorf("emails = %v, want %v", gotEmails, tc.wantEmails)
			}
		})
	}

	t.Run("scope denied is 403", func(t *testing.T) {
		srv := newFixtureServer([]string{scopeReadStoreProfile}) // customers scope withheld
		rec := do(srv, fixtureReq(http.MethodGet, restPath("customers"), nil))
		if rec.Code != http.StatusForbidden {
			t.Fatalf("status = %d, want 403", rec.Code)
		}
		assertErrorBody(t, rec.Body.Bytes(), scopeReadCustomers)
	})
}

// TestFixtures_ItemsEncodesAsArray asserts an empty page serializes as [] rather
// than null, matching the real API (a nil items slice would break typed clients
// that range over it).
func TestFixtures_ItemsEncodesAsArray(t *testing.T) {
	srv := newFixtureServer(nil)
	rec := do(srv, fixtureReq(http.MethodGet, restPath("customers")+"?email=nobody@example.com", nil))
	if !strings.Contains(rec.Body.String(), `"items":[]`) {
		t.Errorf("empty page body = %s, want items:[]", rec.Body.String())
	}
}

// ── Customers: get by id ─────────────────────────────────────────────────

// TestFixtures_CustomerGet covers GET /customers/{id}: a seeded customer, a
// missing id, a non-numeric id, and scope enforcement.
func TestFixtures_CustomerGet(t *testing.T) {
	tests := []struct {
		name       string
		id         string
		scopes     []string
		wantStatus int
		wantEmail  string
	}{
		{name: "existing customer", id: "1001", wantStatus: http.StatusOK, wantEmail: "ada@example.com"},
		{name: "missing customer is 404", id: "424242", wantStatus: http.StatusNotFound},
		{name: "non-numeric id is 404", id: "abc", wantStatus: http.StatusNotFound},
		{name: "scope denied is 403", id: "1001", scopes: []string{scopeReadStoreProfile}, wantStatus: http.StatusForbidden},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			srv := newFixtureServer(tc.scopes)
			rec := do(srv, fixtureReq(http.MethodGet, restPath("customers/"+tc.id), nil))
			if rec.Code != tc.wantStatus {
				t.Fatalf("status = %d, want %d", rec.Code, tc.wantStatus)
			}
			if tc.wantStatus != http.StatusOK {
				return
			}
			var c customers.Customer
			if err := json.NewDecoder(rec.Body).Decode(&c); err != nil {
				t.Fatalf("decode customer: %v", err)
			}
			if c.Email != tc.wantEmail {
				t.Errorf("email = %q, want %q", c.Email, tc.wantEmail)
			}
		})
	}
}

// ── Customers: update marketing consent (read-after-write) ───────────────

// TestFixtures_CustomerUpdate_ReadAfterWrite proves the PUT is reflected in
// subsequent GETs: flipping acceptMarketing true persists, other fields are
// preserved, and the update requires the update_customers scope.
func TestFixtures_CustomerUpdate_ReadAfterWrite(t *testing.T) {
	srv := newFixtureServer(nil)

	// Grace starts with acceptMarketing=false.
	before := getFixtureCustomer(t, srv, 1002)
	if before.AcceptMarketing == nil || *before.AcceptMarketing {
		t.Fatalf("seed precondition: customer 1002 acceptMarketing = %v, want false", before.AcceptMarketing)
	}

	// Flip it to true via PUT (only acceptMarketing in the body, as SetCustomerMarketing sends).
	rec := do(srv, fixtureReq(http.MethodPut, restPath("customers/1002"), strings.NewReader(`{"acceptMarketing":true}`)))
	if rec.Code != http.StatusOK {
		t.Fatalf("PUT status = %d, want 200", rec.Code)
	}
	var upd customers.UpdateResult
	if err := json.NewDecoder(rec.Body).Decode(&upd); err != nil {
		t.Fatalf("decode update result: %v", err)
	}
	if upd.UpdateCount != 1 {
		t.Errorf("updateCount = %d, want 1", upd.UpdateCount)
	}

	// The next GET must reflect the write, with untouched fields preserved.
	after := getFixtureCustomer(t, srv, 1002)
	if after.AcceptMarketing == nil || !*after.AcceptMarketing {
		t.Errorf("after PUT acceptMarketing = %v, want true (read-after-write)", after.AcceptMarketing)
	}
	if after.Email != before.Email || after.Name != before.Name {
		t.Errorf("PUT clobbered untouched fields: name/email = %q/%q, want %q/%q",
			after.Name, after.Email, before.Name, before.Email)
	}
	if after.ID != 1002 {
		t.Errorf("PUT changed the id to %d, want it fixed at 1002", after.ID)
	}
}

// TestFixtures_CustomerUpdate_Errors covers the unhappy PUT paths: missing
// customer, malformed body, and scope enforcement.
func TestFixtures_CustomerUpdate_Errors(t *testing.T) {
	tests := []struct {
		name       string
		id         string
		scopes     []string
		body       string
		wantStatus int
	}{
		{name: "missing customer is 404", id: "999999", body: `{"acceptMarketing":true}`, wantStatus: http.StatusNotFound},
		{name: "non-numeric id is 404", id: "xyz", body: `{"acceptMarketing":true}`, wantStatus: http.StatusNotFound},
		{name: "malformed body is 400", id: "1001", body: `not json`, wantStatus: http.StatusBadRequest},
		{name: "scope denied is 403", id: "1001", scopes: []string{scopeReadCustomers}, body: `{"acceptMarketing":true}`, wantStatus: http.StatusForbidden},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			srv := newFixtureServer(tc.scopes)
			rec := do(srv, fixtureReq(http.MethodPut, restPath("customers/"+tc.id), strings.NewReader(tc.body)))
			if rec.Code != tc.wantStatus {
				t.Fatalf("status = %d, want %d", rec.Code, tc.wantStatus)
			}
			if tc.wantStatus == http.StatusForbidden {
				assertErrorBody(t, rec.Body.Bytes(), scopeUpdateCustomers)
			}
		})
	}
}

// ── Fixtures control API ─────────────────────────────────────────────────

// TestFixtures_ControlSeed covers the /_mock/fixtures control API: seeding a new
// customer with a known email is then findable via the simulated REST endpoints,
// and a replacement profile is served back.
func TestFixtures_ControlSeed(t *testing.T) {
	t.Run("seed a customer then find it by email", func(t *testing.T) {
		srv := newFixtureServer(nil)

		body := `{"name":"Katherine Johnson","email":"katherine@example.com","acceptMarketing":false}`
		rec := do(srv, httptest.NewRequest(http.MethodPost, "/_mock/fixtures/customers", strings.NewReader(body)))
		if rec.Code != http.StatusOK {
			t.Fatalf("seed customer status = %d, want 200", rec.Code)
		}
		var seeded struct {
			ID int64 `json:"id"`
		}
		if err := json.NewDecoder(rec.Body).Decode(&seeded); err != nil {
			t.Fatalf("decode seed response: %v", err)
		}
		if seeded.ID == 0 {
			t.Fatalf("seed response id = 0, want an assigned id")
		}

		// The simulated REST email filter now finds it.
		find := do(srv, fixtureReq(http.MethodGet, restPath("customers")+"?email=katherine@example.com", nil))
		var res customers.SearchResult
		if err := json.NewDecoder(find.Body).Decode(&res); err != nil {
			t.Fatalf("decode search: %v", err)
		}
		if res.Total != 1 || len(res.Items) != 1 || res.Items[0].ID != seeded.ID {
			t.Errorf("find by email = %+v, want the seeded customer id %d", res, seeded.ID)
		}
	})

	t.Run("seed an array of customers", func(t *testing.T) {
		srv := newFixtureServer(nil)
		body := `[{"email":"a@x.com"},{"email":"b@x.com"}]`
		rec := do(srv, httptest.NewRequest(http.MethodPost, "/_mock/fixtures/customers", strings.NewReader(body)))
		if rec.Code != http.StatusOK {
			t.Fatalf("seed array status = %d, want 200", rec.Code)
		}
		var out struct {
			IDs []int64 `json:"ids"`
		}
		if err := json.NewDecoder(rec.Body).Decode(&out); err != nil {
			t.Fatalf("decode seed response: %v", err)
		}
		if len(out.IDs) != 2 {
			t.Errorf("seeded ids = %v, want 2 ids", out.IDs)
		}
	})

	t.Run("replace the profile", func(t *testing.T) {
		srv := newFixtureServer(nil)
		body := `{"settings":{"storeName":"Reseeded Store"}}`
		rec := do(srv, httptest.NewRequest(http.MethodPut, "/_mock/fixtures/profile", strings.NewReader(body)))
		if rec.Code != http.StatusOK {
			t.Fatalf("seed profile status = %d, want 200", rec.Code)
		}
		get := do(srv, fixtureReq(http.MethodGet, restPath("profile"), nil))
		var p profile.Profile
		if err := json.NewDecoder(get.Body).Decode(&p); err != nil {
			t.Fatalf("decode profile: %v", err)
		}
		if p.Settings == nil || p.Settings.StoreName != "Reseeded Store" {
			t.Errorf("profile storeName = %+v, want Reseeded Store", p.Settings)
		}
	})

	t.Run("seed targets a named store via storeId", func(t *testing.T) {
		srv := newFixtureServer(nil)
		body := `{"email":"tenant@example.com"}`
		rec := do(srv, httptest.NewRequest(http.MethodPost, "/_mock/fixtures/customers?storeId=2222", strings.NewReader(body)))
		if rec.Code != http.StatusOK {
			t.Fatalf("seed status = %d, want 200", rec.Code)
		}
		// It lands under store 2222, not the default store.
		other := do(srv, fixtureReq(http.MethodGet, "/api/v3/2222/customers?email=tenant@example.com", nil))
		var res customers.SearchResult
		if err := json.NewDecoder(other.Body).Decode(&res); err != nil {
			t.Fatalf("decode: %v", err)
		}
		if res.Total != 1 {
			t.Errorf("store 2222 total = %d, want 1", res.Total)
		}
		// The default store is unaffected.
		def := do(srv, fixtureReq(http.MethodGet, restPath("customers")+"?email=tenant@example.com", nil))
		var defRes customers.SearchResult
		if err := json.NewDecoder(def.Body).Decode(&defRes); err != nil {
			t.Fatalf("decode: %v", err)
		}
		if defRes.Total != 0 {
			t.Errorf("default store leaked the tenant customer: total = %d, want 0", defRes.Total)
		}
	})

	t.Run("malformed seed body is 400", func(t *testing.T) {
		srv := newFixtureServer(nil)
		rec := do(srv, httptest.NewRequest(http.MethodPost, "/_mock/fixtures/customers", strings.NewReader("not json")))
		if rec.Code != http.StatusBadRequest {
			t.Errorf("malformed seed status = %d, want 400", rec.Code)
		}
	})
}

// ── Fixtures are never proxied ───────────────────────────────────────────

// TestFixtures_NotProxied asserts the simulated endpoints are served locally
// even with proxying enabled — a real store is never contacted for them.
func TestFixtures_NotProxied(t *testing.T) {
	srv := New(config.Config{
		Port:          0,
		AccessToken:   testToken,
		StoreID:       testStoreID,
		ProxyStore:    "999",
		ProxyToken:    "real-token",
		ProxyReadonly: false,
	}, discardLogger())
	// Point the proxy at a URL that must never be reached.
	srv.upstreamBase = "http://proxy.invalid"

	for _, suffix := range []string{"profile", "customers", "customers/1001"} {
		rec := do(srv, fixtureReq(http.MethodGet, restPath(suffix), nil))
		if rec.Code != http.StatusOK {
			t.Errorf("GET %s status = %d, want 200 (served locally, not proxied)", suffix, rec.Code)
		}
	}
}

// getFixtureCustomer GETs a customer by id through the REST handler and decodes
// it, failing if the status is not 200.
func getFixtureCustomer(t *testing.T, srv *Server, id int64) customers.Customer {
	t.Helper()
	rec := do(srv, fixtureReq(http.MethodGet, restPath("customers/"+strconv.FormatInt(id, 10)), nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("GET customer %d status = %d, want 200", id, rec.Code)
	}
	var c customers.Customer
	if err := json.NewDecoder(rec.Body).Decode(&c); err != nil {
		t.Fatalf("decode customer: %v", err)
	}
	return c
}

// assertErrorBody asserts the response body is the Ecwid error shape
// {"errorMessage":...} and that the message names the expected scope.
func assertErrorBody(t *testing.T, body []byte, wantScope string) {
	t.Helper()
	var e struct {
		ErrorMessage string `json:"errorMessage"`
	}
	if err := json.Unmarshal(body, &e); err != nil {
		t.Fatalf("error body is not JSON: %v (%s)", err, body)
	}
	if e.ErrorMessage == "" {
		t.Errorf("error body has no errorMessage: %s", body)
	}
	if wantScope != "" && !strings.Contains(e.ErrorMessage, wantScope) {
		t.Errorf("error message %q does not name the required scope %q", e.ErrorMessage, wantScope)
	}
}
