.DEFAULT_GOAL := all

all: test lint

tool:
	go install gotest.tools/gotestsum@latest

test:
	gotestsum --format standard-quiet -- ./...

lint:
	golangci-lint run

update:
	for ELEMENT in $$(go list -f "{{if not (or .Main .Indirect)}}{{.Path}}{{end}}" -m all); do echo $${ELEMENT}; go get $${ELEMENT}; done
	@go mod tidy
	@go-update

update-library:
	GOPROXY=direct go get github.com/funtimecoding/go-library
	@go mod tidy
	@go-update
