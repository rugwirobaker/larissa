include .env

VERSION = "unset"
DATE=$(shell date -u +%Y-%m-%d-%H:%M:%S-%Z)

GOFILES=$(shell go list ./... | grep -v /vendor/)
IMAGE_DEV_TAG=dev
IMAGE_TAG:=tag
PROJECTNAME=$(shell basename "$(PWD)"

all: binary

build:
	@echo "> building binaries..."
clean:
	@echo "> cleaning up..."
dev:
	@echo "> starting dev environment..."
image:
	@echo "> building docker image..."
install:
	@echo "> installing cli..."

release: ## build the athens proxy with version number
	CGO_ENABLED=0 GOPROXY="https://proxy.golang.org" go build -ldflags "-X github.com/rugwirobaker/larissa/pkg/larissa.version=$(VERSION) -X github.com/rugwirobaker/larissa/pkg/larissa.buildDate=$(DATE)" -o bin/larissa ./cmd/.

test:
	@echo "> running unit tests..."
	@go test $(GOFILES)
tidy:
	@echo "> downloading dependincies..."


help: ## display help page
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
