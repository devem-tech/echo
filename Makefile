MAIN_PACKAGE_PATH := ./cmd/main.go
BINARY_NAME := ./bin/echo

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## test: run all tests
.PHONY: test
test:
	@go test -v -race -buildvcs ./...

## build: build the application
.PHONY: build
build:
	@go build -ldflags='-s' -o ${BINARY_NAME} ${MAIN_PACKAGE_PATH}

## run: run the application
.PHONY: run
run: build
	@${BINARY_NAME} ./examples/routes.json -v --print=HBhb
