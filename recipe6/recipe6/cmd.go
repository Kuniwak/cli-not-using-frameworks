package recipe6

import (
	"fmt"

	"github.com/Kuniwak/cli-not-using-frameworks/recipe6/cli"
)

func MainCommand(args []string, inout *cli.ProcInout) int {
	opts, err := ParseOptions(args, inout)
	if err != nil {
		return 1
	}

	if opts.Help {
		return 0
	}

	fmt.Fprintf(inout.Stdout, "foo: %s\n", opts.Foo)
	fmt.Fprintf(inout.Stdout, "bar: %s\n", opts.Bar)

	return 0
}
