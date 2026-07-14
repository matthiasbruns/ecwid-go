package instantsite

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/matthiasbruns/ecwid-go/ecwid/internal/api"
)

// Service provides access to the Ecwid Instant Site API.
//
// It routes calls across three requesters because the Instant Site API spans
// two hosts and two tokens (see the package doc): the redirect endpoints use the
// main API host + main token; the v1 endpoints use the Instant Site host + the
// Instant Site token; the token-exchange endpoint uses the auth host.
type Service struct {
	main api.Requester // app.ecwid.com/api/v3 — redirects (main token)
	v1   api.Requester // vuega.ecwid.com/api/v1 — profile/pages/tiles/themes/… (instant-site token)
	auth api.Requester // app.ecwid.com/instantsite — token exchange
}

// NewService creates a new Instant Site service. The main requester serves the
// v3 redirect endpoints, v1 serves the Instant Site v1 endpoints, and auth
// serves the token-exchange endpoint.
func NewService(main, v1, auth api.Requester) *Service {
	return &Service{main: main, v1: v1, auth: auth}
}

// ── Token exchange ──────────────────────────────────────────────────────────

// Token exchanges an app secret token for a short-lived (24h) Instant Site
// access token. Store the returned token in config (InstantSiteToken) to build a
// client that can call the v1 Instant Site endpoints.
//
// API: POST /oauth/token
// Required scope: manage_instant_site
func (s *Service) Token(ctx context.Context, req *TokenRequest) (*TokenResult, error) {
	if req == nil {
		return nil, errors.New("token request must not be nil")
	}
	if req.SiteID == "" {
		return nil, errors.New("token request SiteID must not be empty")
	}
	if req.Code == "" {
		return nil, errors.New("token request Code must not be empty")
	}

	grantType := req.GrantType
	if grantType == "" {
		grantType = GrantTypeAuthorizationCode
	}

	q := url.Values{}
	q.Set("grant_type", grantType)
	q.Set("site_id", req.SiteID)
	q.Set("code", req.Code)

	var result TokenResult
	// The auth requester is built with an empty store ID, so its base URL ends
	// with a slash; the path here is intentionally slash-less.
	if err := s.auth.Post(ctx, "oauth/token?"+q.Encode(), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ── Profile ─────────────────────────────────────────────────────────────────

// GetProfile returns the Instant Site profile.
//
// API: GET /profile
// Required scope: manage_instant_site
func (s *Service) GetProfile(ctx context.Context) (*Profile, error) {
	var result Profile
	if err := s.v1.Get(ctx, "/profile", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateProfile creates the Instant Site profile.
//
// API: POST /profile
// Required scope: manage_instant_site
func (s *Service) CreateProfile(ctx context.Context, profile *Profile) (*Profile, error) {
	var result Profile
	if err := s.v1.Post(ctx, "/profile", profile, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateProfile updates the Instant Site profile.
//
// API: PUT /profile
// Required scope: manage_instant_site
func (s *Service) UpdateProfile(ctx context.Context, profile *Profile) (*Profile, error) {
	var result Profile
	if err := s.v1.Put(ctx, "/profile", profile, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ── Lifecycle ───────────────────────────────────────────────────────────────

// Publish publishes the current Instant Site draft.
//
// API: POST /publish
// Required scope: manage_instant_site
func (s *Service) Publish(ctx context.Context) error {
	return s.v1.Post(ctx, "/publish", struct{}{}, nil)
}

// Discard discards unpublished Instant Site draft changes.
//
// API: POST /discard
// Required scope: manage_instant_site
func (s *Service) Discard(ctx context.Context) error {
	return s.v1.Post(ctx, "/discard", struct{}{}, nil)
}

// Clone clones an Instant Site from another store.
//
// API: POST /clone
// Required scopes: manage_instant_site, clone_stores
func (s *Service) Clone(ctx context.Context, req *CloneRequest) error {
	if req == nil {
		return errors.New("clone request must not be nil")
	}
	return s.v1.Post(ctx, "/clone", req, nil)
}

// ── Pages ───────────────────────────────────────────────────────────────────

// ListPages returns the Instant Site pages. Set published to true for the
// published site or false for the draft.
//
// API: GET /page?published={published}
// Required scope: manage_instant_site
func (s *Service) ListPages(ctx context.Context, published bool) (*PageList, error) {
	q := url.Values{}
	q.Set("published", fmt.Sprintf("%t", published))

	var result PageList
	if err := s.v1.Get(ctx, "/page", q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreatePage creates a new Instant Site page.
//
// API: POST /page
// Required scope: manage_instant_site
func (s *Service) CreatePage(ctx context.Context, page *Page) (*CreatePageResult, error) {
	var result CreatePageResult
	if err := s.v1.Post(ctx, "/page", page, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdatePage updates an Instant Site page.
//
// API: PUT /page/{pageId}
// Required scope: manage_instant_site
func (s *Service) UpdatePage(ctx context.Context, pageID string, page *Page) (*Page, error) {
	if pageID == "" {
		return nil, errors.New("pageID must not be empty")
	}
	path := "/page/" + url.PathEscape(pageID)

	var result Page
	if err := s.v1.Put(ctx, path, page, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeletePage deletes an Instant Site page and returns the deleted page.
//
// API: DELETE /page/{pageId}
// Required scope: manage_instant_site
func (s *Service) DeletePage(ctx context.Context, pageID string) (*Page, error) {
	if pageID == "" {
		return nil, errors.New("pageID must not be empty")
	}
	path := "/page/" + url.PathEscape(pageID)

	var result Page
	if err := s.v1.Delete(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ── Tiles ───────────────────────────────────────────────────────────────────

// ListTiles returns Instant Site tiles, optionally filtered by page and
// language.
//
// API: GET /tile?published={published}
// Required scope: manage_instant_site
func (s *Service) ListTiles(ctx context.Context, opts *TileListOptions) (*TileList, error) {
	q := url.Values{}
	if opts != nil {
		q.Set("published", fmt.Sprintf("%t", opts.Published))
		if opts.PageID != "" {
			q.Set("pageId", opts.PageID)
		}
		if opts.Lang != "" {
			q.Set("lang", opts.Lang)
		}
	} else {
		q.Set("published", "false")
	}

	var result TileList
	if err := s.v1.Get(ctx, "/tile", q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetTile returns a single Instant Site tile by ID.
//
// API: GET /tile/{tileId}
// Required scope: manage_instant_site
func (s *Service) GetTile(ctx context.Context, tileID string) (*Tile, error) {
	if tileID == "" {
		return nil, errors.New("tileID must not be empty")
	}
	path := "/tile/" + url.PathEscape(tileID)

	var result Tile
	if err := s.v1.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateTile creates a new Instant Site tile.
//
// API: POST /tile
// Required scope: manage_instant_site
func (s *Service) CreateTile(ctx context.Context, req *CreateTileRequest) (*Tile, error) {
	var result Tile
	if err := s.v1.Post(ctx, "/tile", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateTile updates a single Instant Site tile.
//
// API: PUT /tile/{tileId}
// Required scope: manage_instant_site
func (s *Service) UpdateTile(ctx context.Context, tileID string, tile *TileUpdate) (*Tile, error) {
	if tileID == "" {
		return nil, errors.New("tileID must not be empty")
	}
	path := "/tile/" + url.PathEscape(tileID)

	var result Tile
	if err := s.v1.Put(ctx, path, tile, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateTiles updates the whole tiles list in one request.
//
// API: PUT /tile
// Required scope: manage_instant_site
func (s *Service) UpdateTiles(ctx context.Context, tiles *TileBulkUpdate) (*TileList, error) {
	var result TileList
	if err := s.v1.Put(ctx, "/tile", tiles, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteTile deletes an Instant Site tile, removing it from every page it was
// on.
//
// API: DELETE /tile/{tileId}
// Required scope: manage_instant_site
func (s *Service) DeleteTile(ctx context.Context, tileID string) (*Tile, error) {
	if tileID == "" {
		return nil, errors.New("tileID must not be empty")
	}
	path := "/tile/" + url.PathEscape(tileID)

	var result Tile
	if err := s.v1.Delete(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// TileShowcase returns the available tile templates grouped by category.
//
// API: GET /tile/showcase
// Required scope: manage_instant_site
func (s *Service) TileShowcase(ctx context.Context) (*TileShowcaseResult, error) {
	var result TileShowcaseResult
	if err := s.v1.Get(ctx, "/tile/showcase", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// TileConfig returns the content/design editor config for a tile category type.
//
// API: GET /tile/config/{configType}
// Required scope: manage_instant_site
func (s *Service) TileConfig(ctx context.Context, configType string) (*TileConfigResult, error) {
	if configType == "" {
		return nil, errors.New("configType must not be empty")
	}
	path := "/tile/config/" + url.PathEscape(configType)

	var result TileConfigResult
	if err := s.v1.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ── Tile images ─────────────────────────────────────────────────────────────

// ReserveTileImage reserves an upload slot for a tile image. Upload the image
// binary as multipart/form-data to the returned URL.
//
// API: POST /tile/{tileId}/image
// Required scope: manage_instant_site
func (s *Service) ReserveTileImage(ctx context.Context, tileID string) (*ReserveImageResult, error) {
	if tileID == "" {
		return nil, errors.New("tileID must not be empty")
	}
	path := "/tile/" + url.PathEscape(tileID) + "/image"

	var result ReserveImageResult
	if err := s.v1.Post(ctx, path, struct{}{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetImage returns the rendered result of an uploaded tile image.
//
// API: GET /image/{imageId}
// Required scope: manage_instant_site
func (s *Service) GetImage(ctx context.Context, imageID string) (*ImageResult, error) {
	if imageID == "" {
		return nil, errors.New("imageID must not be empty")
	}
	path := "/image/" + url.PathEscape(imageID)

	var result ImageResult
	if err := s.v1.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ImageBuckets returns the map of image bucket IDs to their CDN base URLs.
//
// API: GET /image/bucket
// Required scope: manage_instant_site
func (s *Service) ImageBuckets(ctx context.Context) (*ImageBucketsResult, error) {
	var result ImageBucketsResult
	if err := s.v1.Get(ctx, "/image/bucket", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ── Themes ──────────────────────────────────────────────────────────────────

// ListThemes returns the store's saved Instant Site themes.
//
// API: GET /themes
// Required scope: manage_instant_site
func (s *Service) ListThemes(ctx context.Context) (*ThemeList, error) {
	var result ThemeList
	if err := s.v1.Get(ctx, "/themes", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateTheme creates a new Instant Site theme from a color palette.
//
// API: POST /themes
// Required scope: manage_instant_site
func (s *Service) CreateTheme(ctx context.Context, colors *ThemeColors) (*Theme, error) {
	var result Theme
	if err := s.v1.Post(ctx, "/themes", themeBody(colors), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateTheme updates an existing Instant Site theme.
//
// API: PUT /themes/{themeId}
// Required scope: manage_instant_site
func (s *Service) UpdateTheme(ctx context.Context, themeID string, colors *ThemeColors) (*Theme, error) {
	if themeID == "" {
		return nil, errors.New("themeID must not be empty")
	}
	path := "/themes/" + url.PathEscape(themeID)

	var result Theme
	if err := s.v1.Put(ctx, path, themeBody(colors), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteTheme deletes an Instant Site theme and returns the deleted theme.
//
// API: DELETE /themes/{themeId}
// Required scope: manage_instant_site
func (s *Service) DeleteTheme(ctx context.Context, themeID string) (*Theme, error) {
	if themeID == "" {
		return nil, errors.New("themeID must not be empty")
	}
	path := "/themes/" + url.PathEscape(themeID)

	var result Theme
	if err := s.v1.Delete(ctx, path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CurrentTheme returns the store's currently applied theme palette.
//
// API: GET /current_theme
// Required scope: manage_instant_site
func (s *Service) CurrentTheme(ctx context.Context) (*CurrentTheme, error) {
	var result CurrentTheme
	if err := s.v1.Get(ctx, "/current_theme", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateCurrentTheme sets the store's currently applied theme palette.
//
// API: PUT /current_theme
// Required scope: manage_instant_site
func (s *Service) UpdateCurrentTheme(ctx context.Context, colors *ThemeColors) (*Theme, error) {
	var result Theme
	if err := s.v1.Put(ctx, "/current_theme", themeBody(colors), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// themeBody wraps a color palette in the {"colors": …} request envelope the
// theme endpoints expect.
func themeBody(colors *ThemeColors) any {
	var c ThemeColors
	if colors != nil {
		c = *colors
	}
	return struct {
		Colors ThemeColors `json:"colors"`
	}{Colors: c}
}

// ── Text labels ─────────────────────────────────────────────────────────────

// TextLabels returns the Instant Site text labels (i18n translations).
//
// API: GET /translation
// Required scope: manage_instant_site
func (s *Service) TextLabels(ctx context.Context) (*TextLabels, error) {
	var result TextLabels
	if err := s.v1.Get(ctx, "/translation", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ── Redirects (v3, main requester) ──────────────────────────────────────────

// SearchRedirects returns a paginated list of 301 URL redirects.
//
// API: GET /instant-site/redirects
// Required scope: manage_instant_site
func (s *Service) SearchRedirects(ctx context.Context, opts *RedirectSearchOptions) (*RedirectSearchResult, error) {
	q := url.Values{}
	if opts != nil {
		if opts.Keyword != "" {
			q.Set("keyword", opts.Keyword)
		}
		if opts.Offset > 0 {
			q.Set("offset", fmt.Sprintf("%d", opts.Offset))
		}
		if opts.Limit > 0 {
			q.Set("limit", fmt.Sprintf("%d", opts.Limit))
		}
	}

	var result RedirectSearchResult
	if err := s.main.Get(ctx, "/instant-site/redirects", q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetRedirect returns a single redirect by ID.
//
// API: GET /instant-site/redirects/{redirectId}
// Required scope: manage_instant_site
func (s *Service) GetRedirect(ctx context.Context, redirectID string) (*Redirect, error) {
	if redirectID == "" {
		return nil, errors.New("redirectID must not be empty")
	}
	path := "/instant-site/redirects/" + url.PathEscape(redirectID)

	var result Redirect
	if err := s.main.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateRedirect creates a new 301 URL redirect.
//
// API: POST /instant-site/redirects
// Required scope: manage_instant_site
func (s *Service) CreateRedirect(ctx context.Context, redirect *Redirect) (*Redirect, error) {
	var result Redirect
	if err := s.main.Post(ctx, "/instant-site/redirects", redirect, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateRedirect updates an existing 301 URL redirect.
//
// API: PUT /instant-site/redirects/{redirectId}
// Required scope: manage_instant_site
func (s *Service) UpdateRedirect(ctx context.Context, redirectID string, redirect *Redirect) (*Redirect, error) {
	if redirectID == "" {
		return nil, errors.New("redirectID must not be empty")
	}
	path := "/instant-site/redirects/" + url.PathEscape(redirectID)

	var result Redirect
	if err := s.main.Put(ctx, path, redirect, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
