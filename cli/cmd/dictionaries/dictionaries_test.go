package dictionaries

import "testing"

func TestCmd_HasSubcommands(t *testing.T) {
	subs := Cmd.Commands()
	want := map[string]bool{"countries": false, "currencies": false, "states": false, "tax-classes": false}
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

func TestCountriesCmd_Flags(t *testing.T) {
	for _, f := range []string{"lang", "with-states"} {
		if countriesCmd.Flags().Lookup(f) == nil {
			t.Errorf("missing flag: --%s", f)
		}
	}
}

func TestCurrenciesCmd_Flags(t *testing.T) {
	if currenciesCmd.Flags().Lookup("lang") == nil {
		t.Error("missing flag: --lang")
	}
}

func TestStatesCmd_Flags(t *testing.T) {
	for _, f := range []string{"country", "lang"} {
		if statesCmd.Flags().Lookup(f) == nil {
			t.Errorf("missing flag: --%s", f)
		}
	}
}

func TestTaxClassesCmd_Flags(t *testing.T) {
	for _, f := range []string{"country", "lang"} {
		if taxClassesCmd.Flags().Lookup(f) == nil {
			t.Errorf("missing flag: --%s", f)
		}
	}
}
