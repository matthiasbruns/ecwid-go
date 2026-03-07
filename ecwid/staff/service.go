package staff

import (
	"context"
	"errors"
	"net/url"

	"github.com/matthiasbruns/ecwid-go/ecwid/internal/api"
)

// Service provides access to the Ecwid staff accounts API.
type Service struct {
	requester api.Requester
}

// NewService creates a new staff service.
func NewService(requester api.Requester) *Service {
	return &Service{requester: requester}
}

// List returns all staff accounts for the store.
//
// API: GET /staff
// Required scope: read_staff
func (s *Service) List(ctx context.Context) (*ListResult, error) {
	var result ListResult
	if err := s.requester.Get(ctx, "/staff", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get returns a single staff account by ID.
//
// API: GET /staff/{staffAccountId}
// Required scope: read_staff
func (s *Service) Get(ctx context.Context, staffAccountID string) (*GetResult, error) {
	if staffAccountID == "" {
		return nil, errors.New("staffAccountID must not be empty")
	}

	path := "/staff/" + url.PathEscape(staffAccountID)

	var result GetResult
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create invites a new staff account.
//
// API: POST /staff
// Required scopes: create_staff, invite_staff
func (s *Service) Create(ctx context.Context, req *CreateRequest) (*CreateResult, error) {
	var result CreateResult
	if err := s.requester.Post(ctx, "/staff", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update modifies a staff account by ID.
//
// API: PUT /staff/{staffAccountId}
// Required scopes: read_staff, update_staff
func (s *Service) Update(ctx context.Context, staffAccountID string, req *UpdateRequest) (*UpdateResult, error) {
	if staffAccountID == "" {
		return nil, errors.New("staffAccountID must not be empty")
	}

	path := "/staff/" + url.PathEscape(staffAccountID)

	var result UpdateResult
	if err := s.requester.Put(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete removes a staff account by ID.
//
// API: DELETE /staff/{staffAccountId}
// Required scope: delete_staff
func (s *Service) Delete(ctx context.Context, staffAccountID string) (*DeleteResult, error) {
	if staffAccountID == "" {
		return nil, errors.New("staffAccountID must not be empty")
	}

	path := "/staff/" + url.PathEscape(staffAccountID)

	var result DeleteResult
	if err := s.requester.Delete(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
