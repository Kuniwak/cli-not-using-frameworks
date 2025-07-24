package main

import (
	"github.com/Kuniwak/cli-not-using-frameworks/recipe1/cli"
	"github.com/Kuniwak/cli-not-using-frameworks/recipe1/recipe1"
)

func main() {
	cli.Run(recipe1.InteractiveCommand)
}
