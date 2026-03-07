package reviews

import (
	"context"
	"errors"
	"fmt"

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

// List returns all reviews for a product.
//
// API: GET /products/{productId}/reviews
func (s *Service) List(ctx context.Context, productID int64) (*SearchResult, error) {
	if productID == 0 {
		return nil, errors.New("productID must not be zero")
	}

	path := fmt.Sprintf("/products/%d/reviews", productID)

	var result SearchResult
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get returns a single review by product and review ID.
//
// API: GET /products/{productId}/reviews/{reviewId}
func (s *Service) Get(ctx context.Context, productID, reviewID int64) (*Review, error) {
	if productID == 0 {
		return nil, errors.New("productID must not be zero")
	}
	if reviewID == 0 {
		return nil, errors.New("reviewID must not be zero")
	}

	path := fmt.Sprintf("/products/%d/reviews/%d", productID, reviewID)

	var result Review
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new review for a product.
//
// API: POST /products/{productId}/reviews
func (s *Service) Create(ctx context.Context, productID int64, review *Review) (*CreateResult, error) {
	if productID == 0 {
		return nil, errors.New("productID must not be zero")
	}

	path := fmt.Sprintf("/products/%d/reviews", productID)

	var result CreateResult
	if err := s.requester.Post(ctx, path, review, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing review.
//
// API: PUT /products/{productId}/reviews/{reviewId}
func (s *Service) Update(ctx context.Context, productID, reviewID int64, review *Review) (*UpdateResult, error) {
	if productID == 0 {
		return nil, errors.New("productID must not be zero")
	}
	if reviewID == 0 {
		return nil, errors.New("reviewID must not be zero")
	}

	path := fmt.Sprintf("/products/%d/reviews/%d", productID, reviewID)

	var result UpdateResult
	if err := s.requester.Put(ctx, path, review, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a single review.
//
// API: DELETE /products/{productId}/reviews/{reviewId}
func (s *Service) Delete(ctx context.Context, productID, reviewID int64) (*DeleteResult, error) {
	if productID == 0 {
		return nil, errors.New("productID must not be zero")
	}
	if reviewID == 0 {
		return nil, errors.New("reviewID must not be zero")
	}

	path := fmt.Sprintf("/products/%d/reviews/%d", productID, reviewID)

	var result DeleteResult
	if err := s.requester.Delete(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteAll deletes all reviews for a product.
//
// API: DELETE /products/{productId}/reviews
func (s *Service) DeleteAll(ctx context.Context, productID int64) (*DeleteResult, error) {
	if productID == 0 {
		return nil, errors.New("productID must not be zero")
	}

	path := fmt.Sprintf("/products/%d/reviews", productID)

	var result DeleteResult
	if err := s.requester.Delete(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
