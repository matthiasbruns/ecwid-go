package customers

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/matthiasbruns/ecwid-go/ecwid/internal/api"
)

// Service provides access to the Ecwid customers API.
type Service struct {
	requester api.Requester
}

// NewService creates a new customers service.
func NewService(requester api.Requester) *Service {
	return &Service{requester: requester}
}

// Search returns a paginated list of customers.
//
// API: GET /customers
// Required scope: read_customers
func (s *Service) Search(ctx context.Context, opts *SearchOptions) (*SearchResult, error) {
	q := url.Values{}
	if opts != nil {
		if opts.Keyword != "" {
			q.Set("keyword", opts.Keyword)
		}
		if opts.Name != "" {
			q.Set("name", opts.Name)
		}
		if opts.Email != "" {
			q.Set("email", opts.Email)
		}
		if opts.UseExactEmailMatch != nil && *opts.UseExactEmailMatch {
			q.Set("useExactEmailMatch", "true")
		}
		if opts.Phone != "" {
			q.Set("phone", opts.Phone)
		}
		if opts.City != "" {
			q.Set("city", opts.City)
		}
		if opts.PostalCode != "" {
			q.Set("postalCode", opts.PostalCode)
		}
		if opts.StateOrProvinceCode != "" {
			q.Set("stateOrProvinceCode", opts.StateOrProvinceCode)
		}
		if opts.CountryCodes != "" {
			q.Set("countryCodes", opts.CountryCodes)
		}
		if opts.CompanyName != "" {
			q.Set("companyName", opts.CompanyName)
		}
		if opts.AcceptMarketing != nil {
			q.Set("acceptMarketing", fmt.Sprintf("%t", *opts.AcceptMarketing))
		}
		if opts.Lang != "" {
			q.Set("lang", opts.Lang)
		}
		if opts.CustomerGroupIDs != "" {
			q.Set("customerGroupIds", opts.CustomerGroupIDs)
		}
		if opts.MinOrderCount > 0 {
			q.Set("minOrderCount", fmt.Sprintf("%d", opts.MinOrderCount))
		}
		if opts.MaxOrderCount > 0 {
			q.Set("maxOrderCount", fmt.Sprintf("%d", opts.MaxOrderCount))
		}
		if opts.MinSalesValue != nil {
			q.Set("minSalesValue", fmt.Sprintf("%g", *opts.MinSalesValue))
		}
		if opts.MaxSalesValue != nil {
			q.Set("maxSalesValue", fmt.Sprintf("%g", *opts.MaxSalesValue))
		}
		if opts.PurchasedProductIDs != "" {
			q.Set("purchasedProductIds", opts.PurchasedProductIDs)
		}
		if opts.B2BB2C != "" {
			q.Set("b2b_b2c", opts.B2BB2C)
		}
		if opts.TaxExempt != nil {
			q.Set("taxExempt", fmt.Sprintf("%t", *opts.TaxExempt))
		}
		if opts.SortBy != "" {
			q.Set("sortBy", opts.SortBy)
		}
		if opts.Offset > 0 {
			q.Set("offset", fmt.Sprintf("%d", opts.Offset))
		}
		if opts.Limit > 0 {
			q.Set("limit", fmt.Sprintf("%d", opts.Limit))
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
	}

	var result SearchResult
	if err := s.requester.Get(ctx, "/customers", q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get returns a single customer by ID.
//
// API: GET /customers/{customerId}
// Required scope: read_customers
func (s *Service) Get(ctx context.Context, customerID int64) (*Customer, error) {
	if customerID == 0 {
		return nil, errors.New("customerID must not be zero")
	}

	path := fmt.Sprintf("/customers/%d", customerID)

	var result Customer
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new customer.
//
// API: POST /customers
// Required scope: create_customers
func (s *Service) Create(ctx context.Context, cust *Customer) (*CreateResult, error) {
	var result CreateResult
	if err := s.requester.Post(ctx, "/customers", cust, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update modifies an existing customer.
//
// API: PUT /customers/{customerId}
// Required scope: update_customers
func (s *Service) Update(ctx context.Context, customerID int64, cust *Customer) (*UpdateResult, error) {
	if customerID == 0 {
		return nil, errors.New("customerID must not be zero")
	}

	path := fmt.Sprintf("/customers/%d", customerID)

	var result UpdateResult
	if err := s.requester.Put(ctx, path, cust, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete removes a customer by ID.
//
// API: DELETE /customers/{customerId}
// Required scope: update_customers
func (s *Service) Delete(ctx context.Context, customerID int64) (*DeleteResult, error) {
	if customerID == 0 {
		return nil, errors.New("customerID must not be zero")
	}

	path := fmt.Sprintf("/customers/%d", customerID)

	var result DeleteResult
	if err := s.requester.Delete(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetOrders returns orders for a customer.
//
// API: GET /customers/{customerId}/orders
// Required scope: read_orders
func (s *Service) GetOrders(ctx context.Context, customerID int64) (*OrdersResult, error) {
	if customerID == 0 {
		return nil, errors.New("customerID must not be zero")
	}

	path := fmt.Sprintf("/customers/%d/orders", customerID)

	var result OrdersResult
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
