package cli

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type options struct {
	Int    int    `name:"int" description:"int"`
	Bool   bool   `name:"bool" description:"bool"`
	String string `name:"string" description:"string"`
}

func TestAnalyzeFlags(t *testing.T) {
	opts := &options{}
	fs := AnalyzeFlags(opts)

	expected := []Flag{
		{Name: "int", Description: "int", Type: FlagTypeInt, FieldName: "Int"},
		{Name: "bool", Description: "bool", Type: FlagTypeBool, FieldName: "Bool"},
		{Name: "string", Description: "string", Type: FlagTypeString, FieldName: "String"},
	}

	if !reflect.DeepEqual(fs, expected) {
		t.Error(cmp.Diff(expected, fs))
	}
}

func TestNewCommandWithCompletion(t *testing.T) {
	testCases := map[string]struct {
		Input    []string
		Expected []string
	}{
		"":     {[]string{}, []string{"--int", "--bool", "--string"}},
		"--":   {[]string{"--"}, []string{"--int", "--bool", "--string"}},
		"--i":  {[]string{"--i"}, []string{"--int"}},
		"-i":   {[]string{"-i"}, []string{"--int"}},
		"--in": {[]string{"--in"}, []string{"--int"}},
		"-in":  {[]string{"-in"}, []string{"--int"}},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			opts := &options{}
			fs := AnalyzeFlags(opts)
			comps := Completion(tc.Input, fs)

			if !reflect.DeepEqual(comps, tc.Expected) {
				t.Error(cmp.Diff(tc.Expected, comps))
			}
		})
	}
}
