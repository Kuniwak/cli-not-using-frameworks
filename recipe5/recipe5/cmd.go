package recipe5

import (
	"fmt"

	"github.com/Kuniwak/cli-not-using-frameworks/recipe5/cli"
)

var MainCommand = cli.NewCommand("recipe5", []cli.SubCommand{
	{
		Name:        "foo",
		Description: "foo command",
		Run: func(args []string, inout *cli.ProcInout) int {
			fmt.Fprintf(inout.Stdout, "foo command executed successfully\n")
			return 0
		},
	},
	{
		Name:        "bar",
		Description: "bar command",
		Run: func(args []string, inout *cli.ProcInout) int {
			fmt.Fprintf(inout.Stdout, "bar command executed successfully\n")
			return 0
		},
	},
})
