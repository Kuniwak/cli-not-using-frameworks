package main

import (
	"github.com/Kuniwak/cli-not-using-frameworks/recipe6/cli"
	"github.com/Kuniwak/cli-not-using-frameworks/recipe6/recipe6"
)

func main() {
	cli.Run(cli.NewCommandWithCompletion(recipe6.MainCommand, cli.NewCompletionByFlags(cli.AnalyzeFlags(&recipe6.Options{}))))
}
