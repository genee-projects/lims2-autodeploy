package main

import (
	"os"
	"os/exec"
	"strings"

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

	//无需进行内容获取, 错误会被 lims2.go 中 defer recover 获取到, arguments 会自动匹配, 异常不会执行
	docopt.Parse(usage, defaultArgs, true, "Lims2 Autodeploy 0.1", false)

	args := []string{
		"exec",
		containerName,
		"php",
		"/usr/share/lims2/cli/get_all_sphinx.php",
		"/usr/share/lims2",
	}

	c := exec.Command("docker", strings.Join(args, " "))

	c.Run()

	return 0
}
