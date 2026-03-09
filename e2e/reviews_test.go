package e2e

import (
	"testing"
)

func TestReviews_Search(t *testing.T) {
	ctx := testContext(t)

	result, err := testClient.Reviews.Search(ctx, nil)
	if err != nil {
		skipIfForbidden(t, err)
		t.Fatalf("Reviews.Search: %v", err)
	}
	t.Logf("found %d reviews", result.Total)
}
