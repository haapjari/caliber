DOCKER_COMPOSE 			= docker-compose
GO 						= go
NAME 					= caliber
VERSION 				= 0.0.1
DOCKER_IMAGE_NAME 		= $(NAME):$(VERSION)
MAIN                    = cmd/main.go

.PHONY: run run-dev run-docker build compose test lint download-tools coverage coverage-html

run: build
	./bin/$(NAME)

run-dev:
	$(GO) run --race $(MAIN)

run-docker:
	docker run -p 8080:8080 $(DOCKER_IMAGE_NAME)

build:
	$(GO) build -o bin/$(NAME) $(MAIN)

compose:
	$(DOCKER_COMPOSE) up --build

test:
	$(GO) test ./...

lint: fmt
	golangci-lint run 

fmt:
	gofumpt -l -w .

download-tools:
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	$(GO) install mvdan.cc/gofumpt@latest


coverage:
	go test -v -coverprofile cover.out ./...