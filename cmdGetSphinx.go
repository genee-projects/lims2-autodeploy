package main

import (
	"fmt"
	"os"

	"github.com/docopt/docopt-go"
)

var cmdGetSphinx = &Command{}

func init() {
	cmdGetSphinx.Run = getSphinx
}

func getSphinx(cmd *Command) int {

	usage := `
Usage: lims2 get-sphinx`

	var defaultArgs []string

	if len(os.Args) > 2 {
		defaultArgs = []string{
			"-h",
		}
	} else {
		defaultArgs = os.Args[1:]
	}

	args, _ := docopt.Parse(usage, defaultArgs, true, "Lims2 Autodeploy 0.1", false)

	fmt.Println(args)

	//do something here
	return 0
}
