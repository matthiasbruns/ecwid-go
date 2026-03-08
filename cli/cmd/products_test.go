package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestProductsCmd_HasSubcommands(t *testing.T) {
	subs := productsCmd.Commands()
	want := map[string]bool{
		"list":   false,
		"get":    false,
		"create": false,
		"update": false,
		"delete": false,
	}
	for _, c := range subs {
		if _, ok := want[c.Name()]; ok {
			want[c.Name()] = true
		}
	}
	for name, found := range want {
		if !found {
			t.Errorf("missing subcommand: %s", name)
		}
	}
}

func TestProductsListCmd_Flags(t *testing.T) {
	flags := []string{"keyword", "category", "limit", "offset", "enabled", "in-stock", "sku", "sort-by"}
	for _, f := range flags {
		if productsListCmd.Flags().Lookup(f) == nil {
			t.Errorf("missing flag: --%s", f)
		}
	}
}

func TestProductsGetCmd_RequiresArg(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	cmd.AddCommand(productsGetCmd)

	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetErr(&buf)
	cmd.SetArgs([]string{"get"})

	err := cmd.Execute()
	if err == nil {
		t.Fatal("expected error when no ID provided")
	}
}

func TestProductsDeleteCmd_RequiresArg(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	cmd.AddCommand(productsDeleteCmd)

	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetErr(&buf)
	cmd.SetArgs([]string{"delete"})

	err := cmd.Execute()
	if err == nil {
		t.Fatal("expected error when no ID provided")
	}
}

func TestProductsCreateCmd_HasFileFlag(t *testing.T) {
	if productsCreateCmd.Flags().Lookup("file") == nil {
		t.Error("create command missing --file flag")
	}
}

func TestProductsUpdateCmd_HasFileFlag(t *testing.T) {
	if productsUpdateCmd.Flags().Lookup("file") == nil {
		t.Error("update command missing --file flag")
	}
}

func TestProductsUpdateCmd_RequiresArg(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	cmd.AddCommand(productsUpdateCmd)

	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetErr(&buf)
	cmd.SetArgs([]string{"update"})

	err := cmd.Execute()
	if err == nil {
		t.Fatal("expected error when no ID provided")
	}
}

func TestProductsCmd_RegisteredOnRoot(t *testing.T) {
	found := false
	for _, c := range rootCmd.Commands() {
		if c.Name() == "products" {
			found = true
			break
		}
	}
	if !found {
		t.Error("products command not registered on root")
	}
}

func TestProductsGetCmd_InvalidID(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	getCmd := &cobra.Command{
		Use:  "get <id>",
		Args: cobra.ExactArgs(1),
		RunE: productsGetCmd.RunE,
	}
	cmd.AddCommand(getCmd)

	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetErr(&buf)
	cmd.SetArgs([]string{"get", "not-a-number"})

	err := cmd.Execute()
	if err == nil {
		t.Fatal("expected error for invalid ID")
	}
	if !strings.Contains(err.Error(), "invalid product ID") {
		t.Errorf("unexpected error: %v", err)
	}
}
