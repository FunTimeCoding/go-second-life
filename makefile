.DEFAULT_GOAL := all

all: test lint

tool:
	@go install gotest.tools/gotestsum@latest
	@GOPROXY=direct go install github.com/funtimecoding/go-library/cmd/golint@latest
	@GOPROXY=direct go install github.com/funtimecoding/go-library/cmd/goupdate@latest

test:
	@gotestsum --format standard-quiet -- ./...

lint:
	@golint
	@golangci-lint run

update:
	@goupdate

update-library:
	@goupdate --exclusive funtimecoding
