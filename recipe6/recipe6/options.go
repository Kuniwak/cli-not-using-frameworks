package recipe6

import (
	"errors"
	"flag"
	"fmt"

	"github.com/Kuniwak/cli-not-using-frameworks/recipe6/cli"
)

type Options struct {
	Foo  string `name:"foo" description:"foo"`
	Bar  string `name:"bar" description:"bar"`
	Help bool   `name:"help" description:"show help"`
}

func ParseOptions(args []string, inout *cli.ProcInout) (*Options, error) {
	opts := &Options{}

	if err := cli.ParseFlags("recipe6", args, opts); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			opts.Help = true
			return opts, err
		}
		return nil, fmt.Errorf("failed to parse options: %w", err)
	}

	return opts, nil
}
