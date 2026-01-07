.PHONY: help build run test clean docker-build docker-run

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the application
	go build -o main ./cmd/api

run: ## Run the application
	go run ./cmd/api/main.go

test: ## Run tests
	go test -v ./...

clean: ## Remove build artifacts
	rm -f main
	go clean

docker-build: ## Build Docker image
	docker build -t golang-fiber-backend .

docker-run: ## Run Docker container
	docker run -p 8080:8080 golang-fiber-backend

docker-compose-up: ## Start services with docker-compose
	docker-compose up --build

docker-compose-down: ## Stop services with docker-compose
	docker-compose down

deps: ## Download dependencies
	go mod download
	go mod tidy

fmt: ## Format code
	go fmt ./...

vet: ## Run go vet
	go vet ./...

lint: fmt vet ## Run linters
