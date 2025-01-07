.PHONY: all build clean

include .env
export

BUILD_NAME=qrcode-backend

all:
	go run cmd/main.go

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bin/qr-generator cmd/main.go

docker:
	docker build -t $(BUILD_NAME) .
	docker tag $(BUILD_NAME) $(DOCKER_REGISTRY)/$(BUILD_NAME)
	docker push $(DOCKER_REGISTRY)/$(BUILD_NAME)

clean:
	rm bin/*
