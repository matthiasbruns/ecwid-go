package cmdutil

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// OutputResult writes v to cmd's stdout in the format specified by the --output flag.
// It supports "json" (pretty-printed, default) and "table" (human-readable columns).
// v may be a struct, a pointer to a struct, or a slice of structs.
func OutputResult(cmd *cobra.Command, v any) error {
	format, _ := cmd.Flags().GetString("output")
	if format == "" {
		format = "json"
	}

	switch format {
	case "json":
		return outputJSON(cmd, v)
	case "table":
		return outputTable(cmd, v)
	default:
		return fmt.Errorf("unsupported output format: %s", format)
	}
}

// outputJSON writes v as pretty-printed JSON to cmd's stdout.
func outputJSON(cmd *cobra.Command, v any) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal json: %w", err)
	}
	_, err = fmt.Fprintln(cmd.OutOrStdout(), string(data))
	return err
}

// outputTable writes v as a tab-aligned table to cmd's stdout.
// For a slice, each element becomes a row. For a single struct, one row is rendered.
func outputTable(cmd *cobra.Command, v any) error {
	rv := reflect.ValueOf(v)

	// Unwrap pointer.
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			_, err := fmt.Fprintln(cmd.OutOrStdout(), "(nil)")
			return err
		}
		rv = rv.Elem()
	}

	switch rv.Kind() {
	case reflect.Slice:
		return outputTableSlice(cmd, rv)
	case reflect.Struct:
		return outputTableSlice(cmd, reflect.ValueOf([]any{rv.Interface()}))
	default:
		// Fall back to JSON for non-struct types.
		return outputJSON(cmd, v)
	}
}

// outputTableSlice renders a slice as a tab-aligned table with headers derived from
// the json tags of the element type's exported fields.
func outputTableSlice(cmd *cobra.Command, rv reflect.Value) error {
	if rv.Len() == 0 {
		_, err := fmt.Fprintln(cmd.OutOrStdout(), "(empty)")
		return err
	}

	// Determine element type from the first non-nil element.
	var first reflect.Value
	for i := range rv.Len() {
		candidate := rv.Index(i)
		for candidate.Kind() == reflect.Ptr || candidate.Kind() == reflect.Interface {
			if candidate.IsNil() {
				candidate = reflect.Value{}
				break
			}
			candidate = candidate.Elem()
		}
		if candidate.IsValid() && candidate.Kind() == reflect.Struct {
			first = candidate
			break
		}
	}
	if !first.IsValid() || first.Kind() != reflect.Struct {
		return outputJSON(cmd, rv.Interface())
	}

	cols := tableColumns(first.Type())
	if len(cols) == 0 {
		return outputJSON(cmd, rv.Interface())
	}

	w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 2, ' ', 0)

	// Header row.
	headers := make([]string, len(cols))
	for i, c := range cols {
		headers[i] = strings.ToUpper(c.header)
	}
	if _, err := fmt.Fprintln(w, strings.Join(headers, "\t")); err != nil {
		return err
	}

	// Data rows.
	for i := range rv.Len() {
		elem := rv.Index(i)
		for elem.Kind() == reflect.Ptr || elem.Kind() == reflect.Interface {
			if elem.IsNil() {
				elem = reflect.Value{}
				break
			}
			elem = elem.Elem()
		}
		vals := make([]string, len(cols))
		if elem.IsValid() {
			for j, c := range cols {
				vals[j] = formatField(elem.Field(c.index))
			}
		}
		if _, err := fmt.Fprintln(w, strings.Join(vals, "\t")); err != nil {
			return err
		}
	}

	return w.Flush()
}

type column struct {
	header string
	index  int
}

// tableColumns returns the columns for a struct type, using json tag names as headers.
func tableColumns(t reflect.Type) []column {
	var cols []column
	for i := range t.NumField() {
		f := t.Field(i)
		if !f.IsExported() {
			continue
		}
		tag := f.Tag.Get("json")
		if tag == "-" {
			continue
		}
		name, _, _ := strings.Cut(tag, ",")
		if name == "" {
			name = f.Name
		}
		cols = append(cols, column{header: name, index: i})
	}
	return cols
}

// formatField converts a reflect.Value to its string representation for table output.
func formatField(v reflect.Value) string {
	if !v.IsValid() {
		return ""
	}

	// Handle nil pointers.
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return ""
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Slice, reflect.Map:
		if v.IsNil() || v.Len() == 0 {
			return ""
		}
		data, err := json.Marshal(v.Interface())
		if err != nil {
			return fmt.Sprintf("%v", v.Interface())
		}
		return string(data)
	case reflect.Struct:
		data, err := json.Marshal(v.Interface())
		if err != nil {
			return fmt.Sprintf("%v", v.Interface())
		}
		return string(data)
	default:
		return fmt.Sprintf("%v", v.Interface())
	}
}
