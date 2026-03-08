package cmd

import (
	"testing"
)

func TestProfileCmd_HasSubcommands(t *testing.T) {
	subs := profileCmd.Commands()
	want := map[string]bool{
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

func TestProfileCmd_RegisteredOnRoot(t *testing.T) {
	for _, c := range rootCmd.Commands() {
		if c.Name() == "profile" {
			return
		}
	}
	t.Error("profile command not registered on root")
}

func TestProfileUpdateCmd_HasFileFlag(t *testing.T) {
	if profileUpdateCmd.Flags().Lookup("file") == nil {
		t.Error("update command missing --file flag")
	}
}
