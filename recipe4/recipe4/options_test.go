package recipe4

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"github.com/Kuniwak/cli-not-using-frameworks/recipe4/cli"
	"github.com/google/go-cmp/cmp"
)

func TestParseOptions_Success(t *testing.T) {
	testCases := []struct {
		Input    []string
		Expected *Options
	}{
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

func TestFlagUsage(t *testing.T) {
	stdin := strings.NewReader("")
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	_, err := ParseOptions([]string{"-h"}, &cli.ProcInout{
		Stdin:  stdin,
		Stdout: stdout,
		Stderr: stderr,
	})
	if err != nil {
		t.Fatalf("failed to parse options: %v", err)
	}

	expected := `Usage: recipe4 [options]
OPTIONS
  -bar string
    	bar
  -foo string
    	foo
`
	if stderr.String() != expected {
		t.Errorf("want %q, got %q", expected, stderr.String())
	}
}
