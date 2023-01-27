build:
	@go build -o main cmd/app/main.go

test:
	@go test ./... --short

run:
	@go run cmd/app/main.go
