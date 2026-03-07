package domains

import (
	"context"

	"github.com/matthiasbruns/ecwid-go/ecwid/internal/api"
)

// Service provides access to the Ecwid store domains API.
type Service struct {
	requester api.Requester
}

// NewService creates a new domains service.
func NewService(requester api.Requester) *Service {
	return &Service{requester: requester}
}

// Get returns the store domain settings including Instant Site and purchased domains.
//
// API: GET /domains
// Required scopes: read_store_profile, buy_domains
func (s *Service) Get(ctx context.Context) (*DomainsResult, error) {
	var result DomainsResult
	if err := s.requester.Get(ctx, "/domains", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Purchase buys a new domain for the store.
//
// API: POST /domains/purchase
// Required scope: buy_domains
func (s *Service) Purchase(ctx context.Context, req *PurchaseRequest) (*PurchaseResult, error) {
	var result PurchaseResult
	if err := s.requester.Post(ctx, "/domains/purchase", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
