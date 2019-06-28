FROM golang:1.12-alpine3.10

RUN apk update && \
    apk upgrade && \
    apk add --no-cache \
    gcc \
    libc-dev \
    git \
    czmq-dev \
    libzmq \
    libsodium

ENV GO111MODULE=on
EXPOSE 2345

RUN mkdir /go/src/simulator
WORKDIR /go/src/simulator

ENTRYPOINT ["./docker-entrypoint.sh"]
