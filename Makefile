.DEFAULT_GOAL := help

VERSION = "unset"
DATE=$(shell date -u +%Y-%m-%d-%H:%M:%S-%Z)

GOFILES=$(shell go list ./... | grep -v /vendor/)
IMAGE_DEV_TAG=dev
IMAGE_TAG:=tag
PROJECTNAME=$(shell basename "$(PWD)"
GOPROXY =$("https://proxy.golang.org")
BUILD_FLAGS = "-X github.com/rugwirobaker/larissa/pkg/build.version=$(VERSION) -X github.com/rugwirobaker/larissa/pkg/build.buildDate=$(DATE)"

all: help

build:  	## build development larissa binary
	@echo "> building binary..."
	@CGO_ENABLED=0 go build -o bin/larissa ./cmd/.

clean:		## remove build artifacts
	@echo "> removing artifacts..."
	@rm -r bin/*

dev:  		## start development environment
	@echo "> starting dev environment..."

image: 		## build docker image
	@echo "> building docker image..."

install: 	## install client cli
	@echo "> installing cli..."

release:	## build the larissa server with version number
	@echo "> creating release binaries..."
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOPROXY) go build -ldflags $(BUILD_FLAGS) -o bin/larissa_windows ./cmd/.
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOPROXY) go build -ldflags $(BUILD_FLAGS) -o bin/larissa_linux ./cmd/.
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOPROXY) go build -ldflags $(BUILD_FLAGS) -o bin/larissa_darwin ./cmd/.

test:		## run unit tests
	@echo "> running unit tests..."
	@go test $(GOFILES)

tidy:		## install dependencies
	@echo "> downloading dependincies..."

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
