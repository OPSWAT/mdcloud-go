# build stage
FROM opswat/mdcloud-go as build-env
ADD . /go/src/github.com/OPSWAT/mdcloud-go
ARG VERSION
ENV VERSION ${VERSION}
ENV CGO_ENABLED 0
RUN go get github.com/mitchellh/gox
ADD . /go/src/github.com/OPSWAT/mdcloud-go
RUN gox -os="linux darwin windows openbsd" -ldflags "-X main.VERSION=$VERSION"

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/github.com/OPSWAT/mdcloud-go/mdcloud-go_* /app/