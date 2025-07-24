package recipe5

import (
	"bytes"
	"strings"
	"testing"

	"github.com/Kuniwak/cli-not-using-frameworks/recipe5/cli"
)

func TestMainCommandHelp(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	_ = MainCommand([]string{"-h"}, &cli.ProcInout{
		Stdin:  strings.NewReader(""),
		Stdout: stdout,
		Stderr: stderr,
	})

	expected := `Usage: recipe5 [command]

COMMANDS
  foo
    	foo command
  bar
    	bar command
`

	if stderr.String() != expected {
		t.Errorf("want %q, got %q", expected, stderr.String())
	}
}

func TestMainCommandFoo(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	exitStatus := MainCommand([]string{"foo"}, &cli.ProcInout{
		Stdin:  strings.NewReader(""),
		Stdout: stdout,
		Stderr: stderr,
	})

	if exitStatus != 0 {
		t.Errorf("want exit status 0, got %d", exitStatus)
	}
}

func TestMainCommandBar(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	exitStatus := MainCommand([]string{"bar"}, &cli.ProcInout{
		Stdin:  strings.NewReader(""),
		Stdout: stdout,
		Stderr: stderr,
	})

	if exitStatus != 0 {
		t.Errorf("want exit status 0, got %d", exitStatus)
	}
}
