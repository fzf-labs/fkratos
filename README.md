# fkratos

一个Golang的博客系统/CMS。

## 技术栈

* [Kratos](https://go-kratos.dev/) -- B站微服务框架
* [OpenTelemetry](https://opentelemetry.io/) -- 分布式可观察系统
* [Wire](https://github.com/google/wire) -- 依赖注入框架
* [OpenAPI](https://www.openapis.org/) -- RESTful API 文档
* [Redis](https://redis.io/) -- 非关系型数据库
* [PostgreSQL](https://www.postgresql.org/) -- 关系型数据库


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

- 生成服务器配置结构代码

   ```bash
   make conf
   ```

- 生成gorm代码

   ```bash
   make gorm
   ```

- 生成wire代码

   ```bash
   make wire
   ```