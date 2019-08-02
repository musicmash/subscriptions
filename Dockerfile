# Create the intermediate builder image.
FROM golang:latest as builder

# Docker is copying directory contents so we need to copy them in same directories.
WORKDIR /go/src/github.com/musicmash/subscriptions
COPY . .

# Build the static application binary.
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -v -a -installsuffix cgo -gcflags "all=-trimpath=$(GOPATH)" -o bin/subscriptions    ./cmd/subscriptions/main.go

# Create the final small image.
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk add --no-cache \
    ca-certificates vim curl && \
    rm -rf /var/cache/apk/*

WORKDIR /root/
COPY --from=builder /go/src/github.com/musicmash/subscriptions/bin .

ENTRYPOINT ["./subscriptions"]
