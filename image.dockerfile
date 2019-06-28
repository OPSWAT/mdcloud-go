# build stage
FROM golang:alpine
ARG VERSION
ENV VERSION ${VERSION}
ENV CGO_ENABLED 0
WORKDIR /tmp
COPY go.mod go.mod
COPY go.sum go.sum
RUN apk --no-cache add git g++ make \
    && go mod download \
    && go get github.com/mitchellh/gox
