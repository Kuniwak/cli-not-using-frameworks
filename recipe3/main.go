package main

import (
	"github.com/Kuniwak/cli-not-using-frameworks/recipe3/cli"
	"github.com/Kuniwak/cli-not-using-frameworks/recipe3/recipe3"
)

func main() {
	cli.Run(recipe3.MainCommandByArgs)
}
