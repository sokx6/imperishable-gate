# Getting Started | Stage 1: White & Pink Spring

**[📖 简体中文](getting-started.md) | [📘 English](getting-started.en.md)**

> 🌸 *"Let's start with some preparations!"* 🌸

This guide will help you quickly set up and run **Imperishable Gate**. Starting from scratch, complete the foundational framework of **Stage 1 "White & Pink Spring"**, and officially begin managing links as a gardener of Hakugyokurou!

## Prerequisites

Before you begin, ensure your system meets the following requirements:

- Go 1.25.1 or higher
- Database (choose one):
  - **SQLite** (default, no additional installation required) ✨ Recommended for quick start
  - **MySQL** 5.7+ / 8.0+
  - **PostgreSQL** 12.0+
- Git

For detailed environment requirements, please refer to the [Configuration Documentation](configuration.en.md).

## Installation Steps

### 1. Clone the Project

```sh
git clone https://github.com/sokx6/imperishable-gate.git
cd imperishable-gate
```

### 2. Install Dependencies

```sh
go mod download
```

### 3. Configure Database (Optional)

**Default Configuration (SQLite)**: No configuration needed, skip to step 4.

**MySQL Configuration**:
```sh
# Create database
mysql -u root -p -e "CREATE DATABASE gate_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"

# Configure environment variables (step 4)
# DB_TYPE=mysql
# DSN=root:password@tcp(127.0.0.1:3306)/gate_db?charset=utf8mb4&parseTime=True&loc=Local
```

**PostgreSQL Configuration**:
```sh
# Create database
createdb gate_db
# Or use psql
psql -U postgres -c "CREATE DATABASE gate_db;"

# Configure environment variables (step 4)
# DB_TYPE=postgres
# DSN=host=localhost user=postgres password=postgres dbname=gate_db port=5432 sslmode=disable TimeZone=Asia/Shanghai
```

### 4. Configure Environment Variables

```sh
# Copy configuration template (optional, uses SQLite by default)
cp .env.example cmd/gate-server/.env

# Edit configuration file if using MySQL or PostgreSQL
vim cmd/gate-server/.env
```

**Quick Start (Using Default SQLite)**:
No configuration needed, skip to step 5!

**Using MySQL or PostgreSQL**:
Configure in the `.env` file:
- `DB_TYPE`: Database type (`sqlite` / `mysql` / `postgres`)
- `DSN`: Database connection string
- `SERVER_ADDR`: Server listening address (e.g., `:4514`)

**Recommended Configuration**:
- `JWT_SECRET`: Custom JWT secret key (required for production)

Example Configuration (SQLite):
```env
DB_TYPE=sqlite
DSN=gate.db
SERVER_ADDR=:4514
JWT_SECRET=your-super-secret-key-here
```

Example Configuration (PostgreSQL):
```env
DB_TYPE=postgres
DSN=host=localhost user=postgres password=postgres dbname=gate_db port=5432 sslmode=disable TimeZone=Asia/Shanghai
SERVER_ADDR=:4514
JWT_SECRET=your-super-secret-key-here
```

### 5. Build Binaries

```sh
# Build server
go build -o bin/gate-server ./cmd/gate-server

# Build client
go build -o bin/gate ./cmd/gate
```

### 6. Start the Server

```sh
# Start with default configuration (SQLite)
./bin/gate-server start

# Or start with .env file configuration
cd cmd/gate-server
../../bin/gate-server start

# Method 2: Go back to project root and start
cd ../..
./bin/gate-server start

# Method 3: Start with command-line arguments
./bin/gate-server start --addr ":4514" --dsn "host=localhost user=postgres password=postgres dbname=gate_db port=5432 sslmode=disable"
```

After the server starts successfully, you will see:
```
Starting Imperishable Gate server on :4514...
Database connected successfully
Server started successfully
```

## Client Usage

### Configure Client

Configure server address (choose one):

```bash
# Method 1: Environment variable (recommended)
export GATE_SERVER_ADDR=http://localhost:4514

# Method 2: Configuration file
mkdir -p ~/.config/gate
echo '{"server_addr": "http://localhost:4514"}' > ~/.config/gate/config.json

# Method 3: Use command-line parameter each time
gate <command> -a http://localhost:4514
```

> **Important**: When setting the server address, make sure to include the `http://` or `https://` prefix, otherwise the default `https://` may cause local connection failures.

### User Authentication

#### Register New User

```sh
./bin/gate register
```

The system will prompt you to enter:
- Username (3-32 characters)
- Email address
- Password (at least 6 characters)

After successful registration, you will receive a verification email (if email service is configured).

#### Login

```sh
./bin/gate login
```

Enter your username and password to log in. After successful login, the token will be automatically saved to the system keyring.

### Basic Operations

#### Add Link

```sh
# Add a single link
./bin/gate add -l "https://example.com"

# Add link with remark
./bin/gate add -l "https://example.com" -r "My example website"

# Add link with tags and alias
./bin/gate add -l "https://example.com" -t "tech,blog" -N "mysite"
```

#### View Links

```sh
# List all links
./bin/gate list

# Query by alias
./bin/gate list -n "mysite"

# Query by tag
./bin/gate list -t "tech"
```

#### Delete Link

```sh
# Delete by URL
./bin/gate delete -l "https://example.com"

# Delete by alias
./bin/gate delete -n "mysite"
```

#### Open Link

```sh
# Open by alias in browser
./bin/gate open -n "mysite"

# Open by URL
./bin/gate open -l "https://example.com"
```

### Advanced Features

#### Tag Management

```sh
# Add tags to link (by URL)
./bin/gate add -l "https://example.com" -t "tech,news"

# Add tags to link (by alias)
./bin/gate add -n "mysite" -t "tech,news"
```

#### Watch Management

```sh
# Enable link monitoring (by URL)
./bin/gate watch -l "https://example.com" -w true

# Enable link monitoring (by alias)
./bin/gate watch -n "mysite" -w true

# Disable monitoring
./bin/gate watch -n "mysite" -w false
```

#### System Check

```sh
# Test server connection
./bin/gate ping -m "hello"
```

#### View Current User

```sh
./bin/gate whoami
```

#### Logout

```sh
./bin/gate logout
```

## Common Issues

### 1. Client Cannot Connect to Server

**Issue**: Client shows connection failure

**Solution**:
- Confirm the server is running
- Check if the server address includes `http://` or `https://` prefix
- Confirm the port number is correct (default 4514)
- Check firewall settings

```bash
# Correct configuration example
export GATE_SERVER_ADDR=http://localhost:4514

# Wrong configuration (missing protocol)
export GATE_SERVER_ADDR=localhost:4514  # ❌ Will be parsed as https://
```

### 2. Keyring Error on Linux

**Issue**: Client shows keyring-related errors

**Solution**:
```sh
# Ubuntu/Debian
sudo apt-get install gnome-keyring libsecret-1-dev

# Fedora/RHEL
sudo dnf install gnome-keyring libsecret-devel
```

### 3. Database Connection Failure

**Issue**: Server shows database connection failure on startup

**Solution**:
- Confirm PostgreSQL is running
- Check the DSN configuration in `.env` file
- Confirm the database has been created
- Check if username and password are correct

### 4. Token Expired

**Issue**: Token expired error during operations

**Solution**:
The client will automatically refresh expired tokens. If automatic refresh fails, please log in again:
```sh
./bin/gate login
```

## Next Steps

- Check [Complete Client Command Documentation](client-guide.en.md)
- View [API Documentation](api.en.md)
- Learn about [Architecture Design](architecture.en.md)
- Read [Configuration Guide](configuration.en.md)
