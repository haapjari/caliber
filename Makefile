DOCKER_COMPOSE 			= docker-compose
GO 						= go
NAME 					= caliber
VERSION 				= $(shell git describe --tags --abbrev=0) 
MAIN                    = cmd/main.go

.PHONY: run run-dev build compose test lint download-tools coverage coverage-html

info: 
	@echo "Name: $(NAME)"
	@echo "Version: $(VERSION)"
	@echo "Last Commit Date: $(shell git log -1 --format=%cd)"
	@echo "Last Commit Hash: $(shell git rev-parse HEAD)"
	@echo "Total Commits: $(shell git rev-list --all --count)"
	@echo "Last Contributor: $(shell git log -1 --format='%an')"

run: build
	./bin/$(NAME) $(ARGS)

run-dev:
	$(GO) run --race $(MAIN) $(ARGS)

build:
	$(GO) build -o bin/$(NAME) $(MAIN)

test:
	$(GO) test -cover ./...

lint: fmt
	golangci-lint run 

fmt:
	gofumpt -l -w .

download-tools:
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	$(GO) install mvdan.cc/gofumpt@latest

coverage:
	go test -v -coverprofile cover.out ./...
