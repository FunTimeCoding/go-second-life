.DEFAULT_GOAL := all

all: test lint

tool:
	go install gotest.tools/gotestsum@latest

lint:
	golangci-lint run

test:
	gotestsum --format standard-quiet -- ./...

update:
	for ELEMENT in $$(go list -f "{{if not (or .Main .Indirect)}}{{.Path}}{{end}}" -m all); do echo $${ELEMENT}; go get $${ELEMENT}; done
	@go mod tidy
	@go-update

update-library:
	GOPROXY=direct go get github.com/funtimecoding/go-library
	@go mod tidy
	@go-update
