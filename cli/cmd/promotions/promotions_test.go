package promotions

import (
	"strings"
	"testing"
)

func TestCmd_HasSubcommands(t *testing.T) {
	subs := Cmd.Commands()
	want := map[string]bool{"list": false, "create": false, "update": false, "delete": false}
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

func TestUpdateCmd_RequiresArg(t *testing.T) {
	if err := updateCmd.Args(updateCmd, []string{}); err == nil {
		t.Fatal("expected error when no ID provided")
	}
}

func TestDeleteCmd_RequiresArg(t *testing.T) {
	if err := deleteCmd.Args(deleteCmd, []string{}); err == nil {
		t.Fatal("expected error when no ID provided")
	}
}

func TestDeleteCmd_InvalidID(t *testing.T) {
	err := deleteCmd.RunE(deleteCmd, []string{"abc"})
	if err == nil || !strings.Contains(err.Error(), "invalid promotion ID") {
		t.Errorf("expected invalid ID error, got: %v", err)
	}
}

func TestUpdateCmd_InvalidID(t *testing.T) {
	err := updateCmd.RunE(updateCmd, []string{"abc"})
	if err == nil || !strings.Contains(err.Error(), "invalid promotion ID") {
		t.Errorf("expected invalid ID error, got: %v", err)
	}
}
