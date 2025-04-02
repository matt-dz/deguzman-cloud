.PHONY: all run build docker-push clean

# Load environment variables
ifneq (,$(wildcard .env))
	include .env
	export
endif

IMAGE_NAME ?= kube-deploy
DOCKER_REGISTRY ?= matthew10125
TAG ?= latest
BUILD_NAME=$(DOCKER_REGISTRY)/$(IMAGE_NAME)
TARGETOS ?= linux
TARGETARCH ?= arm64
CGO_ENABLED ?= 0


all:
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(TARGETOS) GOARCH=$(TARGETARCH) go build -o bin/kube-deploy cmd/main.go

run:
	go run cmd/main.go

build:
	docker build --platform linux/arm64 -t $(BUILD_NAME) .

docker-push:
	docker tag $(BUILD_NAME) $(BUILD_NAME):$(TAG)
	docker push $(BUILD_NAME):$(TAG)

clean:
	rm bin/*
