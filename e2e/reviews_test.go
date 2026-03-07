package e2e

import (
	"testing"

	"github.com/matthiasbruns/ecwid-go/ecwid/reviews"
)

// TestReviews_List requires a known product ID in the test store.
// It exercises the list endpoint; the product may have zero reviews.
func TestReviews_List(t *testing.T) {
	ctx := testContext(t)

	// Use product ID 1 as a smoke test; adjust if needed.
	result, err := testClient.Reviews.List(ctx, 1)
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	_ = result
}

// TestReviews_CRUD creates, reads, updates, and deletes a review.
// Requires a valid product ID in the test store.
func TestReviews_CRUD(t *testing.T) {
	ctx := testContext(t)

	const productID int64 = 1

	created, err := testClient.Reviews.Create(ctx, productID, &reviews.Review{
		Rating:       5,
		ReviewText:   "E2E test review",
		ReviewerName: "E2E Bot",
		Status:       "APPROVED",
	})
	if err != nil {
		t.Fatalf("Create: %v", err)
	}
	if created.ID == 0 {
		t.Fatal("expected non-zero review ID")
	}

	review, err := testClient.Reviews.Get(ctx, productID, created.ID)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if review.ReviewText != "E2E test review" {
		t.Errorf("expected reviewText=E2E test review, got %s", review.ReviewText)
	}

	updated, err := testClient.Reviews.Update(ctx, productID, created.ID, &reviews.Review{
		ReviewText: "E2E test review updated",
	})
	if err != nil {
		t.Fatalf("Update: %v", err)
	}
	if updated.UpdateCount != 1 {
		t.Errorf("expected updateCount=1, got %d", updated.UpdateCount)
	}

	deleted, err := testClient.Reviews.Delete(ctx, productID, created.ID)
	if err != nil {
		t.Fatalf("Delete: %v", err)
	}
	if deleted.DeleteCount != 1 {
		t.Errorf("expected deleteCount=1, got %d", deleted.DeleteCount)
	}
}
