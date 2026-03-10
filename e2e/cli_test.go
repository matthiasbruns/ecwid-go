package e2e

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// cliResult holds the output of a CLI invocation.
type cliResult struct {
	stdout   string
	stderr   string
	exitCode int
}

// runBinary executes the compiled CLI binary with the given args and environment.
// baseEnv is merged on top of a minimal environment (PATH only).
func runBinary(t *testing.T, env []string, args ...string) cliResult {
	t.Helper()

	cmd := exec.Command(binaryPath, args...)
	// Start with a clean environment to avoid inheriting user config.
	cmd.Env = append([]string{
		"PATH=" + os.Getenv("PATH"),
		"HOME=" + t.TempDir(),
	}, env...)
	cmd.Stdin = strings.NewReader("") // prevent hanging

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

// requireCredentialEnv returns env vars for the real Ecwid API, or skips the test.
func requireCredentialEnv(t *testing.T) []string {
	t.Helper()
	storeID := os.Getenv("ECWID_STORE_ID")
	token := os.Getenv("ECWID_TOKEN")
	if storeID == "" || token == "" {
		t.Skip("skipping: ECWID_STORE_ID and ECWID_TOKEN required")
	}
	return []string{
		"ECWID_STORE_ID=" + storeID,
		"ECWID_TOKEN=" + token,
	}
}

// ── Version / Help (no credentials needed) ──────────────────────────────

func TestCLI_Version(t *testing.T) {
	r := runBinary(t, nil, "version")
	if r.exitCode != 0 {
		t.Fatalf("expected exit 0, got %d: %s", r.exitCode, r.stderr)
	}
	if !strings.HasPrefix(r.stdout, "ecwid-cli ") {
		t.Errorf("expected output starting with 'ecwid-cli ', got %q", r.stdout)
	}
}

func TestCLI_Help(t *testing.T) {
	r := runBinary(t, nil, "--help")
	if r.exitCode != 0 {
		t.Fatalf("expected exit 0, got %d: %s", r.exitCode, r.stderr)
	}
	if !strings.Contains(r.stdout, "Ecwid") {
		t.Errorf("expected help output to mention Ecwid, got %q", r.stdout)
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
			r := runBinary(t, nil, sub, "--help")
			if r.exitCode != 0 {
				t.Fatalf("%s --help: exit %d: %s", sub, r.exitCode, r.stderr)
			}
		})
	}
}

// ── Missing credentials ─────────────────────────────────────────────────

func TestCLI_MissingCredentials(t *testing.T) {
	r := runBinary(t, nil, "profile", "get")
	if r.exitCode == 0 {
		t.Fatal("expected non-zero exit code when credentials are missing")
	}
}

// ── Config file loading ─────────────────────────────────────────────────

func TestCLI_ConfigFile(t *testing.T) {
	env := requireCredentialEnv(t)

	tmpDir := t.TempDir()
	cfgPath := filepath.Join(tmpDir, "config.yaml")
	cfgContent := "store_id: " + os.Getenv("ECWID_STORE_ID") + "\ntoken: " + os.Getenv("ECWID_TOKEN") + "\n"
	if err := os.WriteFile(cfgPath, []byte(cfgContent), 0o600); err != nil {
		t.Fatal(err)
	}

	// Use --config flag; don't pass store-id/token via env so config file is the source.
	r := runBinary(t, nil, "--config", cfgPath, "dictionaries", "countries")
	if r.exitCode != 0 {
		t.Fatalf("config file test: exit %d: %s", r.exitCode, r.stderr)
	}
	if !strings.Contains(r.stdout, "US") {
		t.Errorf("expected US in output, got %q", r.stdout)
	}
	_ = env // credentials validated above
}

// ── Env var loading ─────────────────────────────────────────────────────

func TestCLI_EnvVarLoading(t *testing.T) {
	env := requireCredentialEnv(t)

	r := runBinary(t, env, "dictionaries", "countries")
	if r.exitCode != 0 {
		t.Fatalf("env var loading: exit %d: %s", r.exitCode, r.stderr)
	}
	if !strings.Contains(r.stdout, "US") {
		t.Errorf("expected US in output, got %q", r.stdout)
	}
}

// ── Flag precedence over env ────────────────────────────────────────────

func TestCLI_FlagPrecedence(t *testing.T) {
	env := requireCredentialEnv(t)

	// Pass correct store-id via flag, wrong one via env.
	storeID := os.Getenv("ECWID_STORE_ID")
	flagEnv := []string{
		"ECWID_STORE_ID=wrong-store-id",
		"ECWID_TOKEN=" + os.Getenv("ECWID_TOKEN"),
	}

	r := runBinary(t, flagEnv, "--store-id", storeID, "dictionaries", "countries")
	if r.exitCode != 0 {
		t.Fatalf("flag precedence: exit %d: %s", r.exitCode, r.stderr)
	}
	if !strings.Contains(r.stdout, "US") {
		t.Errorf("expected US in output, got %q", r.stdout)
	}
	_ = env
}

// ── JSON output ─────────────────────────────────────────────────────────

func TestCLI_OutputJSON(t *testing.T) {
	env := requireCredentialEnv(t)

	r := runBinary(t, env, "--output", "json", "dictionaries", "countries")
	if r.exitCode != 0 {
		t.Fatalf("json output: exit %d: %s", r.exitCode, r.stderr)
	}
	// JSON output should contain indented brackets.
	if !strings.Contains(r.stdout, "[\n") {
		t.Errorf("expected JSON array with newlines, got %q", r.stdout)
	}
}

// ── Table output ────────────────────────────────────────────────────────

func TestCLI_OutputTable(t *testing.T) {
	env := requireCredentialEnv(t)

	r := runBinary(t, env, "--output", "table", "dictionaries", "countries")
	if r.exitCode != 0 {
		t.Fatalf("table output: exit %d: %s", r.exitCode, r.stderr)
	}
	// Table output should have uppercase headers.
	if !strings.Contains(r.stdout, "CODE") {
		t.Errorf("expected CODE header in table output, got %q", r.stdout)
	}
}

// ── Real API: dictionaries countries ─────────────────────────────────────

func TestCLI_Dictionaries_Countries(t *testing.T) {
	env := requireCredentialEnv(t)

	r := runBinary(t, env, "dictionaries", "countries")
	if r.exitCode != 0 {
		t.Fatalf("dictionaries countries: exit %d: %s", r.exitCode, r.stderr)
	}
	if !strings.Contains(r.stdout, "United States") {
		t.Errorf("expected 'United States' in output, got %q", r.stdout)
	}
}

// ── Invalid store ID ────────────────────────────────────────────────────

func TestCLI_InvalidStoreID(t *testing.T) {
	env := []string{
		"ECWID_STORE_ID=0",
		"ECWID_TOKEN=invalid-token-for-e2e-test",
	}

	r := runBinary(t, env, "profile", "get")
	if r.exitCode == 0 {
		t.Fatal("expected non-zero exit code for invalid store ID")
	}
}
