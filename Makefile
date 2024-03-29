SHELL := /bin/bash
.DEFAULT_GOAL := default
.PHONY: all

BINARY_NAME_NODE=node
BINARY_NAME_NS=ns
BINARY_NAME_CLIENT=client
BINARY_NAME_WEB=web

IMAGE_TAG=$(shell git describe --tags --always)
GIT_COMMIT=$(shell git rev-parse --short HEAD)
ORG_PREFIX := loqutus

tidy:
	go mod tidy

build:
	CGO_ENABLED=0 GOARCH=arm64 GOOS=darwin go build -ldflags "-X main.version=$(GIT_COMMIT)" -o bin/${BINARY_NAME_NODE}-darwin-arm64 cmd/${BINARY_NAME_NODE}/main.go
	CGO_ENABLED=0 GOARCH=arm64 GOOS=darwin go build -ldflags "-X main.version=$(GIT_COMMIT)" -o bin/${BINARY_NAME_NS}-darwin-arm64 cmd/${BINARY_NAME_NS}/main.go
	CGO_ENABLED=0 GOARCH=arm64 GOOS=darwin go build -ldflags "-X main.version=$(GIT_COMMIT)" -o bin/${BINARY_NAME_CLIENT}-darwin-arm64 cmd/${BINARY_NAME_CLIENT}/main.go
	CGO_ENABLED=0 GOARCH=arm64 GOOS=darwin go build -ldflags "-X main.version=$(GIT_COMMIT)" -o bin/${BINARY_NAME_WEB}-darwin-arm64 cmd/${BINARY_NAME_WEB}/main.go
	chmod +x bin/*

docker:
	docker system prune -a -f
	#docker buildx create --name multiarch --use || true
	docker build -t $(ORG_PREFIX)/ds0-$(BINARY_NAME_NS):$(IMAGE_TAG) -f Dockerfile-ns --push .
	docker build -t $(ORG_PREFIX)/ds0-$(BINARY_NAME_NODE):$(IMAGE_TAG) -f Dockerfile-node --push .
	docker build -t $(ORG_PREFIX)/ds0-$(BINARY_NAME_CLIENT):$(IMAGE_TAG) -f Dockerfile-client --push .
	docker build -t $(ORG_PREFIX)/ds0-$(BINARY_NAME_WEB):$(IMAGE_TAG) -f Dockerfile-web --push .
	docker build -t $(ORG_PREFIX)/ds0-test:$(IMAGE_TAG) -f Dockerfile-test --push .

helminstalltest:
	helm dependency build ./deployment/ds0
	helm install ds0 ./deployment/ds0 -n ds0-test --set image.tag=$(IMAGE_TAG)

testwait:
	kubectl wait --for=condition=complete --timeout=60s job/ds0-test -n ds0-test ; kubectl logs --since=1h job/ds0-test -n ds0-test

helmuninstalltest:
	kubectl delete ns ds0-test || true
	kubectl create ns ds0-test

test:
	go test -v -bench=. -benchtime=100x ./...

default: tidy build