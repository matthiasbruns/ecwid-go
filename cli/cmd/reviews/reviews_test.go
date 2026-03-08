package reviews

import (
	"strings"
	"testing"
)

func TestCmd_HasSubcommands(t *testing.T) {
	subs := Cmd.Commands()
	want := map[string]bool{"list": false, "update": false, "delete": false}
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
	for _, f := range []string{"limit", "offset", "status", "product-id"} {
		if listCmd.Flags().Lookup(f) == nil {
			t.Errorf("missing flag: --%s", f)
		}
	}
}

func TestDeleteCmd_RequiresArg(t *testing.T) {
	if err := deleteCmd.Args(deleteCmd, []string{}); err == nil {
		t.Fatal("expected error when no ID provided")
	}
}

func TestDeleteCmd_InvalidID(t *testing.T) {
	err := deleteCmd.RunE(deleteCmd, []string{"abc"})
	if err == nil || !strings.Contains(err.Error(), "invalid review ID") {
		t.Errorf("expected invalid ID error, got: %v", err)
	}
}

func TestUpdateCmd_RequiresArg(t *testing.T) {
	if err := updateCmd.Args(updateCmd, []string{}); err == nil {
		t.Fatal("expected error when no ID provided")
	}
}

func TestUpdateCmd_InvalidID(t *testing.T) {
	err := updateCmd.RunE(updateCmd, []string{"0"})
	if err == nil || !strings.Contains(err.Error(), "invalid review ID") {
		t.Errorf("expected invalid ID error, got: %v", err)
	}
}

func TestUpdateCmd_MissingStatus(t *testing.T) {
	err := updateCmd.RunE(updateCmd, []string{"1"})
	if err == nil || !strings.Contains(err.Error(), "--status flag is required") {
		t.Errorf("expected missing status error, got: %v", err)
	}
}
