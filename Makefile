.PHONY: build run
.SILENT: build run

build:
	CGO_ENABLED=0 go build -o ./bin/echo ./cmd/main.go

run: build
	./bin/echo -i ./examples/routes.json
