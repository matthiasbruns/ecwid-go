package e2e

import (
	"testing"
)

func TestProfile_Get(t *testing.T) {
	ctx := testContext(t)

	p, err := testClient.Profile.Get(ctx)
	if err != nil {
		skipIfForbidden(t, err)
		t.Fatalf("Profile.Get: %v", err)
	}
	if p.GeneralInfo == nil {
		t.Fatal("expected generalInfo to be non-nil")
	}
	if p.GeneralInfo.StoreID == 0 {
		t.Error("expected non-zero storeId")
	}
	if p.Settings != nil && p.Settings.StoreName == "" {
		t.Log("storeName is empty (may be expected for test stores)")
	}
}
