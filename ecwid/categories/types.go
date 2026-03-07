// Package categories provides access to the Ecwid categories API.
package categories

// Category represents a product category in Ecwid.
type Category struct {
	ID               int64   `json:"id"`
	ParentID         int64   `json:"parentId,omitempty"`
	Name             string  `json:"name,omitempty"`
	Description      string  `json:"description,omitempty"`
	Enabled          *bool   `json:"enabled,omitempty"`
	OrderBy          int     `json:"orderBy,omitempty"`
	ProductCount     int     `json:"productCount,omitempty"`
	URL              string  `json:"url,omitempty"`
	ImageURL         string  `json:"imageUrl,omitempty"`
	HdImageURL       string  `json:"hdImageUrl,omitempty"`
	OriginalImageURL string  `json:"originalImageUrl,omitempty"`
	IsSampleCategory bool    `json:"isSampleCategory,omitempty"`
	ProductIDs       []int64 `json:"productIds,omitempty"`
}

// SearchResult is the paginated response from the categories search API.
type SearchResult struct {
	Total  int        `json:"total"`
	Count  int        `json:"count"`
	Offset int        `json:"offset"`
	Limit  int        `json:"limit"`
	Items  []Category `json:"items"`
}

// SearchOptions holds query parameters for searching categories.
type SearchOptions struct {
	Parent           int64
	HiddenCategories bool
	ProductIds       string
	Offset           int
	Limit            int
}

// ProductsResult is the response from the category products endpoint.
type ProductsResult struct {
	ProductIDs []int64 `json:"productIds"`
}

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
