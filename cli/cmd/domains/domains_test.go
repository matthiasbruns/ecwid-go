package domains

import "testing"

func TestCmd_HasSubcommands(t *testing.T) {
	subs := Cmd.Commands()
	want := map[string]bool{"get": false, "purchase": false}
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

func TestPurchaseCmd_HasFileFlag(t *testing.T) {
	if purchaseCmd.Flags().Lookup("file") == nil {
		t.Error("purchase command missing --file flag")
	}
}
