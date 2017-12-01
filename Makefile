export

# Project specific variables
PROJECT=api
VENDOR=./src/vendor
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
	@for gofile in $$(find ./src/api -name "*.go"); do \
		echo "formatting" $$gofile; \
		gofmt -w $$gofile; \
	done

.PHONY: test
test:
	$(GOCMD) test -v -race ./src/$(PROJECT)/...

.PHONY: integrationtest
integrationtest:
	$(GOCMD) test -v -tags=integration ./src/test/...

multi: build darwin linux windows

darwin:
	GOOS=darwin GOARCH=${GOARCH} $(GOBUILD) -o bin/$(PROJECT)_darwin $(PROJECT)
linux:
	GOOS=linux GOARCH=${GOARCH} $(GOBUILD) -o bin/$(PROJECT)_linux $(PROJECT)
windows:
	GOOS=windows GOARCH=${GOARCH} $(GOBUILD) -o bin/$(PROJECT)_windows.exe $(PROJECT)
