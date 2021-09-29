.PHONY: build
build:
	go build -v main.go
run:
	go run main.go
.DEFAULT_GOAL := build