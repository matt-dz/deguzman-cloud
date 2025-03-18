.PHONY: build docker-push

IMAGE_NAME ?= cop3530
DOCKER_REGISTRY ?= matthew10125
TAG ?= latest
BUILD_NAME=$(DOCKER_REGISTRY)/$(IMAGE_NAME)
PLATFORM ?= linux/arm64


build:
	 docker buildx build --platform $(PLATFORM) -t $(BUILD_NAME) .

docker-push:
	docker tag $(BUILD_NAME) $(BUILD_NAME):$(TAG)
	docker push $(BUILD_NAME):$(TAG)
