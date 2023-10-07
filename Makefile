# Variables
DOCKER_COMPOSE 			= docker-compose
GO 						= go
NAME 					= caliber
VERSION 				= 0.0.1
DOCKER_IMAGE_NAME 		= $(NAME):$(VERSION)
MAIN                    = cmd/main.go

.PHONY: run run-dev run-docker build compose test lint download-tools coverage coverage-html

# Run the app
run: build
	./bin/$(NAME)

# Run the app in dev mode (assuming you have some dev configuration or flags)
run-dev:
	$(GO) run --race $(MAIN)

# Run the app in Docker
run-docker:
	docker run -p 8080:8080 $(DOCKER_IMAGE_NAME)

# Build the app
build:
	$(GO) build -o bin/$(NAME) $(MAIN)

# Run docker-compose
compose:
	$(DOCKER_COMPOSE) up --build

# Run all tests
test:
	$(GO) test ./...

# Run golangci-linter
lint:
	golangci-lint run

# Download tools
download-tools:
	$(GO) get github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Generate and view coverage summary in terminal
coverage:
	go test -v -coverprofile cover.out ./...