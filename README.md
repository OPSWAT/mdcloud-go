![Logo](/images/logo.png?raw=true)

mdcloud cli
------------
Command line tool for metadefender cloud ip scanner designed for scanning amazon security groups.

## Build and install

The simple way of installing the tool:
```
wget ... -O /usr/local/bin/mdcloud
```

Visit [this page]() for a list of alternative downloads.

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

Before running the tool, please make sure you have a metadefender cloud apikey. If not, please go to [metadefender.com](https://www.metadefender.com) and click the "Sign up" button.

After obtaining an apikey, you need to specify it in the command line by setting the "MDCLOUD_APIKEY" environment variable, or by passing it as an argument to the tool with `--apikey` like so:

```
mdcloud --apikey <command>
```

The outputs of the source code are executables compiled for specific platforms.

To see possible options run:
```
mdcloud --help

```

In order to scan a security group:
```
mdcloud sgscan
```

This command relies on the fact that amazon credentials are already configured in `~/.aws/credentials`.
The `--include` flag can be used to specify a comma delimited list of security groups to scan. By default all groups are scanned.

To see a list of possible security groups use:

```
mdcloud list
```

For now this tool is designed specifically for scanning ip addresses in aws security groups. More functionality is yet to come :)

Licensed under the [MIT License](https://opensource.org/licenses/MIT)

