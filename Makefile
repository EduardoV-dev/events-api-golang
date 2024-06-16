build:
	@go build -o bin/events-api cmd/main.go

run: build
	@./bin/events-api
