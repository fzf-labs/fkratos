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
BUF_INSTALLED := $(shell command -v buf 2> /dev/null)

.PHONY: conf
# 生成配置文件
conf:
	protoc --proto_path=./internal/conf/ \
	       --proto_path=../../third_party \
	       --go_out=paths=source_relative:./internal/conf/ \
	       ./internal/conf/conf.proto

.PHONY: wire
# 生成 wire 依赖注入代码
wire:
	@go run -mod=mod github.com/google/wire/cmd/wire ./cmd/server

.PHONY: api
# protobuf 生成 Go 代码
api: buf
	@cd ../../ && files=`find api/${APP_NAME} -name *.proto` && \
	protoc --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api \
 	       --go-http_out=paths=source_relative:./api \
 	       --go-grpc_out=paths=source_relative:./api \
 	       --go-errors_out=paths=source_relative:./api \
 	       --validate_out=paths=source_relative,lang=go:./api \
 	       --openapiv2_out ./api \
 	       --openapiv2_opt logtostderr=true \
 	       --openapiv2_opt json_names_for_fields=false \
	       $$files

.PHONY: apidoc
# apidoc 同步文档
apidoc:
	@go run ../../cmd/apidoc/main.go -f configs/config.yaml

.PHONY: buf
# buf 格式化 proto
buf:
	@if [ -n "$(BUF_INSTALLED)" ]; then \
        cd ../../api/${APP_NAME} && \
        buf format -w && \
        echo "proto format finish"; \
    else \
        echo "please installation buf: https://buf.build/docs/installation"; \
    fi

.PHONY: common
common:
	@cd ../../ && files=`find api/paginator -name *.proto` && \
	protoc --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api \
 	       --validate_out=paths=source_relative,lang=go:./api \
	       $$files

.PHONY: protocode
# 通过 proto 文件，生成对应的 data,biz,service 代码 make protocode
protocode:
	@echo "proto code start";
	@go run ../../cmd/protocode/main.go biz -p ../../api/${APP_NAME} -d internal/biz
	@go run ../../cmd/protocode/main.go service -p ../../api/${APP_NAME} -d internal/service
	@echo "proto code finish";


.PHONY: run
# run
run:
	@kratos run

.PHONY: app
# 多个命令同时执行
app: conf api wire

.PHONY: help
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
