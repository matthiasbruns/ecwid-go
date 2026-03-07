// Package reviews provides access to the Ecwid product reviews API.
package reviews

// CreateResult represents the response from a create operation.
type CreateResult struct {
	ID int64 `json:"id"`
}

// UpdateResult represents the response from an update operation.
type UpdateResult struct {
	UpdateCount int `json:"updateCount"`
}

// DeleteResult represents the response from a delete operation.
type DeleteResult struct {
	DeleteCount int `json:"deleteCount"`
}

// Review represents a product review in Ecwid.
type Review struct {
	ID              int64  `json:"id,omitempty"`
	ProductID       int64  `json:"productId,omitempty"`
	Rating          int    `json:"rating,omitempty"`
	ReviewText      string `json:"reviewText,omitempty"`
	ReviewerName    string `json:"reviewerName,omitempty"`
	Status          string `json:"status,omitempty"`
	CreateDate      string `json:"createDate,omitempty"`
	UpdateDate      string `json:"updateDate,omitempty"`
	CreateTimestamp int64  `json:"createTimestamp,omitempty"`
	UpdateTimestamp int64  `json:"updateTimestamp,omitempty"`
}

// SearchResult is the paginated response from the reviews list API.
type SearchResult struct {
	Total int      `json:"total"`
	Items []Review `json:"items"`
}
