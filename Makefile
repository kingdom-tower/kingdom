.PHONY: vendor

vendor:
	go mod vendor
	go mod tidy

check:
	go fmt ./...

build:
	go build -mod=vendor ./...

install:
	go install -mod=vendor ./...
