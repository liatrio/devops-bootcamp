# makefile for devops api project this is required for building and running

run: main.go build
	go run main.go

build: main.go tidy fmt test
	go build

tidy: main.go
	go mod tidy

fmt: main.go
	go fmt

test: main.go
	go test -v

clean:
	rm devops-api
