# 环境要求与配置

## 环境要求

### 基础要求
- **Go**: 1.25.1 或更高版本
- **数据库**（三选一）:
  - **SQLite**: 3.x+（默认，无需额外安装）
  - **MySQL**: 5.7+ / 8.0+
  - **PostgreSQL**: 12.0+
- **操作系统**: Linux / macOS / Windows

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

## 配置说明

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
| `DB_TYPE` | 数据库类型：`sqlite`（默认）、`mysql`、`postgres` | `sqlite` | ❌（默认 SQLite） |
| `DSN` | 数据库连接字符串（根据数据库类型而定） | 见下方示例 | ✅ |

**SQLite 配置（默认，推荐用于开发/小型项目）**：
```bash
DB_TYPE=sqlite
DSN=gate.db
# 或使用绝对路径
# DSN=/var/lib/gate/gate.db
```

**MySQL 配置**：
```bash
DB_TYPE=mysql
DSN=user:password@tcp(127.0.0.1:3306)/gate_db?charset=utf8mb4&parseTime=True&loc=Local
```

**PostgreSQL 配置**：
```bash
DB_TYPE=postgres
DSN=host=localhost user=postgres password=postgres dbname=gate_db port=5432 sslmode=disable TimeZone=Asia/Shanghai
```

#### 🌐 服务器配置

| 环境变量 | 说明 | 示例值 | 必需 |
|---------|------|--------|------|
| `SERVER_ADDR` | 服务器监听地址 | `localhost:4514` 或 `:4514` | ✅ |

#### 🔐 JWT 安全配置

| 环境变量 | 说明 | 示例值 | 必需 |
|---------|------|--------|------|
| `JWT_SECRET` | JWT 签名密钥（生产环境务必修改！） | 使用 `openssl rand -base64 64` 生成 | ⚠️ 推荐 |

> **安全提示**：
> - 生产环境务必设置强随机 `JWT_SECRET`
> - 使用命令生成安全密钥：`openssl rand -base64 64`
> - 切勿将包含真实密钥的 `.env` 文件提交到版本控制系统

#### 📧 邮件服务配置（可选）

用于邮箱验证和链接监控变化通知功能：

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
.env.example                      # 服务端配置模板（根目录）
cmd/gate-server/.env             # 服务端配置文件（需手动创建）
cmd/gate/.env.example            # 客户端配置模板
cmd/gate/.env                    # 客户端配置文件（可选）
.env                             # 通用配置文件（客户端也会读取，可选）
```

**推荐配置方式**：
- **服务端**：在 `cmd/gate-server/` 目录下创建 `.env` 文件
- **客户端**：
  - 方式 1：在 `cmd/gate/` 目录下创建 `.env` 文件
  - 方式 2：在项目根目录创建 `.env` 文件
  - 方式 3：直接使用环境变量 `export GATE_SERVER_ADDR=...`

### 配置优先级

配置的优先级从高到低为：

1. **命令行参数**（最高优先级）
2. **环境变量**（`.env` 文件）
3. **默认值**（最低优先级）

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

## 客户端配置

### 服务器地址配置

客户端有三种方式配置服务器地址（按优先级排序）：

1. **命令行参数** （最高优先级）
   ```bash
   gate <command> -a http://localhost:4514
   # 或
   gate <command> -a https://api.example.com
   ```

2. **环境变量** `GATE_SERVER_ADDR` 或 `SERVER_ADDR`
   ```bash
   export GATE_SERVER_ADDR=http://localhost:4514
   gate <command>
   ```

3. **.env 文件**（客户端目录下）
   ```bash
   # 复制配置模板
   cp cmd/gate/.env.example cmd/gate/.env
   # 或在项目根目录创建 .env
   ```
   
   配置内容：
   ```bash
   GATE_SERVER_ADDR=http://localhost:4514
   ```

4. **默认值**（最低优先级）
   - 如果以上都未设置，使用 `http://localhost:4514`

> **重要提示**：
> - 设置服务器地址时，**请务必包含协议前缀**（`http://` 或 `https://`）
> - 如果不指定协议，默认会使用 `https://`，这可能导致本地开发时连接失败
> - 优先级：`GATE_SERVER_ADDR` > `SERVER_ADDR`
> - 示例：
>   - ✅ 正确：`http://localhost:4514`
>   - ✅ 正确：`https://api.example.com`
>   - ❌ 错误：`localhost:4514`（会被解析为 `https://localhost:4514`）

### 配置文件示例

**服务端配置** (`cmd/gate-server/.env`)：
```bash
DB_TYPE=sqlite
DSN=gate.db
SERVER_ADDR=:4514
JWT_SECRET=your-secret-key
```

**客户端配置** (`cmd/gate/.env` 或项目根目录)：
```bash
GATE_SERVER_ADDR=http://localhost:4514
```

创建配置文件：
```bash
mkdir -p ~/.config/gate
cat > ~/.config/gate/config.json << EOF
{
  "server_addr": "http://localhost:4514"
}
EOF
```

## 数据库配置

### 创建数据库

```sh
# 方式 1: 使用 createdb 命令
createdb gate_db

# 方式 2: 使用 psql
psql -U postgres -c "CREATE DATABASE gate_db;"

# 方式 3: 使用 SQL 客户端
# 连接到 PostgreSQL 后执行
CREATE DATABASE gate_db;
```

### 数据库连接字符串格式

```
host=<主机地址> user=<用户名> password=<密码> dbname=<数据库名> port=<端口> sslmode=<SSL模式> TimeZone=<时区>
```

**示例**：
```
# 本地开发
host=localhost user=postgres password=postgres dbname=gate_db port=5432 sslmode=disable TimeZone=Asia/Shanghai

# 生产环境（启用 SSL）
host=db.example.com user=gateuser password=securepass dbname=gate_db port=5432 sslmode=require TimeZone=UTC
```

### SSL 模式说明

| 模式 | 说明 | 适用场景 |
|------|------|----------|
| `disable` | 不使用 SSL | 本地开发 |
| `require` | 必须使用 SSL | 生产环境 |
| `verify-ca` | 验证 CA 证书 | 高安全要求 |
| `verify-full` | 完整验证 | 最高安全要求 |

## 安全建议

### 生产环境检查清单

- [ ] 修改默认的 `JWT_SECRET` 为强随机值
- [ ] 使用环境变量或安全的密钥管理服务存储敏感信息
- [ ] 启用数据库 SSL 连接（`sslmode=require`）
- [ ] 设置合适的数据库用户权限
- [ ] 定期更新依赖包
- [ ] 配置防火墙规则
- [ ] 启用 HTTPS（使用反向代理如 Nginx）
- [ ] 定期备份数据库

### 开发环境建议

- 使用 `.env` 文件管理本地配置
- 不要将 `.env` 文件提交到版本控制
- 使用 `.env.example` 作为配置模板
- 本地开发可以禁用 SSL（`sslmode=disable`）
