package webhook

import (
	"strings"
	"testing"

	"github.com/matthiasbruns/ecwid-go/ecwid/webhooks"
)

// wantEventCount is the number of Ecwid webhook events the mock fires. It is
// asserted so an accidental add or drop in the catalog is caught.
const wantEventCount = 42

func TestCatalog_Count(t *testing.T) {
	if got := len(Catalog()); got != wantEventCount {
		t.Fatalf("catalog has %d events, want %d", got, wantEventCount)
	}
}

func TestCatalog_TypesUnique(t *testing.T) {
	seen := map[webhooks.EventType]bool{}
	for _, s := range Catalog() {
		if seen[s.Type] {
			t.Errorf("duplicate event type %q", s.Type)
		}
		seen[s.Type] = true
	}
}

func TestCatalog_LookupRoundTrips(t *testing.T) {
	for _, s := range Catalog() {
		got, ok := Lookup(s.Type)
		if !ok {
			t.Errorf("Lookup(%q) not found", s.Type)
			continue
		}
		if got.EntityID != s.EntityID {
			t.Errorf("Lookup(%q).EntityID = %q, want %q", s.Type, got.EntityID, s.EntityID)
		}
	}
	if _, ok := Lookup("no.such.event"); ok {
		t.Error("Lookup of an unknown event returned ok")
	}
}

func TestCatalog_EntityIDType(t *testing.T) {
	for _, s := range Catalog() {
		want := EntityIDNumber
		if strings.HasPrefix(string(s.Type), applicationPrefix) {
			want = EntityIDString
		}
		if got := s.EntityIDType(); got != want {
			t.Errorf("%q EntityIDType = %q, want %q", s.Type, got, want)
		}
	}
}

// The families Ecwid documents as carrying no data must have a nil fixture so
// they are emitted with no data key at all — never "data":{}.
func TestCatalog_NoDataFamilies(t *testing.T) {
	noData := map[webhooks.EventType]bool{
		webhooks.EventProductCreated: true, webhooks.EventProductUpdated: true, webhooks.EventProductDeleted: true,
		webhooks.EventCategoryCreated: true, webhooks.EventCategoryUpdated: true, webhooks.EventCategoryDeleted: true,
		webhooks.EventDiscountCouponCreated: true, webhooks.EventDiscountCouponUpdated: true, webhooks.EventDiscountCouponDeleted: true,
		webhooks.EventApplicationInstalled: true, webhooks.EventApplicationUninstalled: true,
	}
	for _, s := range Catalog() {
		if noData[s.Type] && s.HasData() {
			t.Errorf("%q must carry no data, got %s", s.Type, s.Data)
		}
	}
}

func TestCatalog_GroupDerivation(t *testing.T) {
	cases := map[webhooks.EventType]string{
		webhooks.EventOrderCreated:           "order",
		webhooks.EventCustomerGroupCreated:   "customer_group",
		webhooks.EventApplicationInstalled:   "application",
		webhooks.EventUnfinishedOrderCreated: "unfinished_order",
	}
	for et, want := range cases {
		s, _ := Lookup(et)
		if got := s.Group(); got != want {
			t.Errorf("%q Group = %q, want %q", et, got, want)
		}
	}
}
