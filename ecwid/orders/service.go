package orders

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/matthiasbruns/ecwid-go/ecwid/internal/api"
)

// Service provides access to the Ecwid orders API.
type Service struct {
	requester api.Requester
}

// NewService creates a new orders service.
func NewService(requester api.Requester) *Service {
	return &Service{requester: requester}
}

// Search returns a paginated list of orders.
//
// API: GET /orders
// Required scope: read_orders
func (s *Service) Search(ctx context.Context, opts *SearchOptions) (*SearchResult, error) {
	q := url.Values{}
	if opts != nil {
		if opts.IDs != "" {
			q.Set("ids", opts.IDs)
		}
		if opts.Keywords != "" {
			q.Set("keywords", opts.Keywords)
		}
		if opts.Email != "" {
			q.Set("email", opts.Email)
		}
		if opts.CustomerID > 0 {
			q.Set("customerId", fmt.Sprintf("%d", opts.CustomerID))
		}
		if opts.ProductID != "" {
			q.Set("productId", opts.ProductID)
		}
		if opts.TotalFrom != nil {
			q.Set("totalFrom", fmt.Sprintf("%.2f", *opts.TotalFrom))
		}
		if opts.TotalTo != nil {
			q.Set("totalTo", fmt.Sprintf("%.2f", *opts.TotalTo))
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
		if opts.PickupTimeFrom != "" {
			q.Set("pickupTimeFrom", opts.PickupTimeFrom)
		}
		if opts.PickupTimeTo != "" {
			q.Set("pickupTimeTo", opts.PickupTimeTo)
		}
		if opts.ShippingMethod != "" {
			q.Set("shippingMethod", opts.ShippingMethod)
		}
		if opts.FulfillmentStatus != "" {
			q.Set("fulfillmentStatus", opts.FulfillmentStatus)
		}
		if opts.PaymentMethod != "" {
			q.Set("paymentMethod", opts.PaymentMethod)
		}
		if opts.PaymentModule != "" {
			q.Set("paymentModule", opts.PaymentModule)
		}
		if opts.PaymentStatus != "" {
			q.Set("paymentStatus", opts.PaymentStatus)
		}
		if opts.AcceptMarketing != nil {
			q.Set("acceptMarketing", fmt.Sprintf("%t", *opts.AcceptMarketing))
		}
		if opts.ContainsPreorderItems != nil {
			q.Set("containsPreorderItems", fmt.Sprintf("%t", *opts.ContainsPreorderItems))
		}
		if opts.CouponCode != "" {
			q.Set("couponCode", opts.CouponCode)
		}
		if opts.SubscriptionID > 0 {
			q.Set("subscriptionId", fmt.Sprintf("%d", opts.SubscriptionID))
		}
		if opts.RefererID > 0 {
			q.Set("refererId", fmt.Sprintf("%d", opts.RefererID))
		}
		if opts.ResponseFields != "" {
			q.Set("responseFields", opts.ResponseFields)
		}
		if opts.Offset > 0 {
			q.Set("offset", fmt.Sprintf("%d", opts.Offset))
		}
		if opts.Limit > 0 {
			q.Set("limit", fmt.Sprintf("%d", opts.Limit))
		}
	}

	var result SearchResult
	if err := s.requester.Get(ctx, "/orders", q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get returns a single order by ID.
//
// API: GET /orders/{orderId}
// Required scope: read_orders
func (s *Service) Get(ctx context.Context, orderID string) (*Order, error) {
	if orderID == "" {
		return nil, errors.New("orderID must not be empty")
	}

	path := "/orders/" + url.PathEscape(orderID)

	var result Order
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new order.
//
// API: POST /orders
// Required scope: create_orders
func (s *Service) Create(ctx context.Context, req *CreateRequest) (*CreateResult, error) {
	var result CreateResult
	if err := s.requester.Post(ctx, "/orders", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update modifies an order by ID.
//
// API: PUT /orders/{orderId}
// Required scope: update_orders
func (s *Service) Update(ctx context.Context, orderID string, req *UpdateRequest) (*UpdateResult, error) {
	if orderID == "" {
		return nil, errors.New("orderID must not be empty")
	}

	path := "/orders/" + url.PathEscape(orderID)

	var result UpdateResult
	if err := s.requester.Put(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes an order by ID.
//
// API: DELETE /orders/{orderId}
// Required scope: update_orders
func (s *Service) Delete(ctx context.Context, orderID string) (*DeleteResult, error) {
	if orderID == "" {
		return nil, errors.New("orderID must not be empty")
	}

	path := "/orders/" + url.PathEscape(orderID)

	var result DeleteResult
	if err := s.requester.Delete(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetLast returns the most recently created order.
//
// API: GET /orders/last
// Required scope: read_orders
func (s *Service) GetLast(ctx context.Context) (*Order, error) {
	var result Order
	if err := s.requester.Get(ctx, "/orders/last", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDeleted returns a paginated list of deleted order references.
//
// API: GET /orders/deleted
// Required scope: read_orders
func (s *Service) GetDeleted(ctx context.Context, opts *DeletedOrdersOptions) (*DeletedOrdersResult, error) {
	q := url.Values{}
	if opts != nil {
		if opts.FromDate != "" {
			q.Set("from_date", opts.FromDate)
		}
		if opts.ToDate != "" {
			q.Set("to_date", opts.ToDate)
		}
		if opts.Offset > 0 {
			q.Set("offset", fmt.Sprintf("%d", opts.Offset))
		}
		if opts.Limit > 0 {
			q.Set("limit", fmt.Sprintf("%d", opts.Limit))
		}
	}

	var result DeletedOrdersResult
	if err := s.requester.Get(ctx, "/orders/deleted", q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetRepeatOrderURL returns a URL to repeat the specified order.
//
// API: GET /orders/{orderId}/repeatURL
// Required scope: read_orders
func (s *Service) GetRepeatOrderURL(ctx context.Context, orderID string) (*RepeatOrderURLResult, error) {
	if orderID == "" {
		return nil, errors.New("orderID must not be empty")
	}

	path := "/orders/" + url.PathEscape(orderID) + "/repeatURL"

	var result RepeatOrderURLResult
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetInvoices returns tax invoices for the specified order.
//
// API: GET /orders/{orderId}/invoices
// Required scope: read_orders
func (s *Service) GetInvoices(ctx context.Context, orderID string) (*InvoicesResult, error) {
	if orderID == "" {
		return nil, errors.New("orderID must not be empty")
	}

	path := "/orders/" + url.PathEscape(orderID) + "/invoices"

	var result InvoicesResult
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateInvoice generates a tax invoice for the specified order.
//
// API: POST /orders/{orderId}/invoices/create-invoice
// Required scope: create_orders
func (s *Service) CreateInvoice(ctx context.Context, orderID string) (*CreateInvoiceResult, error) {
	if orderID == "" {
		return nil, errors.New("orderID must not be empty")
	}

	path := "/orders/" + url.PathEscape(orderID) + "/invoices/create-invoice"

	var result CreateInvoiceResult
	if err := s.requester.Post(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetExtraFields returns extra fields for the specified order.
//
// API: GET /orders/{orderId}/extraFields
// Required scope: read_orders
func (s *Service) GetExtraFields(ctx context.Context, orderID string) ([]ExtraField, error) {
	if orderID == "" {
		return nil, errors.New("orderID must not be empty")
	}

	path := "/orders/" + url.PathEscape(orderID) + "/extraFields"

	var result []ExtraField
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateExtraField adds an extra field to the specified order.
//
// API: POST /orders/{orderId}/extraFields
// Required scope: update_orders
func (s *Service) CreateExtraField(ctx context.Context, orderID string, field *ExtraField) (*CreateExtraFieldResult, error) {
	if orderID == "" {
		return nil, errors.New("orderID must not be empty")
	}

	path := "/orders/" + url.PathEscape(orderID) + "/extraFields"

	var result CreateExtraFieldResult
	if err := s.requester.Post(ctx, path, field, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateExtraField updates an extra field on the specified order.
//
// API: PUT /orders/{orderId}/extraFields/{extraFieldId}
// Required scope: update_orders
func (s *Service) UpdateExtraField(ctx context.Context, orderID, extraFieldID string, field *UpdateExtraFieldRequest) (*UpdateResult, error) {
	if orderID == "" {
		return nil, errors.New("orderID must not be empty")
	}
	if extraFieldID == "" {
		return nil, errors.New("extraFieldID must not be empty")
	}

	path := "/orders/" + url.PathEscape(orderID) + "/extraFields/" + url.PathEscape(extraFieldID)

	var result UpdateResult
	if err := s.requester.Put(ctx, path, field, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteExtraField deletes an extra field from the specified order.
//
// API: DELETE /orders/{orderId}/extraFields/{extraFieldId}
// Required scope: update_orders
func (s *Service) DeleteExtraField(ctx context.Context, orderID, extraFieldID string) (*DeleteResult, error) {
	if orderID == "" {
		return nil, errors.New("orderID must not be empty")
	}
	if extraFieldID == "" {
		return nil, errors.New("extraFieldID must not be empty")
	}

	path := "/orders/" + url.PathEscape(orderID) + "/extraFields/" + url.PathEscape(extraFieldID)

	var result DeleteResult
	if err := s.requester.Delete(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Calculate computes order totals, taxes, and shipping without creating the order.
//
// API: POST /order/calculate
// Required scope: read_store_profile
func (s *Service) Calculate(ctx context.Context, req *CalculateRequest) (*Order, error) {
	var result Order
	if err := s.requester.Post(ctx, "/order/calculate", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
