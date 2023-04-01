SHELL := /bin/bash
.DEFAULT_GOAL := default
.PHONY: all
BINARY_NAME_NODE=node
BINARY_NAME_NS=ns
BINARY_NAME_CLIENT=client
IMAGE_TAG=$(shell git describe --tags --always)
GIT_COMMIT=$(shell git rev-parse --short HEAD)
ORG_PREFIX := loqutus

tidy:
	go mod tidy

build:
	CGO_ENABLED=0 GOARCH=arm64 GOOS=darwin go build -o bin/${BINARY_NAME_NODE}-darwin-arm64 cmd/${BINARY_NAME_NODE}/main.go
	CGO_ENABLED=0 GOARCH=arm64 GOOS=darwin go build -o bin/${BINARY_NAME_NS}-darwin-arm64 cmd/${BINARY_NAME_NS}/main.go
	CGO_ENABLED=0 GOARCH=arm64 GOOS=darwin go build -o bin/${BINARY_NAME_CLIENT}-darwin-arm64 cmd/${BINARY_NAME_CLIENT}/main.go
	chmod +x bin/*

docker:
	docker build -t loqutus/ds0-$(BINARY_NAME_NS):$(IMAGE_TAG) -f Dockerfile-ns .
	docker push loqutus/ds0-$(BINARY_NAME_NS):$(IMAGE_TAG)
	docker build -t loqutus/ds0-$(BINARY_NAME_NODE):$(IMAGE_TAG) -f Dockerfile-node .
	docker push loqutus/ds0-$(BINARY_NAME_NODE):$(IMAGE_TAG)
	docker build -t loqutus/ds0-$(BINARY_NAME_CLIENT):$(IMAGE_TAG) -f Dockerfile-client .
	docker push loqutus/ds0-$(BINARY_NAME_CLIENT):$(IMAGE_TAG)

helminstalltest:
	helm install ds0 ./deployment/ds0 -n ds0-test --set image.tag=$(IMAGE_TAG)

helmuninstalltest:
	helm uninstall ds0 -n ds0-test || true

test:
	go test -v ./...

default: tidy build