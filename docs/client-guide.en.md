# Gate CLI Usage Guide | Complete Command Reference

**[简体中文](client-guide.md) | [English](client-guide.en.md)**

> *"Elegant command line, the gardener's reliable assistant"*

## Table of Contents
- [Introduction](#introduction)
- [Installation & Configuration](#installation--configuration)
- [Global Parameters](#global-parameters)
- [Command List](#command-list)
  - [register - Register User](#register---register-user)
  - [login - Login](#login---login)
  - [logout - Logout](#logout---logout)
  - [whoami - Show Current User](#whoami---show-current-user)
  - [ping - Test Connection](#ping---test-connection)
  - [add - Add Link](#add---add-link)
  - [list - List Links](#list---list-links)
  - [delete - Delete Link](#delete---delete-link)
  - [watch - Monitor Link](#watch---monitor-link)
  - [open - Open Link](#open---open-link)
- [Usage Examples](#usage-examples)
- [FAQ](#faq)

---

## Introduction

`gate` is the command-line client tool for **Imperishable Gate**, enabling the gardeners of Hakugyokurou to elegantly manage their links!

### Feature Highlights

- **Stage 2-3**: CRUD operations for links, alias management, tag categorization
- **Stage 4**: Link change monitoring, automatic metadata crawling
- **Stage 5**: Search and quick link opening
- **Stage 6**: User authentication, automatic token management

### Core Feature: Rich Tag-Based Command System

The client implements powerful link management functionality through a flexible tag system:

- **Multi-dimensional Retrieval**: Quickly locate links through tags, URLs, or aliases
- **Tag Combination Queries**: Search with multiple tag combinations
- **Batch Tag Operations**: Add or remove tags for multiple links at once
- **Smart Tag Commands**:
  - `gate search -t "tag1,tag2"` - Search by tag combination
  - `gate add -l url -t "tag1,tag2"` - Set tags when adding links
  - `gate list -t tag` - List all links with a specific tag
  - `gate add -n name -t "new_tag"` - Add tags to existing links

---

## Installation & Configuration

### Build the Client

```bash
cd cmd/gate
go build -o gate
```

### Configure Server Address

There are three ways to configure the server address (in priority order):

1. **Command Line Arguments** (Highest Priority)
   ```bash
   gate <command> -a http://localhost:4514
   ```

2. **Environment Variable** (via `.env` file)
   Create a `.env` file in the project root:
   ```bash
   SERVER_ADDR=localhost:4514
   ```

3. **Default Value** (Lowest Priority)
   ```
   localhost:4514
   ```

---

## Global Parameters

All commands support the following global parameters:

| Parameter | Short | Description | Example |
|-----------|-------|-------------|---------|
| `--addr` | `-a` | Server address | `-a localhost:4514` |
| `--help` | `-h` | Show help information | `gate add -h` |
| `--version` | | Show version information | `gate --version` |

---

## Command List

### register - Register User

Register a new user account.

**Syntax:**
```bash
gate register [--addr <server-address>]
```

**Examples:**
```bash
# Register using default or configured server address
gate register

# Register with specified server address
gate register -a localhost:4514
```

**Interactive Flow:**
```
Please enter your username: myuser
Please enter your email: user@example.com
Please enter your password: ********
Confirm your password: ********
```

---

### login - Login

Login to the server and securely store the refresh token in the system keyring.

**Syntax:**
```bash
gate login [--addr <server-address>]
```

**Examples:**
```bash
# Login using default or configured server address
gate login

# Login with specified server address
gate login -a localhost:4514
```

**Interactive Flow:**
```
Please enter your username: myuser
Please enter your password: ********
Refresh token saved to system keyring.
Login successful!
```

**Notes:**
- After successful login, the refresh token is stored in the system keyring
- Subsequent commands will automatically use the stored token for authentication
- If the access token expires, it will be automatically refreshed

---

### logout - Logout

Logout and clear locally stored tokens.

**Syntax:**
```bash
gate logout [--addr <server-address>]
```

**Examples:**
```bash
# Logout
gate logout

# Logout with specified server address
gate logout -a localhost:4514
```

**Output Example:**
```
Logged out successfully
Tokens cleared from system keyring.
Logout successful!
```

**Notes:**
- Sends a logout request to the server to invalidate the refresh token
- Clears tokens stored in the local system keyring
- Even if the server request fails, local tokens will be cleared
- Command completes successfully even if no stored token is found

---

### whoami - Show Current User

Display information about the currently authenticated user.

**Syntax:**
```bash
gate whoami [--verbose] [--addr <server-address>]
```

**Parameters:**
| Parameter | Short | Description |
|-----------|-------|-------------|
| `--verbose` | `-v` | Show detailed JSON response |

**Examples:**
```bash
# Show current user information
gate whoami

# Show detailed information
gate whoami -v

# Specify server address
gate whoami -a localhost:4514
```

**Output Example:**
```
Authenticated as:
  User ID:  1
  Username: myuser
```

**Detailed Output Example (with -v flag):**
```
Authenticated as:
  User ID:  1
  Username: myuser

Detailed response:
{
  "message": "Success",
  "user_info": {
    "user_id": 1,
    "username": "myuser"
  }
}
```

**Notes:**
- Requires login before using this command
- Used to confirm current authentication status and user identity
- Can be used to debug authentication issues

---

### ping - Test Connection

Test the connection to the server.

**Syntax:**
```bash
gate ping [--message <message>] [--addr <server-address>]
```

**Parameters:**
| Parameter | Short | Default | Description |
|-----------|-------|---------|-------------|
| `--message` | `-m` | "default message" | Test message to send |

**Examples:**
```bash
# Send default message
gate ping

# Send custom message
gate ping -m "Hello Server"

# Specify server address
gate ping -a localhost:4514 -m "test"
```

---

### add - Add Link

Add links, names, tags, or remarks.

**Syntax:**
```bash
gate add --link <URL> [--name <name>...] [--tag <tag>...] [--remark <remark>]
gate add --name <name> --tag <tag>... [--remark <remark>]
gate add --link <URL> --remark <remark>
gate add --name <name> --remark <remark>
```

**Parameters:**
| Parameter | Short | Description |
|-----------|-------|-------------|
| `--link` | `-l` | Link URL |
| `--name` | `-n` | Link alias (multiple allowed) |
| `--tag` | `-t` | Tag (multiple allowed) |
| `--remark` | `-r` | Remark information |

**Use Cases:**

1. **Add new link**
   ```bash
   gate add -l https://github.com
   ```

2. **Add link with name**
   ```bash
   gate add -l https://github.com -n github -n gh
   ```

3. **Add link with tags**
   ```bash
   gate add -l https://github.com -t dev -t tools
   ```

4. **Add link with name and tags**
   ```bash
   gate add -l https://github.com -n github -t dev -t tools
   ```

5. **Add name to existing link**
   ```bash
   gate add -l https://github.com -n newname
   ```

6. **Add tag to link (by URL)**
   ```bash
   gate add -l https://github.com -t newtag
   ```

7. **Add tag to link (by name)**
   ```bash
   gate add -n github -t newtag
   ```

8. **Add remark to link (by URL)**
   ```bash
   gate add -l https://github.com -r "Code hosting platform"
   ```

9. **Add remark to link (by name)**
   ```bash
   gate add -n github -r "Code hosting platform"
   ```

---

### list - List Links

Query and list link information.

**Syntax:**
```bash
gate list [--tag <tag>] [--name <name>] [--page <page>] [--page-size <size>]
```

**Parameters:**
| Parameter | Short | Default | Description |
|-----------|-------|---------|-------------|
| `--tag` | `-t` | | Filter by tag |
| `--name` | `-n` | | Query by name |
| `--page` | `-p` | 1 | Page number |
| `--page-size` | `-s` | 20 | Items per page |

**Examples:**

1. **List all links**
   ```bash
   gate list
   ```

2. **Query by tag**
   ```bash
   gate list -t dev
   ```

3. **Query by name**
   ```bash
   gate list -n github
   ```

4. **Paginated query**
   ```bash
   gate list -p 2 -s 10
   ```

5. **Query by tag with pagination**
   ```bash
   gate list -t tools -p 1 -s 20
   ```

---

### delete - Delete Link

Delete links, names, or tags.

**Syntax:**
```bash
gate delete --link <URL>... [--name <name>...] [--tag <tag>...]
gate delete --name <name>... [--tag <tag>...]
```

**Parameters:**
| Parameter | Short | Description |
|-----------|-------|-------------|
| `--link` | `-l` | Link URL to delete (multiple allowed) |
| `--name` | `-n` | Name to delete (multiple allowed) |
| `--tag` | `-t` | Tag to delete (multiple allowed) |

**Use Cases:**

1. **Delete link**
   ```bash
   gate delete -l https://github.com
   ```

2. **Batch delete links**
   ```bash
   gate delete -l https://github.com -l https://google.com
   ```

3. **Delete link by name**
   ```bash
   gate delete -n github
   ```

4. **Delete a name from link**
   ```bash
   gate delete -l https://github.com -n oldname
   ```

5. **Delete tag from link (by URL)**
   ```bash
   gate delete -l https://github.com -t oldtag
   ```

6. **Delete tag from link (by name)**
   ```bash
   gate delete -n github -t oldtag
   ```

---

### watch - Monitor Link

Monitor or unmonitor link changes; email notifications are sent when link content changes.

**Syntax:**
```bash
gate watch --link <URL> --watch
gate watch --link <URL> --unwatch
gate watch --name <name> --watch
gate watch --name <name> --unwatch
```

**Parameters:**
| Parameter | Short | Description |
|-----------|-------|-------------|
| `--link` | `-l` | Link URL to monitor |
| `--name` | `-n` | Link name to monitor |
| `--watch` | `-w` | Enable monitoring |
| `--unwatch` | `-u` | Disable monitoring |

**Use Cases:**

1. **Enable monitoring by URL**
   ```bash
   gate watch -l https://github.com -w
   ```

2. **Disable monitoring by URL**
   ```bash
   gate watch -l https://github.com -u
   ```

3. **Enable monitoring by name**
   ```bash
   gate watch -n github -w
   ```

4. **Disable monitoring by name**
   ```bash
   gate watch -n github -u
   ```

**Notes:**
- Cannot specify both `--link` and `--name` simultaneously
- Must specify either `--watch` or `--unwatch`
- Monitoring feature requires email service configuration on the server

---

### open - Open Link

Open a link in the browser.

**Syntax:**
```bash
gate open --name <name>
gate open --tag <tag> [--page <page>] [--page-size <size>]
```

**Parameters:**
| Parameter | Short | Default | Description |
|-----------|-------|---------|-------------|
| `--name` | `-n` | | Link name |
| `--tag` | `-t` | | Tag |
| `--page` | `-p` | 1 | Page number (for tag-based opening) |
| `--page-size` | `-s` | 10 | Page size (for tag-based opening) |

**Use Cases:**

1. **Open link by name**
   ```bash
   gate open -n github
   ```

2. **Open link by tag**
   ```bash
   gate open -t dev
   ```

3. **Open link by tag (specify page)**
   ```bash
   gate open -t tools -p 2 -s 5
   ```

**Note:**
- When opening by name, it directly opens the link associated with that name
- When opening by tag, it lists all matching links for selection

---

## Usage Examples

### Advanced Tag System Applications (Client Feature)

The tag system is a core feature of this client. Here are some advanced usage tips:

#### 1. Tag Combination Queries
```bash
# Search by single tag
gate search -t dev

# Search by multiple tag combinations (comma-separated)
gate search -t "dev,tools"
gate search -t "dev,tools,favorite"
```

#### 2. Hierarchical Tag Management
```bash
# Add hierarchical tags for technical documentation
gate add -l https://go.dev/doc -n godoc -t "dev,golang,docs"
gate add -l https://docs.python.org -n pydoc -t "dev,python,docs"

# Query all development-related documentation
gate list -t dev
gate list -t docs

# Query documentation for specific languages
gate list -t golang
gate list -t python
```

#### 3. Batch Tag Operations
```bash
# Add the same tag to multiple links
gate add -n github -t important
gate add -n stackoverflow -t important
gate add -n docs -t important

# Batch delete tags
gate delete -n github -t outdated
gate delete -n oldsite -t outdated
```

#### 4. Quick Open by Tag
```bash
# Open links with a specific tag
gate open -t dev

# Open links with multiple tags
gate open -t "dev,docs"
```

#### 5. Tag Search and Filtering
```bash
# Search for links containing a specific tag
gate search -t "tutorial"

# List all links under a specific tag
gate list -t "golang"

# View links under a tag with pagination
gate list -t "dev" -p 1 -s 10
```

### Complete Workflow Example

```bash
# 1. Register account
gate register

# 2. Login
gate login

# 3. Test connection
gate ping

# 4. Add link
gate add -l https://github.com -n github -n gh -t dev -t tools -r "Code hosting platform"

# 5. View all links
gate list

# 6. Query by tag
gate list -t dev

# 7. Enable link monitoring
gate watch -n github -w

# 8. Open link
gate open -n github

# 9. Add more tags
gate add -n github -t favorite

# 10. Delete a tag
gate delete -n github -t favorite

# 11. Disable monitoring
gate watch -n github -u

# 12. Delete link
gate delete -n github
```

### Batch Operations Examples

```bash
# Batch add links
gate add -l https://github.com -n github -n gh
gate add -l https://google.com -n google -n gg
gate add -l https://stackoverflow.com -n so -t dev

# Batch add tags
gate add -n github -t dev -t tools -t favorite
gate add -n google -t search -t tools

# Batch delete tags
gate delete -n github -t favorite
gate delete -l https://google.com -t search

# Batch delete links
gate delete -l https://github.com -l https://google.com
```

---

## FAQ

### 1. How to configure server address?

Three methods (in priority order):
1. Command line argument: `gate <command> -a localhost:4514`
2. Environment variable: Set `SERVER_ADDR=localhost:4514` in `.env` file
3. Default value: `localhost:4514`

### 2. Where is login information stored?

Refresh tokens are securely stored in the system keyring:
- macOS: Keychain
- Linux: Secret Service API (gnome-keyring, kwallet, etc.)
- Windows: Windows Credential Manager

### 3. What if the access token expires?

The client automatically detects if the access token has expired. If expired, it will automatically obtain a new access token using the refresh token. If the refresh token has also expired, you'll be prompted to login again.

### 4. How to view help information?

```bash
# View all commands
gate --help

# View help for specific command
gate add --help
gate list --help
```

### 5. How does link monitoring work?

- The server periodically checks monitored links for changes
- If changes are detected, notifications are sent via email
- Requires email service configuration on the server side

### 6. Can a link have multiple names?

Yes. A link can have multiple aliases (names), making it easy to access the same link through different names.

```bash
gate add -l https://github.com -n github -n gh -n git
```

### 7. How to update a link's remark?

Simply use the add command to add a new remark, which will overwrite the old one:

```bash
gate add -n github -r "New remark content"
```

### 8. What's the difference between delete operations?

- `gate delete -l <url>`: Deletes the entire link and all associated data
- `gate delete -n <name>`: Deletes a name from the link (if it's the last name, deletes the entire link)
- `gate delete -l <url> -t <tag>`: Only deletes a tag from the link
- `gate delete -l <url> -n <name>`: Only deletes a name from the link

---

## Version Information

Current Version: **v1.0.0**

Check version:
```bash
gate --version
```

---

## Technical Support

For questions or suggestions, please contact the project maintainers or submit an Issue.
