#!/bin/sh

cd /go/src/github.com/OPSWAT/mdcloud-go \
    && mkdir -p /build/vendor \
    && cp -r /build/vendor . \
    && gox -osarch="linux/amd64 linux/386 darwin/amd64 darwin/386 windows/amd64 windows/386" -ldflags "-X main.VERSION=$VERSION -s -w" -output="dist/{{.Dir}}_{{.OS}}_{{.Arch}}"
