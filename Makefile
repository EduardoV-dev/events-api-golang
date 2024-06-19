build:
	@go build -o bin/app cmd/main.go

prod: build
	@./bin/app

dev:
	@go run cmd/main.go -env=development
