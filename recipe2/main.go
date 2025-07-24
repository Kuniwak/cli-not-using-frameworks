package main

import (
	"github.com/Kuniwak/cli-not-using-frameworks/recipe2/cli"
	"github.com/Kuniwak/cli-not-using-frameworks/recipe2/recipe2"
)

func main() {
	cli.Run(recipe2.MainCommandByArgs)
}
