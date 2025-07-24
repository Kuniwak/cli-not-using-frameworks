package recipe2

import (
	"fmt"

	"github.com/Kuniwak/cli-not-using-frameworks/recipe2/cli"
)

func MainCommandByArgs(args []string, inout *cli.ProcInout) int {
	opts, err := ParseOptions(args, inout)
	if err != nil {
		fmt.Fprintf(inout.Stderr, "error: %v\n", err)
		return 1
	}

	if err := MainCommandByOptions(opts, inout); err != nil {
		fmt.Fprintf(inout.Stderr, "error: %v\n", err)
		return 1
	}

	return 0
}

func MainCommandByOptions(opts *Options, inout *cli.ProcInout) error {
	fmt.Fprintf(inout.Stdout, "foo: %s\n", opts.Foo)
	fmt.Fprintf(inout.Stdout, "bar: %s\n", opts.Bar)
	return nil
}
