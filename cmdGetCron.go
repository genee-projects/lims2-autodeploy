package main

import (
	"fmt"
	"os"

	"github.com/docopt/docopt-go"
)

var cmdGetCron = &Command{}

func init() {
	cmdGetCron.Run = getCron
}

func getCron(cmd *Command) int {

	usage := `
Usage: lims2 get-cron`

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
