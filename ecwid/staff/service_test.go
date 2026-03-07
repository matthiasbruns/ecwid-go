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
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/staff" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"total":1,"items":[{"id":100,"email":"admin@example.com","firstName":"Alice","lastName":"Admin","role":"ADMIN"}]}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != 1 {
		t.Errorf("expected total=1, got %d", result.Total)
	}
	if result.Items[0].Email != "admin@example.com" {
		t.Errorf("expected admin@example.com, got %s", result.Items[0].Email)
	}
}

func TestGet(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/staff/100" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":100,"email":"admin@example.com","firstName":"Alice","lastName":"Admin","role":"ADMIN"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	member, err := svc.Get(context.Background(), 100)
	if err != nil {
		t.Fatal(err)
	}
	if member.ID != 100 {
		t.Errorf("expected id=100, got %d", member.ID)
	}
	if member.Role != "ADMIN" {
		t.Errorf("expected role=ADMIN, got %s", member.Role)
	}
}

func TestGet_ZeroID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		t.Fatal("should not reach server")
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	_, err := svc.Get(context.Background(), 0)
	if err == nil {
		t.Fatal("expected error for zero staffID")
	}
}

func TestCreate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/staff" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":200}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Create(context.Background(), &staff.StaffMember{
		Email:     "new@example.com",
		FirstName: "Bob",
		LastName:  "Builder",
		Role:      "MANAGER",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.ID != 200 {
		t.Errorf("expected id=200, got %d", result.ID)
	}
}

func TestUpdate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/staff/100" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"updateCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Update(context.Background(), 100, &staff.StaffMember{
		FirstName: "Updated",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.UpdateCount != 1 {
		t.Errorf("expected updateCount=1, got %d", result.UpdateCount)
	}
}

func TestUpdate_ZeroID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		t.Fatal("should not reach server")
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	_, err := svc.Update(context.Background(), 0, &staff.StaffMember{})
	if err == nil {
		t.Fatal("expected error for zero staffID")
	}
}

func TestDelete(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/staff/100" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"deleteCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Delete(context.Background(), 100)
	if err != nil {
		t.Fatal(err)
	}
	if result.DeleteCount != 1 {
		t.Errorf("expected deleteCount=1, got %d", result.DeleteCount)
	}
}

func TestDelete_ZeroID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		t.Fatal("should not reach server")
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	_, err := svc.Delete(context.Background(), 0)
	if err == nil {
		t.Fatal("expected error for zero staffID")
	}
}

func TestList_Error(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"errorMessage":"internal error","errorCode":"500"}`))
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
	if apiErr.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", apiErr.StatusCode)
	}
}
