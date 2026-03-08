package cmd

import (
	"strings"
	"testing"
)

func TestCustomersCmd_HasSubcommands(t *testing.T) {
	subs := customersCmd.Commands()
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

func TestCustomersListCmd_Flags(t *testing.T) {
	flags := []string{"keyword", "email", "limit", "offset"}
	for _, f := range flags {
		if customersListCmd.Flags().Lookup(f) == nil {
			t.Errorf("missing flag: --%s", f)
		}
	}
}

func TestCustomersGetCmd_RequiresArg(t *testing.T) {
	err := customersGetCmd.Args(customersGetCmd, []string{})
	if err == nil {
		t.Fatal("expected error when no ID provided")
	}
}

func TestCustomersDeleteCmd_RequiresArg(t *testing.T) {
	err := customersDeleteCmd.Args(customersDeleteCmd, []string{})
	if err == nil {
		t.Fatal("expected error when no ID provided")
	}
}

func TestCustomersUpdateCmd_RequiresArg(t *testing.T) {
	err := customersUpdateCmd.Args(customersUpdateCmd, []string{})
	if err == nil {
		t.Fatal("expected error when no ID provided")
	}
}

func TestCustomersGetCmd_InvalidID(t *testing.T) {
	err := customersGetCmd.RunE(customersGetCmd, []string{"not-a-number"})
	if err == nil {
		t.Fatal("expected error for invalid ID")
	}
	if !strings.Contains(err.Error(), "invalid customer ID") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCustomersGetCmd_NegativeID(t *testing.T) {
	err := customersGetCmd.RunE(customersGetCmd, []string{"-1"})
	if err == nil {
		t.Fatal("expected error for negative ID")
	}
	if !strings.Contains(err.Error(), "must be a positive integer") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCustomersCmd_RegisteredOnRoot(t *testing.T) {
	found := false
	for _, c := range rootCmd.Commands() {
		if c.Name() == "customers" {
			found = true
			break
		}
	}
	if !found {
		t.Error("customers command not registered on root")
	}
}
