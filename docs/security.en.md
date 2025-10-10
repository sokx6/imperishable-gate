# Security Features | Stage 6: The Corpse of the Netherworld Mistress

**[简体中文](security.md) | [English](security.en.md)**

> *"Want to pass through this gate? Prove your identity first!"*

This document details the security mechanisms implemented in **Stage 6「The Corpse of the Netherworld Mistress」**. As the gardener of Hakugyokurou, the cryptographic knowledge I learned in my previous life finally comes in handy!

## Authentication System Overview

This project implements a complete **JWT (JSON Web Token)** authentication system, including:
- Dual-token mechanism (Access Token + Refresh Token)
- bcrypt password encryption (surely no one stores passwords in plaintext?)
- System Keyring secure storage
- Automatic token refresh
- Email verification

## JWT Dual-Token Mechanism

### Token Types

**Access Token**
- Validity: 15 minutes
- Purpose: API request authentication
- Storage: System keyring

**Refresh Token**
- Validity: 7 days
- Purpose: Refresh Access Token
- Features: One-time use, deleted on logout

### Token Refresh Flow

The client handles token refresh automatically:

1. Token expires during API request
2. Automatically use Refresh Token to obtain new token
3. Retry original request
4. Seamless to the user

## Password Security

### Bcrypt Encryption

- Uses bcrypt algorithm for password encryption
- Automatic salting
- Irreversible encryption

### Password Requirements

- Minimum length: 6 characters
- Recommended: Include uppercase, lowercase letters, numbers, and special characters

**Example code**:
```go
// Password encryption
hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

// Password verification
err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
```

## Data Isolation

### User-Level Isolation

- All data is isolated by user ID
- Users can only access their own data
- Database queries automatically filtered

**Implementation**:
```go
// Extract user ID from JWT
userID := getUserIDFromJWT(c)

// Automatically filter during query
db.Where("user_id = ?", userID).Find(&links)
```

## Secure Storage

### Client Token Storage

Tokens are stored in the system's secure storage:

| System | Storage Location |
|--------|------------------|
| Linux | GNOME Keyring / KWallet |
| macOS | Keychain |
| Windows | Credential Manager |

**Usage example**:
```go
import "github.com/zalando/go-keyring"

// Save token
keyring.Set("gate", "access_token", token)

// Read token
token, _ := keyring.Get("gate", "access_token")

// Delete token
keyring.Delete("gate", "access_token")
```

### Configuration File Security

```bash
# Set .env file permissions
chmod 600 .env

# Don't commit to version control
echo ".env" >> .gitignore
```

## Database Security

### SQL Injection Protection

Use GORM parameterized queries:

```go
// Safe
db.Where("url = ?", url).First(&link)

// Dangerous (don't do this)
db.Raw("SELECT * FROM links WHERE url = '" + url + "'")
```

### Principle of Least Privilege

```sql
-- Create dedicated user
CREATE USER gateuser WITH PASSWORD 'password';

-- Grant only necessary permissions
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO gateuser;
```

## Input Validation

### Request Validation

```go
type RegisterRequest struct {
    Username string `json:"username" validate:"required,min=3,max=32"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=6"`
}
```

### URL Validation

```go
// Validate URL format
_, err := url.ParseRequestURI(urlString)
if err != nil {
    return errors.New("invalid URL")
}
```

## Logging Security

Don't log sensitive information:

```go
// Unsafe
log.Printf("User: %s, password: %s", username, password)

// Safe
log.Printf("User logged in: %s", username)
```

## Email Verification

### Verification Flow

1. Generate verification token upon registration
2. Send verification email
3. User clicks link to verify
4. Account activated after successful verification

### Token Security

- Uses secure random number generator
- Token length: 32 bytes
- Validity: 24 hours
- One-time use

```go
func generateVerificationToken() (string, error) {
    b := make([]byte, 32)
    _, err := rand.Read(b)
    if err != nil {
        return "", err
    }
    return base64.URLEncoding.EncodeToString(b), nil
}
```

## Security Recommendations

### Basic Security Checklist

- [ ] Use strong random JWT_SECRET
- [ ] Regularly update dependencies
- [ ] Use parameterized queries
- [ ] Validate all user input
- [ ] Don't log sensitive information
- [ ] Regular data backups

### Generate Secure Keys

```bash
# Generate JWT_SECRET
openssl rand -base64 32
```

### Check Dependency Vulnerabilities

```bash
# Update dependencies
go get -u ./...
go mod tidy

# Check outdated dependencies
go list -u -m all
```

## HTTPS Configuration (Optional)

### Using Let's Encrypt

```bash
# Install certbot
sudo apt-get install certbot python3-certbot-nginx

# Obtain certificate
sudo certbot --nginx -d yourdomain.com
```

### Self-Signed Certificate (Testing Only)

```bash
openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
  -keyout selfsigned.key -out selfsigned.crt
```

---

Following these security practices will help protect your application and user data.
