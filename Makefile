.PHONY: build clean test

BINARY_NAME=charm-js
BUILD_DIR=build
GO_FILES=$(wildcard *.go)

# Build shared library for different platforms
build: build-linux build-windows build-macos

ifeq ($(OS),Windows_NT)
  MKDIR_CMD = if not exist $(BUILD_DIR) mkdir $(BUILD_DIR)
  # for cmd.exe we chain set commands with && then run go
  ENV_CMD = set CGO_ENABLED=1&& set GOOS=windows&& set GOARCH=amd64&&
else
  MKDIR_CMD = mkdir -p $(BUILD_DIR)
  ENV_CMD = CGO_ENABLED=1 GOOS=windows GOARCH=amd64
endif

build-linux:
	@$(MKDIR_CMD)
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -buildmode=c-shared -o $(BUILD_DIR)/$(BINARY_NAME).so $(GO_FILES)

build-windows:
	@$(MKDIR_CMD)
	@$(ENV_CMD) go build -buildmode=c-shared -o $(BUILD_DIR)/$(BINARY_NAME).dll $(GO_FILES)

build-macos:
	@$(MKDIR_CMD)
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -buildmode=c-shared -o $(BUILD_DIR)/$(BINARY_NAME).dylib $(GO_FILES)

# Build for current platform
build-local:
	@$(MKDIR_CMD)
	go build -buildmode=c-shared -o $(BUILD_DIR)/$(BINARY_NAME) $(GO_FILES)

test:
	go test -v ./...

clean:
	rm -rf $(BUILD_DIR)

# Generate C header
header:
	@$(MKDIR_CMD)
	go build -buildmode=c-shared -o $(BUILD_DIR)/$(BINARY_NAME) $(GO_FILES)
	@echo "Header generated: $(BUILD_DIR)/$(BINARY_NAME).h"