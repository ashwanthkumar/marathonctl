# Marathon CLI
CLI tool to access and deploy apps and services to [Marathon](https://mesosphere.github.io/marathon/).

## Configuration
You need to create a configuration file `$HOME/.marathon.json` with the following contents
```
{
  "url": "http://marathon.url:8080"
}
```

If you've maraton running in HA mode you can specify multiple marathon URLs as
```
{
  "url": "http://marathon1.url:8080,marathon2.url:8080,marathon3.url:8080"
}
```

## Usage

```
$ marathonctl
Command line client to Marathon

Usage:
  marathonctl [command]

Available Commands:
  deploy      Deploy an app using Marathon's app definition
  version     Version of the Marathon CLI

Flags:
  -h, --help   help for marathonctl

Use "marathonctl [command] --help" for more information about a command.
```

## Deploy Apps
`marathonctl deploy` helps you deploy applications to your Marathon setup from command line. It takes an app definition and tries to deploy it.

```
$ marathonctl deploy -h
Deploy an app using Marathon's app definition

Usage:
  marathonctl deploy <app.json> [flags]

Flags:
  -d, --dry-run              Print the final application configuration but don't deploy
  -e, --environment string   Environment to deploy (default "test")
  -f, --force                Force deploy the app
  -t, --timeout int          timeout in seconds for deployment to complete, else we'll fail (default 900)
```

### Application Definition
The application definition (`app.json`) that's passed it treated as a [Go Template](https://golang.org/pkg/text/template/) and rendered. The available variables for the template is `{{ .DEPLOY_ENV }}`. You can also access environment variables using the convention `{{ .Env.GO_PIPELINE_LABEL }}`, where `GO_PIPELINE_LABEL` is an environment variable. 

Example app.json file could be something like
```
{
  "id": "{{ .DEPLOY_ENV }}.http",
  "cpus": 0.1,
  "mem": 10,
  "instances": 1,
  "ports": [
    0
  ],
  "cmd": "python -m SimpleHTTPServer $PORT0",
  "uris": [
    "https://github.com/ashwanthkumar/wasp-cli/releases/download/v{{ .Env.WASP_CLI_VERSION }}/wasp-linux-amd64"
  ],
  "upgradeStrategy": {
    "minimumHealthCapacity": 0.9,
    "maximumOverCapacity": 0.1
  },
  "env": {
    "DEPLOY_ENV": "{{ .DEPLOY_ENV }}"
  },
  "healthChecks": [
    {
      "protocol": "COMMAND",
      "command": { "value": "curl -f http://$HOST:$PORT0/" },
      "gracePeriodSeconds": 60,
      "intervalSeconds": 30,
      "maxConsecutiveFailures": 3,
      "timeoutSeconds": 10
    }
  ]
}
```
