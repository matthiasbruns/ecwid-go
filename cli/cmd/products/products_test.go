package products

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
	flags := []string{"keyword", "category", "limit", "offset", "enabled", "in-stock", "sku", "sort-by"}
	for _, f := range flags {
		if listCmd.Flags().Lookup(f) == nil {
			t.Errorf("missing flag: --%s", f)
		}
	}
}

func TestGetCmd_RequiresArg(t *testing.T) {
	err := getCmd.Args(getCmd, []string{})
	if err == nil {
		t.Fatal("expected error when no ID provided")
	}
}

func TestDeleteCmd_RequiresArg(t *testing.T) {
	err := deleteCmd.Args(deleteCmd, []string{})
	if err == nil {
		t.Fatal("expected error when no ID provided")
	}
}

func TestUpdateCmd_RequiresArg(t *testing.T) {
	err := updateCmd.Args(updateCmd, []string{})
	if err == nil {
		t.Fatal("expected error when no ID provided")
	}
}

func TestCreateCmd_HasFileFlag(t *testing.T) {
	if createCmd.Flags().Lookup("file") == nil {
		t.Error("create command missing --file flag")
	}
}

func TestUpdateCmd_HasFileFlag(t *testing.T) {
	if updateCmd.Flags().Lookup("file") == nil {
		t.Error("update command missing --file flag")
	}
}

func TestGetCmd_InvalidID(t *testing.T) {
	err := getCmd.RunE(getCmd, []string{"not-a-number"})
	if err == nil {
		t.Fatal("expected error for invalid ID")
	}
	if !strings.Contains(err.Error(), "invalid product ID") {
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

func TestGetCmd_ZeroID(t *testing.T) {
	err := getCmd.RunE(getCmd, []string{"0"})
	if err == nil {
		t.Fatal("expected error for zero ID")
	}
	if !strings.Contains(err.Error(), "must be a positive integer") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestListCmd_RejectsArgs(t *testing.T) {
	err := listCmd.Args(listCmd, []string{"extra"})
	if err == nil {
		t.Fatal("expected error for extra args")
	}
}

func TestCreateCmd_RejectsArgs(t *testing.T) {
	err := createCmd.Args(createCmd, []string{"extra"})
	if err == nil {
		t.Fatal("expected error for extra args")
	}
}
