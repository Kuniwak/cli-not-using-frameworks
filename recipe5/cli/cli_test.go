package cli

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestNewCommand(t *testing.T) {
	s1 := SubCommand{
		Name:        "one",
		Description: "1st subcommand",
		Run: func(args []string, inout *ProcInout) int {
			fmt.Fprintln(inout.Stdout, "1")
			return 0
		},
	}
	s2 := SubCommand{
		Name:        "two",
		Description: "2nd subcommand",
		Run: func(args []string, inout *ProcInout) int {
			fmt.Fprintln(inout.Stdout, "2")
			return 0
		},
	}

	stdin := strings.NewReader("")
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	cmd := NewCommand("test", []SubCommand{s1, s2})
	exitStatus := cmd([]string{"two"}, &ProcInout{
		Stdin:  stdin,
		Stdout: stdout,
		Stderr: stderr,
	})

	if exitStatus != 0 {
		t.Errorf("expected exit status 0, got %d", exitStatus)
	}

	if stdout.String() != "2\n" {
		t.Errorf("expected output '2', got %q", stdout.String())
	}
}
