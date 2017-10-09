# build stage
FROM golang:alpine
ARG VERSION
ENV VERSION ${VERSION}
ENV CGO_ENABLED 0
RUN mkdir -p /go/src/github.com/OPSWAT/mdcloud-go
WORKDIR /go/src/github.com/OPSWAT/mdcloud-go
COPY glide.yaml glide.yaml
COPY glide.lock glide.lock
RUN apk --no-cache add curl git && curl https://glide.sh/get | sh && glide i -v && go get github.com/mitchellh/gox
