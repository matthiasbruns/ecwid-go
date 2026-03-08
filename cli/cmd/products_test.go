package cmd

import (
	"strings"
	"testing"
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
	err := productsGetCmd.Args(productsGetCmd, []string{})
	if err == nil {
		t.Fatal("expected error when no ID provided")
	}
}

func TestProductsDeleteCmd_RequiresArg(t *testing.T) {
	err := productsDeleteCmd.Args(productsDeleteCmd, []string{})
	if err == nil {
		t.Fatal("expected error when no ID provided")
	}
}

func TestProductsUpdateCmd_RequiresArg(t *testing.T) {
	err := productsUpdateCmd.Args(productsUpdateCmd, []string{})
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
	err := productsGetCmd.RunE(productsGetCmd, []string{"not-a-number"})
	if err == nil {
		t.Fatal("expected error for invalid ID")
	}
	if !strings.Contains(err.Error(), "invalid product ID") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestProductsGetCmd_NegativeID(t *testing.T) {
	err := productsGetCmd.RunE(productsGetCmd, []string{"-1"})
	if err == nil {
		t.Fatal("expected error for negative ID")
	}
	if !strings.Contains(err.Error(), "must be a positive integer") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestProductsGetCmd_ZeroID(t *testing.T) {
	err := productsGetCmd.RunE(productsGetCmd, []string{"0"})
	if err == nil {
		t.Fatal("expected error for zero ID")
	}
	if !strings.Contains(err.Error(), "must be a positive integer") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestProductsListCmd_RejectsArgs(t *testing.T) {
	err := productsListCmd.Args(productsListCmd, []string{"extra"})
	if err == nil {
		t.Fatal("expected error for extra args")
	}
}

func TestProductsCreateCmd_RejectsArgs(t *testing.T) {
	err := productsCreateCmd.Args(productsCreateCmd, []string{"extra"})
	if err == nil {
		t.Fatal("expected error for extra args")
	}
}
