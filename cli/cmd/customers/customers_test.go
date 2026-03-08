package customers

import (
	"strings"
	"testing"
)

func TestCmd_HasSubcommands(t *testing.T) {
	subs := Cmd.Commands()
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

func TestListCmd_Flags(t *testing.T) {
	flags := []string{"keyword", "email", "limit", "offset"}
	for _, f := range flags {
		if listCmd.Flags().Lookup(f) == nil {
			t.Errorf("missing flag: --%s", f)
		}
	}
}

func TestGetCmd_RequiresArg(t *testing.T) {
	if err := getCmd.Args(getCmd, []string{}); err == nil {
		t.Fatal("expected error when no ID provided")
	}
}

func TestDeleteCmd_RequiresArg(t *testing.T) {
	if err := deleteCmd.Args(deleteCmd, []string{}); err == nil {
		t.Fatal("expected error when no ID provided")
	}
}

func TestUpdateCmd_RequiresArg(t *testing.T) {
	if err := updateCmd.Args(updateCmd, []string{}); err == nil {
		t.Fatal("expected error when no ID provided")
	}
}

func TestGetCmd_InvalidID(t *testing.T) {
	err := getCmd.RunE(getCmd, []string{"not-a-number"})
	if err == nil {
		t.Fatal("expected error for invalid ID")
	}
	if !strings.Contains(err.Error(), "invalid customer ID") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestGetCmd_NegativeID(t *testing.T) {
	err := getCmd.RunE(getCmd, []string{"-1"})
	if err == nil {
		t.Fatal("expected error for negative ID")
	}
	if !strings.Contains(err.Error(), "must be a positive integer") {
		t.Errorf("unexpected error: %v", err)
	}
}
