package profile

import "testing"

func TestCmd_HasSubcommands(t *testing.T) {
	subs := Cmd.Commands()
	want := map[string]bool{"get": false, "update": false}
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

func TestUpdateCmd_HasFileFlag(t *testing.T) {
	if updateCmd.Flags().Lookup("file") == nil {
		t.Error("update command missing --file flag")
	}
}
