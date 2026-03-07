package subscriptions

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/matthiasbruns/ecwid-go/ecwid/internal/api"
)

// Service provides access to the Ecwid subscriptions API.
type Service struct {
	requester api.Requester
}

// NewService creates a new subscriptions service.
func NewService(requester api.Requester) *Service {
	return &Service{requester: requester}
}

// Search returns a paginated list of recurring subscriptions.
//
// API: GET /subscriptions
// Required scope: read_subscriptions
func (s *Service) Search(ctx context.Context, opts *SearchOptions) (*SearchResult, error) {
	q := url.Values{}
	if opts != nil {
		if opts.ID > 0 {
			q.Set("id", fmt.Sprintf("%d", opts.ID))
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
		chargeFrom := opts.NextChargeFrom
		if chargeFrom == "" {
			chargeFrom = opts.ChargeFrom
		}
		if chargeFrom != "" {
			q.Set("nextChargeFrom", chargeFrom)
		}
		chargeTo := opts.NextChargeTo
		if chargeTo == "" {
			chargeTo = opts.ChargeTo
		}
		if chargeTo != "" {
			q.Set("nextChargeTo", chargeTo)
		}
		if opts.CancelledFrom != "" {
			q.Set("cancelledFrom", opts.CancelledFrom)
		}
		if opts.CancelledTo != "" {
			q.Set("cancelledTo", opts.CancelledTo)
		}
		if opts.CustomerID > 0 {
			q.Set("customerId", fmt.Sprintf("%d", opts.CustomerID))
		}
		if opts.ProductID > 0 {
			q.Set("productId", fmt.Sprintf("%d", opts.ProductID))
		}
		if opts.RecurringInterval != "" {
			q.Set("recurringInterval", opts.RecurringInterval)
		}
		if opts.Status != "" {
			q.Set("status", opts.Status)
		}
		if opts.Offset > 0 {
			q.Set("offset", fmt.Sprintf("%d", opts.Offset))
		}
		if opts.Limit > 0 {
			q.Set("limit", fmt.Sprintf("%d", opts.Limit))
		}
	}

	var result SearchResult
	if err := s.requester.Get(ctx, "/subscriptions", q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get returns a single subscription by ID.
//
// API: GET /subscriptions/{subscriptionId}
// Required scope: read_subscriptions
func (s *Service) Get(ctx context.Context, subscriptionID int64) (*Subscription, error) {
	if subscriptionID <= 0 {
		return nil, errors.New("subscriptionID must be positive")
	}

	path := fmt.Sprintf("/subscriptions/%d", subscriptionID)

	var result Subscription
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update modifies a subscription by ID.
//
// API: PUT /subscriptions/{subscriptionId}
// Required scope: update_subscriptions
func (s *Service) Update(ctx context.Context, subscriptionID int64, req *UpdateRequest) (*UpdateResult, error) {
	if subscriptionID <= 0 {
		return nil, errors.New("subscriptionID must be positive")
	}

	path := fmt.Sprintf("/subscriptions/%d", subscriptionID)

	var result UpdateResult
	if err := s.requester.Put(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
