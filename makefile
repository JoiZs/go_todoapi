.Default_goal:=run
.PHONY: fmt vet build run

fmt:
	@go fmt ./...

vet:fmt
	@go vet -v ./...

build:vet
	@go build -o bin/go_todoapi

run:build
	@./bin/go_todoapi
