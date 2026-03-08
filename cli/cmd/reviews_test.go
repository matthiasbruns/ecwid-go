package cmd

import (
	"strings"
	"testing"
)

func TestReviewsCmd_HasSubcommands(t *testing.T) {
	subs := reviewsCmd.Commands()
	want := map[string]bool{
		"list":   false,
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

func TestReviewsListCmd_Flags(t *testing.T) {
	flags := []string{"limit", "offset", "status", "product-id"}
	for _, f := range flags {
		if reviewsListCmd.Flags().Lookup(f) == nil {
			t.Errorf("missing flag: --%s", f)
		}
	}
}

func TestReviewsDeleteCmd_RequiresArg(t *testing.T) {
	if err := reviewsDeleteCmd.Args(reviewsDeleteCmd, []string{}); err == nil {
		t.Fatal("expected error when no ID provided")
	}
}

func TestReviewsDeleteCmd_InvalidID(t *testing.T) {
	err := reviewsDeleteCmd.RunE(reviewsDeleteCmd, []string{"abc"})
	if err == nil || !strings.Contains(err.Error(), "invalid review ID") {
		t.Errorf("expected invalid ID error, got: %v", err)
	}
}

func TestReviewsCmd_RegisteredOnRoot(t *testing.T) {
	for _, c := range rootCmd.Commands() {
		if c.Name() == "reviews" {
			return
		}
	}
	t.Error("reviews command not registered on root")
}
