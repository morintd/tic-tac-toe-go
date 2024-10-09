install:
	go get ./...
start:
	./bin/application
dev:
	go run *.go
build:
	go build -o bin/application main.go
build_linux:
	GOARCH=amd64 GOOS=linux go build -o bin/application main.go
lint:
	golangci-lint run --timeout=10m
install_tools:
	"$(CURDIR)/scripts/install_tools.sh"

.NOTPARALLEL:

.PHONY: install start dev build build_linux install_tools