package recipe2

import (
	"bytes"
	"strings"
	"testing"

	"github.com/Kuniwak/cli-not-using-frameworks/recipe2/cli"
)

func TestMainCommand(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	err := MainCommandByOptions(&Options{
		Foo: "foo",
		Bar: "bar",
	}, &cli.ProcInout{
		Stdin:  strings.NewReader(""),
		Stdout: stdout,
		Stderr: stderr,
	})
	if err != nil {
		t.Fatalf("failed to run task: %v", err)
	}

	expected := "foo: foo\nbar: bar\n"
	if stdout.String() != expected {
		t.Errorf("want %q, got %q", expected, stdout.String())
	}
}
