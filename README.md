# fkratos

一个Golang的博客系统/CMS。

## 技术栈

* [Kratos](https://go-kratos.dev/) -- B站微服务框架
* [OpenTelemetry](https://opentelemetry.io/) -- 分布式可观察系统
* [Wire](https://github.com/google/wire) -- 依赖注入框架
* [OpenAPI](https://www.openapis.org/) -- RESTful API 文档
* [Redis](https://redis.io/) -- 非关系型数据库
* [PostgreSQL](https://www.postgresql.org/) -- 关系型数据库

## 生成Protobuf API

使用[buf.build](https://buf.build/)进行Protobuf API的构建。

相关命令行工具和插件的具体安装方法请参见：[Kratos微服务框架API工程化指南](https://juejin.cn/post/7191095845096259641)

在`blog-backend`根目录下执行命令：

- 更新buf.lock

    ```bash
    buf mod update
    ```

- 生成GO代码

    ```bash
    buf generate
    ```

- 生成OpenAPI v3文档

    ```bash
    buf generate --path api/admin/service/v1 --template api/admin/service/v1/buf.openapi.gen.yaml
    ```

## Make构建

请在`app/{服务名}/service`下执行：

- 初始化开发环境

   ```bash
   make init
   ```

- 生成API的go代码

   ```bash
   make api
   ```

- 生成API的OpenAPI v3 文档

   ```bash
   make openapi
   ```

- 生成服务器配置结构代码

   ```bash
   make conf
   ```

- 生成ent代码

   ```bash
   make ent
   ```

- 生成wire代码

   ```bash
   make wire
   ```

- 构建程序

   ```bash
   make build
   ```

- 构建Docker镜像

   ```bash
   make docker
   ```