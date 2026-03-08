package cmd

import (
	"testing"
)

func TestReportsCmd_HasSubcommands(t *testing.T) {
	subs := reportsCmd.Commands()
	want := map[string]bool{
		"get":   false,
		"stats": false,
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

func TestReportsGetCmd_RequiresArg(t *testing.T) {
	if err := reportsGetCmd.Args(reportsGetCmd, []string{}); err == nil {
		t.Fatal("expected error when no report type provided")
	}
}

func TestReportsGetCmd_Flags(t *testing.T) {
	flags := []string{"started-from", "ended-at", "time-scale", "compare-period"}
	for _, f := range flags {
		if reportsGetCmd.Flags().Lookup(f) == nil {
			t.Errorf("missing flag: --%s", f)
		}
	}
}

func TestReportsCmd_RegisteredOnRoot(t *testing.T) {
	for _, c := range rootCmd.Commands() {
		if c.Name() == "reports" {
			return
		}
	}
	t.Error("reports command not registered on root")
}
