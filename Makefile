build:
	@go build -o bin/check cmd/main.go

run: build
	@./bin/check