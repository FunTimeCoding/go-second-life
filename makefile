.DEFAULT_GOAL := all

all: lint test

tool:
	@go install gotest.tools/gotestsum@latest
	@GOPROXY=direct go install github.com/funtimecoding/go-library/cmd/golint@latest
	@GOPROXY=direct go install github.com/funtimecoding/go-library/cmd/goupdate@latest

lint:
	@golint --fix
	@golangci-lint run

test:
	@gotestsum --format standard-quiet -- ./...

update:
	@goupdate

update-library:
	@goupdate --exclusive funtimecoding
