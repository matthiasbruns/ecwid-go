package e2e

import (
	"os"
	"testing"
)

// requireInstantSiteToken skips a test that hits the Instant Site v1 host when no
// Instant Site token is configured (the v1 endpoints use a separate 24h token).
func requireInstantSiteToken(t *testing.T) {
	t.Helper()
	if os.Getenv("ECWID_INSTANT_SITE_TOKEN") == "" {
		t.Skip("skipping: ECWID_INSTANT_SITE_TOKEN required for Instant Site v1 endpoints")
	}
}

// Redirects live on the main API host and use the main token, so they only need
// the manage_instant_site scope (skipIfForbidden covers a missing scope).
func TestInstantSite_SearchRedirects(t *testing.T) {
	requireClient(t)
	ctx := testContext(t)

	result, err := testClient.InstantSite.SearchRedirects(ctx, nil)
	if err != nil {
		skipIfForbidden(t, err)
		t.Fatalf("InstantSite.SearchRedirects: %v", err)
	}
	t.Logf("found %d redirects", result.Total)
}

func TestInstantSite_GetProfile(t *testing.T) {
	requireClient(t)
	requireInstantSiteToken(t)
	ctx := testContext(t)

	profile, err := testClient.InstantSite.GetProfile(ctx)
	if err != nil {
		skipIfForbidden(t, err)
		t.Fatalf("InstantSite.GetProfile: %v", err)
	}
	if profile.SiteURL == "" {
		t.Log("profile returned without a siteUrl")
	}
}

func TestInstantSite_ListPages(t *testing.T) {
	requireClient(t)
	requireInstantSiteToken(t)
	ctx := testContext(t)

	result, err := testClient.InstantSite.ListPages(ctx, true)
	if err != nil {
		skipIfForbidden(t, err)
		t.Fatalf("InstantSite.ListPages: %v", err)
	}
	t.Logf("found %d published pages", len(result.Pages))
}

func TestInstantSite_ListThemes(t *testing.T) {
	requireClient(t)
	requireInstantSiteToken(t)
	ctx := testContext(t)

	result, err := testClient.InstantSite.ListThemes(ctx)
	if err != nil {
		skipIfForbidden(t, err)
		t.Fatalf("InstantSite.ListThemes: %v", err)
	}
	t.Logf("found %d themes", len(result.Themes))
}
