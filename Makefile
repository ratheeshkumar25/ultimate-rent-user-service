# Makefile for ultimate-rent-user-service

# Variables
BINARY_NAME=user-service
BUILD_DIR=bin
PROTO_DIR=internal/pb
PROTO_FILES=$(wildcard $(PROTO_DIR)/*/*.proto)

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Protoc parameters
PROTOC=protoc
PROTOC_GEN_GO=protoc-gen-go
PROTOC_GEN_GO_GRPC=protoc-gen-go-grpc

.PHONY: all build clean test run proto deps help

# Default target
all: deps proto build

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) .

# Run the service
run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BUILD_DIR)/$(BINARY_NAME)

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

# Generate protobuf files
proto:
	@echo "Generating protobuf files..."
	@for proto in $(PROTO_FILES); do \
		$(PROTOC) --go_out=. --go_opt=paths=source_relative \
			--go-grpc_out=. --go-grpc_opt=paths=source_relative \
			$$proto; \
	done

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

# Install protoc plugins (if needed)
install-protoc-plugins:
	@echo "Installing protoc plugins..."
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Help
help:
	@echo "Available targets:"
	@echo "  all          - Download deps, generate proto, and build"
	@echo "  build        - Build the binary"
	@echo "  run          - Build and run the service"
	@echo "  test         - Run tests"
	@echo "  clean        - Clean build artifacts"
	@echo "  proto        - Generate protobuf files"
	@echo "  deps         - Download and tidy dependencies"
	@echo "  install-protoc-plugins - Install protoc-gen-go and protoc-gen-go-grpc"
	@echo "  help         - Show this help"
