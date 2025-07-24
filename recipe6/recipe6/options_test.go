package recipe6

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"github.com/Kuniwak/cli-not-using-frameworks/recipe6/cli"
	"github.com/google/go-cmp/cmp"
)

func TestParseOptions(t *testing.T) {
	testCases := map[string]struct {
		Input    []string
		Expected *Options
	}{
		"--foo": {[]string{"--foo", "foo"}, &Options{Foo: "foo"}},
		"--bar": {[]string{"--bar", "bar"}, &Options{Bar: "bar"}},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}

			opts, err := ParseOptions(tc.Input, &cli.ProcInout{
				Stdin:  strings.NewReader(""),
				Stdout: stdout,
				Stderr: stderr,
			})
			if err != nil {
				t.Errorf("want no error, got %v", err)
			}

			if !reflect.DeepEqual(tc.Expected, opts) {
				t.Error(cmp.Diff(tc.Expected, opts))
			}
		})
	}
}
