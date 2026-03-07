package carts

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/matthiasbruns/ecwid-go/ecwid/internal/api"
)

// Service provides access to the Ecwid abandoned carts API.
type Service struct {
	requester api.Requester
}

// NewService creates a new carts service.
func NewService(requester api.Requester) *Service {
	return &Service{requester: requester}
}

// Search returns a paginated list of abandoned carts.
//
// API: GET /carts
// Required scope: read_orders
func (s *Service) Search(ctx context.Context, opts *SearchOptions) (*SearchResult, error) {
	q := url.Values{}
	if opts != nil {
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
		if opts.CustomerID > 0 {
			q.Set("customerId", fmt.Sprintf("%d", opts.CustomerID))
		}
		if opts.TotalFrom != nil {
			q.Set("totalFrom", fmt.Sprintf("%.2f", *opts.TotalFrom))
		}
		if opts.TotalTo != nil {
			q.Set("totalTo", fmt.Sprintf("%.2f", *opts.TotalTo))
		}
		if opts.Offset > 0 {
			q.Set("offset", fmt.Sprintf("%d", opts.Offset))
		}
		if opts.Limit > 0 {
			q.Set("limit", fmt.Sprintf("%d", opts.Limit))
		}
	}

	var result SearchResult
	if err := s.requester.Get(ctx, "/carts", q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get returns a single abandoned cart by ID.
//
// API: GET /carts/{cartId}
// Required scope: read_orders
func (s *Service) Get(ctx context.Context, cartID string) (*Cart, error) {
	if cartID == "" {
		return nil, errors.New("cartID must not be empty")
	}

	path := "/carts/" + url.PathEscape(cartID)

	var result Cart
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update modifies an abandoned cart.
//
// API: PUT /carts/{cartId}
// Required scope: update_orders
func (s *Service) Update(ctx context.Context, cartID string, req *UpdateRequest) (*UpdateResult, error) {
	if cartID == "" {
		return nil, errors.New("cartID must not be empty")
	}

	path := "/carts/" + url.PathEscape(cartID)

	var result UpdateResult
	if err := s.requester.Put(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Place converts an abandoned cart into an order.
//
// API: POST /carts/{cartId}/place
// Required scope: create_orders
func (s *Service) Place(ctx context.Context, cartID string) (*PlaceResult, error) {
	if cartID == "" {
		return nil, errors.New("cartID must not be empty")
	}

	path := "/carts/" + url.PathEscape(cartID) + "/place"

	var result PlaceResult
	if err := s.requester.Post(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
