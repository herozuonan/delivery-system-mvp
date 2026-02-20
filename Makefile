.PHONY: dev test build docker-build

APP_NAME=delivery-system-mvp
IMAGE=ghcr.io/herozuonan/$(APP_NAME)
TAG?=dev

dev:
	go run ./cmd/delivery

test:
	go test ./...

build:
	go build ./cmd/delivery

docker-build:
	docker build -t $(IMAGE):$(TAG) .
