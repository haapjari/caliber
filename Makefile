include .env

run: build
	./bin/$(NAME)

build:
	go build -o bin/$(NAME) cmd/main.go	
