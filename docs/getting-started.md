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

## 安装方式

### 方式一：使用预编译版本（推荐）⭐

如果您不需要修改源码，推荐直接下载预编译的可执行文件，无需安装 Go 环境。

#### 1. 下载可执行文件

访问 [GitHub Releases](https://github.com/locxl/imperishable-gate/releases) 页面，根据您的操作系统下载对应的文件：

**客户端 (gate)**：
- **Linux AMD64**: `gate-linux-amd64`
- **Linux ARM64**: `gate-linux-arm64` (适用于树莓派等 ARM 设备)
- **Windows AMD64**: `gate-windows-amd64.exe`
- **macOS Intel**: `gate-darwin-amd64`
- **macOS Apple Silicon**: `gate-darwin-arm64` (M1/M2/M3 芯片)

**服务端 (gate-server)**：
- **Linux AMD64**: `gate-server-linux-amd64`
- **Linux ARM64**: `gate-server-linux-arm64`
- **Windows AMD64**: `gate-server-windows-amd64.exe`
- **macOS Intel**: `gate-server-darwin-amd64`
- **macOS Apple Silicon**: `gate-server-darwin-arm64`

#### 2. 配置到系统环境变量

为了在任何目录都能直接使用命令，需要将可执行文件配置到系统 PATH 中。

##### Linux / macOS

```bash
# 1. 创建存放目录
mkdir -p ~/.local/bin

# 2. 移动下载的文件到该目录（以 Linux AMD64 为例）
mv ~/Downloads/gate-linux-amd64 ~/.local/bin/gate
mv ~/Downloads/gate-server-linux-amd64 ~/.local/bin/gate-server

# 3. 添加执行权限
chmod +x ~/.local/bin/gate
chmod +x ~/.local/bin/gate-server

# 4. 添加到 PATH（根据您使用的 Shell 选择）
# Bash 用户：
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc

# Zsh 用户：
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc

# Fish 用户：
fish_add_path ~/.local/bin

# 5. 验证安装
gate --version
gate-server --version
```

##### Windows

**方法 1：使用用户环境变量（推荐）**

```powershell
# 1. 创建存放目录
mkdir "$env:USERPROFILE\bin"

# 2. 移动下载的文件到该目录
move "$env:USERPROFILE\Downloads\gate-windows-amd64.exe" "$env:USERPROFILE\bin\gate.exe"
move "$env:USERPROFILE\Downloads\gate-server-windows-amd64.exe" "$env:USERPROFILE\bin\gate-server.exe"

# 3. 添加到 PATH（PowerShell）
$oldPath = [Environment]::GetEnvironmentVariable("Path", "User")
$newPath = "$oldPath;$env:USERPROFILE\bin"
[Environment]::SetEnvironmentVariable("Path", $newPath, "User")

# 4. 重启 PowerShell/CMD，然后验证
gate --version
gate-server --version
```

**方法 2：使用图形界面**

1. 创建文件夹 `C:\Program Files\Gate`（或任意位置）
2. 将下载的 `.exe` 文件重命名并移动到该文件夹：
   - `gate-windows-amd64.exe` → `gate.exe`
   - `gate-server-windows-amd64.exe` → `gate-server.exe`
3. 右键点击 "此电脑" → "属性" → "高级系统设置" → "环境变量"
4. 在 "用户变量" 中找到 `Path`，点击 "编辑"
5. 点击 "新建"，添加路径 `C:\Program Files\Gate`
6. 点击 "确定" 保存
7. 重启命令提示符或 PowerShell，验证：`gate --version`

#### 3. 开始使用

现在您可以在任何目录直接使用命令了：

```bash
# 启动服务端
gate-server start

# 使用客户端
gate register
gate login
gate add https://example.com
```

跳过"方式二"，直接前往 [配置数据库](#3-配置数据库可选) 和 [配置环境变量](#4-配置环境变量) 完成服务端配置。

---

### 方式二：从源码编译

如果您需要修改源码或进行开发，可以从源码编译。

##### 1. 获取项目

```sh
git clone https://github.com/locxl/imperishable-gate.git
cd imperishable-gate
```

#### 2. 安装依赖

```sh
go mod download
```

#### 3. 配置数据库（可选）

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

#### 4. 配置环境变量

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

#### 5. 构建二进制文件

```sh
# 构建服务端
go build -o bin/gate-server ./cmd/gate-server

# 构建客户端
go build -o bin/gate ./cmd/gate
```

#### 6. 启动服务端

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
