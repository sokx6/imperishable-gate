# Gate 客户端命令使用文档

## 目录
- [简介](#简介)
- [安装与配置](#安装与配置)
- [全局参数](#全局参数)
- [命令列表](#命令列表)
  - [register - 注册用户](#register---注册用户)
  - [login - 登录](#login---登录)
  - [ping - 测试连接](#ping---测试连接)
  - [add - 添加链接](#add---添加链接)
  - [list - 列出链接](#list---列出链接)
  - [delete - 删除链接](#delete---删除链接)
  - [watch - 监控链接](#watch---监控链接)
  - [open - 打开链接](#open---打开链接)
- [使用示例](#使用示例)
- [常见问题](#常见问题)

---

## 简介

`gate` 是一个命令行链接管理工具，支持：
- 链接的增删改查
- 链接别名管理
- 标签分类
- 链接变化监控
- 快速打开链接

---

## 安装与配置

### 构建客户端

```bash
cd cmd/gate
go build -o gate
```

### 配置服务器地址

有三种方式配置服务器地址（按优先级排序）：

1. **命令行参数** （最高优先级）
   ```bash
   gate <command> -a localhost:8080
   ```

2. **环境变量**（通过 `.env` 文件）
   在项目根目录创建 `.env` 文件：
   ```bash
   SERVER_ADDR=localhost:8080
   ```

3. **默认值** （最低优先级）
   ```
   127.0.0.1:8080
   ```

---

## 全局参数

所有命令都支持以下全局参数：

| 参数 | 简写 | 说明 | 示例 |
|------|------|------|------|
| `--addr` | `-a` | 服务器地址 | `-a localhost:8080` |
| `--help` | `-h` | 显示帮助信息 | `gate add -h` |
| `--version` | | 显示版本信息 | `gate --version` |

---

## 命令列表

### register - 注册用户

注册新用户账号。

**语法：**
```bash
gate register [--addr <服务器地址>]
```

**示例：**
```bash
# 使用默认或配置的服务器地址注册
gate register

# 指定服务器地址注册
gate register -a localhost:8080
```

**交互流程：**
```
Please enter your username: myuser
Please enter your email: user@example.com
Please enter your password: ********
Confirm your password: ********
```

---

### login - 登录

登录到服务器并将刷新令牌安全存储到系统密钥链。

**语法：**
```bash
gate login [--addr <服务器地址>]
```

**示例：**
```bash
# 使用默认或配置的服务器地址登录
gate login

# 指定服务器地址登录
gate login -a localhost:8080
```

**交互流程：**
```
Please enter your username: myuser
Please enter your password: ********
Refresh token saved to system keyring.
Login successful!
```

**说明：**
- 登录成功后，刷新令牌会存储在系统密钥链中
- 后续命令会自动使用存储的令牌进行身份验证
- 如果访问令牌过期，会自动刷新

---

### ping - 测试连接

测试与服务器的连接。

**语法：**
```bash
gate ping [--message <消息>] [--addr <服务器地址>]
```

**参数：**
| 参数 | 简写 | 默认值 | 说明 |
|------|------|--------|------|
| `--message` | `-m` | "default message" | 要发送的测试消息 |

**示例：**
```bash
# 发送默认消息
gate ping

# 发送自定义消息
gate ping -m "Hello Server"

# 指定服务器地址
gate ping -a localhost:8080 -m "test"
```

---

### add - 添加链接

添加链接、名称、标签或备注。

**语法：**
```bash
gate add --link <URL> [--name <名称>...] [--tag <标签>...] [--remark <备注>]
gate add --name <名称> --tag <标签>... [--remark <备注>]
gate add --link <URL> --remark <备注>
gate add --name <名称> --remark <备注>
```

**参数：**
| 参数 | 简写 | 说明 |
|------|------|------|
| `--link` | `-l` | 链接 URL |
| `--name` | `-n` | 链接的别名（可多个） |
| `--tag` | `-t` | 标签（可多个） |
| `--remark` | `-r` | 备注信息 |

**使用场景：**

1. **添加新链接**
   ```bash
   gate add -l https://github.com
   ```

2. **添加链接并指定名称**
   ```bash
   gate add -l https://github.com -n github -n gh
   ```

3. **添加链接并指定标签**
   ```bash
   gate add -l https://github.com -t dev -t tools
   ```

4. **添加链接、名称和标签**
   ```bash
   gate add -l https://github.com -n github -t dev -t tools
   ```

5. **为现有链接添加名称**
   ```bash
   gate add -l https://github.com -n newname
   ```

6. **为链接添加标签（通过 URL）**
   ```bash
   gate add -l https://github.com -t newtag
   ```

7. **为链接添加标签（通过名称）**
   ```bash
   gate add -n github -t newtag
   ```

8. **为链接添加备注（通过 URL）**
   ```bash
   gate add -l https://github.com -r "代码托管平台"
   ```

9. **为链接添加备注（通过名称）**
   ```bash
   gate add -n github -r "代码托管平台"
   ```

---

### list - 列出链接

查询和列出链接信息。

**语法：**
```bash
gate list [--tag <标签>] [--name <名称>] [--page <页码>] [--page-size <每页数量>]
```

**参数：**
| 参数 | 简写 | 默认值 | 说明 |
|------|------|--------|------|
| `--tag` | `-t` | | 按标签过滤 |
| `--name` | `-n` | | 按名称查询 |
| `--page` | `-p` | 1 | 页码 |
| `--page-size` | `-s` | 20 | 每页显示数量 |

**示例：**

1. **列出所有链接**
   ```bash
   gate list
   ```

2. **按标签查询**
   ```bash
   gate list -t dev
   ```

3. **按名称查询**
   ```bash
   gate list -n github
   ```

4. **分页查询**
   ```bash
   gate list -p 2 -s 10
   ```

5. **按标签查询并分页**
   ```bash
   gate list -t tools -p 1 -s 20
   ```

---

### delete - 删除链接

删除链接、名称或标签。

**语法：**
```bash
gate delete --link <URL>... [--name <名称>...] [--tag <标签>...]
gate delete --name <名称>... [--tag <标签>...]
```

**参数：**
| 参数 | 简写 | 说明 |
|------|------|------|
| `--link` | `-l` | 要删除的链接 URL（可多个） |
| `--name` | `-n` | 要删除的名称（可多个） |
| `--tag` | `-t` | 要删除的标签（可多个） |

**使用场景：**

1. **删除链接**
   ```bash
   gate delete -l https://github.com
   ```

2. **批量删除链接**
   ```bash
   gate delete -l https://github.com -l https://google.com
   ```

3. **通过名称删除链接**
   ```bash
   gate delete -n github
   ```

4. **删除链接的某个名称**
   ```bash
   gate delete -l https://github.com -n oldname
   ```

5. **删除链接的标签（通过 URL）**
   ```bash
   gate delete -l https://github.com -t oldtag
   ```

6. **删除链接的标签（通过名称）**
   ```bash
   gate delete -n github -t oldtag
   ```

---

### watch - 监控链接

监控或取消监控链接的变化，当链接内容发生变化时会收到邮件通知。

**语法：**
```bash
gate watch --link <URL> --watch
gate watch --link <URL> --unwatch
gate watch --name <名称> --watch
gate watch --name <名称> --unwatch
```

**参数：**
| 参数 | 简写 | 说明 |
|------|------|------|
| `--link` | `-l` | 要监控的链接 URL |
| `--name` | `-n` | 要监控的链接名称 |
| `--watch` | `-w` | 开启监控 |
| `--unwatch` | `-u` | 关闭监控 |

**使用场景：**

1. **通过 URL 开启监控**
   ```bash
   gate watch -l https://github.com -w
   ```

2. **通过 URL 关闭监控**
   ```bash
   gate watch -l https://github.com -u
   ```

3. **通过名称开启监控**
   ```bash
   gate watch -n github -w
   ```

4. **通过名称关闭监控**
   ```bash
   gate watch -n github -u
   ```

**注意事项：**
- 不能同时指定 `--link` 和 `--name`
- 必须指定 `--watch` 或 `--unwatch` 之一
- 监控功能需要服务器配置邮件服务

---

### open - 打开链接

在浏览器中打开链接。

**语法：**
```bash
gate open --name <名称>
gate open --tag <标签> [--page <页码>] [--page-size <每页数量>]
```

**参数：**
| 参数 | 简写 | 默认值 | 说明 |
|------|------|--------|------|
| `--name` | `-n` | | 链接名称 |
| `--tag` | `-t` | | 标签 |
| `--page` | `-p` | 1 | 页码（用于按标签打开） |
| `--page-size` | `-s` | 10 | 每页数量（用于按标签打开） |

**使用场景：**

1. **通过名称打开链接**
   ```bash
   gate open -n github
   ```

2. **通过标签打开链接**
   ```bash
   gate open -t dev
   ```

3. **通过标签打开链接（指定页码）**
   ```bash
   gate open -t tools -p 2 -s 5
   ```

**说明：**
- 通过名称打开时，直接打开该名称对应的链接
- 通过标签打开时，会列出所有符合条件的链接供选择

---

## 使用示例

### 完整工作流程

```bash
# 1. 注册账号
gate register

# 2. 登录
gate login

# 3. 测试连接
gate ping

# 4. 添加链接
gate add -l https://github.com -n github -n gh -t dev -t tools -r "代码托管平台"

# 5. 查看所有链接
gate list

# 6. 按标签查询
gate list -t dev

# 7. 开启链接监控
gate watch -n github -w

# 8. 打开链接
gate open -n github

# 9. 添加更多标签
gate add -n github -t favorite

# 10. 删除某个标签
gate delete -n github -t favorite

# 11. 关闭监控
gate watch -n github -u

# 12. 删除链接
gate delete -n github
```

### 批量操作示例

```bash
# 批量添加链接
gate add -l https://github.com -n github -n gh
gate add -l https://google.com -n google -n gg
gate add -l https://stackoverflow.com -n so -t dev

# 批量添加标签
gate add -n github -t dev -t tools -t favorite
gate add -n google -t search -t tools

# 批量删除标签
gate delete -n github -t favorite
gate delete -l https://google.com -t search

# 批量删除链接
gate delete -l https://github.com -l https://google.com
```

---

## 常见问题

### 1. 如何配置服务器地址？

有三种方式（按优先级）：
1. 命令行参数：`gate <command> -a localhost:8080`
2. 环境变量：在 `.env` 文件中设置 `SERVER_ADDR=localhost:8080`
3. 默认值：`127.0.0.1:8080`

### 2. 登录信息存储在哪里？

刷新令牌安全存储在系统密钥链中：
- macOS: Keychain
- Linux: Secret Service API (gnome-keyring, kwallet 等)
- Windows: Windows Credential Manager

### 3. 访问令牌过期怎么办？

客户端会自动检测访问令牌是否过期，如果过期会使用刷新令牌自动获取新的访问令牌。如果刷新令牌也过期，会提示重新登录。

### 4. 如何查看帮助信息？

```bash
# 查看所有命令
gate --help

# 查看特定命令的帮助
gate add --help
gate list --help
```

### 5. 链接监控如何工作？

- 服务器会定期检查被监控的链接是否发生变化
- 如果检测到变化，会通过邮件发送通知
- 需要服务器端配置邮件服务

### 6. 一个链接可以有多个名称吗？

可以。一个链接可以有多个别名（名称），方便通过不同的名称访问同一个链接。

```bash
gate add -l https://github.com -n github -n gh -n git
```

### 7. 如何更新链接的备注？

直接使用 add 命令添加新的备注即可，会覆盖旧备注：

```bash
gate add -n github -r "新的备注内容"
```

### 8. 删除操作的区别？

- `gate delete -l <url>`：删除整个链接及其所有关联数据
- `gate delete -n <name>`：删除链接的某个名称（如果是最后一个名称，则删除整个链接）
- `gate delete -l <url> -t <tag>`：只删除链接的某个标签
- `gate delete -l <url> -n <name>`：只删除链接的某个名称

---

## 版本信息

当前版本: **v1.0.0**

查看版本：
```bash
gate --version
```

---

## 技术支持

如有问题或建议，请联系项目维护者或提交 Issue。
