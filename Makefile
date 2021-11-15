.PHONY: build

build:
		go build -v ./cmd/apiserver

.PHONY: test

test:
		go test -v -race -timeout 60s ./ ...

.DEFAULT_GOAL := build