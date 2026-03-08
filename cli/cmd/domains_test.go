package cmd

import (
	"testing"
)

func TestDomainsCmd_HasSubcommands(t *testing.T) {
	subs := domainsCmd.Commands()
	want := map[string]bool{
		"get":      false,
		"purchase": false,
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

func TestDomainsCmd_RegisteredOnRoot(t *testing.T) {
	for _, c := range rootCmd.Commands() {
		if c.Name() == "domains" {
			return
		}
	}
	t.Error("domains command not registered on root")
}

func TestDomainsPurchaseCmd_HasFileFlag(t *testing.T) {
	if domainsPurchaseCmd.Flags().Lookup("file") == nil {
		t.Error("purchase command missing --file flag")
	}
}
