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
ARG cname
WORKDIR /go/src/${cname}

ENTRYPOINT ["./docker-entrypoint.sh"]
