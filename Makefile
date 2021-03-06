export

# Project specific variables
PROJECT=api
GOMAIN=api.go
OS=$(shell uname)
GOARCH = amd64

# GO env
GOPATH=$(shell pwd)
GO=go
GOCMD=GOPATH=$(GOPATH) $(GO)

GOBUILD = $(GOCMD) build

# Build the project
.PHONY: all
all:	build

.PHONY: build
build: format test compile

.PHONY: compile
compile: darwin linux windows

.PHONY: format
format:
	@for gofile in $$(find ./src -name "*.go"); do \
		echo "formatting" $$gofile; \
		gofmt -w $$gofile; \
	done

.PHONY: test
test:
	$(GOCMD) test -v -race ./src/...

.PHONY: integrationtest
integrationtest:
	$(GOCMD) test -v -tags=integration ./test/...

multi: build darwin linux windows

darwin:
	GOOS=darwin GOARCH=${GOARCH} $(GOBUILD) -o bin/$(PROJECT)_darwin ./src/$(GOMAIN)
linux:
	GOOS=linux GOARCH=${GOARCH} $(GOBUILD) -o bin/$(PROJECT)_linux ./src/$(GOMAIN)
windows:
	GOOS=windows GOARCH=${GOARCH} $(GOBUILD) -o bin/$(PROJECT)_windows.exe ./src/$(GOMAIN)
