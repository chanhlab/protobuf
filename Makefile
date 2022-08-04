SHELL := /bin/bash -o pipefail

UNAME_OS := $(shell uname -s)
UNAME_ARCH := $(shell uname -m)

.PHONY: env
env:
	@echo "GOPATH: $(GOPATH)"

.PHONY: lint
lint:
	@echo "## Run lint"
	prototool all --fix proto
	prototool lint proto

.PHONY: generate
generate:
	@echo "## Generate"
	prototool generate proto

.PHONY: install
install:
	brew install prototool
	go get google.golang.org/genproto/...
	go get github.com/grpc-ecosystem/grpc-gateway
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go mod tidy
