package main

import (
	"os"

	"github.com/docopt/docopt-go"
)

var cmdUp = &Command{}

func init() {
	cmdUp.Run = up
}

func up(cmd *Command) int {

	usage := `
Usage: lims2 up [--all] [COMMAND]

Options:
  --all              Create and run all containers, not only lims2.`

	var defaultArgs []string

	if len(os.Args) > 3 {
		defaultArgs = []string{
			"-h",
		}
	} else {
		defaultArgs = os.Args[1:]
	}

	args, _ := docopt.Parse(usage, defaultArgs, true, "Lims2 Autodeploy 0.1", false)

	return 0
}
