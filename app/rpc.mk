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

.PHONY: gorm
# 生成 GORM 数据库代码
gorm:
	@go run ../../cmd/gormgen/main.go -f configs/config.yaml

.PHONY: sqldump
# 导出sql文件
sqldump:
	@go run ../../cmd/sqldump/main.go -f configs/config.yaml

.PHONY: wire
# 生成 wire 依赖注入代码
wire:
	@go run -mod=mod github.com/google/wire/cmd/wire ./cmd/${APP_NAME}

.PHONY: proto
# 新增 protobuf 文件 make proto PROTO_NAME=demo
proto:
	@cd ../../ && kratos proto add api/${APP_NAME}/v1/${PROTO_NAME}.proto

.PHONY: api
# protobuf 生成 Go 代码
api:
	@cd ../../ && files=`find api/${APP_NAME} -name *.proto` && \
	protoc --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api \
 	       --go-http_out=paths=source_relative:./api \
 	       --go-grpc_out=paths=source_relative:./api \
 	       --go-errors_out=paths=source_relative:./api \
 	       --validate_out=paths=source_relative,lang=go:./api \
	       $$files
# protobuf 生成common Go 代码
common:
	@cd ../../ && files=`find api/paginator -name *.proto` && \
	protoc --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api \
 	       --validate_out=paths=source_relative,lang=go:./api \
	       $$files
.PHONY: service
# 通过 proto 文件，生成对应的 Service 实现代码 make service PROTO_NAME=demo
service:
	@kratos proto server ../../api/${APP_NAME}/v1/${PROTO_NAME}.proto -t internal/service

.PHONY: run
# run
run:
	@kratos run

.PHONY: app
# 多个命令同时执行
app: conf api  wire gorm

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
