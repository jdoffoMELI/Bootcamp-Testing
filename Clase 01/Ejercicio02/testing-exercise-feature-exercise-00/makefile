test:
	echo "Running tests..."
	go test ./... -cover -v -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

lint:
	echo "Running linter..."
	staticcheck ./...

all: test lint