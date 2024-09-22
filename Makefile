.PHONY: clean lint build.local

BINARY               ?= ready-wait-controller
VERSION              ?= $(shell git describe --tags --always --dirty)
SOURCES              = $(shell find . -name '*.go')
GO                   ?= go
GOPKGS               = $(shell $(GO) list ./...)
BUILD_FLAGS          ?= -v
LDFLAGS              ?= -X main.version=$(VERSION) -w -s

default: build.local

clean:
	rm -rf build

test:
	$(GO) vet -v $(GOPKGS)

lint:$(SOURCES)
	$(GO) mod download
	golangci-lint -v run ./...

fmt:
	$(GO) fmt $(GOPKGS)

build.local: build/$(BINARY)

build/$(BINARY): $(SOURCES)
	CGO_ENABLED=0 $(GO) build -o build/$(BINARY) $(BUILD_FLAGS) -ldflags "$(LDFLAGS)" ./cmd/$(BINARY)
