start: build
	./bin/app

build:
	@go build -o ./bin/app ./cmd/http/main.go
dev:
	go run ./cmd/http/main.go

test:
	go test ./...

dev-watch:
	air
