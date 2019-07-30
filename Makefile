include Makefile.ledger
all: lint install

install: go.sum
		go install  $(BUILD_FLAGS) ./cmd/nsd
		go install  $(BUILD_FLAGS) ./cmd/nscli

lint:
	golangci-lint run
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -d -s
	go mod verify