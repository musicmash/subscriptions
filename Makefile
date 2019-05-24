all:

build:
	go build -v ./cmd/artists/...

install:
	go install ./cmd/artists/...

run:
	go run ./cmd/artists/...
