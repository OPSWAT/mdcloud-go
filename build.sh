#!/bin/sh

cd /go/src/github.com/OPSWAT/mdcloud-go \
    && cp -r /tmp/vendor . \
    && gox -osarch="linux/amd64 linux/386 darwin/amd64 darwin/386 windows/arm windows/386" -ldflags "-X main.VERSION=$VERSION -s -w" -output="dist/{{.Dir}}_{{.OS}}_{{.Arch}}"
