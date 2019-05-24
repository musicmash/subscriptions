all:

build:
	go build -v ./cmd/artists/...

install:
	go install ./cmd/artists/...

run:
	go run ./cmd/artists/...

tests t:
	go test -v ./internal/...

lint-all l:
	bash ./scripts/golangci-lint.sh
	bash ./scripts/consistent.sh
