build:
	@go build -o bin/okeledb cmd/main.go

run: build
	@./bin/okeledb

test: 
	@go test -v ./...