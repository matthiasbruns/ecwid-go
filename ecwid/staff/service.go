package staff

import (
	"context"
	"errors"
	"fmt"

	"github.com/matthiasbruns/ecwid-go/ecwid/internal/api"
)

// Service provides access to the Ecwid staff members API.
type Service struct {
	requester api.Requester
}

// NewService creates a new staff service.
func NewService(requester api.Requester) *Service {
	return &Service{requester: requester}
}

// List returns all staff members.
//
// API: GET /staff
func (s *Service) List(ctx context.Context) (*ListResult, error) {
	var result ListResult
	if err := s.requester.Get(ctx, "/staff", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get returns a single staff member by ID.
//
// API: GET /staff/{staffId}
func (s *Service) Get(ctx context.Context, staffID int64) (*StaffMember, error) {
	if staffID == 0 {
		return nil, errors.New("staffID must not be zero")
	}

	path := fmt.Sprintf("/staff/%d", staffID)

	var result StaffMember
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new staff member.
//
// API: POST /staff
func (s *Service) Create(ctx context.Context, member *StaffMember) (*CreateResult, error) {
	var result CreateResult
	if err := s.requester.Post(ctx, "/staff", member, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update modifies an existing staff member.
//
// API: PUT /staff/{staffId}
func (s *Service) Update(ctx context.Context, staffID int64, member *StaffMember) (*UpdateResult, error) {
	if staffID == 0 {
		return nil, errors.New("staffID must not be zero")
	}

	path := fmt.Sprintf("/staff/%d", staffID)

	var result UpdateResult
	if err := s.requester.Put(ctx, path, member, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete removes a staff member.
//
// API: DELETE /staff/{staffId}
func (s *Service) Delete(ctx context.Context, staffID int64) (*DeleteResult, error) {
	if staffID == 0 {
		return nil, errors.New("staffID must not be zero")
	}

	path := fmt.Sprintf("/staff/%d", staffID)

	var result DeleteResult
	if err := s.requester.Delete(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
