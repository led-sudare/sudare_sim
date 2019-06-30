FROM golang:1.12-alpine3.10 as Builder
RUN apk --no-cache add \
    gcc \
    libc-dev \
    git \
    czmq-dev \
    libzmq \
    libsodium
ENV GO111MODULE=on
WORKDIR /app
COPY . /app
RUN go build -a -o sim

FROM alpine:3.10 as Runner
RUN apk --no-cache add ca-certificates czmq
EXPOSE 2345
WORKDIR /app
COPY --from=Builder /app/sim /app/
CMD /app/sim
