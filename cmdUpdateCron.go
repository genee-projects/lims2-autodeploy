package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/docopt/docopt-go"
)

var cmdUpdateCron = &Command{}

func init() {
	cmdUpdateCron.Run = updateCron
}

func updateCron(cmd *Command) int {

	usage := `
Usage: lims2 update-cron`

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
		"bash -c 'php /usr/share/lims2/cli/get_all_cron.php -u=www-data /usr/share/lims2 > /etc/cron.d/lims2'",
	}

	c, err := exec.Command("docker", strings.Join(args, " ")).Output()

	if err != nil {
		panic(err)
	}

	fmt.Println(string(c))

	return 0
}
