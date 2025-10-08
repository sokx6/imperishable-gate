# Architecture Design | Stage 1-6 System Evolution

**[📖 简体中文](architecture.md) | [📘 English](architecture.en.md)**

> 🏯 *"From simple client-server framework to complete user authentication system..."*

## Project Background and Design Philosophy

**Imperishable Gate** adopts a **frontend-backend separation architecture**. The entire system's design is inspired by the Stage structure of Touhou Perfect Cherry Blossom, evolving progressively from **Stage 1 "White & Pink Spring"** to **Stage 6 "The Corpse of the Netherworld Mistress"**.

### 🌸 Architecture Evolution Roadmap

- **Stage 1**: White & Pink Spring - Basic client/server communication framework (Ping)
- **Stage 2**: Black Cat of the Lost Home - Database integration, supporting SQLite/MySQL/PostgreSQL
- **Stage 3**: Night of the Doll's Judgment - Tag, alias, and remark system
- **Stage 4**: Cherry Blossom Barrier on Snow - Metadata crawler, intelligent monitoring
- **Stage 5**: Phantom Expanse of Hakugyokurou Stairs - Search, quick open functionality
- **Stage 6**: The Corpse of the Netherworld Mistress - Complete user system, JWT authentication

## System Architecture Overview

### Core Components

- **Backend Service (gate-server)**: High-performance RESTful API service based on Go + Echo + GORM
- **CLI Client (gate)**: Command-line tool based on Cobra framework with cross-platform support
- **Database Layer**: Supports SQLite (default) / MySQL / PostgreSQL

## Architecture Diagram

```
┌─────────────────────────────────────────────────────────────┐
│                      CLI Client (gate)                      │
│  ┌──────────────┐   ┌──────────────┐  ┌──────────────────┐  │
│  │   Commands   │   │   Services   │  │  System Keyring  │  │
│  └──────┬───────┘   └──────┬───────┘  └────────┬─────────┘  │
│         │                  │                   │            │
│         └──────────────────┴───────────────────┘            │
│                            │                                │
│                       HTTP/JSON                             │
└────────────────────────────┼────────────────────────────────┘
                             │
┌────────────────────────────┼────────────────────────────────┐
│                       RESTful API                           │
│  ┌─────────────────────────┴─────────────────────────────┐  │
│  │              Echo Web Framework                       │  │
│  │  ┌──────────┐  ┌────────────┐  ┌─────────────────┐    │  │
│  │  │  Routes  │→ │ Middleware │→ │    Handlers     │    │  │
│  │  └──────────┘  └────────────┘  └────────┬────────┘    │  │
│  │                                         │             │  │
│  └─────────────────────────────────────────┼─────────────┘  │
│                                            │                │
│  ┌─────────────────────────────────────────┼─────────────┐  │
│  │                  Services               │             │  │
│  │  ┌────────────┐  ┌──────────┐  ┌────────┴─────────┐   │  │
│  │  │    JWT     │  │ Metadata │  │   Link Monitor   │   │  │
│  │  │  Service   │  │  Crawler │  │  (Goroutines)    │   │  │
│  │  └────────────┘  └──────────┘  └──────────────────┘   │  │
│  └───────────────────────────────────────────────────────┘  │
│                             │                               │
│                         GORM ORM                            │
└─────────────────────────────┼───────────────────────────────┘
                              │
┌─────────────────────────────┼───────────────────────────────┐
│                       PostgreSQL                            │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────┐     │
│  │  users   │  │  links   │  │   tags   │  │  names   │     │
│  └──────────┘  └──────────┘  └──────────┘  └──────────┘     │
│  ┌──────────────────────┐  ┌──────────────────────┐         │
│  │   refresh_tokens     │  │     link_tags        │         │
│  └──────────────────────┘  └──────────────────────┘         │
└─────────────────────────────────────────────────────────────┘
```

## Directory Structure

```
cmd/
  gate/             # CLI client entry point
  gate-server/      # Server entry point
internal/
  client/           # Client implementation
    cmd/            # CLI command definitions
    service/        # Client business logic
    utils/          # Client utility functions
  model/            # Data models and entities
  server/           # Server core logic
    database/       # Database initialization and migration
    handlers/       # HTTP request handlers
    middlewares/    # Middleware (JWT authentication, etc.)
    routes/         # Route registration
    service/        # Server business logic
    utils/          # Server utility functions
  types/            # Type definitions
    request/        # Request types
    response/       # Response types
    jwt/            # JWT-related types
    data/           # Data types
```

## Core Components

### Backend Components

#### 1. Echo Web Framework
- High-performance HTTP routing and middleware
- RESTful API design
- Request/response handling

#### 2. GORM + PostgreSQL
- Data persistence
- Object-relational mapping
- Database migration

#### 3. JWT Authentication Service
- Access Token (short-term)
- Refresh Token (long-term)
- Automatic token refresh mechanism

#### 4. Metadata Crawler
- Automatic webpage title, description, and keyword extraction
- Concurrent crawling support

#### 5. Link Monitoring Service
- Scheduled link status checks
- Content change detection
- Email notifications (optional)

### Client Components

#### 1. Cobra CLI Framework
- Command-line interface
- Parameter parsing
- Subcommand management

#### 2. System Keyring Integration
- Secure token storage
- Cross-platform support (Linux/macOS/Windows)

#### 3. HTTP Client
- RESTful API calls
- Automatic token management
- Error handling

## Data Models

### User
- ID: Primary key
- Username: Username (unique)
- Password: Password (bcrypt encrypted)
- Email: Email address
- EmailVerified: Email verification status
- Links: Associated link list
- Tags: Associated tag list

### Link
- ID: Primary key
- UserID: Owner user ID
- Url: Link address (unique per user)
- Tags: Many-to-many associated tags
- Names: One-to-many associated aliases
- Remark: Remark information
- Title: Webpage title (auto-crawled)
- Description: Webpage description (auto-crawled)
- Keywords: Webpage keywords (auto-crawled)
- Watching: Monitoring status
- StatusCode: HTTP status code
- ContentHash: Content hash (for change detection)

### Tag
- ID: Primary key
- UserID: Owner user ID
- Name: Tag name (unique per user)
- Links: Many-to-many associated links

### Name (Alias)
- ID: Primary key
- LinkID: Owner link ID
- Name: Alias (globally unique)

### RefreshToken
- ID: Primary key
- UserID: Owner user ID
- Token: Token string (unique)
- ExpiresAt: Expiration time
- CreatedAt: Creation time

## Technology Stack

### Backend Technologies
- **Go 1.25.1+**: Programming language
- **Echo v4**: Web framework
- **GORM**: ORM framework
- **PostgreSQL**: Relational database
- **golang-jwt/jwt v5**: JWT authentication
- **goquery**: Web scraping
- **bcrypt**: Password encryption

### Client Technologies
- **Cobra**: CLI framework
- **go-keyring**: Credential storage
- **net/http**: HTTP client
