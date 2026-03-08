package reports

import "testing"

func TestCmd_HasSubcommands(t *testing.T) {
	subs := Cmd.Commands()
	want := map[string]bool{"get": false, "stats": false}
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
		t.Fatal("expected error when no report type provided")
	}
}

func TestGetCmd_Flags(t *testing.T) {
	for _, f := range []string{"started-from", "ended-at", "time-scale", "compare-period"} {
		if getCmd.Flags().Lookup(f) == nil {
			t.Errorf("missing flag: --%s", f)
		}
	}
}
