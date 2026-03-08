package staff

import "testing"

func TestCmd_HasSubcommands(t *testing.T) {
	subs := Cmd.Commands()
	want := map[string]bool{"list": false, "get": false, "create": false, "update": false, "delete": false}
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

func TestGetCmd_RequiresArg(t *testing.T) {
	if err := getCmd.Args(getCmd, []string{}); err == nil {
		t.Fatal("expected error when no ID provided")
	}
}

func TestDeleteCmd_RequiresArg(t *testing.T) {
	if err := deleteCmd.Args(deleteCmd, []string{}); err == nil {
		t.Fatal("expected error when no ID provided")
	}
}

func TestUpdateCmd_RequiresArg(t *testing.T) {
	if err := updateCmd.Args(updateCmd, []string{}); err == nil {
		t.Fatal("expected error when no ID provided")
	}
}
