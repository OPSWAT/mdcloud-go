<a href="https://metadefender.com">
  <img src="MD-Cloud-logo-black.svg" width="50%" height="50%">
</a>
<br/>
<br/>

[![mdcloud-result](https://api.metadefender.com/v4/hash/16716AA47017A93D5DF00C860457732BA60ABB75/badge?size=small&type=svg)](https://metadefender.opswat.com/results#!/file/16716AA47017A93D5DF00C860457732BA60ABB75/hash/overview)
[![GoDoc](https://godoc.org/github.com/OPSWAT/mdcloud-go?status.svg)](https://godoc.org/github.com/OPSWAT/mdcloud-go) [![Go Report Card](https://goreportcard.com/badge/github.com/OPSWAT/mdcloud-go)](https://goreportcard.com/report/github.com/OPSWAT/mdcloud-go)

# mdcloud cli

Command line tool for metadefender cloud ip scanner designed for scanning amazon security groups.

## Build and install

The simple way of installing the tool:

```bash
sudo wget -q https://github.com/OPSWAT/mdcloud-go/pkg/releases/download/1.2.0/mdcloud-go_linux_amd64 -O /usr/local/bin/mdcloud && sudo chmod +x /usr/local/bin/mdcloud
```

Visit [this page](https://github.com/OPSWAT/mdcloud-go/pkg/releases) for a list of alternative downloads.

For building we use a docker image with all the dependencies installed. The image is built from `image.dockerfile` file

The docker image is hosted on a public repo on docker hub and can be downloaded locally.

Just in case, here is how to build the docker image:

```bash
make image VERSION=<new_version_of_docker_image>
```

For compiling the source code run:

```bash
make build VERSION=<version of executable>
```

This will produce a folder `dist` which contains all executables.

## Usage

Before running the tool, please make sure you have:

- a metadefender cloud apikey. If not, please go to [metadefender.com](https://www.metadefender.com) and click the "Sign up" button.
- an amazon account configured (config file used by the tool is `~/.aws/credentials`)

After obtaining an apikey, you need to specify it in the command line by setting the `MDCLOUD_APIKEY` environment variable, or by passing it as an argument to the tool with `--apikey` like so:

```bash
mdcloud --apikey <command>
```

The outputs of the source code are executables compiled for specific platforms.

To see possible options run:

```bash
$ mdcloud
Metadefender Cloud API wrapper

Usage:
  mdcloud [command]

Available Commands:
  appinfo       Appinfo for hash
  feed          Feed of hashes, infected or false-positives
  help          Help about any command
  lookup        Lookup or download file or IPs/Domains/URLs
  rescan        Rescan file
  sanitized     Sanitized result by file_id
  scan          Scan file or path
  sglist        List security groups IPs
  sgscan        Scan security groups using IP scan API
  version       Print the version number of mdcloud
  vulnerability Vulnerability for hash

Flags:
  -a, --apikey string      set apikey token (default is MDCLOUD_APIKEY env variable)
  -f, --formatter string   set formatter type to  json or text (default "text")
  -h, --help               help for mdcloud

Use "mdcloud [command] --help" for more information about a command.
```

This command relies on the fact that amazon credentials are already configured in `~/.aws/credentials`.

Licensed under the [MIT License](https://opensource.org/licenses/MIT)
