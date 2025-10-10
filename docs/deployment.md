# 简易部署指南 | 在幻想乡部署你的服务

**[简体中文](deployment.md) | [English](deployment.en.md)**

> *"是时候让白玉楼的链接管理系统服务整个幻想乡了！"*

本文档提供简单的服务器部署方法，适合学习和小型项目使用。如果你要建个"大型数据中心"来服务全幻想乡，这份指南会帮到你！

## 系统要求

- **Go**: 1.25.1 或更高版本
- **数据库**（三选一）:
  - **SQLite**: 3.x+（推荐用于单机部署）
  - **MySQL**: 5.7+ / 8.0+
  - **PostgreSQL**: 12.0+
- **操作系统**: Linux / macOS（推荐 Linux 服务器）

## Stage 1-2 | 基础部署步骤

### 1. 服务器准备

```bash
# 更新系统（Ubuntu/Debian）
sudo apt update && sudo apt upgrade -y

# 安装必要工具
sudo apt install git -y

# 安装数据库（可选，默认使用 SQLite）
# PostgreSQL: sudo apt install postgresql -y
# MySQL: sudo apt install mysql-server -y
```

### 2. 创建数据库（使用 PostgreSQL 或 MySQL 时）

**PostgreSQL**:
```bash
# 切换到 postgres 用户
sudo -u postgres psql

# 在 psql 中执行
CREATE DATABASE gate_db;
CREATE USER gateuser WITH PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE gate_db TO gateuser;
\q
```

### 3. 部署应用

```bash
# 克隆项目
git clone https://github.com/sokx6/imperishable-gate.git
cd imperishable-gate

# 安装依赖
go mod download

# 配置环境变量
cp .env.example cmd/gate-server/.env
vim cmd/gate-server/.env
```

编辑 `.env` 文件：
```env
DSN=host=localhost user=gateuser password=your_password dbname=gate_db port=5432 sslmode=disable TimeZone=Asia/Shanghai
SERVER_ADDR=:4514
JWT_SECRET=your-random-secret-key
```

### 4. 构建和运行

```bash
# 构建
go build -o gate-server ./cmd/gate-server

# 运行
./gate-server start
```

### 5. 后台运行（可选）

使用 `nohup` 或 `screen`：

```bash
# 使用 nohup
nohup ./gate-server start > server.log 2>&1 &

# 或使用 screen
screen -S gate-server
./gate-server start
# 按 Ctrl+A 然后按 D 来分离会话
```

## 使用 systemd 管理（推荐）

创建服务文件：

```bash
sudo vim /etc/systemd/system/gate-server.service
```

添加以下内容：

```ini
[Unit]
Description=Imperishable Gate Server
After=network.target postgresql.service

[Service]
Type=simple
User=your-username
WorkingDirectory=/path/to/imperishable-gate/cmd/gate-server
EnvironmentFile=/path/to/imperishable-gate/cmd/gate-server/.env
ExecStart=/path/to/imperishable-gate/gate-server start
Restart=on-failure

[Install]
WantedBy=multi-user.target
```

启动服务：

```bash
# 重新加载 systemd
sudo systemctl daemon-reload

# 启动服务
sudo systemctl start gate-server

# 设置开机自启
sudo systemctl enable gate-server

# 查看状态
sudo systemctl status gate-server

# 查看日志
sudo journalctl -u gate-server -f
```

## 简单的反向代理（可选）

如果需要通过域名访问，可以使用 Nginx：

```bash
# 安装 Nginx
sudo apt install nginx -y

# 创建配置
sudo vim /etc/nginx/sites-available/gate
```

添加配置：

```nginx
server {
    listen 80;
    server_name your-domain.com;

    location / {
        proxy_pass http://localhost:4514;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

启用配置：

```bash
sudo ln -s /etc/nginx/sites-available/gate /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl reload nginx
```

## 防火墙配置

```bash
# Ubuntu (UFW)
sudo ufw allow 22    # SSH
sudo ufw allow 80    # HTTP
sudo ufw allow 4514  # 应用端口（如果直接访问）
sudo ufw enable
```

## 数据备份

简单的备份脚本：

```bash
#!/bin/bash
# backup.sh

DATE=$(date +%Y%m%d_%H%M%S)
BACKUP_DIR="$HOME/backups"
mkdir -p $BACKUP_DIR

# 备份数据库
pg_dump -U gateuser gate_db > $BACKUP_DIR/gate_db_$DATE.sql

# 保留最近7天的备份
find $BACKUP_DIR -name "gate_db_*.sql" -mtime +7 -delete

echo "Backup completed: gate_db_$DATE.sql"
```

设置定时备份：

```bash
# 编辑 crontab
crontab -e

# 每天凌晨2点备份
0 2 * * * /path/to/backup.sh
```

## 更新应用

```bash
# 进入项目目录
cd imperishable-gate

# 拉取最新代码
git pull

# 重新构建
go build -o gate-server ./cmd/gate-server

# 重启服务
sudo systemctl restart gate-server
```

## 常见问题

### 端口被占用

```bash
# 查找占用端口的进程
sudo lsof -i :4514
# 或
sudo netstat -tulpn | grep 4514
```

### 查看应用日志

```bash
# systemd 服务日志
sudo journalctl -u gate-server -n 50

# nohup 日志
tail -f server.log
```

### 数据库连接失败

```bash
# 检查 PostgreSQL 状态
sudo systemctl status postgresql

# 测试连接
psql -h localhost -U gateuser -d gate_db
```

## 安全建议

1. **修改默认密码**：使用强密码
2. **设置 JWT_SECRET**：使用随机字符串
   ```bash
   openssl rand -base64 32
   ```
3. **限制防火墙**：只开放必要端口
4. **定期备份**：设置自动备份任务
5. **定期更新**：保持系统和依赖更新

## 监控建议

### 简单的健康检查脚本

```bash
#!/bin/bash
# health_check.sh

URL="http://localhost:4514/api/v1/ping"
RESPONSE=$(curl -s -o /dev/null -w "%{http_code}" $URL)

if [ $RESPONSE -eq 200 ]; then
    echo "Service is healthy"
else
    echo "Service is down, restarting..."
    sudo systemctl restart gate-server
fi
```

设置定期检查：

```bash
# 每5分钟检查一次
*/5 * * * * /path/to/health_check.sh
```

---

这个部署指南提供了基础的部署方法，适合学习和开发环境使用。
