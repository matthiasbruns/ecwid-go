package reports_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matthiasbruns/ecwid-go/ecwid/internal/api"
	"github.com/matthiasbruns/ecwid-go/ecwid/reports"
)

func newTestService(t *testing.T, srv *httptest.Server) *reports.Service {
	t.Helper()
	requester := api.NewHTTPClient(api.HTTPClientConfig{
		BaseURL: srv.URL + "/api/v3",
		StoreID: "12345",
		Token:   "secret_test",
	})
	return reports.NewService(requester)
}

func TestGetReport(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/reports/allOrders" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("startedFrom") != "1591646400" {
			t.Errorf("expected startedFrom=1591646400, got %s", r.URL.Query().Get("startedFrom"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"reportType":"allOrders","firstDayOfWeek":"MONDAY","aggregatedData":[{"metricName":"orders","value":42}],"dataset":[{"timestamp":1591646400}]}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	report, err := svc.GetReport(context.Background(), "allOrders", &reports.ReportOptions{
		StartedFrom:    1591646400,
		TimeScaleValue: "day",
	})
	if err != nil {
		t.Fatal(err)
	}
	if report.ReportType != "allOrders" {
		t.Errorf("expected allOrders, got %s", report.ReportType)
	}
	if report.AggregatedData == nil {
		t.Error("expected aggregatedData to be non-nil")
	}
}

func TestGetReportNoOpts(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.RawQuery != "" {
			t.Errorf("expected no query params, got %s", r.URL.RawQuery)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"reportType":"allTraffic"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	report, err := svc.GetReport(context.Background(), "allTraffic", nil)
	if err != nil {
		t.Fatal(err)
	}
	if report.ReportType != "allTraffic" {
		t.Errorf("expected allTraffic, got %s", report.ReportType)
	}
}

func TestLatestStats(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/latest-stats" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("productCountRequired") != "true" {
			t.Error("expected productCountRequired=true")
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"productsUpdated":"2026-03-01 10:00:00 +0100","ordersUpdated":"2026-03-06 14:30:00 +0100","productCountRequired":156,"categoryCountRequired":12}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	stats, err := svc.LatestStats(context.Background(), &reports.LatestStatsOptions{
		ProductCountRequired:  true,
		CategoryCountRequired: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	if stats.ProductCountRequired != 156 {
		t.Errorf("expected 156, got %d", stats.ProductCountRequired)
	}
}

func TestLatestStatsNoOpts(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.RawQuery != "" {
			t.Errorf("expected no query params, got %s", r.URL.RawQuery)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"productsUpdated":"2026-03-01 10:00:00 +0100"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	stats, err := svc.LatestStats(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if stats.ProductsUpdated == "" {
		t.Error("expected productsUpdated to be non-empty")
	}
}

func TestGetReport_Error(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(`{"errorMessage":"access denied","errorCode":403}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	_, err := svc.GetReport(context.Background(), "allOrders", nil)
	if err == nil {
		t.Fatal("expected error")
	}
}
