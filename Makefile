.PHONY: build test vet fmt-check framework-check engineering-test

GOCACHE ?= /tmp/qwsg-go-cache
GOMODCACHE ?= /tmp/qwsg-go-modcache

build:
	mkdir -p build
	GOCACHE=$(GOCACHE) GOMODCACHE=$(GOMODCACHE) go build -trimpath -o build/qwsg ./cmd/qwsg

test:
	GOCACHE=$(GOCACHE) GOMODCACHE=$(GOMODCACHE) go test ./...

vet:
	GOCACHE=$(GOCACHE) GOMODCACHE=$(GOMODCACHE) go vet ./...

fmt-check:
	test -z "$$(gofmt -l cmd internal)"

framework-check:
	./ai/scripts/framework-check.sh

engineering-test: framework-check
	./ai/tests/test-framework.sh
	./ai/tests/test-divert-task-to-test.sh
	./ai/tests/test-next-task.sh
	./ai/tests/test-task-builder.sh
