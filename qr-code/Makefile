.PHONY: build push

DOCKER_REGISTRY?=matthew10125
IMAGE_NAME?=qrcode-frontend
BUILD_NAME?=$(DOCKER_REGISTRY)/$(IMAGE_NAME)
TAG?=latest

build:
	docker build -t $(BUILD_NAME) .

push:
	docker tag $(BUILD_NAME) $(BUILD_NAME)$(TAG)
	docker push $(BUILD_NAME):$(TAG)
