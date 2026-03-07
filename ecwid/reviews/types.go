// Package reviews provides access to the Ecwid product reviews API.
package reviews

// UpdateResult represents the response from an update operation.
type UpdateResult struct {
	UpdateCount int `json:"updateCount"`
}

// DeleteResult represents the response from a delete operation.
type DeleteResult struct {
	DeleteCount int `json:"deleteCount"`
}

// Review represents a product review in Ecwid.
// Docs: https://docs.ecwid.com/api-reference/rest-api/products/product-reviews/search-product-reviews.md
type Review struct {
	ID              int64         `json:"id,omitempty"`
	Status          string        `json:"status,omitempty"`
	CustomerID      int64         `json:"customerId,omitempty"`
	ProductID       int64         `json:"productId,omitempty"`
	OrderID         string        `json:"orderId,omitempty"`
	Rating          int           `json:"rating,omitempty"`
	Review          string        `json:"review,omitempty"`
	ReviewerInfo    *ReviewerInfo `json:"reviewerInfo,omitempty"`
	CreateDate      string        `json:"createDate,omitempty"`
	UpdateDate      string        `json:"updateDate,omitempty"`
	CreateTimestamp int64         `json:"createTimestamp,omitempty"`
	UpdateTimestamp int64         `json:"updateTimestamp,omitempty"`
}

// ReviewerInfo holds details about the review author.
type ReviewerInfo struct {
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
	City   string `json:"city,omitempty"`
	Orders int    `json:"orders,omitempty"`
}

// SearchOptions holds query parameters for searching reviews.
type SearchOptions struct {
	Status      string
	Rating      int
	OrderID     string
	ProductID   int64
	ReviewID    int64
	CreatedFrom string
	CreatedTo   string
	UpdatedFrom string
	UpdatedTo   string
	SortBy      string
	Keyword     string
	Offset      int
	Limit       int
}

// SearchResult is the paginated response from the reviews search API.
type SearchResult struct {
	Total  int      `json:"total"`
	Count  int      `json:"count"`
	Offset int      `json:"offset"`
	Limit  int      `json:"limit"`
	Items  []Review `json:"items"`
}

// UpdateStatusRequest is the body for updating a review's status.
type UpdateStatusRequest struct {
	Status string `json:"status"`
}

// BulkUpdateRequest is the body for bulk update/delete of reviews.
// Docs: https://docs.ecwid.com/api-reference/rest-api/products/product-reviews/bulk-update-delete-product-reviews.md
type BulkUpdateRequest struct {
	SelectMode     string          `json:"selectMode"`
	Delete         *bool           `json:"delete,omitempty"`
	NewStatus      string          `json:"newStatus,omitempty"`
	ReviewIDs      []int64         `json:"reviewIds,omitempty"`
	CurrentFilters *CurrentFilters `json:"currentFilters,omitempty"`
}

// CurrentFilters specifies search criteria for bulk operations.
type CurrentFilters struct {
	ReviewID  []int64 `json:"reviewId,omitempty"`
	ProductID []int64 `json:"productId,omitempty"`
	OrderID   []int64 `json:"orderId,omitempty"`
	Rating    []int   `json:"rating,omitempty"`
	Status    string  `json:"status,omitempty"`
}
