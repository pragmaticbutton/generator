lint:
	@golangci-lint run

build:
	@go build ./...

run:
	@go run ./... $(location) $(name)