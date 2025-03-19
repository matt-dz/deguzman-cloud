.PHONY: all run build docker-push clean

# Load environment variables
ifneq (,$(wildcard .env))
    include .env
    export
endif

IMAGE_NAME ?= dga
DOCKER_REGISTRY ?= matthew10125
TAG ?= latest
BUILD_NAME=$(DOCKER_REGISTRY)/$(IMAGE_NAME)
PLATFORM ?= linux/arm64


all:
	go build -o bin/dga cmd/main.go

run:
	go run cmd/main.go

build:
	 docker buildx build --platform $(PLATFORM) -t $(BUILD_NAME) .

docker-push:
	docker tag $(BUILD_NAME) $(BUILD_NAME):$(TAG)
	docker push $(BUILD_NAME):$(TAG)

clean:
	rm bin/*
