# 📧 邮箱验证功能实现 - 客户端部分

## ✅ 已完成的客户端实现

### 1. 注册流程增强

**文件：** `internal/client/service/register/handler.go`

**流程：**
```
1. 用户输入注册信息（用户名、邮箱、密码）
2. 发送注册请求到服务器
3. 服务器创建账号并发送验证码邮件
4. 客户端提示用户输入验证码
5. 支持以下功能：
   - 输入6位验证码进行验证
   - 输入 "resend" 重新发送验证邮件
   - 最多3次验证尝试
6. 验证成功后完成注册
```

**特性：**
- ✅ 一站式注册体验（无需额外命令）
- ✅ 支持重新发送验证码
- ✅ 友好的错误提示
- ✅ 多次尝试机会（3次）
- ✅ 清晰的进度提示

### 2. 验证码输入功能

**文件：** `internal/client/service/register/input.go`

**新增函数：**
```go
func ReadVerificationCode(attempt, maxAttempts int) (string, error)
```

**功能：**
- 读取6位验证码
- 支持输入 "resend" 触发重发
- 显示当前尝试次数
- 输入验证

### 3. 邮箱验证服务

**文件：** `internal/client/service/verify_email.go`

**新增函数：**
```go
// 验证邮箱
func VerifyEmail(addr, email, code string) error

// 重新发送验证码
func ResendVerificationEmail(addr, email string) error
```

### 4. 请求类型定义

**文件：** `internal/types/request/email_verification_request.go`

```go
type EmailVerificationRequest struct {
    Email string `json:"email"`
    Code  string `json:"code"`
}

type ResendVerificationRequest struct {
    Email string `json:"email"`
}
```

---

## 📋 需要实现的服务端部分

### 1. 数据库模型修改

**修改 User 模型：**
```go
type User struct {
    ID              uint      `gorm:"primarykey"`
    Username        string    `gorm:"not null;uniqueIndex;size:40"`
    Password        string    `gorm:"not null"`
    Email           string    `gorm:"not null;uniqueIndex"`
    EmailVerified   bool      `gorm:"default:false"`           // 新增
    EmailVerifiedAt *time.Time                                 // 新增（可选）
    // ... 其他字段
}
```

**创建 EmailVerification 模型：**
```go
type EmailVerification struct {
    ID        uint      `gorm:"primarykey"`
    UserID    uint      `gorm:"not null;index"`
    Email     string    `gorm:"not null"`
    Code      string    `gorm:"not null;size:6"`
    Token     string    `gorm:"not null;uniqueIndex"`
    ExpiresAt time.Time `gorm:"not null"`
    Used      bool      `gorm:"default:false"`
    CreatedAt time.Time
    User      User      `gorm:"foreignKey:UserID"`
}
```

### 2. 注册处理器修改

**文件：** `internal/server/handlers/register_user_handler.go`

需要在注册成功后发送验证邮件：
```go
func RegisterUserHandler(c echo.Context) error {
    // ... 原有注册逻辑
    
    // 注册成功后发送验证邮件
    _, err := service.SendVerificationEmail(user.ID, user.Email)
    if err != nil {
        // 记录日志但不阻断注册
        log.Error("Failed to send verification email:", err)
    }
    
    return c.JSON(http.StatusOK, response.Response{
        Message: "注册成功，请查收邮箱验证码",
    })
}
```

### 3. 邮箱验证处理器（新增）

**文件：** `internal/server/handlers/verify_email_handler.go`

```go
func VerifyEmailHandler(c echo.Context) error {
    var req request.EmailVerificationRequest
    if err := c.Bind(&req); err != nil {
        return response.InvalidRequestResponse
    }
    
    err := service.VerifyEmail(req.Email, req.Code)
    if err != nil {
        // 错误处理
    }
    
    return c.JSON(http.StatusOK, response.Response{
        Message: "邮箱验证成功",
    })
}

func ResendVerificationEmailHandler(c echo.Context) error {
    // 实现重发逻辑
}
```

### 4. 邮箱验证服务（新增）

**文件：** `internal/server/service/send_verification_email.go`
- 生成6位验证码
- 保存到数据库（15分钟有效期）
- 发送邮件

**文件：** `internal/server/service/verify_email.go`
- 验证码校验
- 过期检查
- 更新用户验证状态

### 5. 登录验证增强

**文件：** `internal/server/handlers/login_handler.go`

添加邮箱验证检查：
```go
if !user.EmailVerified {
    return c.JSON(http.StatusForbidden, response.Response{
        Message: "请先验证邮箱后再登录",
    })
}
```

### 6. 路由注册

**文件：** `internal/server/routes/router.go`

```go
// 公共路由（不需要认证）
v1.POST("/register", handlers.RegisterUserHandler)
v1.POST("/verify-email", handlers.VerifyEmailHandler)              // 新增
v1.POST("/resend-verification", handlers.ResendVerificationEmailHandler) // 新增
v1.POST("/login", handlers.LoginHandler)
```

### 7. 邮件工具（如果还没有）

**文件：** `internal/server/utils/email.go`

实现发送邮件的功能（使用 SMTP）

---

## 🎯 用户使用体验

### 注册流程示例：

```bash
$ gate register

Please enter your username (3-32 characters): alice
Please enter your email: alice@example.com
Please enter your password (minimum 6 characters): ******
Please confirm your password: ******

📤 Sending registration request...
✓ Account created successfully!
📧 A verification code has been sent to your email: alice@example.com
💡 The code is valid for 15 minutes.

Please enter the verification code (or 'resend' to get a new code): 123456

🔍 Verifying your email...

✨ Email verification successful!
✓ Registration completed!
🎉 You can now login with your credentials using 'gate login' command.
```

### 重发验证码示例：

```bash
Please enter the verification code (or 'resend' to get a new code): resend

📧 Resending verification email...
✓ Verification email has been resent!

Please enter the verification code (or 'resend' to get a new code): 789012

🔍 Verifying your email...

✨ Email verification successful!
```

### 验证失败示例：

```bash
Please enter the verification code (or 'resend' to get a new code): 111111

🔍 Verifying your email...
❌ Verification failed: invalid verification code
You have 2 attempt(s) remaining.
💡 Tip: Enter 'resend' to get a new verification code.

[Attempt 2/3] Please enter the verification code (or 'resend' to get a new code):
```

---

## 📝 下一步：服务端实现清单

- [ ] 修改 User 模型，添加 EmailVerified 字段
- [ ] 创建 EmailVerification 模型
- [ ] 实现 SendVerificationEmail 服务
- [ ] 实现 VerifyEmail 服务
- [ ] 创建 VerifyEmailHandler
- [ ] 创建 ResendVerificationEmailHandler
- [ ] 修改 RegisterUserHandler（发送验证邮件）
- [ ] 修改 LoginHandler（检查邮箱验证状态）
- [ ] 注册新路由
- [ ] 配置邮件服务（SMTP）
- [ ] 测试完整流程

---

## 🔧 配置要求

服务端需要配置邮件服务，环境变量示例：

```env
# SMTP 配置
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=your-email@gmail.com
SMTP_PASSWORD=your-app-password
SMTP_FROM=noreply@imperishable-gate.com
```

---

## ✨ 功能亮点

1. **用户体验优化**
   - 一键注册，无需切换命令
   - 清晰的步骤提示
   - 支持重新发送验证码
   - 多次尝试机会

2. **安全性**
   - 6位随机验证码
   - 15分钟有效期
   - 验证码一次性使用
   - 登录前强制验证邮箱

3. **容错性**
   - 3次验证尝试
   - 友好的错误提示
   - 支持重发验证码
   - 详细的失败原因说明

准备好实现服务端了吗？我可以帮您逐步完成！
