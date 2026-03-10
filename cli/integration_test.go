package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

var binaryPath string

func TestMain(m *testing.M) {
	// Build the CLI binary once for all tests.
	dir, err := os.MkdirTemp("", "ecwid-cli-test-*")
	if err != nil {
		panic(err)
	}

	binaryPath = filepath.Join(dir, "ecwid")
	buildCmd := exec.Command("go", "build", "-o", binaryPath, ".")
	buildCmd.Dir = "."
	buildCmd.Stderr = os.Stderr
	if err := buildCmd.Run(); err != nil {
		panic("failed to build CLI binary: " + err.Error())
	}

	code := m.Run()
	_ = os.RemoveAll(dir)
	os.Exit(code)
}

// runCLI executes the compiled CLI binary with the given args and env vars.
// It returns stdout, stderr, and the exit code.
type cliResult struct {
	stdout   string
	stderr   string
	exitCode int
}

func runCLI(t *testing.T, baseURL string, args ...string) cliResult {
	t.Helper()
	return runCLIWithStdin(t, baseURL, "", args...)
}

func runCLIWithStdin(t *testing.T, baseURL, stdin string, args ...string) cliResult {
	t.Helper()

	cmd := exec.Command(binaryPath, args...)
	cmd.Env = append(os.Environ(),
		"ECWID_STORE_ID="+testStoreID,
		"ECWID_TOKEN=test-token",
		"ECWID_BASE_URL="+baseURL,
		"HOME="+t.TempDir(), // Prevent loading user's ~/.ecwid.yaml
	)
	// Always set stdin to prevent hanging if a command unexpectedly reads.
	cmd.Stdin = strings.NewReader(stdin)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	exitCode := 0
	if exitErr, ok := err.(*exec.ExitError); ok {
		exitCode = exitErr.ExitCode()
	} else if err != nil {
		t.Fatalf("failed to run CLI: %v", err)
	}

	return cliResult{
		stdout:   stdout.String(),
		stderr:   stderr.String(),
		exitCode: exitCode,
	}
}

// ── Config / General ────────────────────────────────────────────────────

func TestCLI_Version(t *testing.T) {
	cmd := exec.Command(binaryPath, "version")
	cmd.Env = append(os.Environ(), "HOME="+t.TempDir())
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("version command failed: %v", err)
	}
	if !strings.HasPrefix(string(out), "ecwid-cli ") {
		t.Errorf("expected output starting with 'ecwid-cli ', got %q", string(out))
	}
}

func TestCLI_Help(t *testing.T) {
	cmd := exec.Command(binaryPath, "--help")
	cmd.Env = append(os.Environ(), "HOME="+t.TempDir())
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("help command failed: %v", err)
	}
	if !strings.Contains(string(out), "Ecwid") {
		t.Errorf("expected help text, got %q", string(out))
	}
}

func TestCLI_SubcommandHelp(t *testing.T) {
	t.Parallel()
	for _, sub := range []string{
		"products", "orders", "categories", "customers",
		"dictionaries", "profile", "carts", "subscriptions",
		"promotions", "coupons", "reviews", "staff", "domains", "reports",
	} {
		t.Run(sub, func(t *testing.T) {
			t.Parallel()
			cmd := exec.Command(binaryPath, sub, "--help")
			cmd.Env = append(os.Environ(), "HOME="+t.TempDir())
			if err := cmd.Run(); err != nil {
				t.Fatalf("%s --help failed: %v", sub, err)
			}
		})
	}
}

func TestCLI_MissingCredentials(t *testing.T) {
	cmd := exec.Command(binaryPath, "profile", "get")
	cmd.Env = append(os.Environ(), "HOME="+t.TempDir())
	err := cmd.Run()
	if err == nil {
		t.Fatal("expected error when credentials are missing")
	}
}

func TestCLI_EnvVars(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "profile", "get")
	if r.exitCode != 0 {
		t.Fatalf("expected exit 0, got %d: %s", r.exitCode, r.stderr)
	}
	if !strings.Contains(r.stdout, "storeId") {
		t.Errorf("expected storeId in output, got %q", r.stdout)
	}
}

func TestCLI_ConfigFile(t *testing.T) {
	srv := newMockServer(t)
	tmpDir := t.TempDir()
	cfgPath := filepath.Join(tmpDir, "config.yaml")
	cfgContent := "store_id: " + testStoreID + "\ntoken: test-token\nbase_url: " + srv.URL + "\n"
	if err := os.WriteFile(cfgPath, []byte(cfgContent), 0o600); err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command(binaryPath, "--config", cfgPath, "profile", "get")
	cmd.Env = append(os.Environ(), "HOME="+tmpDir)
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("config file test failed: %v", err)
	}
	if !strings.Contains(string(out), "storeId") {
		t.Errorf("expected storeId in output, got %q", string(out))
	}
}

func TestCLI_FlagPrecedence(t *testing.T) {
	srv := newMockServer(t)
	// Set env vars but override store-id via flag (should still work).
	cmd := exec.Command(binaryPath, "--store-id", testStoreID, "profile", "get")
	cmd.Env = append(os.Environ(),
		"ECWID_TOKEN=test-token",
		"ECWID_BASE_URL="+srv.URL,
		"ECWID_STORE_ID=wrong-id", // Flag should override this
		"HOME="+t.TempDir(),
	)
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("flag precedence test failed: %v", err)
	}
	if !strings.Contains(string(out), "storeId") {
		t.Errorf("expected storeId in output, got %q", string(out))
	}
}

func TestCLI_BaseURLFlag(t *testing.T) {
	srv := newMockServer(t)
	// Use --base-url flag instead of env var.
	cmd := exec.Command(binaryPath, "--base-url", srv.URL, "profile", "get")
	cmd.Env = append(os.Environ(),
		"ECWID_STORE_ID="+testStoreID,
		"ECWID_TOKEN=test-token",
		"HOME="+t.TempDir(),
	)
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("--base-url flag test failed: %v", err)
	}
	assertContains(t, string(out), "storeId")
}

func TestCLI_MaxRetriesFlag(t *testing.T) {
	srv := newMockServer(t)
	// Use --max-retries flag; the /retry endpoint returns 429 first, then 200.
	cmd := exec.Command(binaryPath, "--base-url", srv.URL+"/retry", "--max-retries", "2", "profile", "get")
	cmd.Env = append(os.Environ(),
		"ECWID_STORE_ID="+testStoreID,
		"ECWID_TOKEN=test-token",
		"HOME="+t.TempDir(),
	)
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("--max-retries flag test failed: %v", err)
	}
	assertContains(t, string(out), "Retry Store")
}

// ── Dictionaries ────────────────────────────────────────────────────────

func TestCLI_Dictionaries_Countries(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "dictionaries", "countries")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, `"US"`)
}

func TestCLI_Dictionaries_Currencies(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "dictionaries", "currencies")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, `"USD"`)
}

// ── Profile ─────────────────────────────────────────────────────────────

func TestCLI_Profile_Get(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "profile", "get")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "Test Store")
}

func TestCLI_Profile_Update(t *testing.T) {
	srv := newMockServer(t)
	r := runCLIWithStdin(t, srv.URL, `{"settings":{"storeName":"Updated"}}`, "profile", "update")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "updateCount")
}

// ── Categories ──────────────────────────────────────────────────────────

func TestCLI_Categories_List(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "categories", "list")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "Test Category")
}

func TestCLI_Categories_Get(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "categories", "get", "1001")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "Test Category")
}

func TestCLI_Categories_Create(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "categories", "create", "--name", "New Cat")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "1001")
}

func TestCLI_Categories_Update(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "categories", "update", "1001", "--name", "Updated Cat")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "updateCount")
}

func TestCLI_Categories_Delete(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "categories", "delete", "1001")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "deleteCount")
}

// ── Products ────────────────────────────────────────────────────────────

func TestCLI_Products_List(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "products", "list")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "Test Product")
}

func TestCLI_Products_Get(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "products", "get", "2001")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "Test Product")
}

func TestCLI_Products_Create(t *testing.T) {
	srv := newMockServer(t)
	r := runCLIWithStdin(t, srv.URL, `{"name":"New Product","price":9.99}`, "products", "create")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "2001")
}

func TestCLI_Products_Update(t *testing.T) {
	srv := newMockServer(t)
	r := runCLIWithStdin(t, srv.URL, `{"name":"Updated"}`, "products", "update", "2001")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "updateCount")
}

func TestCLI_Products_Delete(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "products", "delete", "2001")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "deleteCount")
}

// ── Orders ──────────────────────────────────────────────────────────────

func TestCLI_Orders_List(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "orders", "list")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "ORD-001")
}

func TestCLI_Orders_Get(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "orders", "get", "ORD-001")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "test@example.com")
}

func TestCLI_Orders_Create(t *testing.T) {
	srv := newMockServer(t)
	r := runCLIWithStdin(t, srv.URL,
		`{"subtotal":9.99,"total":9.99,"email":"test@example.com","paymentStatus":"PAID","fulfillmentStatus":"AWAITING_PROCESSING"}`,
		"orders", "create")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "3001")
}

func TestCLI_Orders_Update(t *testing.T) {
	srv := newMockServer(t)
	r := runCLIWithStdin(t, srv.URL, `{"fulfillmentStatus":"SHIPPED"}`, "orders", "update", "ORD-001")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "updateCount")
}

func TestCLI_Orders_Delete(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "orders", "delete", "ORD-001")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "deleteCount")
}

// ── Customers ───────────────────────────────────────────────────────────

func TestCLI_Customers_List(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "customers", "list")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "Test Customer")
}

func TestCLI_Customers_Get(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "customers", "get", "4001")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "customer@example.com")
}

func TestCLI_Customers_Create(t *testing.T) {
	srv := newMockServer(t)
	r := runCLIWithStdin(t, srv.URL, `{"email":"new@example.com","name":"New"}`, "customers", "create")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "4001")
}

func TestCLI_Customers_Update(t *testing.T) {
	srv := newMockServer(t)
	r := runCLIWithStdin(t, srv.URL, `{"name":"Updated"}`, "customers", "update", "4001")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "updateCount")
}

func TestCLI_Customers_Delete(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "customers", "delete", "4001")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "deleteCount")
}

// ── Promotions ──────────────────────────────────────────────────────────

func TestCLI_Promotions_List(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "promotions", "list")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "Test Promo")
}

func TestCLI_Promotions_Create(t *testing.T) {
	srv := newMockServer(t)
	r := runCLIWithStdin(t, srv.URL, `{"name":"New Promo"}`, "promotions", "create")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "5001")
}

func TestCLI_Promotions_Update(t *testing.T) {
	srv := newMockServer(t)
	r := runCLIWithStdin(t, srv.URL, `{"name":"Updated Promo"}`, "promotions", "update", "5001")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "updateCount")
}

func TestCLI_Promotions_Delete(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "promotions", "delete", "5001")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "deleteCount")
}

// ── Coupons ─────────────────────────────────────────────────────────────

func TestCLI_Coupons_List(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "coupons", "list")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "TEST10")
}

func TestCLI_Coupons_Get(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "coupons", "get", "6001")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "Test Coupon")
}

func TestCLI_Coupons_Create(t *testing.T) {
	srv := newMockServer(t)
	r := runCLIWithStdin(t, srv.URL, `{"code":"NEW20","name":"New Coupon"}`, "coupons", "create")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "6001")
}

func TestCLI_Coupons_Update(t *testing.T) {
	srv := newMockServer(t)
	r := runCLIWithStdin(t, srv.URL, `{"name":"Updated Coupon"}`, "coupons", "update", "6001")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "updateCount")
}

func TestCLI_Coupons_Delete(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "coupons", "delete", "6001")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "deleteCount")
}

// ── Reviews ─────────────────────────────────────────────────────────────

func TestCLI_Reviews_List(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "reviews", "list")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "7001")
}

func TestCLI_Reviews_Update(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "reviews", "update", "7001", "--status", "APPROVED")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "updateCount")
}

func TestCLI_Reviews_Delete(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "reviews", "delete", "7001")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "deleteCount")
}

// ── Staff ───────────────────────────────────────────────────────────────

func TestCLI_Staff_List(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "staff", "list")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "staff@example.com")
}

func TestCLI_Staff_Get(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "staff", "get", "staff-001")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "staff@example.com")
}

func TestCLI_Staff_Create(t *testing.T) {
	srv := newMockServer(t)
	r := runCLIWithStdin(t, srv.URL, `{"email":"new@example.com"}`, "staff", "create")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "success")
}

func TestCLI_Staff_Update(t *testing.T) {
	srv := newMockServer(t)
	r := runCLIWithStdin(t, srv.URL, `{"firstName":"Updated"}`, "staff", "update", "staff-001")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "updateCount")
}

func TestCLI_Staff_Delete(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "staff", "delete", "staff-001")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "deleteCount")
}

// ── Domains ─────────────────────────────────────────────────────────────

func TestCLI_Domains_Get(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "domains", "get")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "test.ecwid.com")
}

func TestCLI_Domains_Purchase(t *testing.T) {
	srv := newMockServer(t)
	r := runCLIWithStdin(t, srv.URL, `{"domainName":"mydomain.com"}`, "domains", "purchase")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "mydomain.com")
}

// ── Reports ─────────────────────────────────────────────────────────────

func TestCLI_Reports_Get(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "reports", "get", "allOrders")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "allOrders")
}

func TestCLI_Reports_Stats(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "reports", "stats")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "productsUpdated")
}

// ── Carts ───────────────────────────────────────────────────────────────

func TestCLI_Carts_List(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "carts", "list")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "cart-001")
}

func TestCLI_Carts_Get(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "carts", "get", "cart-001")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "cart-001")
}

// ── Subscriptions ───────────────────────────────────────────────────────

func TestCLI_Subscriptions_List(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "subscriptions", "list")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "ACTIVE")
}

func TestCLI_Subscriptions_Get(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "subscriptions", "get", "8001")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	assertContains(t, r.stdout, "8001")
}

// ── Error Handling ──────────────────────────────────────────────────────

func TestCLI_Error_401(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL+"/error-401", "profile", "get")
	if r.exitCode == 0 {
		t.Fatal("expected non-zero exit code for 401")
	}
	assertContains(t, r.stderr, "401")
}

func TestCLI_Error_404(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL+"/error-404", "profile", "get")
	if r.exitCode == 0 {
		t.Fatal("expected non-zero exit code for 404")
	}
	assertContains(t, r.stderr, "404")
}

func TestCLI_Error_429(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL+"/error-429", "profile", "get")
	if r.exitCode == 0 {
		t.Fatal("expected non-zero exit code for 429")
	}
	assertContains(t, r.stderr, "429")
}

func TestCLI_Error_MalformedJSON(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL+"/malformed", "profile", "get")
	if r.exitCode == 0 {
		t.Fatal("expected non-zero exit code for malformed JSON")
	}
}

// ── Output Format ───────────────────────────────────────────────────────

func TestCLI_OutputJSON(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "--output", "json", "profile", "get")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	// JSON output should have indentation.
	assertContains(t, r.stdout, "{\n")
}

func TestCLI_OutputTable(t *testing.T) {
	srv := newMockServer(t)
	r := runCLI(t, srv.URL, "--output", "table", "dictionaries", "countries")
	if r.exitCode != 0 {
		t.Fatalf("exit %d: %s", r.exitCode, r.stderr)
	}
	// Table output should have uppercase headers.
	assertContains(t, r.stdout, "CODE")
}

// ── Helpers ─────────────────────────────────────────────────────────────

func assertContains(t *testing.T, got, want string) {
	t.Helper()
	if !strings.Contains(got, want) {
		t.Errorf("output does not contain %q:\n%s", want, got)
	}
}
