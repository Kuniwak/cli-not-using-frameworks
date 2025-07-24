package recipe1

import (
	"bufio"
	"fmt"
	"github.com/Kuniwak/cli-not-using-frameworks/recipe1/cli"
)

func MainCommand(_ []string, inout *cli.ProcInout) int {
	fmt.Fprintln(inout.Stdout, "Hello, World!")
	return 0
}

func InteractiveCommand(_ []string, inout *cli.ProcInout) int {
	scanner := bufio.NewScanner(inout.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Fprintf(inout.Stdout, "Hello, %s!\n", line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(inout.Stderr, err)
		return 1
	}

	return 0
}
