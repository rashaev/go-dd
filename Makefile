GOBASE := $(shell pwd)
GOPATH := $(GOBASE)/vendor:$(GOBASE)
GOBIN := /usr/local/bin

build:
	go build -o bin/go-dd go-dd.go

format:
	gofmt -w go-dd.go

install:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go install

all: format build