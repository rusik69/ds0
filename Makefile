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

default: tidy build