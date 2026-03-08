package products

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/matthiasbruns/ecwid-go/ecwid/internal/api"
)

// Service provides access to the Ecwid products API.
type Service struct {
	requester api.Requester
}

// NewService creates a new products service.
func NewService(requester api.Requester) *Service {
	return &Service{requester: requester}
}

// ── Core ────────────────────────────────────────────────────────────────

// Search returns a paginated list of products.
//
// API: GET /products
// Required scope: read_catalog
func (s *Service) Search(ctx context.Context, opts *SearchOptions) (*SearchResult, error) {
	q := url.Values{}
	if opts != nil {
		if opts.Keyword != "" {
			q.Set("keyword", opts.Keyword)
		}
		if opts.PriceFrom != nil {
			q.Set("priceFrom", fmt.Sprintf("%.2f", *opts.PriceFrom))
		}
		if opts.PriceTo != nil {
			q.Set("priceTo", fmt.Sprintf("%.2f", *opts.PriceTo))
		}
		if opts.Category > 0 {
			q.Set("category", fmt.Sprintf("%d", opts.Category))
		}
		if opts.IncludeProductsFromSubcategories != nil && *opts.IncludeProductsFromSubcategories {
			q.Set("includeProductsFromSubcategories", "true")
		}
		if opts.SortBy != "" {
			q.Set("sortBy", opts.SortBy)
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
		if opts.Enabled != nil {
			q.Set("enabled", fmt.Sprintf("%t", *opts.Enabled))
		}
		if opts.InStock != nil {
			q.Set("inStock", fmt.Sprintf("%t", *opts.InStock))
		}
		if opts.Sku != "" {
			q.Set("sku", opts.Sku)
		}
		if opts.ProductID != "" {
			q.Set("productId", opts.ProductID)
		}
		if opts.BaseURL != "" {
			q.Set("baseUrl", opts.BaseURL)
		}
		if opts.CleanURLs != nil {
			q.Set("cleanUrls", fmt.Sprintf("%t", *opts.CleanURLs))
		}
		if opts.Lang != "" {
			q.Set("lang", opts.Lang)
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
	if err := s.requester.Get(ctx, "/products", q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get returns a single product by ID.
//
// API: GET /products/{productId}
// Required scope: read_catalog
func (s *Service) Get(ctx context.Context, productID int64) (*Product, error) {
	if productID == 0 {
		return nil, errors.New("productID must not be zero")
	}

	path := fmt.Sprintf("/products/%d", productID)

	var result Product
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new product.
//
// API: POST /products
// Required scope: create_catalog
func (s *Service) Create(ctx context.Context, product *Product) (*CreateResult, error) {
	var result CreateResult
	if err := s.requester.Post(ctx, "/products", product, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update modifies a product by ID.
//
// API: PUT /products/{productId}
// Required scope: update_catalog
func (s *Service) Update(ctx context.Context, productID int64, product *Product) (*UpdateResult, error) {
	if productID == 0 {
		return nil, errors.New("productID must not be zero")
	}

	path := fmt.Sprintf("/products/%d", productID)

	var result UpdateResult
	if err := s.requester.Put(ctx, path, product, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a product by ID.
//
// API: DELETE /products/{productId}
// Required scope: update_catalog
func (s *Service) Delete(ctx context.Context, productID int64) (*DeleteResult, error) {
	if productID == 0 {
		return nil, errors.New("productID must not be zero")
	}

	path := fmt.Sprintf("/products/%d", productID)

	var result DeleteResult
	if err := s.requester.Delete(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteAll deletes all products in the store.
//
// API: DELETE /products
// Required scope: update_catalog_batch_delete
func (s *Service) DeleteAll(ctx context.Context) error {
	return s.requester.Delete(ctx, "/products", nil)
}

// ── Variations (Combinations) ───────────────────────────────────────────

// ListCombinations returns all variations for a product.
//
// API: GET /products/{productId}/combinations
// Required scope: read_catalog
func (s *Service) ListCombinations(ctx context.Context, productID int64) ([]Combination, error) {
	if productID == 0 {
		return nil, errors.New("productID must not be zero")
	}

	path := fmt.Sprintf("/products/%d/combinations", productID)

	var result []Combination
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetCombination returns a single variation by ID.
//
// API: GET /products/{productId}/combinations/{combinationId}
// Required scope: read_catalog
func (s *Service) GetCombination(ctx context.Context, productID, combinationID int64) (*Combination, error) {
	if productID == 0 {
		return nil, errors.New("productID must not be zero")
	}
	if combinationID == 0 {
		return nil, errors.New("combinationID must not be zero")
	}

	path := fmt.Sprintf("/products/%d/combinations/%d", productID, combinationID)

	var result Combination
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateCombination creates a new variation for a product.
//
// API: POST /products/{productId}/combinations
// Required scope: create_catalog
func (s *Service) CreateCombination(ctx context.Context, productID int64, combo *Combination) (*CombinationCreateResult, error) {
	if productID == 0 {
		return nil, errors.New("productID must not be zero")
	}

	path := fmt.Sprintf("/products/%d/combinations", productID)

	var result CombinationCreateResult
	if err := s.requester.Post(ctx, path, combo, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateCombination modifies a variation.
//
// API: PUT /products/{productId}/combinations/{combinationId}
// Required scope: update_catalog
func (s *Service) UpdateCombination(ctx context.Context, productID, combinationID int64, combo *Combination) (*UpdateResult, error) {
	if productID == 0 {
		return nil, errors.New("productID must not be zero")
	}
	if combinationID == 0 {
		return nil, errors.New("combinationID must not be zero")
	}

	path := fmt.Sprintf("/products/%d/combinations/%d", productID, combinationID)

	var result UpdateResult
	if err := s.requester.Put(ctx, path, combo, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteCombination deletes a single variation.
//
// API: DELETE /products/{productId}/combinations/{combinationId}
// Required scope: update_catalog
func (s *Service) DeleteCombination(ctx context.Context, productID, combinationID int64) (*DeleteResult, error) {
	if productID == 0 {
		return nil, errors.New("productID must not be zero")
	}
	if combinationID == 0 {
		return nil, errors.New("combinationID must not be zero")
	}

	path := fmt.Sprintf("/products/%d/combinations/%d", productID, combinationID)

	var result DeleteResult
	if err := s.requester.Delete(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteAllCombinations deletes all variations for a product.
//
// API: DELETE /products/{productId}/combinations
// Required scope: update_catalog
func (s *Service) DeleteAllCombinations(ctx context.Context, productID int64) (*DeleteResult, error) {
	if productID == 0 {
		return nil, errors.New("productID must not be zero")
	}

	path := fmt.Sprintf("/products/%d/combinations", productID)

	var result DeleteResult
	if err := s.requester.Delete(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ── Inventory ───────────────────────────────────────────────────────────

// AdjustInventory adjusts stock for a product by a delta value.
//
// API: PUT /products/{productId}/inventory
// Required scope: update_catalog
func (s *Service) AdjustInventory(ctx context.Context, productID int64, adj *InventoryAdjust) (*InventoryAdjustResult, error) {
	if productID == 0 {
		return nil, errors.New("productID must not be zero")
	}

	path := fmt.Sprintf("/products/%d/inventory", productID)

	var result InventoryAdjustResult
	if err := s.requester.Put(ctx, path, adj, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// AdjustCombinationInventory adjusts stock for a variation by a delta value.
//
// API: PUT /products/{productId}/combinations/{combinationId}/inventory
// Required scope: update_catalog
func (s *Service) AdjustCombinationInventory(ctx context.Context, productID, combinationID int64, adj *InventoryAdjust) (*InventoryAdjustResult, error) {
	if productID == 0 {
		return nil, errors.New("productID must not be zero")
	}
	if combinationID == 0 {
		return nil, errors.New("combinationID must not be zero")
	}

	path := fmt.Sprintf("/products/%d/combinations/%d/inventory", productID, combinationID)

	var result InventoryAdjustResult
	if err := s.requester.Put(ctx, path, adj, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ── Images & Gallery (delete only — upload requires multipart transport) ──

// DeleteImage deletes the main product image.
//
// API: DELETE /products/{productId}/image
// Required scope: update_catalog
func (s *Service) DeleteImage(ctx context.Context, productID int64) (*DeleteResult, error) {
	if productID == 0 {
		return nil, errors.New("productID must not be zero")
	}

	path := fmt.Sprintf("/products/%d/image", productID)

	var result DeleteResult
	if err := s.requester.Delete(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteAllGalleryImages deletes all gallery images for a product.
//
// API: DELETE /products/{productId}/gallery
// Required scope: update_catalog
func (s *Service) DeleteAllGalleryImages(ctx context.Context, productID int64) (*DeleteResult, error) {
	if productID == 0 {
		return nil, errors.New("productID must not be zero")
	}

	path := fmt.Sprintf("/products/%d/gallery", productID)

	var result DeleteResult
	if err := s.requester.Delete(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteGalleryImage deletes a specific gallery image.
//
// API: DELETE /products/{productId}/gallery/{galleryImageId}
// Required scope: update_catalog
func (s *Service) DeleteGalleryImage(ctx context.Context, productID, galleryImageID int64) (*DeleteResult, error) {
	if productID == 0 {
		return nil, errors.New("productID must not be zero")
	}
	if galleryImageID == 0 {
		return nil, errors.New("galleryImageID must not be zero")
	}

	path := fmt.Sprintf("/products/%d/gallery/%d", productID, galleryImageID)

	var result DeleteResult
	if err := s.requester.Delete(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteCombinationImage deletes a variation image.
//
// API: DELETE /products/{productId}/combinations/{combinationId}/image
// Required scope: update_catalog
func (s *Service) DeleteCombinationImage(ctx context.Context, productID, combinationID int64) (*DeleteResult, error) {
	if productID == 0 {
		return nil, errors.New("productID must not be zero")
	}
	if combinationID == 0 {
		return nil, errors.New("combinationID must not be zero")
	}

	path := fmt.Sprintf("/products/%d/combinations/%d/image", productID, combinationID)

	var result DeleteResult
	if err := s.requester.Delete(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ── Video ───────────────────────────────────────────────────────────────

// DeleteVideo deletes the main product video.
//
// API: DELETE /products/{productId}/video
// Required scope: update_catalog
func (s *Service) DeleteVideo(ctx context.Context, productID int64) (*DeleteResult, error) {
	if productID == 0 {
		return nil, errors.New("productID must not be zero")
	}

	path := fmt.Sprintf("/products/%d/video", productID)

	var result DeleteResult
	if err := s.requester.Delete(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetGalleryVideo returns a single gallery video.
//
// API: GET /products/{productId}/gallery/video/{galleryVideoId}
// Required scope: read_catalog
func (s *Service) GetGalleryVideo(ctx context.Context, productID, galleryVideoID int64) (*GalleryVideo, error) {
	if productID == 0 {
		return nil, errors.New("productID must not be zero")
	}
	if galleryVideoID == 0 {
		return nil, errors.New("galleryVideoID must not be zero")
	}

	path := fmt.Sprintf("/products/%d/gallery/video/%d", productID, galleryVideoID)

	var result GalleryVideo
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateGalleryVideo updates a gallery video.
//
// API: PUT /products/{productId}/gallery/video/{galleryVideoId}
// Required scope: update_catalog
func (s *Service) UpdateGalleryVideo(ctx context.Context, productID, galleryVideoID int64, video *GalleryVideoUpdate) (*UpdateResult, error) {
	if productID == 0 {
		return nil, errors.New("productID must not be zero")
	}
	if galleryVideoID == 0 {
		return nil, errors.New("galleryVideoID must not be zero")
	}

	path := fmt.Sprintf("/products/%d/gallery/video/%d", productID, galleryVideoID)

	var result UpdateResult
	if err := s.requester.Put(ctx, path, video, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteGalleryVideo deletes a gallery video.
//
// API: DELETE /products/{productId}/gallery/video/{galleryVideoId}
// Required scope: update_catalog
func (s *Service) DeleteGalleryVideo(ctx context.Context, productID, galleryVideoID int64) (*DeleteResult, error) {
	if productID == 0 {
		return nil, errors.New("productID must not be zero")
	}
	if galleryVideoID == 0 {
		return nil, errors.New("galleryVideoID must not be zero")
	}

	path := fmt.Sprintf("/products/%d/gallery/video/%d", productID, galleryVideoID)

	var result DeleteResult
	if err := s.requester.Delete(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ── Files ───────────────────────────────────────────────────────────────

// DeleteAllFiles deletes all downloadable files for a product.
//
// API: DELETE /products/{productId}/files
// Required scope: update_catalog
func (s *Service) DeleteAllFiles(ctx context.Context, productID int64) (*DeleteResult, error) {
	if productID == 0 {
		return nil, errors.New("productID must not be zero")
	}

	path := fmt.Sprintf("/products/%d/files", productID)

	var result DeleteResult
	if err := s.requester.Delete(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetFile returns a single product file by ID.
//
// API: GET /products/{productId}/files/{fileId}
// Required scope: read_catalog
func (s *Service) GetFile(ctx context.Context, productID, fileID int64) (*ProductFile, error) {
	if productID == 0 {
		return nil, errors.New("productID must not be zero")
	}
	if fileID == 0 {
		return nil, errors.New("fileID must not be zero")
	}

	path := fmt.Sprintf("/products/%d/files/%d", productID, fileID)

	var result ProductFile
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateFile updates a product file's metadata.
//
// API: PUT /products/{productId}/files/{fileId}
// Required scope: update_catalog
func (s *Service) UpdateFile(ctx context.Context, productID, fileID int64, update *ProductFileUpdate) (*UpdateResult, error) {
	if productID == 0 {
		return nil, errors.New("productID must not be zero")
	}
	if fileID == 0 {
		return nil, errors.New("fileID must not be zero")
	}

	path := fmt.Sprintf("/products/%d/files/%d", productID, fileID)

	var result UpdateResult
	if err := s.requester.Put(ctx, path, update, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteFile deletes a single product file.
//
// API: DELETE /products/{productId}/files/{fileId}
// Required scope: update_catalog
func (s *Service) DeleteFile(ctx context.Context, productID, fileID int64) (*DeleteResult, error) {
	if productID == 0 {
		return nil, errors.New("productID must not be zero")
	}
	if fileID == 0 {
		return nil, errors.New("fileID must not be zero")
	}

	path := fmt.Sprintf("/products/%d/files/%d", productID, fileID)

	var result DeleteResult
	if err := s.requester.Delete(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ── Other ───────────────────────────────────────────────────────────────

// UpdateMedia updates the product media (images/videos reordering).
//
// API: PUT /products/{productId}/media
// Required scope: update_catalog
func (s *Service) UpdateMedia(ctx context.Context, productID int64, media *MediaUpdate) (*UpdateResult, error) {
	if productID == 0 {
		return nil, errors.New("productID must not be zero")
	}

	path := fmt.Sprintf("/products/%d/media", productID)

	var result UpdateResult
	if err := s.requester.Put(ctx, path, media, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetSortOrder returns the product sort order.
//
// API: GET /products/sort
// Required scope: read_catalog
func (s *Service) GetSortOrder(ctx context.Context, parentCategory int64) (*SortOrder, error) {
	q := url.Values{}
	if parentCategory > 0 {
		q.Set("parentCategory", fmt.Sprintf("%d", parentCategory))
	}

	var result SortOrder
	if err := s.requester.Get(ctx, "/products/sort", q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateSortOrder updates the product sort order.
//
// API: PUT /products/sort
// Required scope: update_catalog
func (s *Service) UpdateSortOrder(ctx context.Context, sort *SortOrderUpdate) (*UpdateResult, error) {
	var result UpdateResult
	if err := s.requester.Put(ctx, "/products/sort", sort, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDeleted returns a paginated list of deleted product references.
//
// API: GET /products/deleted
// Required scope: read_catalog
func (s *Service) GetDeleted(ctx context.Context, opts *DeletedProductsOptions) (*DeletedProductsResult, error) {
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

	var result DeletedProductsResult
	if err := s.requester.Get(ctx, "/products/deleted", q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetFilters returns available product filters.
//
// API: POST /products/filters
// Required scope: read_catalog
func (s *Service) GetFilters(ctx context.Context, req *FiltersRequest) (*FiltersResult, error) {
	var result FiltersResult
	if err := s.requester.Post(ctx, "/products/filters", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ── Product Classes ─────────────────────────────────────────────────────

// ListClasses returns all product classes/types.
//
// API: GET /classes
// Required scope: read_catalog
func (s *Service) ListClasses(ctx context.Context) ([]ProductClass, error) {
	var result []ProductClass
	if err := s.requester.Get(ctx, "/classes", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetClass returns a single product class by ID.
//
// API: GET /classes/{classId}
// Required scope: read_catalog
func (s *Service) GetClass(ctx context.Context, classID int64) (*ProductClass, error) {
	if classID == 0 {
		return nil, errors.New("classID must not be zero")
	}

	path := fmt.Sprintf("/classes/%d", classID)

	var result ProductClass
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateClass creates a new product class.
//
// API: POST /classes
// Required scope: create_catalog
func (s *Service) CreateClass(ctx context.Context, class *ProductClass) (*ProductClassCreateResult, error) {
	var result ProductClassCreateResult
	if err := s.requester.Post(ctx, "/classes", class, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateClass modifies an existing product class.
//
// API: PUT /classes/{classId}
// Required scope: update_catalog
func (s *Service) UpdateClass(ctx context.Context, classID int64, class *ProductClass) (*UpdateResult, error) {
	if classID == 0 {
		return nil, errors.New("classID must not be zero")
	}

	path := fmt.Sprintf("/classes/%d", classID)

	var result UpdateResult
	if err := s.requester.Put(ctx, path, class, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteClass deletes a product class by ID.
//
// API: DELETE /classes/{classId}
// Required scope: update_catalog
func (s *Service) DeleteClass(ctx context.Context, classID int64) (*DeleteResult, error) {
	if classID == 0 {
		return nil, errors.New("classID must not be zero")
	}

	path := fmt.Sprintf("/classes/%d", classID)

	var result DeleteResult
	if err := s.requester.Delete(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ── Brands & Swatches ───────────────────────────────────────────────────

// ListBrands returns all product brands.
//
// API: GET /brands
// Required scope: read_brands
func (s *Service) ListBrands(ctx context.Context) (*BrandsResult, error) {
	var result BrandsResult
	if err := s.requester.Get(ctx, "/brands", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ListSwatches returns recently used product swatches.
//
// API: GET /swatches
// Required scope: read_catalog
func (s *Service) ListSwatches(ctx context.Context) (*SwatchesResult, error) {
	var result SwatchesResult
	if err := s.requester.Get(ctx, "/swatches", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
