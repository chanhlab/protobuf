SHELL := /bin/bash -o pipefail

UNAME_OS := $(shell uname -s)
UNAME_ARCH := $(shell uname -m)

.PHONY: env
env:
	@echo "GOPATH: $(GOPATH)"

.PHONY: lint
lint:
	@echo "## Run lint"
	cd proto/ && go run github.com/bufbuild/buf/cmd/buf mod update
	cd proto/ && go run github.com/bufbuild/buf/cmd/buf lint

.PHONY: generate
generate:
	@echo "## Generate"
	rm -rf gen
	cd proto/ && go run github.com/bufbuild/buf/cmd/buf mod update
	go run github.com/bufbuild/buf/cmd/buf generate proto

.PHONY: install
install:
	go get github.com/bufbuild/buf/cmd/buf
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
