# Lims2 Autodeploy.

## How To Use

```
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
  --version     Show version.
```

## How To Build

### Docker

```
$ docker run \
    --rm \
    -v "$PWD":/go/src/github.com/lims2-tools/autodeploy \
    -w /go/src/github.com/lims2-tools/autodeploy \
    golang:latest \
    bash -c 'go get "github.com/docopt/docopt-go" && cd /go/src/github.com/lims2-tools/autodeploy && go build -o lims2'
```

### Local
```
$ go build -o lims2
```
