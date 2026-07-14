// Package instantsite provides access to the Ecwid Instant Site API.
//
// The Instant Site API spans two hosts and two tokens:
//
//   - Redirects (301 URL redirects) live on the main API host
//     (app.ecwid.com/api/v3/{storeId}/instant-site/redirects) and use the main
//     store access token — the same transport as every other domain service.
//   - Everything else (profile, pages, tiles, tile images, themes, text labels,
//     and the publish/discard/clone lifecycle) lives on the Instant Site v1 host
//     (vuega.ecwid.com/api/v1/{storeId}/…) and uses a separate 24h token obtained
//     via the token-exchange endpoint (app.ecwid.com/instantsite/oauth/token).
//
// Docs: https://docs.ecwid.com/api-reference/rest-api/instant-site
package instantsite

import "encoding/json"

// ── Token exchange ──────────────────────────────────────────────────────────

// GrantType values for the Instant Site token exchange.
const (
	// GrantTypeAuthorizationCode grants full Instant Site API access.
	GrantTypeAuthorizationCode = "authorization_code"
	// GrantTypeDraftCode grants editor/preview-only access.
	GrantTypeDraftCode = "draft_code"
)

// TokenRequest holds the parameters for exchanging an app secret token for an
// Instant Site access token. All fields are sent as query parameters.
type TokenRequest struct {
	// GrantType is "authorization_code" (full access) or "draft_code"
	// (editor/preview only). Defaults to authorization_code when empty.
	GrantType string
	// SiteID is the Ecwid store ID.
	SiteID string
	// Code is the app secret token carrying the manage_instant_site scope.
	Code string
}

// TokenResult is the response from the Instant Site token-exchange endpoint.
type TokenResult struct {
	AccessToken string `json:"accessToken"`
	TokenType   string `json:"tokenType"`
	ExpiresIn   int    `json:"expiresIn"`
}

// ── Profile ─────────────────────────────────────────────────────────────────

// Profile is the Instant Site profile. GET returns the full object; create and
// update accept the subset of mutable fields (locale, enabledLanguages,
// storeName, countryCode, postalCode, email, dateFormat, timeFormat,
// timezoneOffsetInMinutes).
type Profile struct {
	SiteID                                 string               `json:"siteId,omitempty"`
	Locale                                 string               `json:"locale,omitempty"`
	EnabledLanguages                       []string             `json:"enabledLanguages,omitempty"`
	StoreName                              string               `json:"storeName,omitempty"`
	Tracking                               *ProfileTracking     `json:"tracking,omitempty"`
	CountryCode                            string               `json:"countryCode,omitempty"`
	PostalCode                             string               `json:"postalCode,omitempty"`
	Email                                  string               `json:"email,omitempty"`
	DateFormat                             string               `json:"dateFormat,omitempty"`
	TimeFormat                             string               `json:"timeFormat,omitempty"`
	TimezoneOffsetInMinutes                int                  `json:"timezoneOffsetInMinutes,omitempty"`
	StoreClosed                            *bool                `json:"storeClosed,omitempty"`
	StoreSuspended                         *bool                `json:"storeSuspended,omitempty"`
	IsTemplateSite                         *bool                `json:"isTemplateSite,omitempty"`
	SiteURL                                string               `json:"siteUrl,omitempty"`
	Subscription                           *ProfileSubscription `json:"subscription,omitempty"`
	LatestPublishTimestamp                 int64                `json:"latestPublishTimestamp,omitempty"`
	Onboarding                             *json.RawMessage     `json:"onboarding,omitempty"`
	Vertical                               string               `json:"vertical,omitempty"`
	PreviewTemplateInsideEditor            *bool                `json:"previewTemplateInsideEditor,omitempty"`
	FeatureFlags                           *json.RawMessage     `json:"featureFlags,omitempty"`
	IsDraftChanged                         *bool                `json:"isDraftChanged,omitempty"`
	HideEcwidLinks                         *bool                `json:"hideEcwidLinks,omitempty"`
	SelectedSiteTemplateID                 string               `json:"selectedSiteTemplateId,omitempty"`
	EcwidAPIURL                            string               `json:"ecwidApiUrl,omitempty"`
	StorefrontFiltersEnabled               *bool                `json:"storefrontFiltersEnabled,omitempty"`
	StorefrontProductReviewsFeatureEnabled *bool                `json:"storefrontProductReviewsFeatureEnabled,omitempty"`
}

// ProfileTracking holds analytics tracking configuration on the profile.
type ProfileTracking struct {
	GoogleUniversalAnalyticsID string `json:"googleUniversalAnalyticsId,omitempty"`
	HeapEnabled                *bool  `json:"heapEnabled,omitempty"`
}

// ProfileSubscription describes the store's plan/subscription as seen by the
// Instant Site editor.
type ProfileSubscription struct {
	ChannelID                                 string `json:"channelId,omitempty"`
	ChannelType                               string `json:"channelType,omitempty"`
	PlanName                                  string `json:"planName,omitempty"`
	PlanPeriod                                string `json:"planPeriod,omitempty"`
	IsPaid                                    *bool  `json:"isPaid,omitempty"`
	IsAllowNewCookieBanner                    *bool  `json:"isAllowNewCookieBanner,omitempty"`
	MaxPageNumber                             int    `json:"maxPageNumber,omitempty"`
	IsMultilingualStoreFeatureEnabled         *bool  `json:"isMultilingualStoreFeatureEnabled,omitempty"`
	IsAdvancedDiscountsFeatureAvailable       *bool  `json:"isAdvancedDiscountsFeatureAvailable,omitempty"`
	IsBasicEcommerceFeatureEnabled            *bool  `json:"isBasicEcommerceFeatureEnabled,omitempty"`
	IsRichTextEditorEnabled                   *bool  `json:"isRichTextEditorEnabled,omitempty"`
	IsTemplateMarketFeatureEnabled            *bool  `json:"isTemplateMarketFeatureEnabled,omitempty"`
	IsAccessToControlPanel                    *bool  `json:"isAccessToControlPanel,omitempty"`
	IsStorefrontAgeConfirmationFeatureEnabled *bool  `json:"isStorefrontAgeConfirmationFeatureEnabled,omitempty"`
}

// ── Lifecycle ───────────────────────────────────────────────────────────────

// CloneRequest is the body for cloning an Instant Site from another store.
type CloneRequest struct {
	// Source is the store ID to copy settings from.
	Source int `json:"source"`
	// Draft, when false, publishes immediately; when true, saves as a draft.
	Draft bool `json:"draft"`
}

// ── Pages ───────────────────────────────────────────────────────────────────

// Page represents an Instant Site page.
type Page struct {
	PageID                 string `json:"pageId,omitempty"`
	Title                  string `json:"title,omitempty"`
	URLPath                string `json:"urlPath,omitempty"`
	Visible                *bool  `json:"visible,omitempty"`
	VisibleHeader          *bool  `json:"visibleHeader,omitempty"`
	VisibleFooter          *bool  `json:"visibleFooter,omitempty"`
	VisibleAnnouncementBar *bool  `json:"visibleAnnouncementBar,omitempty"`
	SEOTitle               string `json:"seoTitle,omitempty"`
	SEODescription         string `json:"seoDescription,omitempty"`
	ShareImage             string `json:"shareImage,omitempty"`
	Indexed                *bool  `json:"indexed,omitempty"`
	IsAvailableToEdit      *bool  `json:"isAvailableToEdit,omitempty"`
	// TileIDs is the ordered (top→bottom) list of tiles enabled on the page.
	TileIDs []string `json:"tileIds,omitempty"`
}

// PageList is the response envelope for the list-pages endpoint.
type PageList struct {
	Pages []Page `json:"pages"`
}

// CreatePageResult is the response from creating a page.
type CreatePageResult struct {
	PageID string `json:"pageId"`
}

// ── Tiles ───────────────────────────────────────────────────────────────────

// Tile represents an Instant Site tile (a section/block of a page). The
// content, externalContent, design, and featuresEnabled fields are opaque and
// vary by tile Type; their shape is described by the tile config endpoint
// (see [Service.TileConfig]).
type Tile struct {
	ID              string           `json:"id,omitempty"`
	Type            string           `json:"type,omitempty"`
	Role            string           `json:"role,omitempty"`
	TileName        string           `json:"tileName,omitempty"`
	SourceID        string           `json:"sourceId,omitempty"`
	Content         *json.RawMessage `json:"content,omitempty"`
	ExternalContent *json.RawMessage `json:"externalContent,omitempty"`
	Design          *json.RawMessage `json:"design,omitempty"`
	Visibility      *bool            `json:"visibility,omitempty"`
	Order           int              `json:"order,omitempty"`
	HasChanges      *bool            `json:"hasChanges,omitempty"`
	FeaturesEnabled *json.RawMessage `json:"featuresEnabled,omitempty"`
}

// TileList is the response envelope for the list-tiles endpoint.
type TileList struct {
	Tiles []Tile `json:"tiles"`
}

// TileListOptions holds query parameters for listing tiles.
type TileListOptions struct {
	// Published selects published (true) or draft (false) tiles. Required.
	Published bool
	// PageID optionally filters tiles to a single page.
	PageID string
	// Lang optionally requests translated text (2-letter language code).
	Lang string
}

// TileUpdate is the mutable subset of a tile used by update operations.
type TileUpdate struct {
	ID         string           `json:"id,omitempty"`
	TileName   string           `json:"tileName,omitempty"`
	Content    *json.RawMessage `json:"content,omitempty"`
	Design     *json.RawMessage `json:"design,omitempty"`
	Visibility *bool            `json:"visibility,omitempty"`
	Order      int              `json:"order,omitempty"`
	HasChanges *bool            `json:"hasChanges,omitempty"`
}

// TileBulkUpdate is the body for updating the whole tiles list at once.
type TileBulkUpdate struct {
	Tiles []TileUpdate `json:"tiles"`
}

// TileCreateBody is the tile payload nested inside a CreateTileRequest.
type TileCreateBody struct {
	Type       string           `json:"type,omitempty"`
	TileName   string           `json:"tileName,omitempty"`
	Content    *json.RawMessage `json:"content,omitempty"`
	Design     *json.RawMessage `json:"design,omitempty"`
	Visibility *bool            `json:"visibility,omitempty"`
	Order      int              `json:"order,omitempty"`
	HasChanges *bool            `json:"hasChanges,omitempty"`
}

// CreateTileRequest is the body for creating a tile.
type CreateTileRequest struct {
	TileShowcaseItemID string          `json:"tileShowcaseItemId,omitempty"`
	TileCategoryType   string          `json:"tileCategoryType,omitempty"`
	TileOrder          int             `json:"tileOrder,omitempty"`
	UseTileShowcase    *bool           `json:"useTileShowcase,omitempty"`
	PageID             string          `json:"pageId,omitempty"`
	Tile               *TileCreateBody `json:"tile,omitempty"`
}

// TileShowcaseResult is the response for the available tile templates.
type TileShowcaseResult struct {
	Categories []TileShowcaseCategory `json:"categories"`
}

// TileShowcaseCategory groups showcase items by tile category type.
type TileShowcaseCategory struct {
	Type  string             `json:"type"`
	Items []TileShowcaseItem `json:"items"`
}

// TileShowcaseItem is a single available tile template.
type TileShowcaseItem struct {
	ID              string `json:"id"`
	PreviewImageURL string `json:"previewImageUrl,omitempty"`
	PreviewHeight   int    `json:"previewHeight,omitempty"`
	PreviewWidth    int    `json:"previewWidth,omitempty"`
	FeatureProperty string `json:"featureProperty,omitempty"`
	IsDeprecated    *bool  `json:"isDeprecated,omitempty"`
}

// TileConfigResult is the response for a tile config by category type. The
// config payload is opaque (layoutConfigList with per-type editor configs).
type TileConfigResult struct {
	Type   string           `json:"type"`
	Config *json.RawMessage `json:"config,omitempty"`
}

// ── Tile images ─────────────────────────────────────────────────────────────

// ReserveImageResult is the response from reserving a tile image upload. Clients
// then POST the image binary as multipart/form-data to URL.
type ReserveImageResult struct {
	URL string `json:"url"`
	ID  string `json:"id"`
}

// ImageResult is the response describing an uploaded tile image.
type ImageResult struct {
	Bucket     string           `json:"bucket,omitempty"`
	Set        []ImageRendition `json:"set,omitempty"`
	BorderInfo *json.RawMessage `json:"borderInfo,omitempty"`
}

// ImageRendition is a single rendered size of an uploaded image.
type ImageRendition struct {
	URL    string `json:"url"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

// ImageBucketsResult maps image bucket IDs to their CDN base URLs.
type ImageBucketsResult struct {
	URLs map[string]string `json:"urls"`
}

// ── Themes ──────────────────────────────────────────────────────────────────

// ThemeColors is the six-color palette of an Instant Site theme.
type ThemeColors struct {
	ColorA string `json:"colorA,omitempty"`
	ColorB string `json:"colorB,omitempty"`
	ColorC string `json:"colorC,omitempty"`
	ColorD string `json:"colorD,omitempty"`
	ColorE string `json:"colorE,omitempty"`
	ColorF string `json:"colorF,omitempty"`
}

// Theme is a named color palette.
type Theme struct {
	ThemeID string      `json:"themeId,omitempty"`
	Colors  ThemeColors `json:"colors"`
}

// ThemeList is the response envelope for the list-themes endpoint.
type ThemeList struct {
	Themes []Theme `json:"themes"`
}

// CurrentTheme is the response from the current-theme endpoint.
type CurrentTheme struct {
	Colors ThemeColors `json:"colors"`
}

// ── Text labels ─────────────────────────────────────────────────────────────

// TextLabels is the response for the Instant Site text labels (i18n) endpoint.
type TextLabels struct {
	EditorTranslations   map[string]string            `json:"editorTranslations,omitempty"`
	WebsiteTranslations  map[string]string            `json:"websiteTranslations,omitempty"`
	LanguageTranslations map[string]map[string]string `json:"languageTranslations,omitempty"`
}

// ── Redirects (v3) ──────────────────────────────────────────────────────────

// Redirect is a single 301 URL redirect.
type Redirect struct {
	ID      string `json:"id,omitempty"`
	FromURL string `json:"fromUrl"`
	ToURL   string `json:"toUrl"`
}

// RedirectSearchResult is the paginated response from the redirects search API.
type RedirectSearchResult struct {
	Total  int        `json:"total"`
	Count  int        `json:"count"`
	Offset int        `json:"offset"`
	Limit  int        `json:"limit"`
	Items  []Redirect `json:"items"`
}

// RedirectSearchOptions holds query parameters for searching redirects.
type RedirectSearchOptions struct {
	Keyword string
	Offset  int
	Limit   int
}
