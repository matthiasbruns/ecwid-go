package e2e

import (
	"testing"
)

func TestDomains_Get(t *testing.T) {
	ctx := testContext(t)

	settings, err := testClient.Domains.Get(ctx)
	if err != nil {
		t.Fatalf("Domains.Get: %v", err)
	}
	// Just verify it returns without error; fields may be empty for test stores.
	_ = settings
}

func TestDomains_Templates(t *testing.T) {
	ctx := testContext(t)

	templates, err := testClient.Domains.Templates(ctx)
	if err != nil {
		t.Fatalf("Domains.Templates: %v", err)
	}
	_ = templates
}
