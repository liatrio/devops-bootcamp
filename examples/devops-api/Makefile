# makefile for devops api project this is required for building and running

run: main.go build
	go run main.go

build: main.go fmt test
	go build

fmt: main.go
	go fmt

test: main.go
	go test -v

clean:
	rm devops-api
