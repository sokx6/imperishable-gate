# Simple Deployment Guide | Deploying Your Service in Gensokyo

**[简体中文](deployment.md) | [English](deployment.en.md)**

> *"It's time to let Hakugyokurou's link management system serve all of Gensokyo!"*

This document provides simple server deployment methods suitable for learning and small-scale projects. If you're building a "massive data center" to serve all of Gensokyo, this guide will help you!

## System Requirements

- **Go**: 1.25.1 or higher
- **Database** (choose one):
  - **SQLite**: 3.x+ (recommended for single-machine deployment)
  - **MySQL**: 5.7+ / 8.0+
  - **PostgreSQL**: 12.0+
- **Operating System**: Linux / macOS (Linux server recommended)

## Stage 1-2 | Basic Deployment Steps

### 1. Server Preparation

```bash
# Update system (Ubuntu/Debian)
sudo apt update && sudo apt upgrade -y

# Install necessary tools
sudo apt install git -y

# Install database (optional, uses SQLite by default)
# PostgreSQL: sudo apt install postgresql -y
# MySQL: sudo apt install mysql-server -y
```

### 2. Create Database (when using PostgreSQL or MySQL)

**PostgreSQL**:
```bash
# Switch to postgres user
sudo -u postgres psql

# Execute in psql
CREATE DATABASE gate_db;
CREATE USER gateuser WITH PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE gate_db TO gateuser;
\q
```

### 3. Deploy Application

```bash
# Clone project
git clone https://github.com/sokx6/imperishable-gate.git
cd imperishable-gate

# Install dependencies
go mod download

# Configure environment variables
cp .env.example cmd/gate-server/.env
vim cmd/gate-server/.env
```

Edit the `.env` file:
```env
DSN=host=localhost user=gateuser password=your_password dbname=gate_db port=5432 sslmode=disable TimeZone=Asia/Shanghai
SERVER_ADDR=:4514
JWT_SECRET=your-random-secret-key
```

### 4. Build and Run

```bash
# Build
go build -o gate-server ./cmd/gate-server

# Run
./gate-server start
```

### 5. Run in Background (Optional)

Using `nohup` or `screen`:

```bash
# Using nohup
nohup ./gate-server start > server.log 2>&1 &

# Or using screen
screen -S gate-server
./gate-server start
# Press Ctrl+A then D to detach session
```

## Managing with systemd (Recommended)

Create service file:

```bash
sudo vim /etc/systemd/system/gate-server.service
```

Add the following content:

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

Start the service:

```bash
# Reload systemd
sudo systemctl daemon-reload

# Start service
sudo systemctl start gate-server

# Enable auto-start on boot
sudo systemctl enable gate-server

# Check status
sudo systemctl status gate-server

# View logs
sudo journalctl -u gate-server -f
```

## Simple Reverse Proxy (Optional)

If you need to access via domain name, you can use Nginx:

```bash
# Install Nginx
sudo apt install nginx -y

# Create configuration
sudo vim /etc/nginx/sites-available/gate
```

Add configuration:

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

Enable configuration:

```bash
sudo ln -s /etc/nginx/sites-available/gate /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl reload nginx
```

## Firewall Configuration

```bash
# Ubuntu (UFW)
sudo ufw allow 22    # SSH
sudo ufw allow 80    # HTTP
sudo ufw allow 4514  # Application port (if accessing directly)
sudo ufw enable
```

## Data Backup

Simple backup script:

```bash
#!/bin/bash
# backup.sh

DATE=$(date +%Y%m%d_%H%M%S)
BACKUP_DIR="$HOME/backups"
mkdir -p $BACKUP_DIR

# Backup database
pg_dump -U gateuser gate_db > $BACKUP_DIR/gate_db_$DATE.sql

# Keep backups for the last 7 days
find $BACKUP_DIR -name "gate_db_*.sql" -mtime +7 -delete

echo "Backup completed: gate_db_$DATE.sql"
```

Set up scheduled backup:

```bash
# Edit crontab
crontab -e

# Backup daily at 2 AM
0 2 * * * /path/to/backup.sh
```

## Updating Application

```bash
# Navigate to project directory
cd imperishable-gate

# Pull latest code
git pull

# Rebuild
go build -o gate-server ./cmd/gate-server

# Restart service
sudo systemctl restart gate-server
```

## Common Issues

### Port Already in Use

```bash
# Find process using the port
sudo lsof -i :4514
# or
sudo netstat -tulpn | grep 4514
```

### View Application Logs

```bash
# systemd service logs
sudo journalctl -u gate-server -n 50

# nohup logs
tail -f server.log
```

### Database Connection Failed

```bash
# Check PostgreSQL status
sudo systemctl status postgresql

# Test connection
psql -h localhost -U gateuser -d gate_db
```

## Security Recommendations

1. **Change Default Passwords**: Use strong passwords
2. **Set JWT_SECRET**: Use a random string
   ```bash
   openssl rand -base64 32
   ```
3. **Restrict Firewall**: Only open necessary ports
4. **Regular Backups**: Set up automatic backup tasks
5. **Regular Updates**: Keep system and dependencies updated

## Monitoring Recommendations

### Simple Health Check Script

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

Set up periodic checks:

```bash
# Check every 5 minutes
*/5 * * * * /path/to/health_check.sh
```

---

This deployment guide provides basic deployment methods suitable for learning and development environments.
