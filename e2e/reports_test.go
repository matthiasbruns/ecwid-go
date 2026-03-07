package e2e

import (
	"context"
	"testing"

	"github.com/matthiasbruns/ecwid-go/ecwid/reports"
)

func TestReports_GetReport(t *testing.T) {
	ctx := context.Background()

	report, err := testClient.Reports.GetReport(ctx, "allOrders", &reports.ReportOptions{
		TimeScaleValue: "month",
	})
	if err != nil {
		t.Fatalf("GetReport(allOrders): %v", err)
	}
	if report.ReportType != "allOrders" {
		t.Errorf("expected reportType=allOrders, got %s", report.ReportType)
	}
}

func TestReports_LatestStats(t *testing.T) {
	ctx := context.Background()

	stats, err := testClient.Reports.LatestStats(ctx, &reports.LatestStatsOptions{
		ProductCountRequired:  true,
		CategoryCountRequired: true,
	})
	if err != nil {
		t.Fatalf("LatestStats: %v", err)
	}
	// Just verify we got a response without error.
	// Timestamps may be empty for new stores.
	_ = stats
}
