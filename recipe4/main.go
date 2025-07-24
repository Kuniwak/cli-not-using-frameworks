package main

import (
	"github.com/Kuniwak/cli-not-using-frameworks/recipe4/cli"
	"github.com/Kuniwak/cli-not-using-frameworks/recipe4/recipe4"
)

func main() {
	cli.Run(recipe4.MainCommandByArgs)
}
