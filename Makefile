.PHONY: build test vet fmt-check

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
