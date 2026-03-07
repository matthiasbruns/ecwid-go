package categories

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/matthiasbruns/ecwid-go/ecwid/internal/api"
)

// Service provides access to the Ecwid categories API.
type Service struct {
	requester api.Requester
}

// NewService creates a new categories service.
func NewService(requester api.Requester) *Service {
	return &Service{requester: requester}
}

// Search returns a paginated list of categories.
//
// API: GET /categories
// Required scope: read_catalog
func (s *Service) Search(ctx context.Context, opts *SearchOptions) (*SearchResult, error) {
	q := url.Values{}
	if opts != nil {
		if opts.Parent > 0 {
			q.Set("parent", fmt.Sprintf("%d", opts.Parent))
		}
		if opts.HiddenCategories {
			q.Set("hidden_categories", "true")
		}
		if opts.ProductIds != "" {
			q.Set("productIds", opts.ProductIds)
		}
		if opts.Offset > 0 {
			q.Set("offset", fmt.Sprintf("%d", opts.Offset))
		}
		if opts.Limit > 0 {
			q.Set("limit", fmt.Sprintf("%d", opts.Limit))
		}
	}

	var result SearchResult
	if err := s.requester.Get(ctx, "/categories", q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get returns a single category by ID.
//
// API: GET /categories/{categoryId}
// Required scope: read_catalog
func (s *Service) Get(ctx context.Context, categoryID int64) (*Category, error) {
	if categoryID == 0 {
		return nil, errors.New("categoryID must not be zero")
	}

	path := fmt.Sprintf("/categories/%d", categoryID)

	var result Category
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new category.
//
// API: POST /categories
// Required scope: create_catalog
func (s *Service) Create(ctx context.Context, cat *Category) (*CreateResult, error) {
	var result CreateResult
	if err := s.requester.Post(ctx, "/categories", cat, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update modifies an existing category.
//
// API: PUT /categories/{categoryId}
// Required scope: update_catalog
func (s *Service) Update(ctx context.Context, categoryID int64, cat *Category) (*UpdateResult, error) {
	if categoryID == 0 {
		return nil, errors.New("categoryID must not be zero")
	}

	path := fmt.Sprintf("/categories/%d", categoryID)

	var result UpdateResult
	if err := s.requester.Put(ctx, path, cat, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete removes a category by ID.
//
// API: DELETE /categories/{categoryId}
// Required scope: update_catalog
func (s *Service) Delete(ctx context.Context, categoryID int64) (*DeleteResult, error) {
	if categoryID == 0 {
		return nil, errors.New("categoryID must not be zero")
	}

	path := fmt.Sprintf("/categories/%d", categoryID)

	var result DeleteResult
	if err := s.requester.Delete(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetProducts returns the product IDs assigned to a category.
//
// API: GET /categories/{categoryId}/products
// Required scope: read_catalog
func (s *Service) GetProducts(ctx context.Context, categoryID int64) (*ProductsResult, error) {
	if categoryID == 0 {
		return nil, errors.New("categoryID must not be zero")
	}

	path := fmt.Sprintf("/categories/%d/products", categoryID)

	var result ProductsResult
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
