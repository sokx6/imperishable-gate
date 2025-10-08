# Environment Requirements & Configuration | Stage 1-6 Configuration Guide

**[📖 简体中文](configuration.md) | [📘 English](configuration.en.md)**

> ⚙️ *"The operational experience from the previous life can finally be put to good use in this one!"*

This document introduces the complete configuration methods for **Imperishable Gate** from Stage 1 to Stage 6.

## Environment Requirements

### Basic Requirements
- **Go**: 1.25.1 or higher
- **Database** (choose one, implemented in Stage 2):
  - **SQLite**: 3.x+ (default, no additional installation required) ✨ Recommended for beginners
  - **MySQL**: 5.7+ / 8.0+
  - **PostgreSQL**: 12.0+
- **Operating System**: Linux / macOS / Windows

### Stage 6 | Linux System Keyring Requirements

The client uses the system keyring to securely store tokens (implemented in Stage 6 "Remains of the Underworld Princess"). On Linux systems, you need to install the corresponding keyring service:

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

> **Note**:
> - If you're using a desktop environment (such as GNOME, KDE), keyring service is usually pre-installed
> - For headless server environments, you may need to manually start the keyring daemon
> - macOS and Windows systems require no additional installation and will automatically use the system's Keychain and Credential Manager

## Configuration Instructions

### Environment Variable Configuration

The project uses `.env` files for configuration management. Please follow these steps:

1. **Copy the environment variable template file**
   ```sh
   cp .env.example cmd/gate-server/.env
   ```

2. **Edit the configuration file**
   ```sh
   vim cmd/gate-server/.env
   ```

### Configuration Options

#### 📊 Database Configuration

| Environment Variable | Description | Example Value | Required |
|---------------------|-------------|---------------|----------|
| `DB_TYPE` | Database type: `sqlite` (default), `mysql`, `postgres` | `sqlite` | ❌ (defaults to SQLite) |
| `DSN` | Database connection string (varies by database type) | See examples below | ✅ |

**SQLite Configuration (default, recommended for development/small projects)**:
```bash
DB_TYPE=sqlite
DSN=gate.db
# Or use absolute path
# DSN=/var/lib/gate/gate.db
```

**MySQL Configuration**:
```bash
DB_TYPE=mysql
DSN=user:password@tcp(127.0.0.1:3306)/gate_db?charset=utf8mb4&parseTime=True&loc=Local
```

**PostgreSQL Configuration**:
```bash
DB_TYPE=postgres
DSN=host=localhost user=postgres password=postgres dbname=gate_db port=5432 sslmode=disable TimeZone=Asia/Shanghai
```

#### 🌐 Server Configuration

| Environment Variable | Description | Example Value | Required |
|---------------------|-------------|---------------|----------|
| `SERVER_ADDR` | Server listening address | `localhost:4514` or `:4514` | ✅ |

#### 🔐 JWT Security Configuration

| Environment Variable | Description | Example Value | Required |
|---------------------|-------------|---------------|----------|
| `JWT_SECRET` | JWT signing key (must be changed in production!) | Generate with `openssl rand -base64 64` | ⚠️ Recommended |

> **Security Tips**:
> - Always set a strong random `JWT_SECRET` in production
> - Generate a secure key with: `openssl rand -base64 64`
> - Never commit `.env` files containing real credentials to version control

#### 📧 Email Service Configuration (Optional)

For email verification and link monitoring change notifications:

| Environment Variable | Description | Example Value | Required |
|---------------------|-------------|---------------|----------|
| `EMAIL_HOST` | SMTP server address | `smtp.gmail.com` | 📧 |
| `EMAIL_PORT` | SMTP server port | `587` (TLS) or `465` (SSL) | 📧 |
| `EMAIL_FROM` | Sender email address | `noreply@example.com` | 📧 |
| `EMAIL_PASSWORD` | Email password or app password | `your-app-password` | 📧 |

> **Email Configuration Tips**:
> - Gmail: Use app-specific password ([How to get one](https://support.google.com/accounts/answer/185833))
> - 163/QQ Mail: Use authorization code instead of login password
> - Port selection: 587 (STARTTLS) or 465 (SSL/TLS)

### Configuration File Locations

```
.env.example                      # Server config template (root directory)
cmd/gate-server/.env             # Server config file (create manually)
cmd/gate/.env.example            # Client config template
cmd/gate/.env                    # Client config file (optional)
.env                             # Common config file (client will also read, optional)
```

**Recommended Configuration Method**:
- **Server**: Create `.env` file in `cmd/gate-server/` directory
- **Client**:
  - Method 1: Create `.env` file in `cmd/gate/` directory
  - Method 2: Create `.env` file in project root directory
  - Method 3: Use environment variable directly `export GATE_SERVER_ADDR=...`

### Configuration Priority

Configuration priority from highest to lowest:

1. **Command-line arguments** (highest priority)
2. **Environment variables** (`.env` file)
3. **Default values** (lowest priority)

**Examples**:
```sh
# Method 1: Use .env file configuration
./bin/gate-server start

# Method 2: Use command-line arguments to override configuration
./bin/gate-server start --addr ":9090" --dsn "host=localhost..."

# Method 3: Use environment variables
export DSN="host=localhost..."
./bin/gate-server start
```

## Client Configuration

### Server Address Configuration

The client has three ways to configure the server address (in priority order):

1. **Command-line arguments** (highest priority)
   ```bash
   gate <command> -a http://localhost:4514
   # or
   gate <command> -a https://api.example.com
   ```

2. **Environment variables** `GATE_SERVER_ADDR` or `SERVER_ADDR`
   ```bash
   export GATE_SERVER_ADDR=http://localhost:4514
   gate <command>
   ```

3. **.env file** (in client directory)
   ```bash
   # Copy config template
   cp cmd/gate/.env.example cmd/gate/.env
   # Or create .env in project root
   ```
   
   Configuration content:
   ```bash
   GATE_SERVER_ADDR=http://localhost:4514
   ```

4. **Default value** (lowest priority)
   - If none of the above are set, uses `http://localhost:4514`

> **Important Note**:
> - When setting the server address, **always include the protocol prefix** (`http://` or `https://`)
> - If no protocol is specified, `https://` will be used by default, which may cause connection failures during local development
> - Priority: `GATE_SERVER_ADDR` > `SERVER_ADDR`
> - Examples:
>   - ✅ Correct: `http://localhost:4514`
>   - ✅ Correct: `https://api.example.com`
>   - ❌ Wrong: `localhost:4514` (will be parsed as `https://localhost:4514`)

### Configuration File Examples

**Server Configuration** (`cmd/gate-server/.env`):
```bash
DB_TYPE=sqlite
DSN=gate.db
SERVER_ADDR=:4514
JWT_SECRET=your-secret-key
```

**Client Configuration** (`cmd/gate/.env` or project root):
```bash
GATE_SERVER_ADDR=http://localhost:4514
```

Create configuration file:
```bash
mkdir -p ~/.config/gate
cat > ~/.config/gate/config.json << EOF
{
  "server_addr": "http://localhost:4514"
}
EOF
```

## Database Configuration

### Create Database

```sh
# Method 1: Use createdb command
createdb gate_db

# Method 2: Use psql
psql -U postgres -c "CREATE DATABASE gate_db;"

# Method 3: Use SQL client
# After connecting to PostgreSQL, execute
CREATE DATABASE gate_db;
```

### Database Connection String Format

```
host=<host_address> user=<username> password=<password> dbname=<database_name> port=<port> sslmode=<ssl_mode> TimeZone=<timezone>
```

**Examples**:
```
# Local development
host=localhost user=postgres password=postgres dbname=gate_db port=5432 sslmode=disable TimeZone=Asia/Shanghai

# Production (with SSL enabled)
host=db.example.com user=gateuser password=securepass dbname=gate_db port=5432 sslmode=require TimeZone=UTC
```

### SSL Mode Description

| Mode | Description | Use Case |
|------|-------------|----------|
| `disable` | Do not use SSL | Local development |
| `require` | SSL required | Production |
| `verify-ca` | Verify CA certificate | High security requirements |
| `verify-full` | Full verification | Highest security requirements |

## Security Recommendations

### Production Environment Checklist

- [ ] Change the default `JWT_SECRET` to a strong random value
- [ ] Use environment variables or secure key management services to store sensitive information
- [ ] Enable database SSL connection (`sslmode=require`)
- [ ] Set appropriate database user permissions
- [ ] Regularly update dependencies
- [ ] Configure firewall rules
- [ ] Enable HTTPS (using reverse proxy like Nginx)
- [ ] Regularly backup database

### Development Environment Recommendations

- Use `.env` files to manage local configuration
- Do not commit `.env` files to version control
- Use `.env.example` as configuration template
- SSL can be disabled for local development (`sslmode=disable`)
