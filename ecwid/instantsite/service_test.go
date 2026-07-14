package instantsite_test

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/matthiasbruns/ecwid-go/ecwid/instantsite"
	"github.com/matthiasbruns/ecwid-go/ecwid/internal/api"
)

// newTestService builds a Service whose three requesters all point at the single
// test server, using distinct base-URL prefixes so the v3 (main), v1, and auth
// hosts don't collide on one mux:
//
//	main → /api/v3/12345/…
//	v1   → /api/v1/12345/…
//	auth → /instantsite/… (no store-ID segment)
func newTestService(t *testing.T, srv *httptest.Server) *instantsite.Service {
	t.Helper()
	main := api.NewHTTPClient(api.HTTPClientConfig{
		BaseURL: srv.URL + "/api/v3",
		StoreID: "12345",
		Token:   "secret_test",
	})
	v1 := api.NewHTTPClient(api.HTTPClientConfig{
		BaseURL: srv.URL + "/api/v1",
		StoreID: "12345",
		Token:   "is_test",
	})
	auth := api.NewHTTPClient(api.HTTPClientConfig{
		BaseURL: srv.URL + "/instantsite",
		StoreID: "",
		Token:   "",
	})
	return instantsite.NewService(main, v1, auth)
}

// ── Token ─────────────────────────────────────────────────────────────────

func TestToken(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/instantsite/oauth/token" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		q := r.URL.Query()
		if q.Get("grant_type") != "authorization_code" {
			t.Errorf("expected grant_type=authorization_code, got %s", q.Get("grant_type"))
		}
		if q.Get("site_id") != "12345" {
			t.Errorf("expected site_id=12345, got %s", q.Get("site_id"))
		}
		if q.Get("code") != "app_secret" {
			t.Errorf("expected code=app_secret, got %s", q.Get("code"))
		}
		// The auth requester carries no bearer token.
		if got := r.Header.Get("Authorization"); got != "" {
			t.Errorf("expected no Authorization header, got %q", got)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"accessToken":"tok_abc","tokenType":"bearer","expiresIn":86400}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Token(context.Background(), &instantsite.TokenRequest{SiteID: "12345", Code: "app_secret"})
	if err != nil {
		t.Fatal(err)
	}
	if result.AccessToken != "tok_abc" {
		t.Errorf("expected accessToken=tok_abc, got %s", result.AccessToken)
	}
	if result.ExpiresIn != 86400 {
		t.Errorf("expected expiresIn=86400, got %d", result.ExpiresIn)
	}
}

func TestToken_Guards(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		t.Fatal("should not reach server")
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	if _, err := svc.Token(context.Background(), nil); err == nil {
		t.Error("expected error for nil request")
	}
	if _, err := svc.Token(context.Background(), &instantsite.TokenRequest{Code: "x"}); err == nil {
		t.Error("expected error for empty SiteID")
	}
	if _, err := svc.Token(context.Background(), &instantsite.TokenRequest{SiteID: "12345"}); err == nil {
		t.Error("expected error for empty Code")
	}
}

// ── Profile ───────────────────────────────────────────────────────────────

func TestGetProfile(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/12345/profile" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer is_test" {
			t.Errorf("expected instant-site bearer token, got %q", got)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"siteId":"s1","storeName":"My Store","ecwidApiUrl":"https://vuega.ecwid.com/api/v1"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	p, err := svc.GetProfile(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if p.StoreName != "My Store" {
		t.Errorf("expected storeName=My Store, got %s", p.StoreName)
	}
	if p.EcwidAPIURL != "https://vuega.ecwid.com/api/v1" {
		t.Errorf("unexpected ecwidApiUrl: %s", p.EcwidAPIURL)
	}
}

func TestCreateProfile(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/api/v1/12345/profile" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"storeName":"New Store"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	p, err := svc.CreateProfile(context.Background(), &instantsite.Profile{StoreName: "New Store"})
	if err != nil {
		t.Fatal(err)
	}
	if p.StoreName != "New Store" {
		t.Errorf("expected storeName=New Store, got %s", p.StoreName)
	}
}

func TestUpdateProfile(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"storeName":"Updated"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	p, err := svc.UpdateProfile(context.Background(), &instantsite.Profile{StoreName: "Updated"})
	if err != nil {
		t.Fatal(err)
	}
	if p.StoreName != "Updated" {
		t.Errorf("expected storeName=Updated, got %s", p.StoreName)
	}
}

// ── Lifecycle ─────────────────────────────────────────────────────────────

func TestPublishDiscard(t *testing.T) {
	for _, tc := range []struct {
		name string
		call func(*instantsite.Service) error
		path string
	}{
		{"publish", func(s *instantsite.Service) error { return s.Publish(context.Background()) }, "/api/v1/12345/publish"},
		{"discard", func(s *instantsite.Service) error { return s.Discard(context.Background()) }, "/api/v1/12345/discard"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodPost {
					t.Errorf("expected POST, got %s", r.Method)
				}
				if r.URL.Path != tc.path {
					t.Errorf("unexpected path: %s", r.URL.Path)
				}
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{}`))
			}))
			defer srv.Close()

			if err := tc.call(newTestService(t, srv)); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestClone(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/12345/clone" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		body, _ := io.ReadAll(r.Body)
		if !strings.Contains(string(body), `"source":99`) {
			t.Errorf("expected source=99 in body, got %s", body)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	if err := svc.Clone(context.Background(), &instantsite.CloneRequest{Source: 99, Draft: true}); err != nil {
		t.Fatal(err)
	}
	if err := svc.Clone(context.Background(), nil); err == nil {
		t.Error("expected error for nil clone request")
	}
}

// ── Pages ─────────────────────────────────────────────────────────────────

func TestListPages(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/12345/page" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("published") != "true" {
			t.Errorf("expected published=true, got %s", r.URL.Query().Get("published"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"pages":[{"pageId":"p1","title":"Home","tileIds":["t1","t2"]}]}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.ListPages(context.Background(), true)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Pages) != 1 || result.Pages[0].Title != "Home" {
		t.Errorf("unexpected pages: %+v", result.Pages)
	}
	if len(result.Pages[0].TileIDs) != 2 {
		t.Errorf("expected 2 tileIds, got %d", len(result.Pages[0].TileIDs))
	}
}

func TestCreatePage(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"pageId":"p9"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.CreatePage(context.Background(), &instantsite.Page{Title: "New"})
	if err != nil {
		t.Fatal(err)
	}
	if result.PageID != "p9" {
		t.Errorf("expected pageId=p9, got %s", result.PageID)
	}
}

func TestUpdatePage(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/api/v1/12345/page/p1" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"pageId":"p1","title":"Renamed"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.UpdatePage(context.Background(), "p1", &instantsite.Page{Title: "Renamed"})
	if err != nil {
		t.Fatal(err)
	}
	if result.Title != "Renamed" {
		t.Errorf("expected title=Renamed, got %s", result.Title)
	}
}

func TestPage_EmptyIDGuards(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		t.Fatal("should not reach server")
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	if _, err := svc.UpdatePage(context.Background(), "", &instantsite.Page{}); err == nil {
		t.Error("expected error for empty pageID on update")
	}
	if _, err := svc.DeletePage(context.Background(), ""); err == nil {
		t.Error("expected error for empty pageID on delete")
	}
}

func TestDeletePage(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/api/v1/12345/page/p1" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"pageId":"p1"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.DeletePage(context.Background(), "p1")
	if err != nil {
		t.Fatal(err)
	}
	if result.PageID != "p1" {
		t.Errorf("expected pageId=p1, got %s", result.PageID)
	}
}

// ── Tiles ─────────────────────────────────────────────────────────────────

func TestListTiles(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/12345/tile" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		q := r.URL.Query()
		if q.Get("published") != "true" {
			t.Errorf("expected published=true, got %s", q.Get("published"))
		}
		if q.Get("pageId") != "p1" {
			t.Errorf("expected pageId=p1, got %s", q.Get("pageId"))
		}
		if q.Get("lang") != "en" {
			t.Errorf("expected lang=en, got %s", q.Get("lang"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"tiles":[{"id":"t1","type":"COVER","content":{"title":"Hi"}}]}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.ListTiles(context.Background(), &instantsite.TileListOptions{Published: true, PageID: "p1", Lang: "en"})
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Tiles) != 1 || result.Tiles[0].Type != "COVER" {
		t.Errorf("unexpected tiles: %+v", result.Tiles)
	}
	if result.Tiles[0].Content == nil {
		t.Error("expected opaque content to be preserved")
	}
}

func TestListTiles_NilOptions(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("published") != "false" {
			t.Errorf("expected published=false default, got %s", r.URL.Query().Get("published"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"tiles":[]}`))
	}))
	defer srv.Close()

	if _, err := newTestService(t, srv).ListTiles(context.Background(), nil); err != nil {
		t.Fatal(err)
	}
}

func TestGetTile(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/12345/tile/t1" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":"t1","type":"HEADER"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	tile, err := svc.GetTile(context.Background(), "t1")
	if err != nil {
		t.Fatal(err)
	}
	if tile.Type != "HEADER" {
		t.Errorf("expected type=HEADER, got %s", tile.Type)
	}
}

func TestCreateTile(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/api/v1/12345/tile" {
			t.Errorf("unexpected %s %s", r.Method, r.URL.Path)
		}
		body, _ := io.ReadAll(r.Body)
		if !strings.Contains(string(body), `"tileCategoryType":"COVER"`) {
			t.Errorf("expected tileCategoryType in body, got %s", body)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":"t9","type":"COVER"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	tile, err := svc.CreateTile(context.Background(), &instantsite.CreateTileRequest{
		TileCategoryType: "COVER",
		Tile:             &instantsite.TileCreateBody{Type: "COVER"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if tile.ID != "t9" {
		t.Errorf("expected id=t9, got %s", tile.ID)
	}
}

func TestUpdateTile(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut || r.URL.Path != "/api/v1/12345/tile/t1" {
			t.Errorf("unexpected %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":"t1","tileName":"Renamed"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	tile, err := svc.UpdateTile(context.Background(), "t1", &instantsite.TileUpdate{TileName: "Renamed", Visibility: boolPtr(false)})
	if err != nil {
		t.Fatal(err)
	}
	if tile.TileName != "Renamed" {
		t.Errorf("expected tileName=Renamed, got %s", tile.TileName)
	}
}

func TestUpdateTiles(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut || r.URL.Path != "/api/v1/12345/tile" {
			t.Errorf("unexpected %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"tiles":[{"id":"t1"},{"id":"t2"}]}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.UpdateTiles(context.Background(), &instantsite.TileBulkUpdate{
		Tiles: []instantsite.TileUpdate{{ID: "t1"}, {ID: "t2"}},
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Tiles) != 2 {
		t.Errorf("expected 2 tiles, got %d", len(result.Tiles))
	}
}

func TestDeleteTile(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete || r.URL.Path != "/api/v1/12345/tile/t1" {
			t.Errorf("unexpected %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":"t1"}`))
	}))
	defer srv.Close()

	if _, err := newTestService(t, srv).DeleteTile(context.Background(), "t1"); err != nil {
		t.Fatal(err)
	}
}

func TestTile_EmptyIDGuards(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		t.Fatal("should not reach server")
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	if _, err := svc.GetTile(context.Background(), ""); err == nil {
		t.Error("expected error for empty tileID on get")
	}
	if _, err := svc.UpdateTile(context.Background(), "", &instantsite.TileUpdate{}); err == nil {
		t.Error("expected error for empty tileID on update")
	}
	if _, err := svc.DeleteTile(context.Background(), ""); err == nil {
		t.Error("expected error for empty tileID on delete")
	}
	if _, err := svc.TileConfig(context.Background(), ""); err == nil {
		t.Error("expected error for empty configType")
	}
	if _, err := svc.ReserveTileImage(context.Background(), ""); err == nil {
		t.Error("expected error for empty tileID on reserve image")
	}
	if _, err := svc.GetImage(context.Background(), ""); err == nil {
		t.Error("expected error for empty imageID")
	}
}

func TestTileShowcaseAndConfig(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/api/v1/12345/tile/showcase":
			_, _ = w.Write([]byte(`{"categories":[{"type":"COVER","items":[{"id":"i1"}]}]}`))
		case "/api/v1/12345/tile/config/COVER":
			_, _ = w.Write([]byte(`{"type":"COVER","config":{"layoutConfigList":[]}}`))
		default:
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	sc, err := svc.TileShowcase(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(sc.Categories) != 1 || sc.Categories[0].Items[0].ID != "i1" {
		t.Errorf("unexpected showcase: %+v", sc.Categories)
	}
	cfg, err := svc.TileConfig(context.Background(), "COVER")
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Type != "COVER" || cfg.Config == nil {
		t.Errorf("unexpected config: %+v", cfg)
	}
}

// ── Tile images ───────────────────────────────────────────────────────────

func TestTileImages(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.URL.Path == "/api/v1/12345/tile/t1/image" && r.Method == http.MethodPost:
			_, _ = w.Write([]byte(`{"url":"https://upload","id":"img1"}`))
		case r.URL.Path == "/api/v1/12345/image/img1":
			_, _ = w.Write([]byte(`{"bucket":"eu-fra","set":[{"url":"https://cdn/x.jpg","width":800,"height":600}]}`))
		case r.URL.Path == "/api/v1/12345/image/bucket":
			_, _ = w.Write([]byte(`{"urls":{"eu-fra":"https://d2gt4h1eeousrn.cloudfront.net"}}`))
		default:
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	res, err := svc.ReserveTileImage(context.Background(), "t1")
	if err != nil {
		t.Fatal(err)
	}
	if res.ID != "img1" {
		t.Errorf("expected id=img1, got %s", res.ID)
	}
	img, err := svc.GetImage(context.Background(), "img1")
	if err != nil {
		t.Fatal(err)
	}
	if len(img.Set) != 1 || img.Set[0].Width != 800 {
		t.Errorf("unexpected image renditions: %+v", img.Set)
	}
	buckets, err := svc.ImageBuckets(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if buckets.URLs["eu-fra"] == "" {
		t.Errorf("expected eu-fra bucket URL, got %+v", buckets.URLs)
	}
}

// ── Themes ────────────────────────────────────────────────────────────────

func TestThemes(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.URL.Path == "/api/v1/12345/themes" && r.Method == http.MethodGet:
			_, _ = w.Write([]byte(`{"themes":[{"themeId":"th1","colors":{"colorA":"#fff"}}]}`))
		case r.URL.Path == "/api/v1/12345/themes" && r.Method == http.MethodPost:
			body, _ := io.ReadAll(r.Body)
			if !strings.Contains(string(body), `"colors":{"colorA":"#000"}`) {
				t.Errorf("expected colors envelope in body, got %s", body)
			}
			_, _ = w.Write([]byte(`{"themeId":"th2","colors":{"colorA":"#000"}}`))
		case r.URL.Path == "/api/v1/12345/themes/th2" && r.Method == http.MethodPut:
			_, _ = w.Write([]byte(`{"themeId":"th2","colors":{"colorA":"#111"}}`))
		case r.URL.Path == "/api/v1/12345/themes/th2" && r.Method == http.MethodDelete:
			_, _ = w.Write([]byte(`{"themeId":"th2","colors":{"colorA":"#111"}}`))
		default:
			t.Errorf("unexpected %s %s", r.Method, r.URL.Path)
		}
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	list, err := svc.ListThemes(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(list.Themes) != 1 || list.Themes[0].ThemeID != "th1" {
		t.Errorf("unexpected themes: %+v", list.Themes)
	}
	created, err := svc.CreateTheme(context.Background(), &instantsite.ThemeColors{ColorA: "#000"})
	if err != nil {
		t.Fatal(err)
	}
	if created.ThemeID != "th2" {
		t.Errorf("expected themeId=th2, got %s", created.ThemeID)
	}
	if _, err := svc.UpdateTheme(context.Background(), "th2", &instantsite.ThemeColors{ColorA: "#111"}); err != nil {
		t.Fatal(err)
	}
	if _, err := svc.DeleteTheme(context.Background(), "th2"); err != nil {
		t.Fatal(err)
	}
}

func TestTheme_EmptyIDGuards(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		t.Fatal("should not reach server")
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	if _, err := svc.UpdateTheme(context.Background(), "", &instantsite.ThemeColors{}); err == nil {
		t.Error("expected error for empty themeID on update")
	}
	if _, err := svc.DeleteTheme(context.Background(), ""); err == nil {
		t.Error("expected error for empty themeID on delete")
	}
}

func TestCurrentTheme(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case http.MethodGet:
			if r.URL.Path != "/api/v1/12345/current_theme" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}
			_, _ = w.Write([]byte(`{"colors":{"colorA":"#abc"}}`))
		case http.MethodPut:
			_, _ = w.Write([]byte(`{"themeId":"cur","colors":{"colorA":"#abc"}}`))
		default:
			t.Errorf("unexpected method: %s", r.Method)
		}
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	cur, err := svc.CurrentTheme(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if cur.Colors.ColorA != "#abc" {
		t.Errorf("expected colorA=#abc, got %s", cur.Colors.ColorA)
	}
	updated, err := svc.UpdateCurrentTheme(context.Background(), &instantsite.ThemeColors{ColorA: "#abc"})
	if err != nil {
		t.Fatal(err)
	}
	if updated.ThemeID != "cur" {
		t.Errorf("expected themeId=cur, got %s", updated.ThemeID)
	}
}

// ── Text labels ───────────────────────────────────────────────────────────

func TestTextLabels(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/12345/translation" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"editorTranslations":{"k":"v"},"languageTranslations":{"en":{"Language.en":"English"}}}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	labels, err := svc.TextLabels(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if labels.EditorTranslations["k"] != "v" {
		t.Errorf("unexpected editorTranslations: %+v", labels.EditorTranslations)
	}
	if labels.LanguageTranslations["en"]["Language.en"] != "English" {
		t.Errorf("unexpected languageTranslations: %+v", labels.LanguageTranslations)
	}
}

// ── Redirects (v3, main requester) ────────────────────────────────────────

func TestSearchRedirects(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/instant-site/redirects" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer secret_test" {
			t.Errorf("expected main bearer token, got %q", got)
		}
		if r.URL.Query().Get("keyword") != "old" {
			t.Errorf("expected keyword=old, got %s", r.URL.Query().Get("keyword"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"total":1,"count":1,"offset":0,"limit":10,"items":[{"id":"r1","fromUrl":"/old","toUrl":"/new"}]}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.SearchRedirects(context.Background(), &instantsite.RedirectSearchOptions{Keyword: "old", Limit: 10})
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != 1 || result.Items[0].FromURL != "/old" {
		t.Errorf("unexpected redirects: %+v", result)
	}
}

func TestRedirectCRUD(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.URL.Path == "/api/v3/12345/instant-site/redirects/r1" && r.Method == http.MethodGet:
			_, _ = w.Write([]byte(`{"id":"r1","fromUrl":"/old","toUrl":"/new"}`))
		case r.URL.Path == "/api/v3/12345/instant-site/redirects" && r.Method == http.MethodPost:
			_, _ = w.Write([]byte(`{"id":"r2","fromUrl":"/a","toUrl":"/b"}`))
		case r.URL.Path == "/api/v3/12345/instant-site/redirects/r2" && r.Method == http.MethodPut:
			_, _ = w.Write([]byte(`{"id":"r2","fromUrl":"/a","toUrl":"/c"}`))
		default:
			t.Errorf("unexpected %s %s", r.Method, r.URL.Path)
		}
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	got, err := svc.GetRedirect(context.Background(), "r1")
	if err != nil {
		t.Fatal(err)
	}
	if got.ToURL != "/new" {
		t.Errorf("expected toUrl=/new, got %s", got.ToURL)
	}
	created, err := svc.CreateRedirect(context.Background(), &instantsite.Redirect{FromURL: "/a", ToURL: "/b"})
	if err != nil {
		t.Fatal(err)
	}
	if created.ID != "r2" {
		t.Errorf("expected id=r2, got %s", created.ID)
	}
	updated, err := svc.UpdateRedirect(context.Background(), "r2", &instantsite.Redirect{FromURL: "/a", ToURL: "/c"})
	if err != nil {
		t.Fatal(err)
	}
	if updated.ToURL != "/c" {
		t.Errorf("expected toUrl=/c, got %s", updated.ToURL)
	}
}

func TestRedirect_EmptyIDGuards(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		t.Fatal("should not reach server")
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	if _, err := svc.GetRedirect(context.Background(), ""); err == nil {
		t.Error("expected error for empty redirectID on get")
	}
	if _, err := svc.UpdateRedirect(context.Background(), "", &instantsite.Redirect{}); err == nil {
		t.Error("expected error for empty redirectID on update")
	}
}

// ── Error handling ────────────────────────────────────────────────────────

func TestSearchRedirects_Error(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(`{"errorMessage":"forbidden","errorCode":"403"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	_, err := svc.SearchRedirects(context.Background(), nil)
	if err == nil {
		t.Fatal("expected error")
	}
	var apiErr *api.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *api.APIError, got %T", err)
	}
	if apiErr.StatusCode != http.StatusForbidden {
		t.Errorf("expected 403, got %d", apiErr.StatusCode)
	}
}

// ensure encoding/json import is used for a compile-time sanity check on a
// RawMessage round-trip of opaque tile content.
func TestTileContentRoundTrip(t *testing.T) {
	raw := json.RawMessage(`{"title":"Hi"}`)
	tile := instantsite.Tile{Content: &raw}
	data, err := json.Marshal(tile)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), `"content":{"title":"Hi"}`) {
		t.Errorf("expected content preserved, got %s", data)
	}
}

func boolPtr(v bool) *bool { return &v }
