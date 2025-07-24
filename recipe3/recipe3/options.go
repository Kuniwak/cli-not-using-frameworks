package recipe3

import (
	"errors"
	"flag"

	"github.com/Kuniwak/cli-not-using-frameworks/recipe3/cli"
)

type Options struct {
	SomethingRequired string
	Help              bool
}

func ParseOptions(args []string, inout *cli.ProcInout) (*Options, error) {
	flags := flag.NewFlagSet("recipe3", flag.ContinueOnError)
	flags.SetOutput(inout.Stderr)

	options := &Options{}
	flags.StringVar(&options.SomethingRequired, "something-required", "", "Something required")

	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			options.Help = true
			return options, nil
		}
		return nil, err
	}

	if options.SomethingRequired == "" {
		return nil, errors.New("something-required is required")
	}

	return options, nil
}
