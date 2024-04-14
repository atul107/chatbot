# Makefile for the Go chatbot application

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=chatbot
BINARY_UNIX=$(BINARY_NAME)_unix

# Docker parameters
DOCKER_BUILD=docker build
DOCKER_RUN=docker run
IMAGE_NAME=chatbot_image
CONTAINER_NAME=chatbot_container
PORT=8080

all: test build
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/server
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run: build
	./$(BINARY_NAME)
setup: 
	scripts/./setup.sh
docker-build:
	$(DOCKER_BUILD) -t $(IMAGE_NAME) .
docker-run: docker-build
	$(DOCKER_RUN) --name $(CONTAINER_NAME) -p $(PORT):$(PORT) --rm -d $(IMAGE_NAME)
docker-stop:
	docker stop $(CONTAINER_NAME)
