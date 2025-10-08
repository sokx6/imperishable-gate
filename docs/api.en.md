# Imperishable Gate API Documentation | Complete RESTful API Reference

**[📖 简体中文](api.md) | [📘 English](api.en.md)**

> 🔌 *"The API Gateway to Hakugyokurou Link Management System"*

## Basic Information

- **Base URL**: `/api/v1`
- **Authentication**: JWT Bearer Token (Implemented in Stage 6, except for public routes)
- **Content-Type**: `application/json`
- **Architecture**: RESTful API (Designed in Stage 1-2)

## 📚 Table of Contents

- [Stage 6 | Authentication API](#stage-6--authentication-api)
- [Stage 2-3 | Link Management API](#stage-2-3--link-management-api)
- [Stage 3 | Name (Alias) Management API](#stage-3--name-alias-management-api)
- [Stage 3 | Tag Management API](#stage-3--tag-management-api)
- [Stage 3 | Remark Management API](#stage-3--remark-management-api)
- [Stage 6 | Email Verification API](#stage-6--email-verification-api)
- [Stage 1 | Public API](#stage-1--public-api)

---

## Stage 6 | Authentication API

> 🔐 *"The Netherworld Princess's Remains - Complete User Authentication System"*

### 1. User Registration

**Endpoint**: `POST /api/v1/register`

**Description**: Register a new user account. Email verification is required after registration.

**Authentication**: Not required

**Request Body**:
```json
{
  "username": "string",  // Required, 3-32 characters
  "email": "string",     // Required, valid email address
  "password": "string"   // Required, minimum 6 characters (encrypted with bcrypt, not stored in plain text!)
}
```

**Success Response** (200 OK):
```json
{
  "message": "Registration successful. Please check your email to verify your account."
}
```

**Error Responses**:
- `409 Conflict`: Username already exists
  ```json
  {
    "message": "Username already exists"
  }
  ```
- `409 Conflict`: Email already registered
  ```json
  {
    "message": "Email already registered"
  }
  ```
- `400 Bad Request`: Invalid request data
- `500 Internal Server Error`: Failed to send verification email

**Example**:
```bash
curl -X POST http://localhost:4514/api/v1/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123"
  }'
```

---

### 2. User Login

**Endpoint**: `POST /api/v1/login`

**Description**: User login to obtain access token and refresh token

**Authentication**: Not required

**Request Body**:
```json
{
  "username": "string",  // Required
  "password": "string"   // Required
}
```

**Success Response** (200 OK):
```json
{
  "success": true,
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "message": "Login successful"
}
```

**Error Responses**:
- `401 Unauthorized`: Invalid username or password
  ```json
  {
    "message": "Authentication failed"
  }
  ```
- `403 Forbidden`: Email not verified
  ```json
  {
    "message": "Email not verified"
  }
  ```
- `404 Not Found`: User not found
  ```json
  {
    "message": "User not found"
  }
  ```

**Example**:
```bash
curl -X POST http://localhost:4514/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123"
  }'
```

---

### 3. Refresh Access Token

**Endpoint**: `POST /api/v1/refresh`

**Description**: Use refresh token to obtain a new access token

**Authentication**: Not required

**Request Body**:
```json
{
  "refresh_token": "string"  // Required, valid refresh token
}
```

**Success Response** (200 OK):
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Error Responses**:
- `401 Unauthorized`: Invalid or expired refresh token
- `400 Bad Request`: Invalid request data

**Example**:
```bash
curl -X POST http://localhost:4514/api/v1/refresh \
  -H "Content-Type: application/json" \
  -d '{
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }'
```

---

### 4. User Logout

**Endpoint**: `POST /api/v1/logout`

**Description**: Logout and invalidate the refresh token

**Authentication**: Not required

**Request Body**:
```json
{
  "refresh_token": "string"  // Required
}
```

**Success Response** (200 OK):
```json
{
  "message": "Logged out successfully"
}
```

**Error Responses**:
- `400 Bad Request`: Invalid request data

**Example**:
```bash
curl -X POST http://localhost:4514/api/v1/logout \
  -H "Content-Type: application/json" \
  -d '{
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }'
```

---

### 5. Get Current User Information (Whoami)

**Endpoint**: `GET /api/v1/whoami`

**Description**: Get information of the currently authenticated user

**Authentication**: Required (Bearer Token)

**Request Headers**:
```
Authorization: Bearer <access_token>
```

**Request Body**: None

**Success Response** (200 OK):
```json
{
  "message": "Success",
  "user_info": {
    "user_id": 1,
    "username": "testuser"
  }
}
```

**Error Responses**:
- `401 Unauthorized`: Invalid, expired, or missing token
  ```json
  {
    "message": "Invalid or expired token"
  }
  ```

**Example**:
```bash
curl -X GET http://localhost:4514/api/v1/whoami \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

---

## Link Management API

### 6. Add Link

**Endpoint**: `POST /api/v1/links`

**Description**: Add a new link with optional names, tags, and remark

**Authentication**: Required (Bearer Token)

**Request Headers**:
```
Authorization: Bearer <access_token>
```

**Request Body**:
```json
{
  "link": "string",      // Required, URL to add
  "names": ["string"],   // Optional, list of names
  "tags": ["string"],    // Optional, list of tags
  "remark": "string",    // Optional, remark
  "name": "string"       // Optional, single name
}
```

**Success Response** (200 OK):
```json
{
  "message": "Added successfully"
}
```

**Error Responses**:
- `400 Bad Request`: Invalid URL format
  ```json
  {
    "message": "Invalid URL format"
  }
  ```
- `409 Conflict`: Link already exists
  ```json
  {
    "message": "Link already exists"
  }
  ```
- `409 Conflict`: Name already exists
  ```json
  {
    "message": "Name already exists"
  }
  ```
- `401 Unauthorized`: Not authenticated or invalid token

**Example**:
```bash
curl -X POST http://localhost:4514/api/v1/links \
  -H "Authorization: Bearer <your_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "link": "https://example.com",
    "names": ["example"],
    "tags": ["website", "demo"],
    "remark": "Example website"
  }'
```

---

### 7. Get Link List

**Endpoint**: `GET /api/v1/links`

**Description**: Get all links of the current user with pagination

**Authentication**: Required (Bearer Token)

**Request Headers**:
```
Authorization: Bearer <access_token>
```

**Query Parameters**:
- `page`: Page number (Required, starts from 1)
- `page_size`: Items per page (Required)

**Success Response** (200 OK):
```json
{
  "message": "Links retrieved successfully",
  "links": [
    {
      "id": 1,
      "url": "https://example.com",
      "tags": ["website", "demo"],
      "names": ["example"],
      "remark": "Example website",
      "title": "Example Domain",
      "description": "Example domain description",
      "keywords": "example, domain",
      "status_code": 200,
      "watching": false
    }
  ]
}
```

**Error Responses**:
- `400 Bad Request`: Invalid parameters
- `401 Unauthorized`: Not authenticated or invalid token

**Example**:
```bash
curl -X GET "http://localhost:4514/api/v1/links?page=1&page_size=10" \
  -H "Authorization: Bearer <your_token>"
```

---

### 8. Search Links by Keyword

**Endpoint**: `GET /api/v1/links/search`

**Description**: Search links by keyword (searches URL, title, description, keywords, names, tags)

**Authentication**: Required (Bearer Token)

**Request Headers**:
```
Authorization: Bearer <access_token>
```

**Query Parameters**:
- `keyword`: Search keyword (Required)
- `page`: Page number (Required, starts from 1)
- `page_size`: Items per page (Required)

**Success Response** (200 OK):
```json
{
  "message": "Links retrieved successfully",
  "links": [
    {
      "id": 1,
      "url": "https://example.com",
      "tags": ["website"],
      "names": ["example"],
      "remark": "Example website",
      "title": "Example Domain",
      "description": "Example domain description",
      "keywords": "example, domain",
      "status_code": 200,
      "watching": false
    }
  ]
}
```

**Error Responses**:
- `400 Bad Request`: Invalid parameters
- `401 Unauthorized`: Not authenticated or invalid token

**Example**:
```bash
curl -X GET "http://localhost:4514/api/v1/links/search?keyword=example&page=1&page_size=10" \
  -H "Authorization: Bearer <your_token>"
```

---

### 9. Get Link by Name

**Endpoint**: `GET /api/v1/names/:name`

**Description**: Get link details by name

**Authentication**: Required (Bearer Token)

**Request Headers**:
```
Authorization: Bearer <access_token>
```

**Path Parameters**:
- `name`: Name of the link

**Success Response** (200 OK):
```json
{
  "message": "Link retrieved successfully",
  "links": [
    {
      "id": 1,
      "url": "https://example.com",
      "tags": ["website"],
      "names": ["example"],
      "remark": "Example website",
      "title": "Example Domain",
      "description": "Example domain description",
      "keywords": "example, domain",
      "status_code": 200,
      "watching": false
    }
  ]
}
```

**Error Responses**:
- `404 Not Found`: Name not found
  ```json
  {
    "message": "Name not found"
  }
  ```
- `401 Unauthorized`: Not authenticated or invalid token

**Example**:
```bash
curl -X GET http://localhost:4514/api/v1/names/example \
  -H "Authorization: Bearer <your_token>"
```

---

### 10. Get Links by Tag

**Endpoint**: `GET /api/v1/tags/:tag`

**Description**: Get all links with the specified tag

**Authentication**: Required (Bearer Token)

**Request Headers**:
```
Authorization: Bearer <access_token>
```

**Path Parameters**:
- `tag`: Tag name

**Query Parameters**:
- `page`: Page number (Required, starts from 1)
- `page_size`: Items per page (Required)

**Success Response** (200 OK):
```json
{
  "message": "Links retrieved successfully",
  "links": [
    {
      "id": 1,
      "url": "https://example.com",
      "tags": ["website", "demo"],
      "names": ["example"],
      "remark": "Example website",
      "title": "Example Domain",
      "description": "Example domain description",
      "keywords": "example, domain",
      "status_code": 200,
      "watching": false
    }
  ]
}
```

**Error Responses**:
- `404 Not Found`: Tag not found
  ```json
  {
    "message": "Tag not found"
  }
  ```
- `400 Bad Request`: Invalid parameters
- `401 Unauthorized`: Not authenticated or invalid token

**Example**:
```bash
curl -X GET "http://localhost:4514/api/v1/tags/website?page=1&page_size=10" \
  -H "Authorization: Bearer <your_token>"
```

---

### 11. Delete Link (by URL)

**Endpoint**: `DELETE /api/v1/links`

**Description**: Delete link by URL

**Authentication**: Required (Bearer Token)

**Request Headers**:
```
Authorization: Bearer <access_token>
```

**Request Body**:
```json
{
  "url": "string"  // Required, URL of the link to delete
}
```

**Success Response** (200 OK):
```json
{
  "message": "Links deleted successfully"
}
```

**Error Responses**:
- `404 Not Found`: Link not found
  ```json
  {
    "message": "Link not found"
  }
  ```
- `401 Unauthorized`: Not authenticated or invalid token

**Example**:
```bash
curl -X DELETE http://localhost:4514/api/v1/links \
  -H "Authorization: Bearer <your_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://example.com"
  }'
```

---

### 12. Delete Link (by Name)

**Endpoint**: `DELETE /api/v1/links/name/:name`

**Description**: Delete link by its name

**Authentication**: Required (Bearer Token)

**Request Headers**:
```
Authorization: Bearer <access_token>
```

**Path Parameters**:
- `name`: Name of the link

**Success Response** (200 OK):
```json
{
  "message": "Links deleted successfully"
}
```

**Error Responses**:
- `404 Not Found`: Name not found
  ```json
  {
    "message": "Name not found"
  }
  ```
- `401 Unauthorized`: Not authenticated or invalid token

**Example**:
```bash
curl -X DELETE http://localhost:4514/api/v1/links/name/example \
  -H "Authorization: Bearer <your_token>"
```

---

### 13. Watch Link (by URL)

**Endpoint**: `PATCH /api/v1/links/watch`

**Description**: Enable or disable monitoring for link changes by URL

**Authentication**: Required (Bearer Token)

**Request Headers**:
```
Authorization: Bearer <access_token>
```

**Request Body**:
```json
{
  "url": "string",   // Required, link URL
  "watch": true      // Required, true=start monitoring, false=stop monitoring
}
```

**Success Response** (200 OK):
```json
{
  "message": "Link is now being watched"
}
```
or
```json
{
  "message": "Link is no longer being watched"
}
```

**Error Responses**:
- `404 Not Found`: Link not found
- `401 Unauthorized`: Not authenticated or invalid token

**Example**:
```bash
curl -X PATCH http://localhost:4514/api/v1/links/watch \
  -H "Authorization: Bearer <your_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://example.com",
    "watch": true
  }'
```

---

### 14. Watch Link (by Name)

**Endpoint**: `PATCH /api/v1/name/watch`

**Description**: Enable or disable monitoring for link changes by name

**Authentication**: Required (Bearer Token)

**Request Headers**:
```
Authorization: Bearer <access_token>
```

**Request Body**:
```json
{
  "name": "string",  // Required, link name
  "watch": true      // Required, true=start monitoring, false=stop monitoring
}
```

**Success Response** (200 OK):
```json
{
  "message": "Link is now being watched"
}
```
or
```json
{
  "message": "Link is no longer being watched"
}
```

**Error Responses**:
- `404 Not Found`: Name not found
- `401 Unauthorized`: Not authenticated or invalid token

**Example**:
```bash
curl -X PATCH http://localhost:4514/api/v1/name/watch \
  -H "Authorization: Bearer <your_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "example",
    "watch": true
  }'
```

---

## Name Management API

---

### 15. Add Names to Link

**Endpoint**: `POST /api/v1/names`

**Description**: Add one or more names to a specified link

**Authentication**: Required (Bearer Token)

**Request Headers**:
```
Authorization: Bearer <access_token>
```

**Request Body**:
```json
{
  "url": "string",       // Required, link URL
  "names": ["string"]    // Required, list of names to add
}
```

**Success Response** (200 OK):
```json
{
  "message": "Names added successfully"
}
```

**Error Responses**:
- `404 Not Found`: Link not found
  ```json
  {
    "message": "Link not found"
  }
  ```
- `409 Conflict`: Name already exists
  ```json
  {
    "message": "Name already exists"
  }
  ```
- `401 Unauthorized`: Not authenticated or invalid token

**Example**:
```bash
curl -X POST http://localhost:4514/api/v1/names \
  -H "Authorization: Bearer <your_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://example.com",
    "names": ["example", "demo"]
  }'
```

---

### 16. Remove Names from Link

**Endpoint**: `PATCH /api/v1/links/names/remove`

**Description**: Remove one or more names from a specified link

**Authentication**: Required (Bearer Token)

**Request Headers**:
```
Authorization: Bearer <access_token>
```

**Request Body**:
```json
{
  "url": "string",       // Required, link URL
  "names": ["string"]    // Required, list of names to remove
}
```

**Success Response** (200 OK):
```json
{
  "message": "Names deleted successfully"
}
```

**Error Responses**:
- `404 Not Found`: Link not found
- `404 Not Found`: Name not found
- `401 Unauthorized`: Not authenticated or invalid token

**Example**:
```bash
curl -X PATCH http://localhost:4514/api/v1/links/names/remove \
  -H "Authorization: Bearer <your_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://example.com",
    "names": ["example"]
  }'
```

---

## Tag Management API

---

### 17. Add Tags to Link (by URL)

**Endpoint**: `POST /api/v1/tags`

**Description**: Add one or more tags to a link by URL

**Authentication**: Required (Bearer Token)

**Request Headers**:
```
Authorization: Bearer <access_token>
```

**Request Body**:
```json
{
  "url": "string",      // Required, link URL
  "tags": ["string"]    // Required, list of tags to add
}
```

**Success Response** (200 OK):
```json
{
  "message": "Tags added successfully"
}
```

**Error Responses**:
- `404 Not Found`: Link not found
  ```json
  {
    "message": "Link not found"
  }
  ```
- `401 Unauthorized`: Not authenticated or invalid token

**Example**:
```bash
curl -X POST http://localhost:4514/api/v1/tags \
  -H "Authorization: Bearer <your_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://example.com",
    "tags": ["website", "demo"]
  }'
```

---

### 18. Add Tags to Link (by Name)

**Endpoint**: `POST /api/v1/name/:name/tags`

**Description**: Add tags to a link by its name

**Authentication**: Required (Bearer Token)

**Request Headers**:
```
Authorization: Bearer <access_token>
```

**Path Parameters**:
- `name`: Name of the link

**Request Body**:
```json
{
  "tags": ["string"]    // Required, list of tags to add
}
```

**Success Response** (200 OK):
```json
{
  "message": "Tags added successfully"
}
```

**Error Responses**:
- `404 Not Found`: Name not found
  ```json
  {
    "message": "Name not found"
  }
  ```
- `401 Unauthorized`: Not authenticated or invalid token

**Example**:
```bash
curl -X POST http://localhost:4514/api/v1/name/example/tags \
  -H "Authorization: Bearer <your_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "tags": ["important", "work"]
  }'
```

---

### 19. Remove Tags from Link (by URL)

**Endpoint**: `PATCH /api/v1/links/by-url/tags/remove`

**Description**: Remove one or more tags from a link by URL

**Authentication**: Required (Bearer Token)

**Request Headers**:
```
Authorization: Bearer <access_token>
```

**Request Body**:
```json
{
  "url": "string",      // Required, link URL
  "tags": ["string"]    // Required, list of tags to remove
}
```

**Success Response** (200 OK):
```json
{
  "message": "Tags deleted successfully"
}
```

**Error Responses**:
- `404 Not Found`: Link not found
- `404 Not Found`: Tag not found
- `401 Unauthorized`: Not authenticated or invalid token

**Example**:
```bash
curl -X PATCH http://localhost:4514/api/v1/links/by-url/tags/remove \
  -H "Authorization: Bearer <your_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://example.com",
    "tags": ["demo"]
  }'
```

---

### 20. Remove Tags from Link (by Name)

**Endpoint**: `PATCH /api/v1/:name/tags/remove`

**Description**: Remove tags from a link by its name

**Authentication**: Required (Bearer Token)

**Request Headers**:
```
Authorization: Bearer <access_token>
```

**Path Parameters**:
- `name`: Name of the link

**Request Body**:
```json
{
  "tags": ["string"]    // Required, list of tags to remove
}
```

**Success Response** (200 OK):
```json
{
  "message": "Tags deleted successfully"
}
```

**Error Responses**:
- `404 Not Found`: Name not found
- `404 Not Found`: Tag not found
- `401 Unauthorized`: Not authenticated or invalid token

**Example**:
```bash
curl -X PATCH http://localhost:4514/api/v1/example/tags/remove \
  -H "Authorization: Bearer <your_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "tags": ["work"]
  }'
```

---

## Remark Management API

---

### 21. Add Remark to Link (by URL)

**Endpoint**: `POST /api/v1/remarks`

**Description**: Add or update remark for a link by URL

**Authentication**: Required (Bearer Token)

**Request Headers**:
```
Authorization: Bearer <access_token>
```

**Request Body**:
```json
{
  "url": "string",       // Required, link URL
  "remark": "string"     // Required, remark content
}
```

**Success Response** (200 OK):
```json
{
  "message": "Remark added successfully"
}
```

**Error Responses**:
- `404 Not Found`: Link not found
  ```json
  {
    "message": "Link not found"
  }
  ```
- `409 Conflict`: Remark already exists
  ```json
  {
    "message": "Remark already exists"
  }
  ```
- `401 Unauthorized`: Not authenticated or invalid token

**Example**:
```bash
curl -X POST http://localhost:4514/api/v1/remarks \
  -H "Authorization: Bearer <your_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://example.com",
    "remark": "This is an example website"
  }'
```

---

### 22. Add Remark to Link (by Name)

**Endpoint**: `POST /api/v1/name/:name/remark`

**Description**: Add or update remark for a link by its name

**Authentication**: Required (Bearer Token)

**Request Headers**:
```
Authorization: Bearer <access_token>
```

**Path Parameters**:
- `name`: Name of the link

**Request Body**:
```json
{
  "remark": "string"     // Required, remark content
}
```

**Success Response** (200 OK):
```json
{
  "message": "Remark added successfully"
}
```

**Error Responses**:
- `404 Not Found`: Name not found
  ```json
  {
    "message": "Name not found"
  }
  ```
- `409 Conflict`: Remark already exists
- `401 Unauthorized`: Not authenticated or invalid token

**Example**:
```bash
curl -X POST http://localhost:4514/api/v1/name/example/remark \
  -H "Authorization: Bearer <your_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "remark": "Important example site"
  }'
```

---

## Email Verification API

---

### 23. Verify Email and Complete Registration

**Endpoint**: `POST /api/v1/verify-email`

**Description**: Complete email verification and account activation using email and verification code

**Authentication**: Not required

**Request Body**:
```json
{
  "email": "string",  // Required, email address
  "code": "string"    // Required, 6-digit verification code
}
```

**Success Response** (200 OK):
```json
{
  "message": "Email verified successfully!"
}
```

**Error Responses**:
- `400 Bad Request`: Email or code is empty
  ```json
  {
    "message": "Email or code cannot be empty"
  }
  ```
- `400 Bad Request`: Email not registered
  ```json
  {
    "message": "Email not registered"
  }
  ```
- `400 Bad Request`: Email already verified
  ```json
  {
    "message": "Email is already verified"
  }
  ```
- `400 Bad Request`: Invalid or already used verification code
  ```json
  {
    "message": "Invalid or already used verification code"
  }
  ```
- `400 Bad Request`: Verification code expired
  ```json
  {
    "message": "Verification code has expired, please request a new one"
  }
  ```
- `429 Too Many Requests`: Too many verification attempts
  ```json
  {
    "message": "Too many verification attempts. Please request a new verification code"
  }
  ```

**Example**:
```bash
curl -X POST http://localhost:4514/api/v1/verify-email \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "code": "123456"
  }'
```

---

### 24. Resend Verification Email

**Endpoint**: `POST /api/v1/resend-verification`

**Description**: Resend email verification code

**Authentication**: Not required

**Request Body**:
```json
{
  "email": "string"  // Required, email address
}
```

**Success Response** (200 OK):
```json
{
  "message": "Verification email resent successfully!"
}
```

**Error Responses**:
- `400 Bad Request`: Email is empty
  ```json
  {
    "message": "Email cannot be empty"
  }
  ```
- `400 Bad Request`: Email not registered
  ```json
  {
    "message": "Email not registered"
  }
  ```
- `400 Bad Request`: Email already verified
  ```json
  {
    "message": "Email is already verified"
  }
  ```
- `429 Too Many Requests`: Verification code still valid
  ```json
  {
    "message": "Verification code is still valid. Please wait until it expires (15 minutes) before requesting a new one."
  }
  ```
- `429 Too Many Requests`: Request too frequent
  ```json
  {
    "message": "Please wait at least 2 minutes before requesting a new verification code"
  }
  ```

**Example**:
```bash
curl -X POST http://localhost:4514/api/v1/resend-verification \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com"
  }'
```

---

### 25. Request Password Reset Email (by Email)

**Endpoint**: `PATCH /api/v1/email/password/request`

**Description**: Request to send password reset verification code by email

**Authentication**: Not required

**Request Body**:
```json
{
  "email": "string"  // Required, valid email address
}
```

**Success Response** (200 OK):
```json
{
  "message": "If the email is registered, a reset password email has been sent."
}
```

**Error Responses**:
- `400 Bad Request`: Invalid email format
- `500 Internal Server Error`: Failed to send email

**Example**:
```bash
curl -X PATCH http://localhost:4514/api/v1/email/password/request \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com"
  }'
```

---

### 26. Request Password Reset Email (by Username)

**Endpoint**: `PATCH /api/v1/username/password/request`

**Description**: Request to send password reset verification code to associated email by username

**Authentication**: Not required

**Request Body**:
```json
{
  "username": "string"  // Required, username
}
```

**Success Response** (200 OK):
```json
{
  "message": "If the email is registered, a reset password email has been sent."
}
```

**Error Responses**:
- `400 Bad Request`: Username is empty
  ```json
  {
    "message": "Username cannot be empty"
  }
  ```
- `400 Bad Request`: Username not registered
  ```json
  {
    "message": "Username not registered"
  }
  ```
- `500 Internal Server Error`: Failed to send email

**Example**:
```bash
curl -X PATCH http://localhost:4514/api/v1/username/password/request \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser"
  }'
```

---

### 27. Verify Email and Reset Password (by Email)

**Endpoint**: `PATCH /api/v1/email/password`

**Description**: Reset password using email and verification code

**Authentication**: Not required

**Request Body**:
```json
{
  "email": "string",         // Required, valid email address
  "code": "string",          // Required, 6-digit verification code
  "new_password": "string"   // Required, new password (8-64 characters)
}
```

**Success Response** (200 OK):
```json
{
  "message": "Password has been reset successfully. You can now log in with your new password."
}
```

**Error Responses**:
- `400 Bad Request`: Email or code is empty
  ```json
  {
    "message": "Email or code cannot be empty"
  }
  ```
- `400 Bad Request`: Invalid or already used verification code
  ```json
  {
    "message": "Invalid or already used verification code"
  }
  ```
- `400 Bad Request`: Verification code expired
  ```json
  {
    "message": "Verification code has expired, please request a new one"
  }
  ```
- `429 Too Many Requests`: Too many verification attempts

**Example**:
```bash
curl -X PATCH http://localhost:4514/api/v1/email/password \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "code": "123456",
    "new_password": "newpassword123"
  }'
```

---

### 28. Verify Email and Reset Password (by Username)

**Endpoint**: `PATCH /api/v1/username/password`

**Description**: Reset password using username and verification code

**Authentication**: Not required

**Request Body**:
```json
{
  "username": "string",      // Required, username
  "code": "string",          // Required, 6-digit verification code
  "new_password": "string"   // Required, new password (8-64 characters)
}
```

**Success Response** (200 OK):
```json
{
  "message": "Password has been reset successfully. You can now log in with your new password."
}
```

**Error Responses**:
- `400 Bad Request`: Username or code is empty
  ```json
  {
    "message": "Username or code cannot be empty"
  }
  ```
- `400 Bad Request`: Invalid or already used verification code
- `400 Bad Request`: Verification code expired
- `429 Too Many Requests`: Too many verification attempts

**Example**:
```bash
curl -X PATCH http://localhost:4514/api/v1/username/password \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "code": "123456",
    "new_password": "newpassword123"
  }'
```

---

## Public API

### 29. Ping

**Endpoint**: `POST /api/v1/ping`

**Description**: Test server connection

**Authentication**: Not required

**Request Body**:
```json
{
  "action": "ping",    // Must be "ping"
  "message": "string"  // Optional, client message
}
```

**Success Response** (200 OK):
```json
{
  "message": "pong"
}
```

**Example**:
```bash
curl -X POST http://localhost:4514/api/v1/ping \
  -H "Content-Type: application/json" \
  -d '{
    "action": "ping",
    "message": "Hello server"
  }'
```

---

## Data Models

### Link

```json
{
  "id": 1,                    // Link ID
  "url": "string",            // Link URL
  "tags": ["string"],         // List of tags
  "names": ["string"],        // List of names
  "remark": "string",         // Remark
  "title": "string",          // Page title
  "description": "string",    // Page description
  "keywords": "string",       // Page keywords
  "status_code": 200,         // HTTP status code
  "watching": false           // Whether being monitored
}
```

---

## Error Code Reference

| HTTP Status Code | Description |
|-----------------|-------------|
| 200 | Request successful |
| 400 | Invalid request parameters or format |
| 401 | Not authenticated or invalid/expired token |
| 403 | Forbidden (e.g., email not verified) |
| 404 | Resource not found |
| 409 | Resource conflict (e.g., already exists) |
| 429 | Too many requests |
| 500 | Internal server error |

---

## Authentication Guide

### Obtaining Access Token

1. Register account: `POST /api/v1/register`
2. Verify email: `POST /api/v1/verify-email`
3. Login to get tokens: `POST /api/v1/login`

### Using Access Token

Add to request headers for authenticated APIs:
```
Authorization: Bearer <access_token>
```

### Refreshing Token

When access token expires, use refresh token to get a new access token:
```bash
POST /api/v1/refresh
{
  "refresh_token": "<your_refresh_token>"
}
```

---

## Important Notes

1. **Pagination Parameters**: All paginated APIs require `page` and `page_size` parameters
2. **Token Expiration**: Access tokens have a shorter validity period, refresh tokens have a longer validity period
3. **Verification Code Validity**: Email verification codes are valid for 15 minutes
4. **Resend Limit**: Wait at least 2 minutes before resending verification code
5. **URL Format**: Valid URL format is required when adding links
6. **Name Uniqueness**: Each name must be unique per user
7. **Monitoring Feature**: When monitoring is enabled, the system periodically checks link changes and notifies users

---

**Documentation Generated**: October 8, 2025

**API Version**: v1
