package reviews

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/matthiasbruns/ecwid-go/ecwid/internal/api"
)

// Service provides access to the Ecwid product reviews API.
type Service struct {
	requester api.Requester
}

// NewService creates a new reviews service.
func NewService(requester api.Requester) *Service {
	return &Service{requester: requester}
}

// Search returns a paginated list of product reviews.
//
// API: GET /reviews
// Required scope: read_reviews
func (s *Service) Search(ctx context.Context, opts *SearchOptions) (*SearchResult, error) {
	q := url.Values{}
	if opts != nil {
		if opts.Status != "" {
			q.Set("status", opts.Status)
		}
		if opts.Rating > 0 {
			q.Set("rating", fmt.Sprintf("%d", opts.Rating))
		}
		if opts.OrderID != "" {
			q.Set("orderId", opts.OrderID)
		}
		if opts.ProductID > 0 {
			q.Set("productId", fmt.Sprintf("%d", opts.ProductID))
		}
		if opts.ReviewID > 0 {
			q.Set("reviewId", fmt.Sprintf("%d", opts.ReviewID))
		}
		if opts.CreatedFrom != "" {
			q.Set("createdFrom", opts.CreatedFrom)
		}
		if opts.CreatedTo != "" {
			q.Set("createdTo", opts.CreatedTo)
		}
		if opts.UpdatedFrom != "" {
			q.Set("updatedFrom", opts.UpdatedFrom)
		}
		if opts.UpdatedTo != "" {
			q.Set("updatedTo", opts.UpdatedTo)
		}
		if opts.SortBy != "" {
			q.Set("sortBy", opts.SortBy)
		}
		if opts.Keyword != "" {
			q.Set("keyword", opts.Keyword)
		}
		if opts.Offset > 0 {
			q.Set("offset", fmt.Sprintf("%d", opts.Offset))
		}
		if opts.Limit > 0 {
			q.Set("limit", fmt.Sprintf("%d", opts.Limit))
		}
	}

	var result SearchResult
	if err := s.requester.Get(ctx, "/reviews", q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateStatus updates the status of a review (moderated/published).
//
// API: PUT /reviews/{reviewId}
// Required scope: update_reviews
func (s *Service) UpdateStatus(ctx context.Context, reviewID int64, status string) (*UpdateResult, error) {
	if reviewID == 0 {
		return nil, errors.New("reviewID must not be zero")
	}

	path := fmt.Sprintf("/reviews/%d", reviewID)

	var result UpdateResult
	if err := s.requester.Put(ctx, path, &UpdateStatusRequest{Status: status}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete removes a review by ID.
//
// API: DELETE /reviews/{reviewId}
// Required scopes: read_reviews, update_reviews
func (s *Service) Delete(ctx context.Context, reviewID int64) (*DeleteResult, error) {
	if reviewID == 0 {
		return nil, errors.New("reviewID must not be zero")
	}

	path := fmt.Sprintf("/reviews/%d", reviewID)

	var result DeleteResult
	if err := s.requester.Delete(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BulkUpdate performs bulk status update or deletion of reviews.
//
// API: PUT /reviews/mass_update
// Required scope: update_reviews
func (s *Service) BulkUpdate(ctx context.Context, req *BulkUpdateRequest) (*UpdateResult, error) {
	var result UpdateResult
	if err := s.requester.Put(ctx, "/reviews/mass_update", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
