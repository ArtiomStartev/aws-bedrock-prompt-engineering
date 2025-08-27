# AWS Bedrock Prompt Engineering Project Makefile

.PHONY: help build run test clean install deps check examples

# Default target
help:
	@echo "🚀 AWS Bedrock Prompt Engineering Project"
	@echo ""
	@echo "Available commands:"
	@echo "  make install     - Install Go dependencies"
	@echo "  make build       - Build the main application"
	@echo "  make run         - Run the main application"
	@echo "  make test        - Run tests (when implemented)"
	@echo "  make check       - Check code formatting and run linter"
	@echo "  make clean       - Clean build artifacts"
	@echo "  make deps        - Update dependencies"

# Install dependencies
install:
	@echo "📦 Installing dependencies..."
	go mod download
	go mod tidy

# Build the application
build:
	@echo "🔨 Building application..."
	go build -o bin/prompt-engineering main.go

# Run the main application
run:
	@echo "🚀 Running prompt engineering demo..."
	go run main.go

# Run tests
test:
	@echo "🧪 Running tests..."
	go test ./...

# Check code formatting and linting
check:
	@echo "🔍 Checking code format..."
	go fmt ./...
	@echo "✅ Code formatting complete"

# Clean build artifacts
clean:
	@echo "🧹 Cleaning build artifacts..."
	rm -rf bin/
	go clean

# Update dependencies
deps:
	@echo "🔄 Updating dependencies..."
	go get -u ./...
	go mod tidy

# Create necessary directories
setup:
	@echo "📁 Setting up project directories..."
	mkdir -p bin/
	mkdir -p logs/
	@echo "✅ Project setup complete"
