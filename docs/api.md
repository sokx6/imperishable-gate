# Imperishable Gate API 文档 | RESTful API 完整参考

**[简体中文](api.md) | [English](api.en.md)**

> *"通往白玉楼链接管理系统的API之门"*

## 基础信息

- **基础 URL**: `/api/v1`
- **认证方式**: JWT Bearer Token（Stage 6 实现，除公共路由外）
- **Content-Type**: `application/json`
- **架构风格**: RESTful API（Stage 1-2 设计）

## 目录

- [认证相关 API](#认证相关-api)
- [链接管理 API](#链接管理-api)
- [名称管理 API](#名称管理-api)
- [标签管理 API](#标签管理-api)
- [备注管理 API](#备注管理-api)
- [邮箱验证 API](#邮箱验证-api)
- [公共 API](#公共-api)
- [数据模型](#数据模型)
- [错误码说明](#错误码说明)
- [认证说明](#认证说明)
- [注意事项](#注意事项)

---

## 认证相关 API

> *"冥界大小姐的亡骸 - 完整的用户认证系统"*

### 1. 用户注册

**端点**: `POST /api/v1/register`

**描述**: 注册新用户账号，注册后需要验证邮箱

**认证**: 不需要

**请求体**:
```json
{
  "username": "string",  // 必填，3-32字符
  "email": "string",     // 必填，有效的邮箱地址
  "password": "string"   // 必填，最少6字符（会使用bcrypt加密，不会明文存储！）
}
```

**成功响应** (200 OK):
```json
{
  "message": "Registration successful. Please check your email to verify your account."
}
```

**错误响应**:
- `409 Conflict`: 用户名已存在
  ```json
  {
    "message": "Username already exists"
  }
  ```
- `409 Conflict`: 邮箱已注册
  ```json
  {
    "message": "Email already registered"
  }
  ```
- `400 Bad Request`: 请求数据无效
- `500 Internal Server Error`: 发送验证邮件失败

**示例**:
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

### 2. 用户登录

**端点**: `POST /api/v1/login`

**描述**: 用户登录获取访问令牌和刷新令牌

**认证**: 不需要

**请求体**:
```json
{
  "username": "string",  // 必填
  "password": "string"   // 必填
}
```

**成功响应** (200 OK):
```json
{
  "success": true,
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "message": "Login successful"
}
```

**错误响应**:
- `401 Unauthorized`: 用户名或密码错误
  ```json
  {
    "message": "Authentication failed"
  }
  ```
- `403 Forbidden`: 邮箱未验证
  ```json
  {
    "message": "Email not verified"
  }
  ```
- `404 Not Found`: 用户不存在
  ```json
  {
    "message": "User not found"
  }
  ```

**示例**:
```bash
curl -X POST http://localhost:4514/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123"
  }'
```

---

### 3. 刷新访问令牌

**端点**: `POST /api/v1/refresh`

**描述**: 使用刷新令牌获取新的访问令牌

**认证**: 不需要

**请求体**:
```json
{
  "refresh_token": "string"  // 必填，有效的刷新令牌
}
```

**成功响应** (200 OK):
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**错误响应**:
- `401 Unauthorized`: 刷新令牌无效或已过期
- `400 Bad Request`: 请求数据无效

**示例**:
```bash
curl -X POST http://localhost:4514/api/v1/refresh \
  -H "Content-Type: application/json" \
  -d '{
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }'
```

---

### 4. 用户登出

**端点**: `POST /api/v1/logout`

**描述**: 登出并使刷新令牌失效

**认证**: 不需要

**请求体**:
```json
{
  "refresh_token": "string"  // 必填
}
```

**成功响应** (200 OK):
```json
{
  "message": "Logged out successfully"
}
```

**错误响应**:
- `400 Bad Request`: 请求数据无效

**示例**:
```bash
curl -X POST http://localhost:4514/api/v1/logout \
  -H "Content-Type: application/json" \
  -d '{
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }'
```

---

### 5. 获取当前用户信息 (Whoami)

**端点**: `GET /api/v1/whoami`

**描述**: 获取当前认证用户的信息

**认证**: 需要 (Bearer Token)

**请求头**:
```
Authorization: Bearer <access_token>
```

**请求体**: 无

**成功响应** (200 OK):
```json
{
  "message": "Success",
  "user_info": {
    "user_id": 1,
    "username": "testuser"
  }
}
```

**错误响应**:
- `401 Unauthorized`: 令牌无效、过期或未提供
  ```json
  {
    "message": "Invalid or expired token"
  }
  ```

**示例**:
```bash
curl -X GET http://localhost:4514/api/v1/whoami \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

---

## 链接管理 API

### 6. 添加链接

**端点**: `POST /api/v1/links`

**描述**: 添加一个新链接（仅添加URL，如需添加名称、标签、备注请使用对应的API）

**认证**: 需要 (Bearer Token)

**请求头**:
```
Authorization: Bearer <access_token>
```

**请求体**:
```json
{
  "link": "string"      // 必填，要添加的链接URL
}
```

**成功响应** (200 OK):
```json
{
  "message": "Added successfully"
}
```

**错误响应**:
- `400 Bad Request`: URL格式无效
  ```json
  {
    "message": "Invalid URL format"
  }
  ```
- `409 Conflict`: 链接已存在
  ```json
  {
    "message": "Link already exists"
  }
  ```
- `401 Unauthorized`: 未认证或令牌无效

**示例**:
```bash
curl -X POST http://localhost:4514/api/v1/links \
  -H "Authorization: Bearer <your_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "link": "https://example.com"
  }'
```

---

### 7. 获取链接列表

**端点**: `GET /api/v1/links`

**描述**: 分页获取当前用户的所有链接

**认证**: 需要 (Bearer Token)

**请求头**:
```
Authorization: Bearer <access_token>
```

**查询参数**:
- `page`: 页码 (必填，从1开始)
- `page_size`: 每页数量 (必填)

**成功响应** (200 OK):
```json
{
  "message": "Links retrieved successfully",
  "links": [
    {
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

**错误响应**:
- `400 Bad Request`: 参数无效
- `401 Unauthorized`: 未认证或令牌无效

**示例**:
```bash
curl -X GET "http://localhost:4514/api/v1/links?page=1&page_size=10" \
  -H "Authorization: Bearer <your_token>"
```

---

### 8. 按关键词搜索链接

**端点**: `GET /api/v1/links/search`

**描述**: 根据关键词搜索链接（搜索URL、标题、描述、关键词、名称、标签）

**认证**: 需要 (Bearer Token)

**请求头**:
```
Authorization: Bearer <access_token>
```

**查询参数**:
- `keyword`: 搜索关键词 (必填)
- `page`: 页码 (必填，从1开始)
- `page_size`: 每页数量 (必填)

**成功响应** (200 OK):
```json
{
  "message": "Links retrieved successfully",
  "links": [
    {
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

**错误响应**:
- `400 Bad Request`: 参数无效
- `401 Unauthorized`: 未认证或令牌无效

**示例**:
```bash
curl -X GET "http://localhost:4514/api/v1/links/search?keyword=example&page=1&page_size=10" \
  -H "Authorization: Bearer <your_token>"
```

---

### 9. 按名称获取链接

**端点**: `GET /api/v1/names/:name`

**描述**: 通过名称获取对应的链接详情

**认证**: 需要 (Bearer Token)

**请求头**:
```
Authorization: Bearer <access_token>
```

**路径参数**:
- `name`: 链接的名称

**成功响应** (200 OK):
```json
{
  "message": "Link retrieved successfully",
  "links": [
    {
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

**错误响应**:
- `404 Not Found`: 名称不存在
  ```json
  {
    "message": "Name not found"
  }
  ```
- `401 Unauthorized`: 未认证或令牌无效

**示例**:
```bash
curl -X GET http://localhost:4514/api/v1/names/example \
  -H "Authorization: Bearer <your_token>"
```

---

### 10. 按标签获取链接列表

**端点**: `GET /api/v1/tags/:tag`

**描述**: 获取指定标签下的所有链接

**认证**: 需要 (Bearer Token)

**请求头**:
```
Authorization: Bearer <access_token>
```

**路径参数**:
- `tag`: 标签名称

**查询参数**:
- `page`: 页码 (必填，从1开始)
- `page_size`: 每页数量 (必填)

**成功响应** (200 OK):
```json
{
  "message": "Links retrieved successfully",
  "links": [
    {
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

**错误响应**:
- `404 Not Found`: 标签不存在
  ```json
  {
    "message": "Tag not found"
  }
  ```
- `400 Bad Request`: 参数无效
- `401 Unauthorized`: 未认证或令牌无效

**示例**:
```bash
curl -X GET "http://localhost:4514/api/v1/tags/website?page=1&page_size=10" \
  -H "Authorization: Bearer <your_token>"
```

---

### 11. 删除链接 (按URL)

**端点**: `DELETE /api/v1/links`

**描述**: 根据URL删除一个或多个链接（通过查询参数）

**认证**: 需要 (Bearer Token)

**请求头**:
```
Authorization: Bearer <access_token>
```

**查询参数**:
- `link`: 要删除的链接URL（可以重复多次以删除多个链接）

**成功响应** (200 OK):
```json
{
  "message": "Links deleted successfully"
}
```

**错误响应**:
- `400 Bad Request`: URL格式无效
- `404 Not Found`: 链接不存在
  ```json
  {
    "message": "Link not found"
  }
  ```
- `401 Unauthorized`: 未认证或令牌无效

**示例**:
```bash
# 删除单个链接
curl -X DELETE "http://localhost:4514/api/v1/links?link=https://example.com" \
  -H "Authorization: Bearer <your_token>"

# 删除多个链接
curl -X DELETE "http://localhost:4514/api/v1/links?link=https://example.com&link=https://test.com" \
  -H "Authorization: Bearer <your_token>"
```

---

### 12. 删除链接 (按名称)

**端点**: `DELETE /api/v1/links/name/:name`

**描述**: 根据名称删除对应的链接

**认证**: 需要 (Bearer Token)

**请求头**:
```
Authorization: Bearer <access_token>
```

**路径参数**:
- `name`: 链接的名称

**成功响应** (200 OK):
```json
{
  "message": "Links deleted successfully"
}
```

**错误响应**:
- `404 Not Found`: 名称不存在
  ```json
  {
    "message": "Name not found"
  }
  ```
- `401 Unauthorized`: 未认证或令牌无效

**示例**:
```bash
curl -X DELETE http://localhost:4514/api/v1/links/name/example \
  -H "Authorization: Bearer <your_token>"
```

---

### 13. 监控链接 (按URL)

**端点**: `PATCH /api/v1/links/watch`

**描述**: 设置或取消监控指定URL的链接变化

**认证**: 需要 (Bearer Token)

**请求头**:
```
Authorization: Bearer <access_token>
```

**请求体**:
```json
{
  "url": "string",   // 必填，链接URL
  "watch": true      // 必填，true=开始监控, false=取消监控
}
```

**成功响应** (200 OK):
```json
{
  "message": "Link is now being watched"
}
```
或
```json
{
  "message": "Link is no longer being watched"
}
```

**错误响应**:
- `404 Not Found`: 链接不存在
- `401 Unauthorized`: 未认证或令牌无效

**示例**:
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

### 14. 监控链接 (按名称)

**端点**: `PATCH /api/v1/name/watch`

**描述**: 通过名称设置或取消监控链接变化

**认证**: 需要 (Bearer Token)

**请求头**:
```
Authorization: Bearer <access_token>
```

**请求体**:
```json
{
  "name": "string",  // 必填，链接名称
  "watch": true      // 必填，true=开始监控, false=取消监控
}
```

**成功响应** (200 OK):
```json
{
  "message": "Link is now being watched"
}
```
或
```json
{
  "message": "Link is no longer being watched"
}
```

**错误响应**:
- `404 Not Found`: 名称不存在
- `401 Unauthorized`: 未认证或令牌无效

**示例**:
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

## 名称管理 API

---

### 15. 为链接添加名称

**端点**: `POST /api/v1/names`

**描述**: 为指定的链接添加一个或多个名称

**认证**: 需要 (Bearer Token)

**请求头**:
```
Authorization: Bearer <access_token>
```

**请求体**:
```json
{
  "link": "string",      // 必填，链接URL
  "names": ["string"]    // 必填，要添加的名称列表
}
```

**成功响应** (200 OK):
```json
{
  "message": "Names added successfully"
}
```

**错误响应**:
- `400 Bad Request`: 链接URL格式无效
- `409 Conflict`: 名称已存在
  ```json
  {
    "message": "Name already exists"
  }
  ```
- `401 Unauthorized`: 未认证或令牌无效

**示例**:
```bash
curl -X POST http://localhost:4514/api/v1/names \
  -H "Authorization: Bearer <your_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "link": "https://example.com",
    "names": ["example", "demo"]
  }'
```

---

### 16. 删除链接的名称

**端点**: `PATCH /api/v1/links/names/remove`

**描述**: 删除指定链接的一个或多个名称

**认证**: 需要 (Bearer Token)

**请求头**:
```
Authorization: Bearer <access_token>
```

**请求体**:
```json
{
  "url": "string",       // 必填，链接URL
  "names": ["string"]    // 必填，要删除的名称列表
}
```

**成功响应** (200 OK):
```json
{
  "message": "Names deleted successfully"
}
```

**错误响应**:
- `404 Not Found`: 链接不存在
- `404 Not Found`: 名称不存在
- `401 Unauthorized`: 未认证或令牌无效

**示例**:
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

## 标签管理 API

---

### 17. 为链接添加标签 (按URL)

**端点**: `POST /api/v1/tags`

**描述**: 为指定URL的链接添加一个或多个标签

**认证**: 需要 (Bearer Token)

**请求头**:
```
Authorization: Bearer <access_token>
```

**请求体**:
```json
{
  "link": "string",      // 必填，链接URL
  "tags": ["string"]     // 必填，要添加的标签列表
}
```

**成功响应** (200 OK):
```json
{
  "message": "Tags added successfully"
}
```

**错误响应**:
- `400 Bad Request`: 链接URL格式无效
- `401 Unauthorized`: 未认证或令牌无效

**示例**:
```bash
curl -X POST http://localhost:4514/api/v1/tags \
  -H "Authorization: Bearer <your_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "link": "https://example.com",
    "tags": ["website", "demo"]
  }'
```

---

### 18. 为链接添加标签 (按名称)

**端点**: `POST /api/v1/name/:name/tags`

**描述**: 通过链接的名称为其添加标签

**认证**: 需要 (Bearer Token)

**请求头**:
```
Authorization: Bearer <access_token>
```

**路径参数**:
- `name`: 链接的名称

**请求体**:
```json
{
  "tags": ["string"]    // 必填，要添加的标签列表
}
```

**成功响应** (200 OK):
```json
{
  "message": "Tags added successfully"
}
```

**错误响应**:
- `404 Not Found`: 名称不存在
  ```json
  {
    "message": "Name not found"
  }
  ```
- `401 Unauthorized`: 未认证或令牌无效

**示例**:
```bash
curl -X POST http://localhost:4514/api/v1/name/example/tags \
  -H "Authorization: Bearer <your_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "tags": ["important", "work"]
  }'
```

---

### 19. 删除链接的标签 (按URL)

**端点**: `PATCH /api/v1/links/by-url/tags/remove`

**描述**: 删除指定URL链接的一个或多个标签

**认证**: 需要 (Bearer Token)

**请求头**:
```
Authorization: Bearer <access_token>
```

**请求体**:
```json
{
  "url": "string",      // 必填，链接URL
  "tags": ["string"]    // 必填，要删除的标签列表
}
```

**成功响应** (200 OK):
```json
{
  "message": "Tags deleted successfully"
}
```

**错误响应**:
- `404 Not Found`: 链接不存在
- `404 Not Found`: 标签不存在
- `401 Unauthorized`: 未认证或令牌无效

**示例**:
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

### 20. 删除链接的标签 (按名称)

**端点**: `PATCH /api/v1/:name/tags/remove`

**描述**: 通过链接的名称删除其标签

**认证**: 需要 (Bearer Token)

**请求头**:
```
Authorization: Bearer <access_token>
```

**路径参数**:
- `name`: 链接的名称

**请求体**:
```json
{
  "tags": ["string"]    // 必填，要删除的标签列表
}
```

**成功响应** (200 OK):
```json
{
  "message": "Tags deleted successfully"
}
```

**错误响应**:
- `404 Not Found`: 名称不存在
- `404 Not Found`: 标签不存在
- `401 Unauthorized`: 未认证或令牌无效

**示例**:
```bash
curl -X PATCH http://localhost:4514/api/v1/example/tags/remove \
  -H "Authorization: Bearer <your_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "tags": ["work"]
  }'
```

---

## 备注管理 API

---

### 21. 为链接添加备注 (按URL)

**端点**: `POST /api/v1/remarks`

**描述**: 为指定URL的链接添加或更新备注

**认证**: 需要 (Bearer Token)

**请求头**:
```
Authorization: Bearer <access_token>
```

**请求体**:
```json
{
  "link": "string",      // 必填，链接URL
  "remark": "string"     // 必填，备注内容
}
```

**成功响应** (200 OK):
```json
{
  "message": "Remark added successfully"
}
```

**错误响应**:
- `400 Bad Request`: 链接URL格式无效或备注为空
- `401 Unauthorized`: 未认证或令牌无效

**示例**:
```bash
curl -X POST http://localhost:4514/api/v1/remarks \
  -H "Authorization: Bearer <your_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "link": "https://example.com",
    "remark": "This is an example website"
  }'
```

---

### 22. 为链接添加备注 (按名称)

**端点**: `POST /api/v1/name/:name/remark`

**描述**: 通过链接的名称为其添加或更新备注

**认证**: 需要 (Bearer Token)

**请求头**:
```
Authorization: Bearer <access_token>
```

**路径参数**:
- `name`: 链接的名称

**请求体**:
```json
{
  "remark": "string"     // 必填，备注内容
}
```

**成功响应** (200 OK):
```json
{
  "message": "Remark added successfully"
}
```

**错误响应**:
- `404 Not Found`: 名称不存在
  ```json
  {
    "message": "Name not found"
  }
  ```
- `409 Conflict`: 备注已存在
- `401 Unauthorized`: 未认证或令牌无效

**示例**:
```bash
curl -X POST http://localhost:4514/api/v1/name/example/remark \
  -H "Authorization: Bearer <your_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "remark": "Important example site"
  }'
```

---

## 邮箱验证 API

---

### 23. 验证邮箱并完成注册

**端点**: `POST /api/v1/verify-email`

**描述**: 使用邮箱和验证码完成邮箱验证和账号激活

**认证**: 不需要

**请求体**:
```json
{
  "email": "string",  // 必填，邮箱地址
  "code": "string"    // 必填，6位验证码
}
```

**成功响应** (200 OK):
```json
{
  "message": "Email verified successfully!"
}
```

**错误响应**:
- `400 Bad Request`: 邮箱或验证码为空
  ```json
  {
    "message": "Email or code cannot be empty"
  }
  ```
- `400 Bad Request`: 邮箱未注册
  ```json
  {
    "message": "Email not registered"
  }
  ```
- `400 Bad Request`: 邮箱已验证
  ```json
  {
    "message": "Email is already verified"
  }
  ```
- `400 Bad Request`: 验证码无效或已使用
  ```json
  {
    "message": "Invalid or already used verification code"
  }
  ```
- `400 Bad Request`: 验证码已过期
  ```json
  {
    "message": "Verification code has expired, please request a new one"
  }
  ```
- `429 Too Many Requests`: 验证尝试次数过多
  ```json
  {
    "message": "Too many verification attempts. Please request a new verification code"
  }
  ```

**示例**:
```bash
curl -X POST http://localhost:4514/api/v1/verify-email \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "code": "123456"
  }'
```

---

### 24. 重新发送验证邮件

**端点**: `POST /api/v1/resend-verification`

**描述**: 重新发送邮箱验证码

**认证**: 不需要

**请求体**:
```json
{
  "email": "string"  // 必填，邮箱地址
}
```

**成功响应** (200 OK):
```json
{
  "message": "Verification email resent successfully!"
}
```

**错误响应**:
- `400 Bad Request`: 邮箱为空
  ```json
  {
    "message": "Email cannot be empty"
  }
  ```
- `400 Bad Request`: 邮箱未注册
  ```json
  {
    "message": "Email not registered"
  }
  ```
- `400 Bad Request`: 邮箱已验证
  ```json
  {
    "message": "Email is already verified"
  }
  ```
- `429 Too Many Requests`: 验证码仍然有效
  ```json
  {
    "message": "Verification code is still valid. Please wait until it expires (15 minutes) before requesting a new one."
  }
  ```
- `429 Too Many Requests`: 请求过于频繁
  ```json
  {
    "message": "Please wait at least 2 minutes before requesting a new verification code"
  }
  ```

**示例**:
```bash
curl -X POST http://localhost:4514/api/v1/resend-verification \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com"
  }'
```

---

### 25. 请求重置密码邮件 (通过邮箱)

**端点**: `PATCH /api/v1/email/password/request`

**描述**: 通过邮箱请求发送重置密码验证码

**认证**: 不需要

**请求体**:
```json
{
  "email": "string"  // 必填，有效的邮箱地址
}
```

**成功响应** (200 OK):
```json
{
  "message": "If the email is registered, a reset password email has been sent."
}
```

**错误响应**:
- `400 Bad Request`: 邮箱格式无效
- `500 Internal Server Error`: 发送邮件失败

**示例**:
```bash
curl -X PATCH http://localhost:4514/api/v1/email/password/request \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com"
  }'
```

---

### 26. 请求重置密码邮件 (通过用户名)

**端点**: `PATCH /api/v1/username/password/request`

**描述**: 通过用户名请求发送重置密码验证码到关联邮箱

**认证**: 不需要

**请求体**:
```json
{
  "username": "string"  // 必填，用户名
}
```

**成功响应** (200 OK):
```json
{
  "message": "If the email is registered, a reset password email has been sent."
}
```

**错误响应**:
- `400 Bad Request`: 用户名为空
  ```json
  {
    "message": "Username cannot be empty"
  }
  ```
- `400 Bad Request`: 用户名未注册
  ```json
  {
    "message": "Username not registered"
  }
  ```
- `500 Internal Server Error`: 发送邮件失败

**示例**:
```bash
curl -X PATCH http://localhost:4514/api/v1/username/password/request \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser"
  }'
```

---

### 27. 验证邮箱并重置密码 (通过邮箱)

**端点**: `PATCH /api/v1/email/password`

**描述**: 使用邮箱和验证码重置密码

**认证**: 不需要

**请求体**:
```json
{
  "email": "string",         // 必填，有效的邮箱地址
  "code": "string",          // 必填，6位验证码
  "new_password": "string"   // 必填，新密码 (8-64字符)
}
```

**成功响应** (200 OK):
```json
{
  "message": "Password has been reset successfully. You can now log in with your new password."
}
```

**错误响应**:
- `400 Bad Request`: 邮箱或验证码为空
  ```json
  {
    "message": "Email or code cannot be empty"
  }
  ```
- `400 Bad Request`: 验证码无效或已使用
  ```json
  {
    "message": "Invalid or already used verification code"
  }
  ```
- `400 Bad Request`: 验证码已过期
  ```json
  {
    "message": "Verification code has expired, please request a new one"
  }
  ```
- `429 Too Many Requests`: 验证尝试次数过多

**示例**:
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

### 28. 验证邮箱并重置密码 (通过用户名)

**端点**: `PATCH /api/v1/username/password`

**描述**: 使用用户名和验证码重置密码

**认证**: 不需要

**请求体**:
```json
{
  "username": "string",      // 必填，用户名
  "code": "string",          // 必填，6位验证码
  "new_password": "string"   // 必填，新密码 (8-64字符)
}
```

**成功响应** (200 OK):
```json
{
  "message": "Password has been reset successfully. You can now log in with your new password."
}
```

**错误响应**:
- `400 Bad Request`: 用户名或验证码为空
  ```json
  {
    "message": "Username or code cannot be empty"
  }
  ```
- `400 Bad Request`: 验证码无效或已使用
- `400 Bad Request`: 验证码已过期
- `429 Too Many Requests`: 验证尝试次数过多

**示例**:
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

## 公共 API

### 29. Ping

**端点**: `POST /api/v1/ping`

**描述**: 测试服务器连接

**认证**: 不需要

**请求体**:
```json
{
  "action": "ping",    // 必须是 "ping"
  "message": "string"  // 可选，客户端消息
}
```

**成功响应** (200 OK):
```json
{
  "message": "pong"
}
```

**示例**:
```bash
curl -X POST http://localhost:4514/api/v1/ping \
  -H "Content-Type: application/json" \
  -d '{
    "action": "ping",
    "message": "Hello server"
  }'
```

---

## 数据模型

### Link (链接)

```json
{
  "url": "string",            // 链接URL
  "tags": ["string"],         // 标签列表
  "names": ["string"],        // 名称列表
  "remark": "string",         // 备注
  "title": "string",          // 网页标题
  "description": "string",    // 网页描述
  "keywords": "string",       // 网页关键词
  "status_code": 200,         // HTTP状态码
  "watching": false           // 是否监控中
}
```

---

## 错误码说明

| HTTP状态码 | 说明 |
|-----------|------|
| 200 | 请求成功 |
| 400 | 请求参数错误或格式无效 |
| 401 | 未认证或令牌无效/过期 |
| 403 | 禁止访问（如邮箱未验证） |
| 404 | 资源不存在 |
| 409 | 资源冲突（如已存在） |
| 429 | 请求过于频繁 |
| 500 | 服务器内部错误 |

---

## 认证说明

### 获取访问令牌

1. 注册账号：`POST /api/v1/register`
2. 验证邮箱：`POST /api/v1/verify-email`
3. 登录获取令牌：`POST /api/v1/login`

### 使用访问令牌

在需要认证的API请求头中添加：
```
Authorization: Bearer <access_token>
```

### 刷新令牌

当访问令牌过期时，使用刷新令牌获取新的访问令牌：
```bash
POST /api/v1/refresh
{
  "refresh_token": "<your_refresh_token>"
}
```

---

## 注意事项

1. **分页参数**：所有分页API都需要 `page` 和 `page_size` 参数
2. **令牌过期**：访问令牌有效期较短，刷新令牌有效期较长
3. **验证码有效期**：邮箱验证码有效期为15分钟
4. **重发限制**：验证码重发需要等待至少2分钟
5. **URL格式**：添加链接时需要提供有效的URL格式
6. **名称唯一性**：同一用户下，每个名称必须唯一
7. **监控功能**：开启监控后，系统会定期检查链接变化并通知用户

---

**文档生成时间**: 2025-10-09

**API版本**: v1
