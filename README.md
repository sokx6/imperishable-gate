# Imperishable Gate

Imperishable Gate 是一个受东方 Project 启发的命令行链接管理系统，提供链接的添加、删除、查询、标签管理、健康检查和元数据抓取等功能。后端使用 Go 语言结合 Echo 框架与 PostgreSQL，前端为 CLI 客户端，支持 JWT 身份认证。

## 功能特性

- 🔐 **用户认证系统** - JWT + Refresh Token 双令牌机制
- 🔗 **链接管理** - 添加、删除、查询链接，支持多维度检索
- 🏷️ **标签系统** - 为链接添加标签，按标签分类管理
- 📝 **别名功能** - 为链接设置自定义名称，快速访问
- 📋 **备注功能** - 为链接添加备注信息
- 🔍 **元数据抓取** - 自动获取网页标题、描述和关键词
- 👀 **监控功能** - 定时检查链接状态和内容变化
- 🔑 **安全存储** - 使用系统 keyring 安全存储令牌

## 目录结构

```
cmd/
  gate/             # CLI 客户端入口
  gate-server/      # 服务端入口
internal/
  client/           # 客户端实现
    cmd/            # CLI 命令定义
    service/        # 客户端业务逻辑
    utils/          # 客户端工具函数
  model/            # 数据模型与实体
  server/           # 服务端核心逻辑
    database/       # 数据库初始化与迁移
    handlers/       # HTTP 请求处理
    middlewares/    # 中间件（JWT 认证等）
    routes/         # 路由注册
    service/        # 服务端业务逻辑
    utils/          # 服务端工具函数
  types/            # 类型定义
    request/        # 请求类型
    response/       # 响应类型
    jwt/            # JWT 相关类型
    data/           # 数据类型
```

## 环境要求

- Go 1.25.1+
- PostgreSQL 数据库
- GORM + PostgreSQL 驱动
- Echo Web 框架
- Cobra CLI 框架

## 快速开始

### 1. 获取项目

```sh
git clone https://github.com/sokx6/imperishable-gate.git
cd imperishable-gate
```

### 2. 构建二进制

```sh
mkdir -p bin
go build -o bin/gate-server ./cmd/gate-server
go build -o bin/gate ./cmd/gate
```

### 3. 配置数据库

1. 创建数据库：
   ```sh
   createdb gate_db
   ```
2. 设置环境变量（可选）：
   ```sh
   export GATE_DSN='host=localhost user=postgres password=postgres dbname=gate_db port=5432 sslmode=disable TimeZone=Asia/Shanghai'
   ```

### 4. 启动服务端

```sh
./bin/gate-server start --port 8080 --dsn "$GATE_DSN"
```

### 5. 使用客户端

#### 用户认证

- **注册用户**
  ```sh
  ./bin/gate register
  ```

- **登录**
  ```sh
  ./bin/gate login
  ```

#### 链接管理

- **添加链接**
  ```sh
  ./bin/gate add -l "https://example.com"
  # 添加链接并设置备注
  ./bin/gate add -l "https://example.com" -r "我的示例网站"
  ```

- **列出所有链接**
  ```sh
  ./bin/gate list
  ```

- **通过名称查询链接**
  ```sh
  ./bin/gate list -n "mylink"
  ```

- **通过标签查询链接**
  ```sh
  ./bin/gate list -t "tech"
  ```

- **删除链接**
  ```sh
  ./bin/gate delete -l "https://example.com"
  # 通过名称删除
  ./bin/gate delete -n "mylink"
  ```

#### 标签和名称管理

- **为链接添加标签**
  ```sh
  # 通过 URL 添加
  ./bin/gate add -l "https://example.com" -t "tech,news"
  # 通过名称添加
  ./bin/gate add -n "mylink" -t "tech,news"
  ```

- **为链接添加别名**
  ```sh
  ./bin/gate add -l "https://example.com" -N "mylink"
  ```

- **为链接添加备注**
  ```sh
  # 通过 URL 添加
  ./bin/gate add -l "https://example.com" -r "这是一个备注"
  # 通过名称添加
  ./bin/gate add -n "mylink" -r "这是一个备注"
  ```

#### 监控管理

- **设置链接监控**
  ```sh
  # 通过 URL 设置
  ./bin/gate watch -l "https://example.com" -w true
  # 通过名称设置
  ./bin/gate watch -n "mylink" -w true
  ```

#### 系统检查

- **健康检查**
  ```sh
  ./bin/gate ping -m "hello"
  ```

## HTTP API

### 公共路由（无需认证）

| 方法 | 路径 | 描述 |
| ---- | ---- | ---- |
| POST | `/api/v1/register` | 用户注册 |
| POST | `/api/v1/login` | 用户登录 |
| POST | `/api/v1/ping` | 健康检查 |
| POST | `/api/v1/refresh` | 刷新访问令牌 |
| POST | `/api/v1/logout` | 用户登出 |

### 受保护路由（需要 JWT 认证）

#### 链接查询

| 方法 | 路径 | 描述 |
| ---- | ---- | ---- |
| GET | `/api/v1/links` | 获取当前用户的所有链接 |
| GET | `/api/v1/names/:name` | 通过名称获取链接 |
| GET | `/api/v1/tags/:tag` | 通过标签获取链接 |

#### 链接添加

| 方法 | 路径 | 描述 |
| ---- | ---- | ---- |
| POST | `/api/v1/links` | 添加新链接 |
| POST | `/api/v1/names` | 为链接添加名称 |
| POST | `/api/v1/remarks` | 通过 URL 添加或更新备注 |
| POST | `/api/v1/name/:name/remark` | 通过名称添加或更新备注 |
| POST | `/api/v1/tags` | 通过 URL 添加标签 |
| POST | `/api/v1/name/:name/tags` | 通过名称添加标签 |

#### 链接更新

| 方法 | 路径 | 描述 |
| ---- | ---- | ---- |
| PATCH | `/api/v1/links/watch` | 通过 URL 设置监控状态 |
| PATCH | `/api/v1/name/watch` | 通过名称设置监控状态 |
| PATCH | `/api/v1/links/names/remove` | 通过 URL 移除名称 |
| PATCH | `/api/v1/links/by-url/tags/remove` | 通过 URL 移除标签 |
| PATCH | `/api/v1/:name/tags/remove` | 通过名称移除标签 |

#### 链接删除

| 方法 | 路径 | 描述 |
| ---- | ---- | ---- |
| DELETE | `/api/v1/links` | 通过 URL 批量删除链接 |
| DELETE | `/api/v1/links/name/:name` | 通过名称删除链接 |

## 数据模型

### User（用户）
- ID：主键
- Username：用户名（唯一）
- Password：密码（加密存储）
- Email：邮箱
- Links：关联的链接列表
- Tags：关联的标签列表

### Link（链接）
- ID：主键
- UserID：所属用户 ID
- Url：链接地址（用户内唯一）
- Tags：多对多关联的标签
- Names：一对多关联的别名
- Remark：备注信息
- Title：网页标题（自动抓取）
- Description：网页描述（自动抓取）
- Keywords：网页关键词（自动抓取）
- Watching：是否监控
- StatusCode：HTTP 状态码

### Tag（标签）
- ID：主键
- UserID：所属用户 ID
- Name：标签名（用户内唯一）
- Links：多对多关联的链接

### Name（别名）
- ID：主键
- LinkID：所属链接 ID
- Name：别名（全局唯一）

## 技术特性

### 后端
- **Web 框架**：Echo v4 - 高性能 Go Web 框架
- **ORM**：GORM - 功能强大的 Go ORM
- **数据库**：PostgreSQL - 可靠的关系型数据库
- **认证**：JWT（golang-jwt/jwt v5）+ Refresh Token 机制
- **密码加密**：bcrypt 加密存储
- **网页抓取**：goquery - 自动获取网页元数据
- **定时任务**：支持监控链接的定时检查

### 客户端
- **CLI 框架**：Cobra - 强大的 CLI 应用构建工具
- **凭证存储**：go-keyring - 系统级安全存储（支持 Linux、macOS、Windows）
- **自动认证**：智能令牌管理，自动刷新过期令牌
- **用户友好**：交互式输入，清晰的错误提示

## 安全特性

- 🔒 密码使用 bcrypt 加密存储
- 🎫 JWT 访问令牌（短期有效）+ Refresh Token（长期有效）双令牌机制
- 🔑 令牌存储在系统 keyring，不存储在配置文件
- 🛡️ 所有业务 API 都需要 JWT 认证
- 👤 用户数据隔离，仅能访问自己的数据

## 开发进度

### 已完成功能
- ✅ 用户注册、登录、登出
- ✅ JWT 认证中间件
- ✅ 链接的增删改查
- ✅ 标签管理（添加、删除、查询）
- ✅ 别名管理（添加、删除、查询）
- ✅ 备注功能
- ✅ 网页元数据自动抓取
- ✅ 链接监控状态管理
- ✅ CLI 客户端基础功能
- ✅ 自动令牌刷新

### 待完成功能
- ⏳ 定时任务：监控链接变化通知
- ⏳ 响应格式统一化
- ⏳ 更完善的错误处理
- ⏳ 单元测试和集成测试
- ⏳ 配置文件支持
- ⏳ Docker 部署支持
- ⏳ API 文档（Swagger）

## 贡献指南

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT License
