.PHONY: all

include .build.env
export

BUILD_NAME=qrcode-frontend
BUILD_TAG=$(DOCKER_REGISTRY)/$(BUILD_NAME)

all:
	docker build -t $(BUILD_TAG) .
	docker push $(BUILD_TAG)
