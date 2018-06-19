export ROOT=$(realpath $(dir $(firstword $(MAKEFILE_LIST))))
export APP_NAME=ci
export BIN=$(ROOT)/bin
export GO=$(shell which go)
export GOBIN?=$(BIN)
export GOPATH=$(abspath $(ROOT)/../../..)
export BUILD=cd $(ROOT) && $(GO) install -v $(LD_ARGS)
all:
	GOPATH=$(GOPATH) $(BUILD) test.com/ci/cmd/...
	echo $(USER)
run-server: all
	$(BIN)/server

deploy:

