# Imperishable Gate | 不朽之门

<div align="center">

![Go Version](https://img.shields.io/badge/Go-1.25.1+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-green.svg?style=flat)
![Database](https://img.shields.io/badge/Database-SQLite%20%7C%20MySQL%20%7C%20PostgreSQL-blue?style=flat)
![Echo](https://img.shields.io/badge/Echo-Web_Framework-00C7B7?style=flat)
![GORM](https://img.shields.io/badge/GORM-ORM-red?style=flat)
![Cobra](https://img.shields.io/badge/Cobra-CLI-orange?style=flat)
![JWT](https://img.shields.io/badge/JWT-Auth-black?style=flat)

**Link Management System by a Hakugyokurou Gardener**

*Named after Imperishable Night, with architecture inspired by Perfect Cherry Blossom*

**[简体中文](README.md) | [English](README.en.md)**

[Quick Start](#stage-1--white--pink-spring---quick-start) • [Features](#core-features) • [Documentation](#documentation) • [Contributing](#contributing)

</div>

---

## Background

> *One day years ago, I unfortunately fell into the Netherworld by sheer bad luck. Now I've become an ordinary gardener living at Hakugyokurou. Besides sweeping the yard and cooking for my master every day, I try some interesting things in my spare time...*

**Imperishable Gate** is a modern command-line link management system inspired by the Touhou Project. When internet infrastructure extended to the Hakugyokurou area, as a gardener, I stepped onto the "information superhighway" only to find browser bookmarks too simplistic, PC software too bulky, and documentation too... green-skinned.

So, I decided to use the programming skills learned in my previous life to create an elegant link management system!

### Project Highlights

This project fully implements all features from **Stage 1 "White & Pink Spring"** to **Stage 6 "The Corpse of the Netherworld Mistress"**, providing complete link lifecycle management including adding, deleting, querying, tag categorization, alias management, notes, automatic metadata crawling, and intelligent link monitoring.

**Featured: Rich Tag-Based Command System** - The client implements a powerful tag system that enables multi-dimensional link management commands, supporting tag-based search, batch operations, tag combination queries, and more for efficient and convenient link management.

## Architecture

This project adopts a **client-server separation architecture**, inspired by the layered design of Touhou Youyoumu:

- **Stage 1-2 | Foundation Layer (White & Pink Spring · Black Cat of the Lost Home)**
  - CLI Client: Command-line tool based on the Cobra framework
  - Server: RESTful API service based on Go + Echo
  - Database: Supports SQLite / MySQL / PostgreSQL

- **Stage 3-4 | Feature Enhancement Layer (Night of the Doll's Judgment · Cherry Blossom Barrier on Snow)**
  - Tag system, alias management, notes functionality
  - Automatic metadata crawling (title, description, keywords)
  - Intelligent monitoring system (content change detection, email notifications)

- **Stage 5-6 | Security Authentication Layer (Phantom Expanse of Hakugyokurou Stairs · The Corpse of the Netherworld Mistress)**
  - JWT dual-token authentication (Access Token + Refresh Token)
  - User registration, login, email verification
  - Secure token storage (system Keyring)

## Core Features

### Stage 6 | The Corpse of the Netherworld Mistress - Security Authentication

- **Dual Token Mechanism**: JWT Access Token (short-term) + Refresh Token (long-term)
- **Password Encryption**: Uses bcrypt algorithm for secure storage
- **Secure Storage**: Tokens stored in system keyring (libsecret / Secret Service)
- **Auto Refresh**: Tokens refresh automatically on expiration, seamless experience
- **Email Verification**: Complete email verification flow

### Stage 3 | Night of the Doll's Judgment - Flexible Tag System (Client Feature)

The client implements rich command features through the tag system, which is a core highlight of this project:

- **Many-to-Many Relations**: One link can have multiple tags, one tag can be associated with multiple links
- **Tag Combination Queries**: Search links with multiple tag combination conditions
- **Batch Tag Operations**: Support bulk tag addition/deletion via URL or alias
- **Tag Categorization Management**: Manage and retrieve links by tag categories
- **User Isolation**: Each user has an independent tag namespace
- **Smart Tag Commands**:
  - `gate search -t "tag1,tag2"` - Search by tags
  - `gate add -l url -t "tag1,tag2"` - Set tags when adding links
  - `gate list -t tag` - List all links with a specific tag

### Stage 2-3 | Black Cat of the Lost Home & Night of the Doll's Judgment - Core Link Management

- **Multi-dimensional Search**: Quick queries by URL, tags, or aliases
- **Batch Operations**: Add or delete multiple links, tags, aliases at once
- **Alias System** (Name): Set multiple aliases for links to avoid repetitive long URLs
- **Notes Feature**: Add personalized notes to each link
- **Association Management**: Flexible relationships between links, tags, and aliases

### Stage 4 | Cherry Blossom Barrier on Snow - Intelligent Monitoring

- **Metadata Crawling**: Auto-fetch webpage titles, descriptions, keywords
- **Scheduled Monitoring**: Background scheduled checking of link status and content changes
- **Auto Detection**: Automatically discover webpage content updates
- **Tiered Monitoring**: Two modes - watching (high-frequency) and non-watching (low-frequency)
- **Change Notifications**: Email notifications on content changes (SMTP protocol)

### Stage 5 | Phantom Expanse of Hakugyokurou Stairs - Convenient CLI Experience

- **Interactive Interface**: Friendly command-line interaction
- **Smart Prompts**: Clear error messages and operation guidance
- **Auto Authentication**: Intelligent token management, no frequent logins needed
- **Quick Open**: `gate open` to directly open links in browser
- **Cross-platform Support**: Linux, macOS, Windows

## Stage 1 | White & Pink Spring - Quick Start

### Prerequisites

- Go 1.25.1+
- Database (choose one):
  - **SQLite** (default, no additional installation required)
  - **MySQL** 5.7+ / 8.0+
  - **PostgreSQL** 12.0+

### Installation

```bash
# 1. Clone the repository
git clone https://github.com/sokx6/imperishable-gate.git
cd imperishable-gate

# 2. Install dependencies
go mod download

# 3. Configure environment variables (optional)
# Server configuration (uses SQLite by default, no extra config needed)
# Copy .env.example from project root to cmd/gate-server/
cp .env.example cmd/gate-server/.env

# Client configuration (optional, to set server address)
# .env.example is already in cmd/gate/ directory
cp cmd/gate/.env.example cmd/gate/.env
# Or use environment variable directly
export GATE_SERVER_ADDR=http://localhost:4514

# 4. Build
go build -o bin/gate-server ./cmd/gate-server
go build -o bin/gate ./cmd/gate

# 5. Start the server
./bin/gate-server start
```

### Client Usage

```bash
# Stage 6 | User System
# Configure server address (IMPORTANT: must include http:// or https:// prefix)
export GATE_SERVER_ADDR=http://localhost:4514

# Register new user
./bin/gate register
# Login
./bin/gate login
# Check current login status
./bin/gate whoami

# Stage 2-3 | Link Management
# Add link (auto-crawl metadata)
./bin/gate add -l "https://thwiki.cc" -t "touhou,wiki" -N "thwiki"

# Add notes to link
./bin/gate add -l "https://thwiki.cc" --remark "Touhou Project Chinese Wiki"

# View all links
./bin/gate list

# Search by tag
./bin/gate search -t "touhou"

# Stage 5 | Open Links
# Open directly by alias
./bin/gate open -n "thwiki"
# Open multiple at once
./bin/gate open -n "thwiki,pixiv"

# Stage 4 | Monitoring
# Enable monitoring for a link (high-frequency checking)
./bin/gate watch -n "thwiki"

# Tag-related operations (Client feature)
# Search by tag combination
./bin/gate search -t "touhou,wiki"
# Add tags to existing link
./bin/gate add -n "thwiki" -t "reference"
# List all links with a specific tag
./bin/gate list -t "touhou"
```

**Tips**:
- When setting `GATE_SERVER_ADDR`, always include the protocol prefix (`http://` or `https://`)
- Using aliases (Names) avoids repetitive long URL input
- Tokens are automatically stored in system Keyring, secure and reliable
- **Tag system is a core feature of the client**, making full use of tags enables efficient link management

For detailed usage, please refer to the [Getting Started Guide](docs/getting-started.en.md) and [Complete Client Documentation](docs/client-guide.en.md).

## Documentation

### User Documentation (Beginner's Guide)
- [Getting Started](docs/getting-started.en.md) - Stage 1: White & Pink Spring - Installation and basic usage
- [Configuration](docs/configuration.en.md) - Complete configuration for database, server, and client
- [Client Guide](docs/client-guide.en.md) - Detailed explanation of all CLI commands (focus on tag commands)
- [API Documentation](docs/api.en.md) - Complete RESTful API reference

### Developer Documentation (Advanced Content)
- [Architecture](docs/architecture.en.md) - System architecture evolution from Stage 1 to Stage 6
- [Contributing Guide](docs/contributing.en.md) - How to contribute code to the project
- [Deployment Guide](docs/deployment.en.md) - Production environment deployment best practices
- [Security Features](docs/security.en.md) - JWT, bcrypt, Keyring security mechanisms explained

### Development Log
- [Devlog](docs/devlog.md) - Development process records and insights

## Tech Stack

### Backend Server
- **Go 1.25.1+**: Programming language
- **Echo v4**: Lightweight web framework, RESTful API design
- **GORM**: ORM framework, elegant database operations
- **Database Support**:
  - **SQLite** (default) - Stage 2 basic implementation, zero configuration
  - **MySQL** - Stage 2 extended support
  - **PostgreSQL** - Stage 2 extended support
- **JWT (golang-jwt/jwt)**: Stage 6 authentication mechanism
- **bcrypt**: Stage 6 password encryption
- **goquery**: Stage 4 webpage metadata crawling
- **SMTP**: Stage 4 email notification functionality

### CLI Client
- **Cobra**: CLI framework, elegant command-line design
- **go-keyring**: Stage 6 credential secure storage (libsecret / Secret Service)
- **Cross-platform Support**: Linux / macOS / Windows

## Development Progress

### Completed (v1.0.0)

#### Stage 1 | White & Pink Spring
- Basic client & server framework
- RESTful API design
- Ping test functionality

#### Stage 2 | Black Cat of the Lost Home
- Database integration (SQLite/MySQL/PostgreSQL)
- Link CRUD operations
- Complete API routing design

#### Stage 3 | Night of the Doll's Judgment
- Tag management system (many-to-many relations)
- Alias management (Name → Link mapping)
- Notes functionality
- Query links by tags/aliases

#### Stage 4 | Cherry Blossom Barrier on Snow
- Automatic webpage metadata crawling (title, description, keywords)
- Scheduled polling mechanism (goroutine implementation)
- Link monitoring system (watching/non-watching)
- Email notification functionality (SMTP protocol)

#### Stage 5 | Phantom Expanse of Hakugyokurou Stairs
- Link search functionality (fuzzy search)
- `gate open` command (open in browser)
- Batch open multiple links

#### Stage 6 | The Corpse of the Netherworld Mistress
- Complete user system (register/login/logout/deactivate)
- JWT dual-token authentication (Access Token + Refresh Token)
- Email verification functionality
- Token secure storage (system Keyring)
- Automatic token refresh
- `whoami` command

### Planned (v2.0)
- ElasticSearch integration (Stage 5 extension)
- Collection/View objects (Stage 5 extension)
- Administrator system (Stage 6 extension)
- Group system (Stage 6 extension)
- Audit log system (Stage 6 extension)
- Link import/export functionality
- Unit test coverage

## Contributing

Contributions, bug reports, and feature requests are welcome!

### Contribution Workflow

1. **Fork** this project to your account
2. Create a feature branch: `git checkout -b feature/AmazingFeature`
3. Write code following project conventions:
   - Use `feat:`, `fix:`, `docs:`, `refactor:` etc. in commit messages
   - Pay attention to code organization, avoid writing all code in one file
   - Add necessary error handling
4. Commit changes: `git commit -m 'feat: Add some AmazingFeature'`
5. Push to branch: `git push origin feature/AmazingFeature`
6. Create a **Pull Request**

### Suggested Directions

- **Stage 5 Extensions**: ElasticSearch integration, Collection system
- **Stage 6 Extensions**: Administrator system, group sharing, audit logs
- **Test Improvements**: Unit tests, integration tests
- **Documentation**: More examples, best practices

See [Contributing Guide](docs/contributing.en.md) for details.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to all contributors and friends who starred this project
- Special thanks to the **Touhou Project (東方Project)** game series for inspiration
  - Project name origin: **Imperishable Night (永夜抄)**
  - Architecture design inspiration: **Perfect Cherry Blossom (妖妖梦)** Stage structure
  - Theme atmosphere: Hakugyokurou, Netherworld, spring snow cherry blossoms
- Thanks to THBWiki (Touhou Project Chinese Wiki) for providing rich resources

## Contact

- **GitHub Issues**: [Submit an issue](https://github.com/sokx6/imperishable-gate/issues)
- **Creator**: QQ 2841929072

---

<div align="center">

**[Back to Top](#imperishable-gate--不朽之门)**

Made with love and Go by [locxl](https://github.com/sokx6)

*Inspired by Touhou Project 上海アリス幻樂団 (Team Shanghai Alice)*

</div>
