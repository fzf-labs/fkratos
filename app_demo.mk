GOPATH ?= $(shell go env GOPATH)

# 确保在运行构建过程之前设置 GOPATH。
ifeq "$(GOPATH)" ""
  $(error Please set the environment variable GOPATH before running `make`)
endif
FAIL_ON_STDOUT := awk '{ print } END { if (NR > 0) { exit 1 } }'

GO              := GO111MODULE=on go

ARCH      := "`uname -s`"
LINUX     := "Linux"
MAC       := "Darwin"

ifeq ($(OS),Windows_NT)
    IS_WINDOWS:=1
endif

APP_VERSION=$(shell git describe --tags --always)
APP_RELATIVE_PATH=$(shell a=`basename $$PWD` && cd .. && b=`basename $$PWD` && echo $$b/$$a)
APP_NAME=$(shell echo $(APP_RELATIVE_PATH) | rev |cut -d '/' -f 1 | rev)
APP_DOCKER_IMAGE=$(shell echo $(APP_NAME) |awk -F '@' '{print "fkratos/" $$0 ":0.1.0"}')


.PHONY:  dep vendor build clean docker conf ent wire api openapi run test cover vet lint app



# build golang application
build:
ifeq ("$(wildcard ./bin/)","")
	mkdir bin
endif
	@go build -ldflags "-X main.Service.Version=$(APP_VERSION)" -o ./bin/ ./...

# clean build files
clean:
	@go clean
	$(if $(IS_WINDOWS), del "coverage.out", rm -f "coverage.out")

# build docker image
docker:
	@docker build -t $(APP_DOCKER_IMAGE) . \
				  -f ../../../.docker/Dockerfile \
				  --build-arg APP_RELATIVE_PATH=$(APP_RELATIVE_PATH) GRPC_PORT=9000 REST_PORT=8000

# generate config define code
conf:
	protoc --proto_path=./internal/conf/ \
	       --proto_path=../../third_party \
	       --go_out=paths=source_relative:./internal/conf/ \
	       ./internal/conf/conf.proto

# generate gorm code
gorm:
	@go run ../../cmd/gormgen/main.go -f configs/config.yaml -s ${APP_NAME}

# generate wire code
wire:
	@go run -mod=mod github.com/google/wire/cmd/wire ./cmd/${APP_NAME}

# generate protobuf api go code
api:
	@cd ../../../ && \
	buf generate

# generate OpenAPI v3 doc
openapi:
	@cd ../../../ && \
	buf generate --path api/admin/service/v1 --template api/admin/service/v1/buf.openapi.gen.yaml

# run application
run:
	@go run ./cmd/server -conf ./configs

# run tests
test:
	@go test ./...

# run coverage tests
cover:
	@go test -v ./... -coverprofile=coverage.out

# run static analysis
vet:
	@go vet

# run lint
lint:
	@golangci-lint run

# build service app
app: api wire conf ent build

# show help
help:
	@echo ""
	@echo "Usage:"
	@echo " make [target]"
	@echo ""
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
