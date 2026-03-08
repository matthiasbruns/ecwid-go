package coupons

import (
	"strings"
	"testing"
)

func TestCmd_HasSubcommands(t *testing.T) {
	subs := Cmd.Commands()
	want := map[string]bool{"list": false, "get": false, "create": false, "update": false, "delete": false}
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
	for _, f := range []string{"limit", "offset", "code"} {
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

func TestGetCmd_InvalidID(t *testing.T) {
	err := getCmd.RunE(getCmd, []string{"abc"})
	if err == nil || !strings.Contains(err.Error(), "invalid coupon ID") {
		t.Errorf("expected invalid ID error, got: %v", err)
	}
}

func TestUpdateCmd_InvalidID(t *testing.T) {
	err := updateCmd.RunE(updateCmd, []string{"abc"})
	if err == nil || !strings.Contains(err.Error(), "invalid coupon ID") {
		t.Errorf("expected invalid ID error, got: %v", err)
	}
}

func TestDeleteCmd_InvalidID(t *testing.T) {
	err := deleteCmd.RunE(deleteCmd, []string{"abc"})
	if err == nil || !strings.Contains(err.Error(), "invalid coupon ID") {
		t.Errorf("expected invalid ID error, got: %v", err)
	}
}
