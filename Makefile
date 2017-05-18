SRCS := $(shell find . -type f -name '*.go' ! -name '*_test.go')
TESTS := $(shell find . -type f -name '*_test.go')

.PHONY: build
build:
	go build -o slack-tail $(SRCS)

.PHONY: run
run:
	go run $(SRCS)

.PHONY: test
test:
	go test -cover -v $(TESTS) $(SRCS)
