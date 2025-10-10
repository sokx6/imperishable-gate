# Imperishable Gate Documentation Center | 不朽之门

**[简体中文](INDEX.md) | [English](INDEX.en.md)**

> *"Want to deeply understand this system? Let's open the scrolls of Hakugyokurou..."* 🌸

Welcome to the **Imperishable Gate** documentation center! This project is a modern command-line link management system inspired by the Touhou Project, fully implementing all features from **Stage 1 "White & Pink Spring"** to **Stage 6 "The Corpse of the Netherworld Mistress"**.

---

## Stage 1-2 | White & Pink Spring & Black Cat of the Lost Home - Quick Start

### Must-Read for Beginners

- **[Getting Started Guide](getting-started.en.md)**  
  *Stage 1 Basics: From installation to running your first command*
  - Environment setup (Go, database selection: SQLite/MySQL/PostgreSQL)
  - Server and client installation
  - Basic command usage examples (add / list / ping)
  
- **[Configuration](configuration.en.md)**  
  *Stage 2 Database: Multi-database support configuration*
  - Server environment variable configuration (DB_TYPE, DSN, JWT_SECRET)
  - Client configuration (GATE_SERVER_ADDR)
  - Database connection string examples

---

## Stage 3-5 | Night of the Doll's Judgment & Cherry Blossom Barrier on Snow & Phantom Expanse of Hakugyokurou Stairs - Features

### Client Usage Guide

- **[Complete Client Manual](client-guide.en.md)**  
  *Detailed explanation of all CLI commands*
  - **Stage 2-3**: Link, tag, alias, and notes management
  - **Stage 4**: Monitoring (watch) and metadata crawling
  - **Stage 5**: Search and quick open
  - **Stage 6**: User system (register / login / whoami)

### API Development Documentation

- **[RESTful API Documentation](api.en.md)**  
  *Complete HTTP API interface documentation*
  - Authentication-related APIs (JWT Token)
  - Link management APIs (CRUD)
  - Tag and alias APIs
  - Monitoring and notification APIs

---

## Stage 1-6 | Architecture & Development - Technical Deep Dive

### Developer Documentation

- **[Architecture Design](architecture.en.md)**  
  *System evolution from Stage 1 to Stage 6*
  - Client-server separation architecture
  - Database design (ER diagrams)
  - Technology stack selection rationale
  - Module division for each Stage

- **[Contributing Guide](contributing.en.md)**  
  *How to contribute code to the project*
  - Fork and Pull Request workflow
  - Commit message conventions
  - Code review standards

---

## Deployment & Security - Production Environment

### Operations Documentation

- **[Deployment Guide](deployment.en.md)**  
  *Production environment deployment best practices*
  - Server configuration
  - systemd service setup
  - Nginx reverse proxy
  - Backup strategies

- **[Security Features Explained](security.en.md)**  
  *Stage 6: Deep dive into security mechanisms*
  - JWT dual-token authentication principle (Access Token + Refresh Token)
  - bcrypt password encryption (surely no one stores passwords in plaintext?)
  - Keyring secure storage (libsecret / Secret Service)
  - Data isolation and access control

---

## Development Log & Records

### Project Journey

- **[Devlog](devlog.md)**  
  *Development process records and insights*
  - Implementation process for each Stage
  - Problems encountered and solutions
  - Reflections on technology selection

---

## Touhou Project Elements

This project is deeply inspired by the Touhou Project series:

### Name Origins
- **Project Name**: Imperishable Night (永夜抄)
- **Stage Architecture**: Perfect Cherry Blossom (妖妖梦)

### Stage Titles (from Perfect Cherry Blossom)
1. **Stage 1**: White & Pink Spring (白銀之春)
2. **Stage 2**: Black Cat of the Lost Home (迷途之家の黒猫)
3. **Stage 3**: Night of the Doll's Judgment (人偶裁判の夜)
4. **Stage 4**: Cherry Blossom Barrier on Snow (雪上の櫻花結界)
5. **Stage 5**: Phantom Expanse of Hakugyokurou Stairs (白玉樓階梯の幻闊)
6. **Stage 6**: The Corpse of the Netherworld Mistress (冥界大小姐の亡骸)

### Worldview Setting
- **Location**: Hakugyokurou (the garden of the Saigyouji family)
- **Identity**: Gardener (sweeping the yard, cooking for the master)
- **Goal**: Managing links collected while surfing the internet

---

## FAQ - Frequently Asked Questions

### Q: Where should I start?
**A**: Beginners should start with the **[Getting Started Guide](getting-started.en.md)**! It will guide you through the Stage 1 basic setup.

### Q: How do I switch databases?
**A**: Check the database configuration section in **[Configuration](configuration.en.md)**. Supports SQLite (default) / MySQL / PostgreSQL.

### Q: Client can't connect to the server?
**A**: Make sure to include the `http://` or `https://` prefix when setting `GATE_SERVER_ADDR`!

### Q: Is token storage secure?
**A**: Very secure! Check out **[Security Features](security.en.md)** to learn about the Keyring mechanism (Stage 6 implementation).

### Q: How can I contribute code?
**A**: Read the **[Contributing Guide](contributing.en.md)**, then submit your Pull Request! Follow the `feat:`, `fix:`, `docs:` commit conventions.

### Q: What features are supported?
**A**: 
- **Stage 1-2**: Basic client/server architecture, database integration
- **Stage 3**: Tag, alias (Name), and notes systems
- **Stage 4**: Metadata crawling, intelligent monitoring, email notifications
- **Stage 5**: Search, quick open
- **Stage 6**: User system, JWT authentication, email verification

---

<div align="center">

## Quick Links

| I want to... | View Documentation |
|-------------|-------------------|
| Quick install and run | [Getting Started](getting-started.en.md) |
| Configure database and server | [Configuration](configuration.en.md) |
| Learn all commands | [Client Guide](client-guide.en.md) |
| Understand API interfaces | [API Documentation](api.en.md) |
| Understand system architecture | [Architecture Design](architecture.en.md) |
| Participate in development | [Contributing Guide](contributing.en.md) |
| Deploy to production | [Deployment Guide](deployment.en.md) |
| Understand security mechanisms | [Security Features](security.en.md) |

---

### *"After seven days of intense development, my mysterious wonderful software has finally released version 1.0"* 🌸

**[Back to Home](../README.en.md)** | **[GitHub Repository](https://github.com/sokx6/imperishable-gate)**

*Made with love and Go | Inspired by Touhou Project 上海アリス幻樂団*

</div>
