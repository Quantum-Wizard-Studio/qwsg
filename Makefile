.PHONY: build build-contract-check install install-service release release-check release-authority-check release-authority-tools test vet fmt-check framework-check engineering-test

GOCACHE ?= /tmp/qwsg-go-cache
GOMODCACHE ?= /tmp/qwsg-go-modcache
PREFIX ?= /usr/local
DESTDIR ?=
BINDIR ?= $(PREFIX)/bin
INSTALL ?= install
SYSTEMD_USER_UNIT_DIR ?= $(PREFIX)/lib/systemd/user
VERSION := $(shell tr -d '\r\n' < VERSION)
BUILD_COMMIT ?= $(shell git rev-parse --short=12 HEAD 2>/dev/null || printf unknown)
BUILD_DATE ?= unknown
LDFLAGS := -X main.version=$(VERSION) -X main.buildCommit=$(BUILD_COMMIT) -X main.buildDate=$(BUILD_DATE)

build:
	mkdir -p build
	GOCACHE=$(GOCACHE) GOMODCACHE=$(GOMODCACHE) go build -trimpath -buildvcs=false -ldflags "$(LDFLAGS)" -o build/qwsg ./cmd/qwsg

build-contract-check:
	./scripts/test-build-contract.sh

install:
	@test -f build/qwsg -a -x build/qwsg || { \
		printf '%s\n' 'Error: build/qwsg is missing or not executable; run make build as the normal user first.' >&2; \
		exit 1; \
	}
	$(INSTALL) -d -m 0755 "$(DESTDIR)$(BINDIR)"
	$(INSTALL) -m 0755 build/qwsg "$(DESTDIR)$(BINDIR)/qwsg"

install-service: install
	@test "$(PREFIX)" = "/usr/local" || { \
		printf '%s\n' 'Error: QWSG 1.0 supports only PREFIX=/usr/local because the shipped unit has an absolute, audited ExecStart.' >&2; \
		exit 1; \
	}
	$(INSTALL) -d -m 0755 "$(DESTDIR)$(SYSTEMD_USER_UNIT_DIR)"
	$(INSTALL) -m 0644 packaging/systemd/qwsg-guardian.service "$(DESTDIR)$(SYSTEMD_USER_UNIT_DIR)/qwsg-guardian.service"

release:
	./scripts/build-release.sh

release-check:
	./scripts/test-release-plumbing.sh
	./scripts/test-release-reproducibility.sh

release-authority-tools:
	mkdir -p build/release-authority
	GOCACHE=$(GOCACHE) GOMODCACHE=$(GOMODCACHE) go build -trimpath -buildvcs=false -o build/release-authority/qwsg-release-index ./cmd/qwsg-release-index
	GOCACHE=$(GOCACHE) GOMODCACHE=$(GOMODCACHE) GOOS=windows GOARCH=amd64 go build -trimpath -buildvcs=false -o build/release-authority/qwsg-release-sign-offline.exe ./cmd/qwsg-release-sign-offline

release-authority-check:
	./scripts/test-release-authority-reproducibility.sh

test:
	GOCACHE=$(GOCACHE) GOMODCACHE=$(GOMODCACHE) go test ./...

vet:
	GOCACHE=$(GOCACHE) GOMODCACHE=$(GOMODCACHE) go vet ./...

fmt-check:
	test -z "$$(gofmt -l cmd internal)"

framework-check:
	./ai/scripts/framework-check.sh

engineering-test: framework-check build-contract-check
	./ai/tests/test-framework.sh
	./ai/tests/test-divert-task-to-test.sh
	./ai/tests/test-next-task.sh
	./ai/tests/test-task-builder.sh
