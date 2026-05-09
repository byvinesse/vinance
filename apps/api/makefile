include .env
export $(shell sed 's/=.*//' .env)

start:
	go run main.go

lint:
	@echo "Applying linter"
	golangci-lint run ./...

test:
	@echo "Running all tests"
	go test ./...

test-vc:
	@echo "Running tests with verbose coverage"
	go test -v -cover ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated at coverage.html"