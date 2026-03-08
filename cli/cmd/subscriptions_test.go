package cmd

import (
	"strings"
	"testing"
)

func TestSubscriptionsCmd_HasSubcommands(t *testing.T) {
	subs := subscriptionsCmd.Commands()
	want := map[string]bool{
		"list":   false,
		"get":    false,
		"update": false,
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

func TestSubscriptionsListCmd_Flags(t *testing.T) {
	flags := []string{"limit", "offset", "status", "customer-id", "product-id"}
	for _, f := range flags {
		if subscriptionsListCmd.Flags().Lookup(f) == nil {
			t.Errorf("missing flag: --%s", f)
		}
	}
}

func TestSubscriptionsGetCmd_RequiresArg(t *testing.T) {
	if err := subscriptionsGetCmd.Args(subscriptionsGetCmd, []string{}); err == nil {
		t.Fatal("expected error when no ID provided")
	}
}

func TestSubscriptionsGetCmd_InvalidID(t *testing.T) {
	err := subscriptionsGetCmd.RunE(subscriptionsGetCmd, []string{"abc"})
	if err == nil || !strings.Contains(err.Error(), "invalid subscription ID") {
		t.Errorf("expected invalid ID error, got: %v", err)
	}
}

func TestSubscriptionsCmd_RegisteredOnRoot(t *testing.T) {
	for _, c := range rootCmd.Commands() {
		if c.Name() == "subscriptions" {
			return
		}
	}
	t.Error("subscriptions command not registered on root")
}
