package staff_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matthiasbruns/ecwid-go/ecwid/internal/api"
	"github.com/matthiasbruns/ecwid-go/ecwid/staff"
)

func newTestService(t *testing.T, srv *httptest.Server) *staff.Service {
	t.Helper()
	return staff.NewService(api.NewHTTPClient(api.HTTPClientConfig{
		BaseURL: srv.URL + "/api/v3",
		StoreID: "12345",
		Token:   "secret_test",
	}))
}

func TestList(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/staff" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"staffList":[{"id":"p27632593","name":"John Doe","email":"john@example.com","staffScopes":["REPORT_ACCESS","SALES_MANAGEMENT"],"inviteAccepted":true}]}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(result.StaffList) != 1 {
		t.Fatalf("expected 1 staff, got %d", len(result.StaffList))
	}
	if result.StaffList[0].ID != "p27632593" {
		t.Errorf("expected id=p27632593, got %s", result.StaffList[0].ID)
	}
	if result.StaffList[0].Name != "John Doe" {
		t.Errorf("expected name=John Doe, got %s", result.StaffList[0].Name)
	}
	if len(result.StaffList[0].StaffScopes) != 2 {
		t.Errorf("expected 2 scopes, got %d", len(result.StaffList[0].StaffScopes))
	}
}

func TestGet(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/staff/p3855016" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"email":"ec.apps@lightspeedhq.com","staffScopes":["SALES_MANAGEMENT","CATALOG_MANAGEMENT"]}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Get(context.Background(), "p3855016")
	if err != nil {
		t.Fatal(err)
	}
	if result.Email != "ec.apps@lightspeedhq.com" {
		t.Errorf("expected email=ec.apps@lightspeedhq.com, got %s", result.Email)
	}
	if len(result.StaffScopes) != 2 {
		t.Errorf("expected 2 scopes, got %d", len(result.StaffScopes))
	}
}

func TestGet_EmptyID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		t.Fatal("should not reach server")
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	_, err := svc.Get(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty staffAccountID")
	}
}

func TestCreate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"success":true}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Create(context.Background(), &staff.CreateRequest{
		Email:       "new@example.com",
		StaffScopes: []string{"REPORT_ACCESS"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if !result.Success {
		t.Error("expected success=true")
	}
}

func TestUpdate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/staff/p123" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"updateCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Update(context.Background(), "p123", &staff.UpdateRequest{
		StaffScopes: []string{"REPORT_ACCESS", "SALES_MANAGEMENT"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.UpdateCount != 1 {
		t.Errorf("expected updateCount=1, got %d", result.UpdateCount)
	}
}

func TestDelete(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"deleteCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Delete(context.Background(), "p123")
	if err != nil {
		t.Fatal(err)
	}
	if result.DeleteCount != 1 {
		t.Errorf("expected deleteCount=1, got %d", result.DeleteCount)
	}
}

func TestList_Error(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(`{"errorMessage":"access denied","errorCode":"403"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	_, err := svc.List(context.Background())
	if err == nil {
		t.Fatal("expected error")
	}
	var apiErr *api.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *api.APIError, got %T", err)
	}
}
