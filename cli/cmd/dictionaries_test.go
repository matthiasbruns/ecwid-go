package cmd

import (
	"testing"
)

func TestDictionariesCmd_HasSubcommands(t *testing.T) {
	subs := dictionariesCmd.Commands()
	want := map[string]bool{
		"countries":   false,
		"currencies":  false,
		"states":      false,
		"tax-classes": false,
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

func TestDictionariesCountriesCmd_Flags(t *testing.T) {
	flags := []string{"lang", "with-states"}
	for _, f := range flags {
		if dictionariesCountriesCmd.Flags().Lookup(f) == nil {
			t.Errorf("missing flag: --%s", f)
		}
	}
}

func TestDictionariesStatesCmd_Flags(t *testing.T) {
	flags := []string{"country", "lang"}
	for _, f := range flags {
		if dictionariesStatesCmd.Flags().Lookup(f) == nil {
			t.Errorf("missing flag: --%s", f)
		}
	}
}

func TestDictionariesCmd_RegisteredOnRoot(t *testing.T) {
	for _, c := range rootCmd.Commands() {
		if c.Name() == "dictionaries" {
			return
		}
	}
	t.Error("dictionaries command not registered on root")
}
