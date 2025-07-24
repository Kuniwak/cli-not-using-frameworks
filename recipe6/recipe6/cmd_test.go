package recipe6

import (
	"bytes"
	"strings"
	"testing"

	"github.com/Kuniwak/cli-not-using-frameworks/recipe6/cli"
)

func TestMainCommand(t *testing.T) {
	testCases := map[string]struct {
		Input    []string
		Expected string
	}{
		"--foo": {[]string{"--foo", "foo"}, "foo"},
		"--bar": {[]string{"--bar", "bar"}, "bar"},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}

			exitStatus := MainCommand(tc.Input, &cli.ProcInout{
				Stdin:  strings.NewReader(""),
				Stdout: stdout,
				Stderr: stderr,
				Env:    make(map[string]string),
			})

			if exitStatus != 0 {
				t.Errorf("want exit status 0, got %d", exitStatus)
			}
		})
	}
}
