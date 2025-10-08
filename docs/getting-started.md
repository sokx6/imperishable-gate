# 快速开始 | Stage 1: 白銀之春

**[📖 简体中文](getting-started.md) | [📘 English](getting-started.en.md)**

> 🌸 *"首先来做一些准备工作吧！"* 🌸

本指南将帮助您快速搭建和运行 **Imperishable Gate（不朽之门）**。从零开始，完成 **Stage 1「白銀之春」** 的基础框架搭建，让您的白玉楼庭师身份正式开始管理链接！

## 前置要求

在开始之前，请确保您的系统满足以下要求：

- Go 1.25.1 或更高版本
- 数据库（三选一）：
  - **SQLite**（默认，无需额外安装）✨ 推荐用于快速开始
  - **MySQL** 5.7+ / 8.0+
  - **PostgreSQL** 12.0+
- Git

详细的环境要求请参考 [配置文档](configuration.md)。

## 安装步骤

### 1. 获取项目

```sh
git clone https://github.com/sokx6/imperishable-gate.git
cd imperishable-gate
```

### 2. 安装依赖

```sh
go mod download
```

### 3. 配置数据库（可选）

**默认配置（SQLite）**：无需任何配置，直接跳到步骤 4。

**MySQL 配置**：
```sh
# 创建数据库
mysql -u root -p -e "CREATE DATABASE gate_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"

# 配置环境变量（步骤 4）
# DB_TYPE=mysql
# DSN=root:password@tcp(127.0.0.1:3306)/gate_db?charset=utf8mb4&parseTime=True&loc=Local
```

**PostgreSQL 配置**：
```sh
# 创建数据库
createdb gate_db
# 或使用 psql
psql -U postgres -c "CREATE DATABASE gate_db;"

# 配置环境变量（步骤 4）
# DB_TYPE=postgres
# DSN=host=localhost user=postgres password=postgres dbname=gate_db port=5432 sslmode=disable TimeZone=Asia/Shanghai
```

### 4. 配置环境变量

```sh
# 复制配置模板（可选，默认使用 SQLite）
cp .env.example cmd/gate-server/.env

# 如需使用 MySQL 或 PostgreSQL，编辑配置文件
vim cmd/gate-server/.env
```

**快速开始（使用默认 SQLite）**：
无需配置，直接跳到步骤 5！

**使用 MySQL 或 PostgreSQL**：
在 `.env` 文件中配置：
- `DB_TYPE`：数据库类型（`sqlite` / `mysql` / `postgres`）
- `DSN`：数据库连接字符串
- `SERVER_ADDR`：服务器监听地址（如 `:4514`）

**推荐配置项**：
- `JWT_SECRET`：自定义 JWT 密钥（生产环境必需）

示例配置（SQLite）：
```env
DB_TYPE=sqlite
DSN=gate.db
SERVER_ADDR=:4514
JWT_SECRET=your-super-secret-key-here
```

示例配置（PostgreSQL）：
```env
DB_TYPE=postgres
DSN=host=localhost user=postgres password=postgres dbname=gate_db port=5432 sslmode=disable TimeZone=Asia/Shanghai
SERVER_ADDR=:4514
JWT_SECRET=your-super-secret-key-here
```

### 5. 构建二进制文件

```sh
# 构建服务端
go build -o bin/gate-server ./cmd/gate-server

# 构建客户端
go build -o bin/gate ./cmd/gate
```

### 6. 启动服务端

```sh
# 使用默认配置启动（SQLite）
./bin/gate-server start

# 或使用 .env 文件配置启动
cd cmd/gate-server
../../bin/gate-server start

# 方式 2: 返回项目根目录启动
cd ../..
./bin/gate-server start

# 方式 3: 使用命令行参数启动
./bin/gate-server start --addr ":4514" --dsn "host=localhost user=postgres password=postgres dbname=gate_db port=5432 sslmode=disable"
```

服务端成功启动后，会显示：
```
Starting Imperishable Gate server on :4514...
Database connected successfully
Server started successfully
```

## 客户端使用

### 配置客户端

配置服务器地址（三选一）：

```bash
# 方式 1: 环境变量（推荐）
export GATE_SERVER_ADDR=http://localhost:4514

# 方式 2: 配置文件
mkdir -p ~/.config/gate
echo '{"server_addr": "http://localhost:4514"}' > ~/.config/gate/config.json

# 方式 3: 每次使用命令行参数
gate <command> -a http://localhost:4514
```

> **重要**：设置服务器地址时务必加上 `http://` 或 `https://` 前缀，否则默认使用 `https://` 可能导致本地连接失败。

### 用户认证

#### 注册新用户

```sh
./bin/gate register
```

系统会提示输入：
- 用户名（3-32字符）
- 邮箱地址
- 密码（至少6字符）

注册成功后会收到验证邮件（如果配置了邮件服务）。

#### 登录

```sh
./bin/gate login
```

输入用户名和密码即可登录。登录成功后，令牌会自动保存到系统 keyring。

### 基本操作

#### 添加链接

```sh
# 添加单个链接
./bin/gate add -l "https://example.com"

# 添加链接并设置备注
./bin/gate add -l "https://example.com" -r "我的示例网站"

# 添加链接、标签和别名
./bin/gate add -l "https://example.com" -t "tech,blog" -N "mysite"
```

#### 查看链接

```sh
# 列出所有链接
./bin/gate list

# 通过别名查询
./bin/gate list -n "mysite"

# 通过标签查询
./bin/gate list -t "tech"
```

#### 删除链接

```sh
# 通过 URL 删除
./bin/gate delete -l "https://example.com"

# 通过别名删除
./bin/gate delete -n "mysite"
```

#### 打开链接

```sh
# 通过别名在浏览器中打开
./bin/gate open -n "mysite"

# 通过 URL 打开
./bin/gate open -l "https://example.com"
```

### 高级功能

#### 标签管理

```sh
# 为链接添加标签（通过 URL）
./bin/gate add -l "https://example.com" -t "tech,news"

# 为链接添加标签（通过别名）
./bin/gate add -n "mysite" -t "tech,news"
```

#### 监控管理

```sh
# 启用链接监控（通过 URL）
./bin/gate watch -l "https://example.com" -w true

# 启用链接监控（通过别名）
./bin/gate watch -n "mysite" -w true

# 禁用监控
./bin/gate watch -n "mysite" -w false
```

#### 系统检查

```sh
# 测试服务器连接
./bin/gate ping -m "hello"
```

#### 查看当前用户

```sh
./bin/gate whoami
```

#### 登出

```sh
./bin/gate logout
```

## 常见问题

### 1. 客户端无法连接服务器

**问题**：客户端提示连接失败

**解决方案**：
- 确认服务端已启动
- 检查服务器地址是否包含 `http://` 或 `https://` 前缀
- 确认端口号正确（默认 4514）
- 检查防火墙设置

```bash
# 正确的配置示例
export GATE_SERVER_ADDR=http://localhost:4514

# 错误的配置（缺少协议）
export GATE_SERVER_ADDR=localhost:4514  # ❌ 会被解析为 https://
```

### 2. Linux 下 keyring 错误

**问题**：客户端提示 keyring 相关错误

**解决方案**：
```sh
# Ubuntu/Debian
sudo apt-get install gnome-keyring libsecret-1-dev

# Fedora/RHEL
sudo dnf install gnome-keyring libsecret-devel
```

### 3. 数据库连接失败

**问题**：服务端启动时提示数据库连接失败

**解决方案**：
- 确认 PostgreSQL 已启动
- 检查 `.env` 文件中的 DSN 配置
- 确认数据库已创建
- 检查用户名和密码是否正确

### 4. 令牌过期

**问题**：操作时提示令牌过期

**解决方案**：
客户端会自动刷新过期的令牌。如果自动刷新失败，请重新登录：
```sh
./bin/gate login
```

## 下一步

- 查看 [客户端完整命令文档](client-guide.md)
- 查看 [API 文档](api.md)
- 了解 [架构设计](architecture.md)
- 阅读 [配置说明](configuration.md)
