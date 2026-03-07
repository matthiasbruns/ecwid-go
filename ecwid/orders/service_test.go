package orders_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matthiasbruns/ecwid-go/ecwid/internal/api"
	"github.com/matthiasbruns/ecwid-go/ecwid/orders"
)

func newTestService(t *testing.T, srv *httptest.Server) *orders.Service {
	t.Helper()
	return orders.NewService(api.NewHTTPClient(api.HTTPClientConfig{
		BaseURL: srv.URL + "/api/v3",
		StoreID: "12345",
		Token:   "secret_test",
	}))
}

func TestSearch(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/orders" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("paymentStatus") != "PAID" {
			t.Error("expected paymentStatus=PAID")
		}
		if r.URL.Query().Get("fulfillmentStatus") != "SHIPPED" {
			t.Error("expected fulfillmentStatus=SHIPPED")
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"total":1,"count":1,"offset":0,"limit":100,"items":[{"id":"12345","total":99.99,"email":"test@example.com","paymentStatus":"PAID","fulfillmentStatus":"SHIPPED"}]}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Search(context.Background(), &orders.SearchOptions{
		PaymentStatus:     "PAID",
		FulfillmentStatus: "SHIPPED",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != 1 {
		t.Errorf("expected total=1, got %d", result.Total)
	}
	if result.Items[0].ID != "12345" {
		t.Errorf("expected id=12345, got %s", result.Items[0].ID)
	}
}

func TestSearch_WithQueryParams(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if q.Get("email") != "user@test.com" {
			t.Errorf("expected email=user@test.com, got %q", q.Get("email"))
		}
		if q.Get("customerId") != "42" {
			t.Errorf("expected customerId=42, got %q", q.Get("customerId"))
		}
		if q.Get("totalFrom") != "10.00" {
			t.Errorf("expected totalFrom=10.00, got %q", q.Get("totalFrom"))
		}
		if q.Get("limit") != "50" {
			t.Errorf("expected limit=50, got %q", q.Get("limit"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"total":0,"count":0,"offset":0,"limit":50,"items":[]}`))
	}))
	defer srv.Close()

	totalFrom := 10.0
	svc := newTestService(t, srv)
	_, err := svc.Search(context.Background(), &orders.SearchOptions{
		Email:      "user@test.com",
		CustomerID: 42,
		TotalFrom:  &totalFrom,
		Limit:      50,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestSearch_NilOpts(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"total":0,"count":0,"offset":0,"limit":100,"items":[]}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Search(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != 0 {
		t.Errorf("expected total=0, got %d", result.Total)
	}
}

func TestGet(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/orders/100" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":"100","total":49.99,"email":"test@example.com","paymentStatus":"PAID","fulfillmentStatus":"DELIVERED"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	order, err := svc.Get(context.Background(), "100")
	if err != nil {
		t.Fatal(err)
	}
	if order.Total != 49.99 {
		t.Errorf("expected total=49.99, got %f", order.Total)
	}
	if order.FulfillmentStatus != "DELIVERED" {
		t.Errorf("expected DELIVERED, got %s", order.FulfillmentStatus)
	}
}

func TestGet_EmptyID(t *testing.T) {
	svc := orders.NewService(nil)
	_, err := svc.Get(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty orderID")
	}
}

func TestCreate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/orders" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":1001,"orderid":"XJ12345"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Create(context.Background(), &orders.CreateRequest{
		Subtotal:          40.00,
		Total:             49.99,
		Email:             "test@example.com",
		FulfillmentStatus: "AWAITING_PROCESSING",
		PaymentStatus:     "PAID",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.ID != 1001 {
		t.Errorf("expected id=1001, got %d", result.ID)
	}
	if result.OrderID != "XJ12345" {
		t.Errorf("expected orderid=XJ12345, got %s", result.OrderID)
	}
}

func TestUpdate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/orders/100" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"updateCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Update(context.Background(), "100", &orders.UpdateRequest{
		FulfillmentStatus: "SHIPPED",
		TrackingNumber:    "1Z999AA10123456784",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.UpdateCount != 1 {
		t.Errorf("expected updateCount=1, got %d", result.UpdateCount)
	}
}

func TestUpdate_EmptyID(t *testing.T) {
	svc := orders.NewService(nil)
	_, err := svc.Update(context.Background(), "", nil)
	if err == nil {
		t.Fatal("expected error for empty orderID")
	}
}

func TestDelete(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/orders/100" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"deleteCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Delete(context.Background(), "100")
	if err != nil {
		t.Fatal(err)
	}
	if result.DeleteCount != 1 {
		t.Errorf("expected deleteCount=1, got %d", result.DeleteCount)
	}
}

func TestDelete_EmptyID(t *testing.T) {
	svc := orders.NewService(nil)
	_, err := svc.Delete(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty orderID")
	}
}

func TestGetLast(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/orders/last" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":"999","total":25.00,"paymentStatus":"PAID"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	order, err := svc.GetLast(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if order.ID != "999" {
		t.Errorf("expected id=999, got %s", order.ID)
	}
}

func TestGetDeleted(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/orders/deleted" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("from_date") != "2024-01-01" {
			t.Errorf("expected from_date=2024-01-01, got %q", r.URL.Query().Get("from_date"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"total":1,"count":1,"offset":0,"limit":100,"items":[{"id":42,"date":"2024-06-15"}]}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.GetDeleted(context.Background(), &orders.DeletedOrdersOptions{FromDate: "2024-01-01"})
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != 1 {
		t.Errorf("expected total=1, got %d", result.Total)
	}
	if result.Items[0].ID != 42 {
		t.Errorf("expected id=42, got %d", result.Items[0].ID)
	}
}

func TestGetRepeatOrderURL(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/orders/100/repeatURL" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"repeatOrderUrl":"https://store.example.com/repeat/100"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.GetRepeatOrderURL(context.Background(), "100")
	if err != nil {
		t.Fatal(err)
	}
	if result.RepeatOrderURL != "https://store.example.com/repeat/100" {
		t.Errorf("unexpected URL: %s", result.RepeatOrderURL)
	}
}

func TestGetRepeatOrderURL_EmptyID(t *testing.T) {
	svc := orders.NewService(nil)
	_, err := svc.GetRepeatOrderURL(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty orderID")
	}
}

func TestGetInvoices(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/orders/100/invoices" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"invoices":[{"internalId":1,"id":"INV-001","type":"SALE"}]}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.GetInvoices(context.Background(), "100")
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Invoices) != 1 {
		t.Fatalf("expected 1 invoice, got %d", len(result.Invoices))
	}
	if result.Invoices[0].ID != "INV-001" {
		t.Errorf("expected INV-001, got %s", result.Invoices[0].ID)
	}
}

func TestGetInvoices_EmptyID(t *testing.T) {
	svc := orders.NewService(nil)
	_, err := svc.GetInvoices(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty orderID")
	}
}

func TestCreateInvoice(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/orders/100/invoices/create-invoice" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":42}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.CreateInvoice(context.Background(), "100")
	if err != nil {
		t.Fatal(err)
	}
	if result.ID != 42 {
		t.Errorf("expected id=42, got %d", result.ID)
	}
}

func TestCreateInvoice_EmptyID(t *testing.T) {
	svc := orders.NewService(nil)
	_, err := svc.CreateInvoice(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty orderID")
	}
}

func TestGetExtraFields(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/12345/orders/100/extraFields" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[{"id":"field1","value":"val1","title":"Field 1"}]`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	fields, err := svc.GetExtraFields(context.Background(), "100")
	if err != nil {
		t.Fatal(err)
	}
	if len(fields) != 1 {
		t.Fatalf("expected 1 field, got %d", len(fields))
	}
	if fields[0].ID != "field1" {
		t.Errorf("expected field1, got %s", fields[0].ID)
	}
}

func TestGetExtraFields_EmptyID(t *testing.T) {
	svc := orders.NewService(nil)
	_, err := svc.GetExtraFields(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty orderID")
	}
}

func TestCreateExtraField(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/orders/100/extraFields" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"createCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.CreateExtraField(context.Background(), "100", &orders.ExtraField{
		ID:    "custom_field",
		Value: "test_value",
		Title: "Custom Field",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.CreateCount != 1 {
		t.Errorf("expected createCount=1, got %d", result.CreateCount)
	}
}

func TestCreateExtraField_EmptyID(t *testing.T) {
	svc := orders.NewService(nil)
	_, err := svc.CreateExtraField(context.Background(), "", nil)
	if err == nil {
		t.Fatal("expected error for empty orderID")
	}
}

func TestUpdateExtraField(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/orders/100/extraFields/field1" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"updateCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.UpdateExtraField(context.Background(), "100", "field1", &orders.UpdateExtraFieldRequest{
		Value: "updated_value",
		Title: "Updated Field",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.UpdateCount != 1 {
		t.Errorf("expected updateCount=1, got %d", result.UpdateCount)
	}
}

func TestUpdateExtraField_EmptyOrderID(t *testing.T) {
	svc := orders.NewService(nil)
	_, err := svc.UpdateExtraField(context.Background(), "", "field1", nil)
	if err == nil {
		t.Fatal("expected error for empty orderID")
	}
}

func TestUpdateExtraField_EmptyFieldID(t *testing.T) {
	svc := orders.NewService(nil)
	_, err := svc.UpdateExtraField(context.Background(), "100", "", nil)
	if err == nil {
		t.Fatal("expected error for empty extraFieldID")
	}
}

func TestDeleteExtraField(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/orders/100/extraFields/field1" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"deleteCount":1}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.DeleteExtraField(context.Background(), "100", "field1")
	if err != nil {
		t.Fatal(err)
	}
	if result.DeleteCount != 1 {
		t.Errorf("expected deleteCount=1, got %d", result.DeleteCount)
	}
}

func TestDeleteExtraField_EmptyOrderID(t *testing.T) {
	svc := orders.NewService(nil)
	_, err := svc.DeleteExtraField(context.Background(), "", "field1")
	if err == nil {
		t.Fatal("expected error for empty orderID")
	}
}

func TestDeleteExtraField_EmptyFieldID(t *testing.T) {
	svc := orders.NewService(nil)
	_, err := svc.DeleteExtraField(context.Background(), "100", "")
	if err == nil {
		t.Fatal("expected error for empty extraFieldID")
	}
}

func TestCalculate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/12345/order/calculate" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"subtotal":40.00,"total":49.99,"tax":9.99}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	result, err := svc.Calculate(context.Background(), &orders.CalculateRequest{
		Items: []byte(`[{"productId":123,"price":40.00,"quantity":1}]`),
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != 49.99 {
		t.Errorf("expected total=49.99, got %f", result.Total)
	}
}

func TestSearch_Error(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(`{"errorMessage":"access denied","errorCode":"403"}`))
	}))
	defer srv.Close()

	svc := newTestService(t, srv)
	_, err := svc.Search(context.Background(), nil)
	if err == nil {
		t.Fatal("expected error")
	}

	var apiErr *api.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *api.APIError, got %T", err)
	}
}
