# 安全特性 | Stage 6: 冥界大小姐の亡骸

**[📘 English](security.en.md) | [📖 简体中文](security.md)**

> 🔐 *"想要穿过这扇门？先证明你的身份吧！"*

本文档详细介绍 **Stage 6「冥界大小姐の亡骸」** 实现的安全机制。作为白玉楼的庭师，你在上一世学习的密码学知识终于派上用场了！

## 认证系统概述

本项目实现了完整的 **JWT (JSON Web Token)** 认证系统，包括：
- ✅ 双令牌机制（Access Token + Refresh Token）
- ✅ bcrypt 密码加密（不会明文存密码的吧？）
- ✅ 系统 Keyring 安全存储
- ✅ 自动令牌刷新
- ✅ 邮箱验证

## 🔑 JWT 双令牌机制

### 令牌类型

**Access Token（访问令牌）**
- 有效期：15 分钟
- 用途：API 请求认证
- 存储：系统 keyring

**Refresh Token（刷新令牌）**
- 有效期：7 天
- 用途：刷新 Access Token
- 特性：一次性使用，登出时删除

### 令牌刷新流程

客户端会自动处理令牌刷新：

1. API 请求时令牌过期
2. 自动使用 Refresh Token 获取新令牌
3. 重试原请求
4. 用户无感知

## 密码安全

### Bcrypt 加密

- 使用 bcrypt 算法加密密码
- 自动加盐
- 不可逆加密

### 密码要求

- 最小长度：6 字符
- 建议：包含大小写字母、数字、特殊字符

**示例代码**：
```go
// 密码加密
hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

// 密码验证
err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
```

## 数据隔离

### 用户级隔离

- 所有数据按用户 ID 隔离
- 用户只能访问自己的数据
- 数据库查询自动过滤

**实现**：
```go
// 从 JWT 中提取用户 ID
userID := getUserIDFromJWT(c)

// 查询时自动过滤
db.Where("user_id = ?", userID).Find(&links)
```

## 安全存储

### 客户端令牌存储

令牌存储在系统安全存储中：

| 系统 | 存储位置 |
|------|---------|
| Linux | GNOME Keyring / KWallet |
| macOS | Keychain |
| Windows | Credential Manager |

**使用示例**：
```go
import "github.com/zalando/go-keyring"

// 保存令牌
keyring.Set("gate", "access_token", token)

// 读取令牌
token, _ := keyring.Get("gate", "access_token")

// 删除令牌
keyring.Delete("gate", "access_token")
```

### 配置文件安全

```bash
# 设置 .env 文件权限
chmod 600 .env

# 不要提交到版本控制
echo ".env" >> .gitignore
```

## 数据库安全

### SQL 注入防护

使用 GORM 参数化查询：

```go
// ✅ 安全
db.Where("url = ?", url).First(&link)

// ❌ 危险（不要这样做）
db.Raw("SELECT * FROM links WHERE url = '" + url + "'")
```

### 最小权限原则

```sql
-- 创建专用用户
CREATE USER gateuser WITH PASSWORD 'password';

-- 只授予必要权限
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO gateuser;
```

## 输入验证

### 请求验证

```go
type RegisterRequest struct {
    Username string `json:"username" validate:"required,min=3,max=32"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=6"`
}
```

### URL 验证

```go
// 验证 URL 格式
_, err := url.ParseRequestURI(urlString)
if err != nil {
    return errors.New("invalid URL")
}
```

## 日志安全

不要在日志中记录敏感信息：

```go
// ❌ 不安全
log.Printf("User: %s, password: %s", username, password)

// ✅ 安全
log.Printf("User logged in: %s", username)
```

## 邮箱验证

### 验证流程

1. 注册时生成验证令牌
2. 发送验证邮件
3. 用户点击链接验证
4. 验证成功后激活账户

### 令牌安全

- 使用安全随机数生成器
- 令牌长度：32 字节
- 有效期：24 小时
- 一次性使用

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

## 安全建议

### 基本安全检查清单

- [ ] 使用强随机的 JWT_SECRET
- [ ] 定期更新依赖
- [ ] 使用参数化查询
- [ ] 验证所有用户输入
- [ ] 不在日志中记录敏感信息
- [ ] 定期备份数据

### 生成安全密钥

```bash
# 生成 JWT_SECRET
openssl rand -base64 32
```

### 检查依赖漏洞

```bash
# 更新依赖
go get -u ./...
go mod tidy

# 检查过期依赖
go list -u -m all
```

## HTTPS 配置（可选）

### 使用 Let's Encrypt

```bash
# 安装 certbot
sudo apt-get install certbot python3-certbot-nginx

# 获取证书
sudo certbot --nginx -d yourdomain.com
```

### 自签名证书（仅测试用）

```bash
openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
  -keyout selfsigned.key -out selfsigned.crt
```

---

遵循这些安全实践，可以保护你的应用和用户数据安全。
