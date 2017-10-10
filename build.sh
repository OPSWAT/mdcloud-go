#!/bin/sh

cd /go/src/github.com/OPSWAT/mdcloud-go \
    && cp -r /tmp/vendor . \
    && gox -os="linux darwin windows" -ldflags "-X main.VERSION=$VERSION" -output="dist/{{.Dir}}_{{.OS}}_{{.Arch}}"
