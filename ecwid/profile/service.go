package profile

import (
	"context"

	"github.com/matthiasbruns/ecwid-go/ecwid/internal/api"
)

// Service provides access to the Ecwid store profile API.
type Service struct {
	requester api.Requester
}

// NewService creates a new profile service.
func NewService(requester api.Requester) *Service {
	return &Service{requester: requester}
}

// Get returns the full store profile.
//
// API: GET /profile
// Required scope: read_store_profile
func (s *Service) Get(ctx context.Context) (*Profile, error) {
	var result Profile
	if err := s.requester.Get(ctx, "/profile", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update modifies the store profile.
//
// API: PUT /profile
// Required scope: update_store_profile
func (s *Service) Update(ctx context.Context, req *UpdateRequest) (*UpdateResult, error) {
	var result UpdateResult
	if err := s.requester.Put(ctx, "/profile", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
