lint:
	@golangci-lint run

build:
	@go build -o generator -v ./...

run:
	@go run ./... $(location) $(name)