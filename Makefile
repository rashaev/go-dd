GOBASE := $(shell pwd)
GOPATH := $(GOBASE)/vendor:$(GOBASE)
GOBIN := /usr/local/bin

build:
	go build -o bin/go-dd go-dd.go

format:
	gofmt -w go-dd.go

test: clean
	go test -v 

clean:
	@rm -f testdata/testfiledst
	@rm -f testdata/testfiledst_limit
	@rm -f testdata/testfiledst_offset

install:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go install

all: format build
