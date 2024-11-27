# Simple Makefile for a Go project

# Build the application
all: build test

build:
	@echo "Building..."
	
	@go build -o main.exe cmd/api/main.go

# Run the application
run:
	@echo "Server is on running..."

	@go run cmd/api/main.go

# Run DB, API container
docker-run:
	@docker compose up -d

# Create (build, run) DB, API container
docker-build:
	@docker compose up --build -d

docker-db-run:
	@docker-compose up -d postgresdb

# Check Logs
docker-logs:
	@docker compose logs -f

# Shutdown DB container
docker-down:
	@docker compose down

# Shutdown DB, API (remove db volume) container
docker-remove:
	@docker-compose down -v

# Get into Bash/Ash, Psql of container
docker-exec-api:
	@docker exec -it platformservicego-api-1 ash

docker-exec-psql:
	@docker exec -it platformservicego-postgresdb-1 psql -U devdotpep -d PlatformServiceDB

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v
# Integrations Tests for the application
itest:
	@echo "Running integration tests..."
	@go test ./internal/database -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@powershell -ExecutionPolicy Bypass -Command "if (Get-Command air -ErrorAction SilentlyContinue) { \
		air; \
		Write-Output 'Watching...'; \
	} else { \
		Write-Output 'Installing air...'; \
		go install github.com/air-verse/air@latest; \
		air; \
		Write-Output 'Watching...'; \
	}"

.PHONY: all build run test clean watch docker-run docker-down itest
