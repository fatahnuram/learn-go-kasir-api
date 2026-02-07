SHELL := /bin/bash

PROJECT_NAME := "github.com/fatahnuram/learn-go-kasir-api"
PKG := "$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)

.DEFAULT_GOAL := help


.PHONY: test
# Test *_test.go files, the parameter -count=1 means that caching is disabled
test:
	go test -count=1 -short ${PKG_LIST}


.PHONY: build
# Build binary
build:
	CGO_ENABLED=0 go build


.PHONY: start
# Build and run service
start: build
	./learn-go-kasir-api


.PHONY: run
# Build and run service
run: start


.PHONY: help
# Show help
help:
	@echo ''
	@echo 'Usage:'
	@echo '  make <target>'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[1;36m  %-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)
