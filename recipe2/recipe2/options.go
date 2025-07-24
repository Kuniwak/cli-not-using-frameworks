package recipe2

import (
	"errors"
	"flag"

	"github.com/Kuniwak/cli-not-using-frameworks/recipe2/cli"
)

type Options struct {
	Foo  string
	Bar  string
	Help bool
}

func ParseOptions(args []string, inout *cli.ProcInout) (*Options, error) {
	flags := flag.NewFlagSet("recipe2", flag.ContinueOnError)
	flags.SetOutput(inout.Stderr)

	options := &Options{}
	flags.StringVar(&options.Foo, "foo", "", "Foo")
	flags.StringVar(&options.Bar, "bar", "", "Bar")

	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			options.Help = true
			return options, nil
		}
		return nil, err
	}

	return options, nil
}
