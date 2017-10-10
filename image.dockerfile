# build stage
FROM golang:alpine
ARG VERSION
ENV VERSION ${VERSION}
ENV CGO_ENABLED 0
WORKDIR /tmp
COPY glide.yaml glide.yaml
COPY glide.lock glide.lock
RUN apk --no-cache add curl git g++ make \
    && curl https://glide.sh/get | sh \
    && glide i -v \
    && go get github.com/mitchellh/gox
