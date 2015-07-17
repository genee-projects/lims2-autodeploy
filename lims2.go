package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/docopt/docopt-go"
)

var (

	//所有命令
	commandsMap = map[interface{}]*Command{
		"up":          cmdUp,
		"get-cron":    cmdGetCron,
		"get-sphinx":  cmdGetSphinx,
		"update-cron": cmdUpdateCron,
	}

	//定义操作容器的名称
	containerName = "lims2"
)

type Command struct {
	command string
	args    map[string]interface{}
	Run     func(cmd *Command) int
}

func main() {
	usage := `Lims2 Autodeploy.

Usage:
  lims2 <command> [<args>...]
  lims2 -h | --help
  lims2 --version

Options:
  -h --help     Show this screen.
  --version     Show version.

Commands:
  up                   Create and start containers
  get-cron             Output crontab
  get-sphinx           Output sphinxsearch config
  update-cron          Output to /etc/cron.d/lims2 in container.`

	var defaultArgs []string
	if len(os.Args[1:]) == 0 {
		defaultArgs = []string{
			"-h",
		}
	} else {
		defaultArgs = os.Args[1:]
	}

	args, _ := docopt.Parse(usage, defaultArgs, true, "Lims2 Autodeploy 0.1", true)

	command := args["<command>"]

	subCommand, exits := commandsMap[command]

	//如果传值错误
	if !exits {
		docopt.Parse(usage, []string{"-h"}, true, "Lims2 Autodeploy 0.1", true)
	}

	defer func() {

		if r := recover(); r != nil {
			fmt.Println(r)
			os.Exit(1)
		}
	}()

	if _, err := os.Stat("main.yml"); err == nil {

	} else {
		panic(errors.New("main.yml not exists."))
	}

	subCommand.Run(subCommand)
}
