package recipe4

import (
	"errors"
	"flag"

	"github.com/Kuniwak/cli-not-using-frameworks/recipe4/cli"
)

type Options struct {
	Foo  string
	Bar  string
	Help bool
}

func ParseOptions(args []string, inout *cli.ProcInout) (*Options, error) {
	flags := flag.NewFlagSet("recipe4", flag.ContinueOnError)
	flags.SetOutput(inout.Stderr)

	options := &Options{}
	flags.StringVar(&options.Foo, "foo", "", "foo")
	flags.StringVar(&options.Bar, "bar", "", "bar")

	flags.Usage = func() {
		inout.Stderr.Write([]byte("Usage: recipe4 [options]\n"))
		inout.Stderr.Write([]byte("OPTIONS\n"))
		flags.PrintDefaults()
	}

	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			options.Help = true
			return options, nil
		}
		return nil, err
	}

	return options, nil
}
