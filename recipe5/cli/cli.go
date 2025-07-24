package cli

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type ProcInout struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

func NewProcInout() *ProcInout {
	return &ProcInout{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
}

type Command func(args []string, inout *ProcInout) int

func Run(c Command) {
	args := os.Args[1:]
	exitStatus := c(args, NewProcInout())
	os.Exit(exitStatus)
}

type SubCommand struct {
	Name        string
	Description string
	Run         func(args []string, inout *ProcInout) int
}

func NewCommand(name string, cmds []SubCommand) Command {
	return func(args []string, inout *ProcInout) int {
		flags := flag.NewFlagSet(name, flag.ContinueOnError)
		flags.SetOutput(inout.Stderr)

		flags.Usage = func() {
			fmt.Fprintf(inout.Stderr, "Usage: %s [command]\n\n", name)

			fmt.Fprintf(inout.Stderr, "COMMANDS\n")
			for _, cmd := range cmds {
				fmt.Fprintf(inout.Stderr, "  %s\n    \t%s\n", cmd.Name, cmd.Description)
			}
		}

		if err := flags.Parse(args); err != nil {
			if err == flag.ErrHelp {
				return 0
			}
			return 1
		}

		if flags.NArg() == 0 {
			fmt.Fprintf(inout.Stderr, "error: no command provided\n")
			flags.Usage()
			return 1
		}

		for _, cmd := range cmds {
			if cmd.Name == flags.Arg(0) {
				return cmd.Run(flags.Args()[1:], inout)
			}
		}

		fmt.Fprintf(inout.Stderr, "error: unknown command: %s\n", flags.Arg(0))
		flags.Usage()
		return 1
	}
}
