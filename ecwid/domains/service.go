package domains

import (
	"context"
	"errors"
	"net/url"

	"github.com/matthiasbruns/ecwid-go/ecwid/internal/api"
)

// Service provides access to the Ecwid store domain settings API.
type Service struct {
	requester api.Requester
}

// NewService creates a new domains service.
func NewService(requester api.Requester) *Service {
	return &Service{requester: requester}
}

// Get returns the store domain settings.
//
// API: GET /domain
func (s *Service) Get(ctx context.Context) (*DomainSettings, error) {
	var result DomainSettings
	if err := s.requester.Get(ctx, "/domain", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update modifies the store domain settings.
//
// API: PUT /domain
func (s *Service) Update(ctx context.Context, settings *DomainSettings) (*UpdateResult, error) {
	var result UpdateResult
	if err := s.requester.Put(ctx, "/domain", settings, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Templates returns a list of available domain templates.
//
// API: GET /domain/templates
func (s *Service) Templates(ctx context.Context) ([]DomainTemplate, error) {
	var result []DomainTemplate
	if err := s.requester.Get(ctx, "/domain/templates", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// Whois checks the availability of a domain name.
//
// API: GET /domain/whois/{domainName}
func (s *Service) Whois(ctx context.Context, domainName string) (*WhoisResult, error) {
	if domainName == "" {
		return nil, errors.New("domainName must not be empty")
	}

	path := "/domain/whois/" + url.PathEscape(domainName)

	var result WhoisResult
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Purchase initiates a domain purchase.
//
// API: POST /domain/purchase
func (s *Service) Purchase(ctx context.Context, req *PurchaseRequest) (*PurchaseResult, error) {
	var result PurchaseResult
	if err := s.requester.Post(ctx, "/domain/purchase", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Remove removes the custom domain from the store.
//
// API: DELETE /domain
func (s *Service) Remove(ctx context.Context) (*DeleteResult, error) {
	var result DeleteResult
	if err := s.requester.Delete(ctx, "/domain", &result); err != nil {
		return nil, err
	}
	return &result, nil
}
