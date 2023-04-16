GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)


.PHONY: init
# 初始化环境
init:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	@go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	@go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	@go install github.com/envoyproxy/protoc-gen-validate@latest
	@go install github.com/bufbuild/buf/cmd/buf@latest
	@go install github.com/google/gnostic@latest
	@go install entgo.io/ent/cmd/ent@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/go-kratos/kratos/cmd/kratos/v2@latest

.PHONY: dep
# 下载模块依赖
dep:
	@go mod download
	@go mod tidy

.PHONY: new
# 新增加一个服务  make new APP_NAME=rpc_user
new:
	@kratos new app/${APP_NAME} --nomod

# 生成配置文件
conf:
	@files=`find ./internal/bootstrap/conf -name *.proto` && \
	protoc --proto_path=./internal/bootstrap/conf \
	       --proto_path=./third_party \
	       --go_out=paths=source_relative:./internal/bootstrap/conf \
	       --validate_out=paths=source_relative,lang=go:./internal/bootstrap/conf \
	       $$files

.PHONY: fmt
# 格式化代码
fmt:
	@gofmt -s -w .

.PHONY: vet
# 代码检查 vet
vet:
	@go vet ./...

.PHONY: ci-lint
# 代码检查 golangci-lint
ci-lint:
	@golangci-lint run ./...

# git 记录清除
git-clean:
	#清除开始
	@git checkout --orphan latest_branch
	@git add -A
	@git commit -am "clean"
	@git branch -D ${gitBranch}
	@git branch -m ${gitBranch}
	@git push -f origin ${gitBranch}
	#清除结束


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
