.PHONY: codegen test testv lint lintv fmt vet

all: clean codegen fmt lint build test

clean:
	rm -f bin/*

codegen:
	./scripts/run_autogen.sh

build:
	go build -o ./bin cmd/main.go

test:
	go test -race ./...,

testv:
	go test -race -v ./...

lint:
	golangci-lint run -c .golangci.yml
# gofmt -l -s -w .

lintv:
	golangci-lint run -v -c .golangci.yml

fmt:
	go fmt ./...

vet:
	go vet ./...

list:
	go list ./...
