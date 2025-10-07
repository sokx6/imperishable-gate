# Imperishable Gate

<div align="center">

![Go Version](https://img.shields.io/badge/Go-1.25.1+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-green.svg)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Database-336791?style=flat&logo=postgresql)
![Echo](https://img.shields.io/badge/Echo-Web_Framework-00C7B7?style=flat)
![Cobra](https://img.shields.io/badge/Cobra-CLI_Framework-blue?style=flat)

</div>

**Imperishable Gate（不朽之门）** 是一个受东方 Project 启发的现代化命令行链接管理系统。它提供了完整的链接生命周期管理功能，包括添加、删除、查询、标签分类、别名管理、备注、元数据自动抓取和智能链接监控等功能。

## 📖 项目简介

本项目采用**前后端分离架构**：
- **后端服务**：基于 Go + Echo + PostgreSQL + GORM 构建的高性能 RESTful API 服务
- **CLI 客户端**：基于 Cobra 框架的强大命令行工具，支持跨平台使用

## ✨ 核心特性

### 🔐 企业级(?)安全认证系统
- **双令牌机制**：JWT Access Token（短期） + Refresh Token（长期）
- **密码加密**：采用 bcrypt 算法安全存储用户密码
- **安全存储**：令牌存储在系统 keyring（支持 Linux/macOS/Windows）
- **自动刷新**：令牌过期自动刷新，提供无感知的用户体验
- **数据隔离**：完善的用户权限管理，数据完全隔离

### 🔗 强大(?)的链接管理功能
- **多维度检索**：支持通过 URL、标签、别名进行快速查询
- **批量操作**：一次性添加或删除多个链接、标签、别名
- **智能抓取**：自动获取网页元数据（标题、描述、关键词）
- **状态追踪**：实时记录 HTTP 状态码，监控链接健康状况
- **关联管理**：链接、标签、别名之间的灵活关联

### 🏷️ 灵活的标签系统(可能会报错就是了)
- **多对多关联**：一个链接可以有多个标签，一个标签可以关联多个链接
- **标签分类**：按标签分类管理和检索链接
- **批量标签操作**：支持通过 URL 或别名批量添加/删除标签
- **用户隔离**：标签在用户范围内唯一，互不干扰

### 📝 别名与备注管理
- **自定义别名**：为链接设置易记的自定义别名（全局唯一）
- **快速定位**：通过别名快速查找和操作链接
- **详细备注**：支持为每个链接添加详细的备注信息
- **灵活操作**：可通过别名或 URL 执行各种操作

### 👀 智能监控系统（并非智能）
- **定时监控**：后台定时检查链接状态和内容变化
- **自动检测**：自动发现网页内容更新
- **分级监控**：支持 watching（高频）和 non-watching（低频）两种监控模式
- **变化通知**：内容变化时自动通知（支持邮件等多种方式）
- **状态管理**：灵活设置链接的监控状态

### � 便捷的 CLI 体验（CLI对许多人来说真的便捷吗）
- **交互式界面**：友好的命令行交互，简单易用
- **智能提示**：清晰的错误提示和操作指引
- **自动认证**：智能令牌管理，无需频繁登录
- **跨平台支持**：支持 Linux、macOS、Windows 系统
- **丰富命令**：提供完整的 CRUD 操作命令集

## 📋 目录

- [🏗️ 架构设计](#🏗️-架构设计)
- [📂 目录结构](#📂-目录结构)
- [🔧 环境要求](#🔧-环境要求)
- [⚙️ 配置说明](#⚙️-配置说明)
- [🚀 快速开始](#🚀-快速开始)
- [📋 HTTP API](#📋-http-api)
- [📊 数据模型](#📊-数据模型)
- [🛠️ 技术特性](#🛠️-技术特性)
- [🔒 安全特性](#🔒-安全特性)
- [📅 开发进度](#📅-开发进度)
- [🤝 贡献指南](#🤝-贡献指南)
- [📄 许可证](#📄-许可证)

## 🏗️ 架构设计

```
┌─────────────────────────────────────────────────────────────┐
│                      CLI Client (gate)                       │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────────┐  │
│  │   Commands   │  │   Services   │  │  System Keyring  │  │
│  └──────┬───────┘  └──────┬───────┘  └────────┬─────────┘  │
│         │                  │                    │            │
│         └──────────────────┴────────────────────┘            │
│                            │                                 │
│                       HTTP/JSON                              │
└────────────────────────────┼────────────────────────────────┘
                             │
┌────────────────────────────┼────────────────────────────────┐
│                       RESTful API                            │
│  ┌────────────────────────┴──────────────────────────────┐  │
│  │              Echo Web Framework                        │  │
│  │  ┌──────────┐  ┌────────────┐  ┌─────────────────┐   │  │
│  │  │  Routes  │→ │ Middleware │→ │    Handlers     │   │  │
│  │  └──────────┘  └────────────┘  └────────┬────────┘   │  │
│  │                                          │             │  │
│  └──────────────────────────────────────────┼────────────┘  │
│                                             │                │
│  ┌──────────────────────────────────────────┼────────────┐  │
│  │                  Services                │             │  │
│  │  ┌────────────┐  ┌──────────┐  ┌────────┴─────────┐  │  │
│  │  │    JWT     │  │ Metadata │  │   Link Monitor   │  │  │
│  │  │  Service   │  │  Crawler │  │  (Goroutines)    │  │  │
│  │  └────────────┘  └──────────┘  └──────────────────┘  │  │
│  └────────────────────────────────────────────────────────┘ │
│                              │                               │
│                         GORM ORM                             │
└──────────────────────────────┼──────────────────────────────┘
                               │
┌──────────────────────────────┼──────────────────────────────┐
│                       PostgreSQL                             │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────┐    │
│  │  users   │  │  links   │  │   tags   │  │  names   │    │
│  └──────────┘  └──────────┘  └──────────┘  └──────────┘    │
│  ┌──────────────────────┐  ┌──────────────────────┐        │
│  │   refresh_tokens     │  │     link_tags        │        │
│  └──────────────────────┘  └──────────────────────┘        │
└─────────────────────────────────────────────────────────────┘
```

## 📂 目录结构

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

## 🔧 环境要求

- Go 1.25.1+
- PostgreSQL 数据库
- GORM + PostgreSQL 驱动
- Echo Web 框架
- Cobra CLI 框架

### Linux 系统 Keyring 要求

客户端使用系统 keyring 安全存储令牌。在 Linux 系统上，需要安装相应的 keyring 服务：

#### Ubuntu/Debian
```sh
sudo apt-get update
sudo apt-get install gnome-keyring libsecret-1-dev
```

#### Fedora/RHEL/CentOS
```sh
sudo dnf install gnome-keyring libsecret-devel
```

#### Arch Linux
```sh
sudo pacman -S gnome-keyring libsecret
```

> **注意**：
> - 如果您使用的是桌面环境（如 GNOME、KDE），通常已经预装了 keyring 服务
> - 对于无图形界面的服务器环境，可能需要手动启动 keyring 守护进程
> - macOS 和 Windows 系统无需额外安装，会自动使用系统的 Keychain 和 Credential Manager

## ⚙️ 配置说明

### 环境变量配置

项目使用 `.env` 文件管理配置。请按以下步骤配置：

1. **复制环境变量模板文件**
   ```sh
   cp .env.example cmd/gate-server/.env
   ```

2. **编辑配置文件**
   ```sh
   vim cmd/gate-server/.env
   ```

### 配置项说明

#### 📊 数据库配置

| 环境变量 | 说明 | 示例值 | 必需 |
|---------|------|--------|------|
| `DSN` | PostgreSQL 数据库连接字符串 | `host=localhost user=postgres password=postgres dbname=gate_db port=5432 sslmode=disable TimeZone=Asia/Shanghai` | ✅ |

#### 🌐 服务器配置

| 环境变量 | 说明 | 示例值 | 必需 |
|---------|------|--------|------|
| `SERVER_ADDR` | 服务器监听地址 | `localhost:8080` 或 `:8080` | ✅ |

#### 🔐 JWT 安全配置

| 环境变量 | 说明 | 示例值 | 必需 |
|---------|------|--------|------|
| `JWT_SECRET` | JWT 签名密钥（生产环境务必修改！） | 使用 `openssl rand -base64 64` 生成 | ⚠️ 推荐 |

> **安全提示**：
> - 生产环境务必设置强随机 `JWT_SECRET`
> - 使用命令生成安全密钥：`openssl rand -base64 64`
> - 切勿将包含真实密钥的 `.env` 文件提交到版本控制系统

#### 📧 邮件服务配置（可选）

用于链接监控变化通知功能：

| 环境变量 | 说明 | 示例值 | 必需 |
|---------|------|--------|------|
| `EMAIL_HOST` | SMTP 服务器地址 | `smtp.gmail.com` | 📧 |
| `EMAIL_PORT` | SMTP 服务器端口 | `587` (TLS) 或 `465` (SSL) | 📧 |
| `EMAIL_FROM` | 发件人邮箱地址 | `noreply@example.com` | 📧 |
| `EMAIL_PASSWORD` | 邮箱密码或授权码 | `your-app-password` | 📧 |

> **邮箱配置提示**：
> - Gmail：使用应用专用密码（[获取方法](https://support.google.com/accounts/answer/185833)）
> - 163/QQ邮箱：使用授权码而非登录密码
> - 端口选择：587（STARTTLS）或 465（SSL/TLS）

### 配置文件位置

```
cmd/gate-server/.env    # 服务端配置文件（需手动创建）
.env.example            # 配置模板文件（不含敏感信息）
```

### 配置优先级

1. 命令行参数（最高优先级）
2. 环境变量（`.env` 文件）
3. 默认值（最低优先级）

**示例**：
```sh
# 方式 1: 使用 .env 文件配置
./bin/gate-server start

# 方式 2: 使用命令行参数覆盖配置
./bin/gate-server start --addr ":9090" --dsn "host=localhost..."

# 方式 3: 使用环境变量
export DSN="host=localhost..."
./bin/gate-server start
```

## 🚀 快速开始

### 1. 获取项目

```sh
git clone https://github.com/sokx6/imperishable-gate.git
cd imperishable-gate
```

### 2. 安装依赖

```sh
go mod download
```

### 3. 配置数据库

创建 PostgreSQL 数据库：

```sh
# 使用 psql 创建数据库
createdb gate_db

# 或使用 SQL 命令
psql -U postgres -c "CREATE DATABASE gate_db;"
```

### 4. 配置环境变量

```sh
# 复制配置模板
cp .env.example cmd/gate-server/.env

# 编辑配置文件，设置数据库连接等信息
vim cmd/gate-server/.env
```

**必需配置项**：
- `DSN`：数据库连接字符串
- `SERVER_ADDR`：服务器监听地址

**推荐配置项**：
- `JWT_SECRET`：自定义 JWT 密钥（生产环境必需）

### 5. 构建二进制

```sh
# 构建服务端和客户端
go build -o bin/gate-server ./cmd/gate-server
go build -o bin/gate ./cmd/gate

# 或使用简化命令
make build  # 如果项目提供了 Makefile
```

### 6. 启动服务端

```sh
# 使用 .env 文件配置启动
cd cmd/gate-server
../../bin/gate-server start

# 或使用命令行参数启动
./bin/gate-server start --addr ":8080" --dsn "host=localhost user=postgres password=postgres dbname=gate_db port=5432 sslmode=disable"
```

服务端成功启动后，会显示：
```
Starting Imperishable Gate server on :8080...
```

### 7. 使用客户端

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

## 📋 HTTP API

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

## 📊 数据模型

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

## 🛠️ 技术特性

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

## 🔒 安全特性

- 🔒 密码使用 bcrypt 加密存储
- 🎫 JWT 访问令牌（短期有效）+ Refresh Token（长期有效）双令牌机制
- 🔑 令牌存储在系统 keyring，不存储在配置文件
- 🛡️ 所有业务 API 都需要 JWT 认证
- 👤 用户数据隔离，仅能访问自己的数据

## 📅 开发进度

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
- ✅ 定时任务：监控链接变化通知

### 待完成功能
- ⏳ 响应格式统一化
- ⏳ 更完善的错误处理
- ⏳ 单元测试和集成测试
- ⏳ 配置文件支持
- ⏳ Docker 部署支持
- ⏳ API 文档（Swagger）

## 🤝 贡献指南

欢迎提交 Issue 和 Pull Request！

## 📄 许可证

MIT License
