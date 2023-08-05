# fkratos

kratos 微服务脚手架,采用大仓模式。

## 技术栈

### 框架：

* [Kratos](https://go-kratos.dev/) -- B站微服务框架

#### 中间件：

* [Redis](https://redis.io/) -- 非关系型数据库
* [PostgreSQL](https://www.postgresql.org/) -- 关系型数据库

### 配置中心:

### 服务注册与发现:

### 服务监控:

### 链路追踪:

### 日志:

### 熔断器:

* [sre breaker](https://github.com/go-kratos/aegis/tree/main/circuitbreaker/sre) -- 默认客户端熔断

### 限流器:

* [bbr limiter](https://github.com/go-kratos/aegis/tree/main/ratelimit/bbr) -- 默认限流器

### 参数校验:

* [protoc-gen-validate](https://github.com/bufbuild/protoc-gen-validate) -- proto验证器

### 错误处理:

### 接口文档

* [OpenAPI](https://www.openapis.org/) -- RESTful API 文档
* [ApiFox](https://apifox.com/) -- ApiFox 接口文档

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

致谢:

- [zeromicro/go-zero](https://github.com/zeromicro/go-zero)
- [tx7do/kratos-transport](https://github.com/tx7do/kratos-transport)