package cmd

import (
	"testing"
)

func TestStaffCmd_HasSubcommands(t *testing.T) {
	subs := staffCmd.Commands()
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

func TestStaffGetCmd_RequiresArg(t *testing.T) {
	if err := staffGetCmd.Args(staffGetCmd, []string{}); err == nil {
		t.Fatal("expected error when no ID provided")
	}
}

func TestStaffDeleteCmd_RequiresArg(t *testing.T) {
	if err := staffDeleteCmd.Args(staffDeleteCmd, []string{}); err == nil {
		t.Fatal("expected error when no ID provided")
	}
}

func TestStaffCmd_RegisteredOnRoot(t *testing.T) {
	for _, c := range rootCmd.Commands() {
		if c.Name() == "staff" {
			return
		}
	}
	t.Error("staff command not registered on root")
}
