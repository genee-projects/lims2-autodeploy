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
$ docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:latest go build -o lims2
```

### Local
```
$ go build -o lims2
```
