package e2e

import (
	"testing"
)

func TestDomains_Get(t *testing.T) {
	ctx := testContext(t)

	result, err := testClient.Domains.Get(ctx)
	if err != nil {
		skipIfForbidden(t, err)
		t.Fatalf("Domains.Get: %v", err)
	}
	if result.InstantSiteDomain == nil {
		t.Log("no instantSiteDomain returned (may require buy_domains scope)")
	}
}
