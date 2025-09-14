.PHONY: build clean test

BINARY_NAME=charm-js
BUILD_DIR=build
GO_FILES=$(wildcard *.go)

# Build shared library for different platforms
build: build-linux build-windows build-macos

build-linux:
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -buildmode=c-shared -o $(BUILD_DIR)/$(BINARY_NAME).so $(GO_FILES)

build-windows:
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -buildmode=c-shared -o $(BUILD_DIR)/$(BINARY_NAME).dll $(GO_FILES)

build-macos:
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -buildmode=c-shared -o $(BUILD_DIR)/$(BINARY_NAME).dylib $(GO_FILES)

# Build for current platform
build-local:
	@mkdir -p $(BUILD_DIR)
	go build -buildmode=c-shared -o $(BUILD_DIR)/$(BINARY_NAME) $(GO_FILES)

test:
	go test -v ./...

clean:
	rm -rf $(BUILD_DIR)

# Generate C header
header:
	@mkdir -p $(BUILD_DIR)
	go build -buildmode=c-shared -o $(BUILD_DIR)/$(BINARY_NAME) $(GO_FILES)
	@echo "Header generated: $(BUILD_DIR)/$(BINARY_NAME).h"