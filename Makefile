all: lint test

lint:
ifeq (, $(shell which golangci-lint))
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $$GOPATH/bin v1.17.1
endif
	golangci-lint run

test:
	go test -v -cover -coverprofile=coverage.txt -covermode=atomic -race ./...
