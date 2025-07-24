package recipe2

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"github.com/Kuniwak/cli-not-using-frameworks/recipe2/cli"
	"github.com/google/go-cmp/cmp"
)

func TestParseOptions(t *testing.T) {
	testCases := []struct {
		Input    []string
		Expected *Options
	}{
		{
			Input: []string{"-foo", "foo", "-bar", "bar"},
			Expected: &Options{
				Foo: "foo",
				Bar: "bar",
			},
		},
		{
			Input: []string{"-h"},
			Expected: &Options{
				Help: true,
			},
		},
	}

	for _, testCase := range testCases {
		stdout := &bytes.Buffer{}
		stderr := &bytes.Buffer{}

		opts, err := ParseOptions(testCase.Input, &cli.ProcInout{
			Stdin:  strings.NewReader(""),
			Stdout: stdout,
			Stderr: stderr,
		})
		if err != nil {
			t.Fatalf("failed to parse options: %v", err)
		}

		if !reflect.DeepEqual(opts, testCase.Expected) {
			t.Error(cmp.Diff(opts, testCase.Expected))
		}
	}
}
