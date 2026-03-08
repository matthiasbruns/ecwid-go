package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

type testProduct struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type testNested struct {
	ID   int      `json:"id"`
	Tags []string `json:"tags"`
}

func newCmdWithOutput(format string) *cobra.Command {
	cmd := &cobra.Command{Use: "test"}
	cmd.Flags().String("output", format, "")
	return cmd
}

func captureOutput(t *testing.T, cmd *cobra.Command, v any) string {
	t.Helper()
	var buf bytes.Buffer
	cmd.SetOut(&buf)
	if err := outputResult(cmd, v); err != nil {
		t.Fatalf("outputResult error: %v", err)
	}
	return buf.String()
}

func TestOutputResult_JSON_Struct(t *testing.T) {
	cmd := newCmdWithOutput("json")
	out := captureOutput(t, cmd, testProduct{ID: 1, Name: "Widget", Price: 9.99})

	if !strings.Contains(out, `"id": 1`) {
		t.Errorf("expected id field, got:\n%s", out)
	}
	if !strings.Contains(out, `"name": "Widget"`) {
		t.Errorf("expected name field, got:\n%s", out)
	}
}

func TestOutputResult_JSON_Slice(t *testing.T) {
	cmd := newCmdWithOutput("json")
	items := []testProduct{
		{ID: 1, Name: "A", Price: 1.0},
		{ID: 2, Name: "B", Price: 2.0},
	}
	out := captureOutput(t, cmd, items)

	if !strings.HasPrefix(strings.TrimSpace(out), "[") {
		t.Errorf("expected JSON array, got:\n%s", out)
	}
}

func TestOutputResult_JSON_Default(t *testing.T) {
	// Empty output flag should default to json.
	cmd := newCmdWithOutput("")
	out := captureOutput(t, cmd, testProduct{ID: 1, Name: "X", Price: 0})

	if !strings.Contains(out, `"id": 1`) {
		t.Errorf("expected json output by default, got:\n%s", out)
	}
}

func TestOutputResult_Table_Slice(t *testing.T) {
	cmd := newCmdWithOutput("table")
	items := []testProduct{
		{ID: 1, Name: "Widget", Price: 9.99},
		{ID: 2, Name: "Gadget", Price: 19.99},
	}
	out := captureOutput(t, cmd, items)

	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines (header + 2 rows), got %d:\n%s", len(lines), out)
	}

	header := lines[0]
	if !strings.Contains(header, "ID") || !strings.Contains(header, "NAME") || !strings.Contains(header, "PRICE") {
		t.Errorf("expected uppercase headers, got: %s", header)
	}

	if !strings.Contains(lines[1], "Widget") || !strings.Contains(lines[1], "9.99") {
		t.Errorf("expected first row data, got: %s", lines[1])
	}
}

func TestOutputResult_Table_SingleStruct(t *testing.T) {
	cmd := newCmdWithOutput("table")
	out := captureOutput(t, cmd, testProduct{ID: 42, Name: "Solo", Price: 5.0})

	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines (header + 1 row), got %d:\n%s", len(lines), out)
	}
	if !strings.Contains(lines[1], "Solo") {
		t.Errorf("expected data row, got: %s", lines[1])
	}
}

func TestOutputResult_Table_EmptySlice(t *testing.T) {
	cmd := newCmdWithOutput("table")
	out := captureOutput(t, cmd, []testProduct{})

	if !strings.Contains(out, "(empty)") {
		t.Errorf("expected (empty) for empty slice, got:\n%s", out)
	}
}

func TestOutputResult_Table_NestedFields(t *testing.T) {
	cmd := newCmdWithOutput("table")
	items := []testNested{
		{ID: 1, Tags: []string{"a", "b"}},
	}
	out := captureOutput(t, cmd, items)

	if !strings.Contains(out, `["a","b"]`) {
		t.Errorf("expected JSON-encoded tags, got:\n%s", out)
	}
}

func TestOutputResult_Table_NilPointer(t *testing.T) {
	cmd := newCmdWithOutput("table")
	var p *testProduct
	out := captureOutput(t, cmd, p)

	if !strings.Contains(out, "(nil)") {
		t.Errorf("expected (nil) for nil pointer, got:\n%s", out)
	}
}

func TestOutputResult_Table_SliceWithNilElements(t *testing.T) {
	cmd := newCmdWithOutput("table")
	p := &testProduct{ID: 1, Name: "Valid", Price: 5.0}
	items := []*testProduct{nil, p, nil}
	out := captureOutput(t, cmd, items)

	// Should not panic; should contain header and the valid row.
	if !strings.Contains(out, "ID") || !strings.Contains(out, "NAME") {
		t.Errorf("expected table headers, got:\n%s", out)
	}
	if !strings.Contains(out, "Valid") {
		t.Errorf("expected non-nil row data, got:\n%s", out)
	}
}

func TestOutputResult_Table_AllNilSlice(t *testing.T) {
	cmd := newCmdWithOutput("table")
	items := []*testProduct{nil, nil}
	// Should fall back to JSON since no non-nil element found.
	out := captureOutput(t, cmd, items)
	if !strings.Contains(out, "null") {
		t.Errorf("expected JSON fallback for all-nil slice, got:\n%s", out)
	}
}

func TestOutputResult_InvalidFormat(t *testing.T) {
	cmd := newCmdWithOutput("xml")
	var buf bytes.Buffer
	cmd.SetOut(&buf)
	err := outputResult(cmd, testProduct{ID: 1})
	if err == nil {
		t.Fatal("expected error for unsupported format")
	}
	if !strings.Contains(err.Error(), "unsupported output format") {
		t.Errorf("unexpected error: %v", err)
	}
}
