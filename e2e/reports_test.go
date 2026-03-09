package e2e

import (
	"testing"

	"github.com/matthiasbruns/ecwid-go/ecwid/reports"
)

func TestReports_GetReport(t *testing.T) {
	ctx := testContext(t)

	report, err := testClient.Reports.GetReport(ctx, "allOrders", &reports.ReportOptions{
		TimeScaleValue: "month",
	})
	if err != nil {
		skipIfForbidden(t, err)
		t.Fatalf("GetReport(allOrders): %v", err)
	}
	if report.ReportType != "allOrders" {
		t.Errorf("expected reportType=allOrders, got %s", report.ReportType)
	}
}

func TestReports_LatestStats(t *testing.T) {
	ctx := testContext(t)

	stats, err := testClient.Reports.LatestStats(ctx, &reports.LatestStatsOptions{
		ProductCountRequired:  true,
		CategoryCountRequired: true,
	})
	if err != nil {
		skipIfForbidden(t, err)
		t.Fatalf("LatestStats: %v", err)
	}
	_ = stats
}
