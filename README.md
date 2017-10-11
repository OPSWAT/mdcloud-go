![Logo](/images/logo.png?raw=true)

mdcloud cli
------------
Command line tool for metadefender cloud ip scanner designed for scanning amazon security groups.

## Build and install

The simple way of installing the tool:
```
sudo wget -q https://github.com/OPSWAT/mdcloud-go/releases/download/1.0.0/mdcloud-go_linux_amd64 -O /usr/local/bin/mdcloud && sudo chmod +x /usr/local/bin/mdcloud
```

Visit [this page](https://github.com/OPSWAT/mdcloud-go/releases) for a list of alternative downloads.

For building we use a docker image with all the dependencies installed. The image is built from `image.dockerfile` file

The docker image is hosted on a public repo on docker hub and can be downloaded locally.

Just in case, here is how to build the docker image:

```
make image VERSION=<new_version_of_docker_image>
```

For compiling the source code run:

```
make build VERSION=<version of executable>
```

This will produce a folder `dist` which contains all executables.

## Usage

Before running the tool, please make sure you have
- a metadefender cloud apikey. If not, please go to [metadefender.com](https://www.metadefender.com) and click the "Sign up" button.
- an amazon account configured (config file used by the tool is `~/.aws/credentials`)

After obtaining an apikey, you need to specify it in the command line by setting the "MDCLOUD_APIKEY" environment variable, or by passing it as an argument to the tool with `--apikey` like so:

```
$> mdcloud --apikey <command>
```

The outputs of the source code are executables compiled for specific platforms.

To see possible options run:
```
$> mdcloud --help
Metadefender Cloud API wrapper

Usage:
  mdcloud-go [command]

Available Commands:
  help        Help about any command
  sglist      List security groups IPs
  sgscan      Scan security groups using IP Scan API
  version     Print the version number of mdcloud-go

Flags:
  -a, --apikey string   apikey token (default is MDCLOUD_APIKEY env variable)
  -h, --help            help for mdcloud-go

Use "mdcloud-go [command] --help" for more information about a command.

```

In order to scan a security group:
```
$> mdcloud sgscan
109.103.100.28 : OK
35.162.110.83 : OK
35.165.42.65 : OK
54.183.119.45 : OK
62.231.66.135 : OK
79.114.3.225 : OK
113.161.86.113 : OK
14.161.18.101 : OK
195.56.42.178 : OK
90.174.2.6 : OK
54.67.15.72 : OK
90.174.2.50 : OK
95.76.19.167 : OK
76.14.65.6 : OK
52.12.101.151 : OK
52.12.182.222 : OK
52.12.86.29 : OK
52.12.88.107 : OK
```

This command relies on the fact that amazon credentials are already configured in `~/.aws/credentials`.
The `--include` flag can be used to specify a comma delimited list of security groups to scan. By default all groups are scanned.

To see a list of possible security groups use:

```
$> mdcloud sglist
109.103.100.28
35.162.110.83
35.165.42.65
54.183.119.45
62.231.66.135
79.114.3.225
113.161.86.113
14.161.18.101
195.56.42.178
90.174.2.6
54.67.15.72
90.174.2.50
95.76.19.167
76.14.65.6
```

## Notes

For now this tool is designed specifically for scanning ip addresses in aws security groups. More functionality is yet to come :)

Licensed under the [MIT License](https://opensource.org/licenses/MIT)

