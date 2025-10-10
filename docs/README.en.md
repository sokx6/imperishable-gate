# Documentation Index

**[简体中文](README.md) | [English](README.en.md)**

Welcome to the Imperishable Gate documentation!

## Quick Navigation

### Getting Started

1. **[Getting Started](getting-started.en.md)** - Installation and basic usage
2. **[Configuration](configuration.en.md)** - Environment variables and settings
3. **[Client Guide](client-guide.en.md)** - CLI command reference

### User Documentation

- **[API Documentation](api.en.md)** - RESTful API interface reference
- **[Architecture Design](architecture.en.md)** - System architecture and tech stack
- **[Security Features](security.en.md)** - Security mechanisms explained

### 👨‍Development

- **[Architecture Design](architecture.en.md)** - System architecture and tech stack
- **[Contributing Guide](contributing.en.md)** - How to contribute to the project
- **[Deployment Guide](deployment.en.md)** - Server deployment methods

## Frequently Asked Questions

### Installation & Configuration

- [System Requirements](configuration.en.md#system-requirements)
- [Database Configuration](configuration.en.md#database-configuration)
- [Client Configuration](configuration.en.md#client-configuration)

### Usage Issues

- [Client Cannot Connect](getting-started.en.md#1-client-cannot-connect-to-server)
- [Linux Keyring Error](getting-started.en.md#2-linux-keyring-error)
- [Database Connection Failed](getting-started.en.md#3-database-connection-failed)

### Development Issues

- [Project Structure](architecture.en.md#project-structure)
- [Coding Standards](contributing.en.md#coding-standards)
- [How to Contribute](contributing.en.md)

## Important Notes

### Client Configuration

When configuring the server address, **always include the protocol prefix**:

```bash
# Correct
export GATE_SERVER_ADDR=http://localhost:4514

# Wrong (will default to https://)
export GATE_SERVER_ADDR=localhost:4514
```

See: [Client Configuration](configuration.en.md#client-configuration)

## Document List

### User Documentation
- [Getting Started](getting-started.en.md)
- [Configuration](configuration.en.md)
- [Client Guide](client-guide.en.md)
- [API Documentation](api.en.md)

### Technical Documentation
- [Architecture Design](architecture.en.md)
- [Security Features](security.en.md)

### Development Documentation
- [Architecture Design](architecture.en.md)
- [Contributing Guide](contributing.en.md)
- [Deployment Guide](deployment.en.md)

## External Resources

### Go Learning
- [Official Go Documentation](https://go.dev/doc/)
- [Go by Example](https://gobyexample.com/)

### Framework Documentation
- [Echo Documentation](https://echo.labstack.com/)
- [GORM Documentation](https://gorm.io/docs/)
- [Cobra Documentation](https://github.com/spf13/cobra)

## Suggested Reading Order

**New Users**:
1. Getting Started
2. Configuration
3. Client Guide

**Developers**:
1. Architecture Design
2. Contributing Guide
3. Security Features

**DevOps**:
1. Configuration
2. Security Features
3. Deployment Guide

---

Have questions? Check the [FAQ](#-frequently-asked-questions) or ask on [GitHub Issues](https://github.com/sokx6/imperishable-gate/issues).
