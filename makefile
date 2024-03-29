GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
API_PROTO_FILES=$(shell find api -name *.proto)

.PHONY: init
# init env
init:
	go mod tidy

.PHONY: config
# generate internal proto
config:
	./proto.sh conf

.PHONY: api
# generate api proto
api:
	./proto.sh api

.PHONY: build-interface
# build
build:
	#mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...
	cd ./app/interface/service && ./build.sh

.PHONY: build-interface
# build
build:
	#mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...
	cd ./app/interface/service && ./build.sh

.PHONY: all
# generate all
all:
	make api;
	make errors;
	make config;
	make generate;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
