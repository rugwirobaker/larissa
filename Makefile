include .env

VERSION = "unset"
DATE=$(shell date -u +%Y-%m-%d-%H:%M:%S-%Z)

GOFILES=$(shell go list ./... | grep -v /vendor/)
IMAGE_DEV_TAG=dev
IMAGE_TAG:=tag
PROJECTNAME=$(shell basename "$(PWD)"
GOPROXY =$("https://proxy.golang.org")
BUILD_FLAGS = "-X github.com/rugwirobaker/larissa/pkg/build.version=$(VERSION) -X github.com/rugwirobaker/larissa/pkg/build.buildDate=$(DATE)"

all: binary

build:
	@echo "> building binaries..."
	CGO_ENABLED=0 go build -o bin/larissa ./cmd/.
clean:
	@echo "> cleaning up..."
dev:
	@echo "> starting dev environment..."
image:
	@echo "> building docker image..."
install:
	@echo "> installing cli..."

release: ## build the larissa server with version number
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOPROXY) go build -ldflags $(BUILD_FLAGS) -o bin/larissa_windows ./cmd/.
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOPROXY) go build -ldflags $(BUILD_FLAGS) -o bin/larissa_linux ./cmd/.
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOPROXY) go build -ldflags $(BUILD_FLAGS) -o bin/larissa_darwin ./cmd/.

test:
	@echo "> running unit tests..."
	@go test $(GOFILES)
tidy:
	@echo "> downloading dependincies..."


help: ## display help page
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
