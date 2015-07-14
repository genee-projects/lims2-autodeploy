package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
)

func main() {
	usage := `Lims2 Autodeploy.

Usage:
  lims2 <command> [<args>...]
  lims2 -h | --help
  lims2 --version

Commands:
  up                   Create and start containers
  get-cron             Output crontab
  get-sphinx           Output sphinxsearch config
  update-cron          Output to /etc/cron.d/lims2 in container.
    
Options:
  -h --help     Show this screen.
  --version     Show version.`

	arguments, _ := docopt.Parse(usage, nil, true, "Lims2 Autodeploy 0.1", false)
	fmt.Println(arguments)
}
