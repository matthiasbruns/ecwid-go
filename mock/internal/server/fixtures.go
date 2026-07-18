package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"maps"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/matthiasbruns/ecwid-go/ecwid/customers"
	"github.com/matthiasbruns/ecwid-go/ecwid/profile"
)

// Ecwid access scopes gated by the simulated REST endpoints. A request whose
// token was not granted the required scope gets the same 403 a real store
// returns, so a consumer can exercise the unhappy path (see Config.HasScope).
const (
	scopeReadStoreProfile = "read_store_profile"
	scopeReadCustomers    = "read_customers"
	scopeUpdateCustomers  = "update_customers"
)

// Customers paging bounds, matching the real API: the default page size when no
// limit is given and the hard cap Ecwid enforces on limit.
const (
	defaultCustomersLimit = 100
	maxCustomersLimit     = 100

	// maxCustomerBodyBytes bounds a customer-update request body so a runaway
	// client cannot make the mock buffer without limit. A customer record is
	// small; 1 MB is generous.
	maxCustomerBodyBytes = 1 << 20
)

// fixtureStore is the mock's in-memory backing for the simulated customer and
// store-profile endpoints. It is the default source of truth for those routes
// so the mock needs no real store and no proxy to answer them. It is keyed by
// store ID (the {storeId} path segment) so a multi-tenant consumer can seed and
// exercise several stores against one mock. Safe for concurrent use.
type fixtureStore struct {
	mu        sync.RWMutex
	profiles  map[string]*profile.Profile             // storeID -> profile
	customers map[string]map[int64]customers.Customer // storeID -> id -> customer
}

// newFixtureStore returns an empty fixture store.
func newFixtureStore() *fixtureStore {
	return &fixtureStore{
		profiles:  make(map[string]*profile.Profile),
		customers: make(map[string]map[int64]customers.Customer),
	}
}

// setProfile installs the store profile for storeID, replacing any existing one.
func (fs *fixtureStore) setProfile(storeID string, p *profile.Profile) {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	fs.profiles[storeID] = p
}

// getProfile returns the seeded profile for storeID and whether one is present.
func (fs *fixtureStore) getProfile(storeID string) (*profile.Profile, bool) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()
	p, ok := fs.profiles[storeID]
	return p, ok
}

// putCustomer upserts a customer under storeID. A customer with a zero ID is
// assigned the next free ID; the assigned (or supplied) ID is returned.
func (fs *fixtureStore) putCustomer(storeID string, c customers.Customer) int64 {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	if fs.customers[storeID] == nil {
		fs.customers[storeID] = make(map[int64]customers.Customer)
	}
	if c.ID == 0 {
		c.ID = fs.allocIDLocked(storeID)
	}
	fs.customers[storeID][c.ID] = c
	return c.ID
}

// allocIDLocked returns an unused customer ID for storeID: one past the highest
// existing ID, floored so generated IDs stay in a plausible range. Callers must
// hold the write lock.
func (fs *fixtureStore) allocIDLocked(storeID string) int64 {
	var maxID int64 = 1000
	for id := range fs.customers[storeID] {
		if id > maxID {
			maxID = id
		}
	}
	return maxID + 1
}

// getCustomer returns a customer by ID for storeID and whether it exists.
func (fs *fixtureStore) getCustomer(storeID string, id int64) (customers.Customer, bool) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()
	c, ok := fs.customers[storeID][id]
	return c, ok
}

// searchCustomers returns a page of customers for storeID in the paged envelope
// the real API uses (items/count/total/offset/limit). When email is non-empty
// only customers whose email matches case-insensitively are returned, which is
// how FindCustomerByEmail narrows the list. Results are ordered by ID so paging
// is deterministic.
func (fs *fixtureStore) searchCustomers(storeID, email string, offset, limit int) customers.SearchResult {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	store := fs.customers[storeID]
	matched := make([]customers.Customer, 0, len(store))
	for id := range store {
		if email != "" && !strings.EqualFold(store[id].Email, email) {
			continue
		}
		matched = append(matched, store[id])
	}
	sort.Slice(matched, func(i, j int) bool { return matched[i].ID < matched[j].ID })

	total := len(matched)
	lo := min(offset, total)
	hi := min(lo+limit, total)
	page := matched[lo:hi]

	return customers.SearchResult{
		Total:  total,
		Count:  len(page),
		Offset: offset,
		Limit:  limit,
		Items:  page,
	}
}

// patchCustomer applies a partial update to the customer identified by id under
// storeID and returns the merged customer. The patch is a set of raw JSON fields
// (as the SDK sends on PUT /customers/{id}); every supplied field overwrites the
// stored one while untouched fields are preserved, so a subsequent GET reflects
// the write (e.g. flipping acceptMarketing). The ID is immutable. The bool
// reports whether the customer existed; a non-nil error means the merge produced
// invalid JSON.
func (fs *fixtureStore) patchCustomer(storeID string, id int64, patch map[string]json.RawMessage) (customers.Customer, bool, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	existing, ok := fs.customers[storeID][id]
	if !ok {
		return customers.Customer{}, false, nil
	}

	raw, err := json.Marshal(existing)
	if err != nil {
		return customers.Customer{}, true, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(raw, &fields); err != nil {
		return customers.Customer{}, true, err
	}
	maps.Copy(fields, patch)
	merged, err := json.Marshal(fields)
	if err != nil {
		return customers.Customer{}, true, err
	}
	var updated customers.Customer
	if err := json.Unmarshal(merged, &updated); err != nil {
		return customers.Customer{}, true, err
	}
	updated.ID = id // the path ID is authoritative; a body cannot re-key a customer.
	fs.customers[storeID][id] = updated
	return updated, true, nil
}

// seedDefaults installs a ready-to-use store profile and a small deterministic
// customer set for storeID, so the fixtures answer the happy path out of the box
// with no seeding required by the consumer.
func seedDefaults(fs *fixtureStore, storeID string) {
	fs.setProfile(storeID, defaultProfile(storeID))
	seeds := defaultCustomers()
	for i := range seeds {
		fs.putCustomer(storeID, seeds[i])
	}
}

// defaultProfile builds a realistic, minimal store profile for storeID using the
// ecwid/profile types, so the mock's shape and field names match the typed
// client exactly.
func defaultProfile(storeID string) *profile.Profile {
	id, _ := strconv.ParseInt(storeID, 10, 64) // non-numeric store IDs simply yield 0.
	return &profile.Profile{
		GeneralInfo: &profile.GeneralInfo{
			StoreID:  id,
			StoreURL: "https://mystore.example.com",
		},
		Account: &profile.Account{
			AccountName:  "Mock Store Owner",
			AccountEmail: "owner@example.com",
		},
		Settings: &profile.Settings{
			StoreName: "Mock Store",
		},
		Company: &profile.Company{
			CompanyName: "Mock Store",
			Email:       "store@example.com",
			CountryCode: "US",
		},
		FormatsAndUnits: &profile.FormatsAndUnits{
			Currency:       "USD",
			CurrencyPrefix: "$",
			WeightUnit:     "LB",
			Timezone:       "America/New_York",
		},
		Languages: &profile.Languages{
			EnabledLanguages: []string{"en"},
			DefaultLanguage:  "en",
		},
	}
}

// defaultCustomers is the deterministic starter customer set. It deliberately
// covers the shapes a consumer tests against: a known email to find, and both
// acceptMarketing states (including one set to false, ready to be flipped true).
func defaultCustomers() []customers.Customer {
	return []customers.Customer{
		{
			ID:              1001,
			Name:            "Ada Lovelace",
			Email:           "ada@example.com",
			Registered:      "2024-01-15 10:30:00 +0000",
			Updated:         "2024-01-15 10:30:00 +0000",
			AcceptMarketing: boolPtr(true),
			Lang:            "en",
		},
		{
			ID:              1002,
			Name:            "Grace Hopper",
			Email:           "grace@example.com",
			Registered:      "2024-02-20 09:00:00 +0000",
			Updated:         "2024-02-20 09:00:00 +0000",
			AcceptMarketing: boolPtr(false),
			Lang:            "en",
		},
		{
			ID:              1003,
			Name:            "Alan Turing",
			Email:           "alan@example.com",
			Registered:      "2024-03-10 14:45:00 +0000",
			Updated:         "2024-03-10 14:45:00 +0000",
			AcceptMarketing: boolPtr(false),
			Lang:            "en",
		},
	}
}

// boolPtr returns a pointer to b, for the *bool fields in the ecwid types.
func boolPtr(b bool) *bool { return &b }

// requireAuthScope gates a simulated REST request on both the mock's bearer
// token and the required Ecwid scope. A missing/wrong token yields 401 (as the
// storage routes do); a valid token that lacks the scope yields the 403 shape a
// real store returns. It returns true only when the request may proceed.
func (s *Server) requireAuthScope(w http.ResponseWriter, r *http.Request, scope string) bool {
	if !s.authorized(r) {
		writeJSONError(w, http.StatusUnauthorized, "invalid or missing access token")
		return false
	}
	if !s.cfg.HasScope(scope) {
		writeJSONError(w, http.StatusForbidden,
			"This method requires the '"+scope+"' access scope, which the access token was not granted.")
		return false
	}
	return true
}

// handleProfileGet serves GET /api/v3/{storeId}/profile (scope
// read_store_profile). A store with no seeded profile falls back to a generated
// default so the endpoint is never empty.
func (s *Server) handleProfileGet(w http.ResponseWriter, r *http.Request) {
	if !s.requireAuthScope(w, r, scopeReadStoreProfile) {
		return
	}
	storeID := r.PathValue("storeId")
	p, ok := s.fixtures.getProfile(storeID)
	if !ok {
		p = defaultProfile(storeID)
	}
	writeJSON(w, http.StatusOK, p)
}

// handleCustomersSearch serves GET /api/v3/{storeId}/customers (scope
// read_customers): a paged customer list, optionally narrowed by ?email= (the
// FindCustomerByEmail path) and windowed by ?offset= / ?limit=.
func (s *Server) handleCustomersSearch(w http.ResponseWriter, r *http.Request) {
	if !s.requireAuthScope(w, r, scopeReadCustomers) {
		return
	}
	q := r.URL.Query()
	res := s.fixtures.searchCustomers(
		r.PathValue("storeId"),
		q.Get("email"),
		parseOffset(q.Get("offset")),
		parseLimit(q.Get("limit")),
	)
	writeJSON(w, http.StatusOK, res)
}

// handleCustomerGet serves GET /api/v3/{storeId}/customers/{id} (scope
// read_customers): a single customer, or 404 when absent.
func (s *Server) handleCustomerGet(w http.ResponseWriter, r *http.Request) {
	if !s.requireAuthScope(w, r, scopeReadCustomers) {
		return
	}
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		writeJSONError(w, http.StatusNotFound, "customer not found")
		return
	}
	c, ok := s.fixtures.getCustomer(r.PathValue("storeId"), id)
	if !ok {
		writeJSONError(w, http.StatusNotFound, "customer not found")
		return
	}
	writeJSON(w, http.StatusOK, c)
}

// handleCustomerUpdate serves PUT /api/v3/{storeId}/customers/{id} (scope
// update_customers). The body is a partial customer; supplied fields (e.g.
// acceptMarketing) are merged into the stored record so a later GET reflects the
// write. It answers with {"updateCount":1}, matching the real API, or 404 when
// the customer does not exist.
func (s *Server) handleCustomerUpdate(w http.ResponseWriter, r *http.Request) {
	if !s.requireAuthScope(w, r, scopeUpdateCustomers) {
		return
	}
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		writeJSONError(w, http.StatusNotFound, "customer not found")
		return
	}

	body, ok := readFixtureBody(w, r, maxCustomerBodyBytes)
	if !ok {
		return
	}
	var patch map[string]json.RawMessage
	if err := json.Unmarshal(body, &patch); err != nil {
		writeJSONError(w, http.StatusBadRequest, "request body is not a JSON object")
		return
	}

	_, ok, err = s.fixtures.patchCustomer(r.PathValue("storeId"), id, patch)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid customer update: "+err.Error())
		return
	}
	if !ok {
		writeJSONError(w, http.StatusNotFound, "customer not found")
		return
	}
	writeJSON(w, http.StatusOK, customers.UpdateResult{UpdateCount: 1})
}

// readFixtureBody reads a fixture-mutation request body under a size limit. It
// reads one byte past the limit (like the storage routes) so an over-limit body
// is reported as 413 rather than silently truncated into a different, possibly
// valid payload; and it rejects a bare JSON null (400), which the decoders would
// otherwise accept as an empty map/slice/struct and treat as success. It returns
// the body and whether the caller may proceed; on false it has already written
// the error response.
func readFixtureBody(w http.ResponseWriter, r *http.Request, limit int) ([]byte, bool) {
	body, err := io.ReadAll(io.LimitReader(r.Body, int64(limit)+1))
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "failed to read request body")
		return nil, false
	}
	if len(body) > limit {
		writeJSONError(w, http.StatusRequestEntityTooLarge,
			fmt.Sprintf("request body exceeds the %d-byte limit", limit))
		return nil, false
	}
	if bytes.Equal(bytes.TrimSpace(body), []byte("null")) {
		writeJSONError(w, http.StatusBadRequest, "request body must not be null")
		return nil, false
	}
	return body, true
}

// parseOffset parses a non-negative offset, defaulting to 0 for an absent or
// invalid value.
func parseOffset(raw string) int {
	if raw == "" {
		return 0
	}
	n, err := strconv.Atoi(raw)
	if err != nil || n < 0 {
		return 0
	}
	return n
}

// parseLimit parses the page size, defaulting to defaultCustomersLimit when
// absent or invalid and clamping to maxCustomersLimit, matching the real API.
func parseLimit(raw string) int {
	if raw == "" {
		return defaultCustomersLimit
	}
	n, err := strconv.Atoi(raw)
	if err != nil || n <= 0 {
		return defaultCustomersLimit
	}
	if n > maxCustomersLimit {
		return maxCustomersLimit
	}
	return n
}
