package e2e

import (
	"testing"
)

func TestStaff_List(t *testing.T) {
	ctx := testContext(t)

	result, err := testClient.Staff.List(ctx)
	if err != nil {
		skipIfForbidden(t, err)
		t.Fatalf("Staff.List: %v", err)
	}
	if len(result.StaffList) == 0 {
		t.Log("no staff accounts found (may be expected)")
	}
}
