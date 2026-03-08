package cmd

import (
	"strings"
	"testing"
)

func TestPromotionsCmd_HasSubcommands(t *testing.T) {
	subs := promotionsCmd.Commands()
	want := map[string]bool{
		"list":   false,
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

func TestPromotionsUpdateCmd_RequiresArg(t *testing.T) {
	if err := promotionsUpdateCmd.Args(promotionsUpdateCmd, []string{}); err == nil {
		t.Fatal("expected error when no ID provided")
	}
}

func TestPromotionsDeleteCmd_RequiresArg(t *testing.T) {
	if err := promotionsDeleteCmd.Args(promotionsDeleteCmd, []string{}); err == nil {
		t.Fatal("expected error when no ID provided")
	}
}

func TestPromotionsDeleteCmd_InvalidID(t *testing.T) {
	err := promotionsDeleteCmd.RunE(promotionsDeleteCmd, []string{"abc"})
	if err == nil || !strings.Contains(err.Error(), "invalid promotion ID") {
		t.Errorf("expected invalid ID error, got: %v", err)
	}
}

func TestPromotionsCmd_RegisteredOnRoot(t *testing.T) {
	for _, c := range rootCmd.Commands() {
		if c.Name() == "promotions" {
			return
		}
	}
	t.Error("promotions command not registered on root")
}
