package cmd

import (
	"strings"
	"testing"
)

func TestExamplesAvoidKnownFailingQueries(t *testing.T) {
	broken := map[string]string{
		"shop":         `--keyword "居酒屋"`,
		"special list": "--category SPC0",
	}
	for name, bad := range broken {
		c, _, err := rootCmd.Find(strings.Fields(name))
		if err != nil {
			t.Fatal(err)
		}
		if strings.Contains(c.Example, bad) {
			t.Errorf("%s example uses %s, which the API rejects or matches nothing", name, bad)
		}
	}
}

func TestExamplesUseValidFlags(t *testing.T) {
	for _, c := range rootCmd.Commands() {
		for _, sub := range append(c.Commands(), c) {
			for _, line := range strings.Split(sub.Example, "\n") {
				fields := strings.Fields(line)
				if len(fields) < 2 || fields[0] != "hpp" {
					continue
				}
				found, args, err := rootCmd.Find(fields[1:])
				if err != nil {
					t.Fatalf("%q: %v", line, err)
				}
				if err := found.ParseFlags(args); err != nil {
					t.Errorf("%q: %v", line, err)
				}
			}
		}
	}
}
