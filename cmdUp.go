package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

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

	docopt.Parse(usage, defaultArgs, true, "Lims2 Autodeploy 0.1", false)

	//需要做的事情如下:
	// 1. 当前目录创建 volumes
	// 2. volumes/info.json 写入 main.yml 配置
	// 3. lims2 容器运行
	// 4. 其他服务运行

	// 1. 当前目录创建 volumes

	if _, err := os.Stat("/opt/lims2/volumes"); os.IsNotExist(err) {
		os.Mkdir("/opt/lims2/volumes", os.ModePerm)
	}

	// 2. volumes/info.json 写入 main.yml 配置
	// 为何需要配置一份 yaml, 配置一份 json
	// yaml 配置更易读? json 不易读

	//配置 revision
	var revision string

	if r := cmd.info.Get("revision"); r == nil {

		revision = "latest"
	}

	for name, _ := range cmd.info.Get("services").(map[interface{}]interface{}) {

		switch name {
		case "node-lims2":

			var nodeLims2 = map[string]interface{}{
				"host":      cmd.info.Get("services", "node-lims2", "host"),
				"port":      cmd.info.Get("services", "node-lims2", "port"),
				"salt":      cmd.info.Get("services", "node-lims2", "salt"),
				"rpc_token": cmd.info.Get("services", "node-lims2", "rpc_token"),
			}

			fmt.Println(nodeLims2)

		case "sphinxsearch":

			var sphinxSearch = map[string]interface{}{
				"host": cmd.info.Get("services", "sphinxsearch", "host"),
				"port": cmd.info.Get("services", "sphinxsearch", "port"),
			}
			fmt.Println(sphinxSearch)

		case "database":

			var database = map[string]interface{}{
				"host":     cmd.info.Get("services", "database", "host"),
				"password": cmd.info.Get("services", "database", "password"),
			}

			fmt.Println(database)
		case "redis":

			var redis = map[string]interface{}{
				"host": cmd.info.Get("services", "redis", "host"),
				"port": cmd.info.Get("services", "redis", "port"),
			}

			fmt.Println(redis)
		case "genee-updater":

			var geneeUpdater = map[string]interface{}{
				"site_url": cmd.info.Get("services", "genee-updater", "site_url"),
				"port":     cmd.info.Get("services", "genee-updater", "port"),
			}
			fmt.Println(geneeUpdater)
		case "glogon":
			var glogon = map[string]interface{}{
				"port": cmd.info.Get("services", "glogon", "port"),
				"host": cmd.info.Get("services", "glogon", "host"),
			}

			fmt.Println(glogon)

		case "casc":

			var cacs = map[string]interface{}{
				"port": cmd.info.Get("services", "cacs", "port"),
			}

			fmt.Println(cacs)
		case "icco":

			var icco = map[string]interface{}{
				"port": cmd.info.Get("services", "icco", "port"),
			}

			fmt.Println(icco)

		case "icco-agent":

			var iccoAgent = map[string]interface{}{
				"port":         cmd.info.Get("services", "icco-agent", "port"),
				"device_port":  cmd.info.Get("services", "icco-agent", "device_port"),
				"service_port": cmd.info.Get("services", "icco-agent", "service_port"),
			}

			fmt.Println(iccoAgent)

		case "epc":

			var epc = map[string]interface{}{
				"port": cmd.info.Get("services", "epc", "port"),
			}

			fmt.Println(epc)
		default:
		}

	}

	//未使用 strings.Replace 是因为 strings.Replace 无法进行 map 批量替换
	//未使用 template, 是因为 template 太复杂了
	//故使用 fmt.Sprintf
	args := []string{
		"run",
		"--detach",
		"--privileged",
		"--restart=always",
		fmt.Sprintf("--name=%s", containerName),
		fmt.Sprintf("--publish=%d:%d", 3007, 3007),
	}
	//挂载的 volumes
	var volumes = map[string]string{
		"volumes":                  "/volumes",
		"/etc/lims2":               "/etc/lims2",
		"/var/lib/lims2":           "/var/lib/lims2",
		"/home/disk":               "/home/disk",
		"/var/lib/lims2_vidcam":    "/var/lib/lims2_vidcam",
		"data/run/lock/lims2":      "/run/lock/lims2",
		"data/etc/php5":            "/etc/php5",
		"data/etc/nginx":           "/etc/nginx",
		"data/var/www":             "/var/www",
		"volumes/var/log/nginx":    "/var/log/nginx",
		"/dev/log":                 "/dev/log",
		"/etc/sphinxsearch/conf.d": "/etc/sphinxsearch/conf.d",
		"/etc/msmtprc":             "/etc/msmtprc",
		"/tmp/genee-nodejs-ipc":    "/tmp/genee-nodejs-ipc",
		"/var/lib/php5":            "/var/lib/php5",
	}

	for k, v := range volumes {
		args = append(args, fmt.Sprintf("-v %s:%s", k, v))
	}

	var fqdn = cmd.info.Get("fqdn").(string)
	var version = cmd.info.Get("version").(string)

	args = append(args, fmt.Sprintf("docker.genee.in/genee/lims2-%s_%s:%s", fqdn, version, revision))

	c, err := exec.Command("docker", strings.Join(args, " ")).Output()

	if err != nil {
		panic(err)
	}

	fmt.Println(string(c))

	return 0
}
