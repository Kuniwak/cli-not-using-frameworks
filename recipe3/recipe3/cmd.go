package recipe3

import (
	"fmt"

	"github.com/Kuniwak/cli-not-using-frameworks/recipe3/cli"
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
	fmt.Fprintf(inout.Stdout, "something-required: %s\n", opts.SomethingRequired)
	return nil
}
