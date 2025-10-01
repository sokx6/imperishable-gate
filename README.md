# Imperishable Gate

Imperishable Gate 是一个受东方 Project 启发的命令行链接管理系统，支持链接的添加、删除、查询与健康检查。后端使用 Go 语言结合 Echo 框架与 PostgreSQL，前端为 CLI 客户端。

## 目录结构

```
cmd/
  gate/             # CLI 客户端入口
  gate-server/      # 服务端入口
internal/
  client/cmd        # 客户端命令实现
  model/            # 数据模型与实体
  server/           # 服务端核心逻辑
    database/       # 数据库初始化与迁移
    handlers/       # HTTP 请求处理
    routes/         # 路由注册
pkg/
  protocol/         # 协议定义（预留）
```

## 环境要求

- Go 1.25.1
- gorm + postgres
- Echo
- Cobra

## 快速开始

### 1. 配置数据库

1. 创建数据库：
   ```sh
   createdb gate_db
   ```
2. 设置环境变量（可选）：
   ```sh
   export GATE_DSN='host=localhost user=postgres password=postgres dbname=gate_db port=5432 sslmode=disable TimeZone=Asia/Shanghai'
   ```

### 2. 启动服务端

```sh
go run cmd/gate-server/main.go start --port 8080 --dsn "$GATE_DSN"
```

### 3. 使用客户端

- **添加链接**
  ```sh
  go run cmd/gate/main.go add -H localhost:8080/api/v1/links/add -l "https://example.com"
  ```

- **列出链接**
  ```sh
  go run cmd/gate/main.go list -H localhost:8080/api/v1/links/list
  ```

- **删除链接**
  ```sh
  go run cmd/gate/main.go delete -H localhost:8080/api/v1/links/delete -l "https://example.com"
  ```

- **健康检查**
  ```sh
  go run cmd/gate/main.go ping -H localhost:8080/api/v1/ping -m "hello"
  ```

## HTTP API

| 方法 | 路径                       | 描述           |
| ---- | -------------------------- | -------------- |
| POST | `/api/v1/ping`             | 健康检查       |
| POST | `/api/v1/links/add`        | 添加链接       |
| DELETE | `/api/v1/links/delete`   | 删除链接（支持批量） |
| GET  | `/api/v1/links/list`       | 获取全部链接   |


© 2025 Imperishable Gate
