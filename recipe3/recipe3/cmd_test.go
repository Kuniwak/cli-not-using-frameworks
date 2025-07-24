package recipe3

import (
	"bytes"
	"strings"
	"testing"

	"github.com/Kuniwak/cli-not-using-frameworks/recipe3/cli"
)

func TestMainCommandByOptions(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	err := MainCommandByOptions(&Options{
		SomethingRequired: "something",
	}, &cli.ProcInout{
		Stdin:  strings.NewReader(""),
		Stdout: stdout,
		Stderr: stderr,
	})
	if err != nil {
		t.Fatalf("failed to run task: %v", err)
	}

	expected := "something-required: something\n"
	if stdout.String() != expected {
		t.Errorf("want %q, got %q", expected, stdout.String())
	}
}
