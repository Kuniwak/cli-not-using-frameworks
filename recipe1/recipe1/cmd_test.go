package recipe1

import (
	"bytes"
	"strings"
	"testing"

	"github.com/Kuniwak/cli-not-using-frameworks/recipe1/cli"
)

func TestMainCommand(t *testing.T) {
	stdin := strings.NewReader("")
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	exitStatus := MainCommand([]string{}, &cli.ProcInout{
		Stdin:  stdin,
		Stdout: stdout,
		Stderr: stderr,
	})

	if exitStatus != 0 {
		t.Errorf("unexpected exit status: %d", exitStatus)
	}

	expected := "Hello, World!\n"
	if stdout.String() != expected {
		t.Errorf("want %q, got %q", expected, stdout.String())
	}
}

func TestInteractiveCommand(t *testing.T) {
	stdin := strings.NewReader("foo\nbar\n")
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	exitStatus := InteractiveCommand([]string{}, &cli.ProcInout{
		Stdin:  stdin,
		Stdout: stdout,
		Stderr: stderr,
	})

	if exitStatus != 0 {
		t.Errorf("unexpected exit status: %d", exitStatus)
	}

	expected := "Hello, foo!\nHello, bar!\n"
	if stdout.String() != expected {
		t.Errorf("want %q, got %q", expected, stdout.String())
	}
}
