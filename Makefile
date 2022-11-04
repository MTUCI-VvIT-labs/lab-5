run:
	GIN_MODE=release ./bin/main

build:
	 go build -o bin/main ./cmd/app/main.go