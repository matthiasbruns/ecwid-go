package cmd

import (
	"strings"
	"testing"
)

func TestCouponsCmd_HasSubcommands(t *testing.T) {
	subs := couponsCmd.Commands()
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

func TestCouponsListCmd_Flags(t *testing.T) {
	flags := []string{"limit", "offset", "code"}
	for _, f := range flags {
		if couponsListCmd.Flags().Lookup(f) == nil {
			t.Errorf("missing flag: --%s", f)
		}
	}
}

func TestCouponsGetCmd_RequiresArg(t *testing.T) {
	if err := couponsGetCmd.Args(couponsGetCmd, []string{}); err == nil {
		t.Fatal("expected error when no ID provided")
	}
}

func TestCouponsGetCmd_InvalidID(t *testing.T) {
	err := couponsGetCmd.RunE(couponsGetCmd, []string{"abc"})
	if err == nil || !strings.Contains(err.Error(), "invalid coupon ID") {
		t.Errorf("expected invalid ID error, got: %v", err)
	}
}

func TestCouponsCmd_RegisteredOnRoot(t *testing.T) {
	for _, c := range rootCmd.Commands() {
		if c.Name() == "coupons" {
			return
		}
	}
	t.Error("coupons command not registered on root")
}
