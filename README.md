# Imperishable Gate | 不朽之门

<div align="center">

![Go Version](https://img.shields.io/badge/Go-1.25.1+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-green.svg?style=flat)
![Database](https://img.shields.io/badge/Database-SQLite%20%7C%20MySQL%20%7C%20PostgreSQL-blue?style=flat)
![Echo](https://img.shields.io/badge/Echo-Web_Framework-00C7B7?style=flat)
![GORM](https://img.shields.io/badge/GORM-ORM-red?style=flat)
![Cobra](https://img.shields.io/badge/Cobra-CLI-orange?style=flat)
![JWT](https://img.shields.io/badge/JWT-Auth-black?style=flat)

**白玉楼庭师的链接管理系统**

*以永夜抄为名，妖妖梦架构为灵感*

**[简体中文](README.md) | [English](README.en.md)**

[快速开始](#stage-1--白銀之春---快速开始) • [功能特性](#核心特性) • [文档](#文档) • [贡献](#贡献)

</div>

---

## 项目背景

> *多年前的某一天，我因为运气不好，不幸掉进了冥界。如今我成为了一名生活在白玉楼的普通庭师。除了每天扫院子、给主人做饭外，我还会在空闲时间尝试一些有趣的事情...*

**Imperishable Gate（不朽之门）** 是一个受东方 Project 启发的现代化命令行链接管理系统。当互联网基础设施延伸到白玉楼地区后，作为庭师的我踏上了"信息高速公路"，却发现浏览器书签过于简陋，PC 软件过于臃肿，文档又太过... 绿皮。

于是，我决定用前世学到的编程技能，创造一个优雅的链接管理系统！

### 项目亮点

本项目完整实现了从 **Stage 1 "白銀之春"** 到 **Stage 6 "冥界大小姐の亡骸"** 的全部功能，提供完整的链接生命周期管理，包括添加、删除、查询、标签分类、别名管理、备注、元数据自动抓取和智能链接监控等功能。

**特色功能：基于标签的丰富命令系统** - 客户端通过灵活的标签系统，实现了多维度的链接管理命令，支持按标签搜索、批量操作、标签组合查询等功能，让链接管理更加高效便捷。

## 架构设计

本项目采用**前后端分离架构**，借鉴了东方妖妖梦的分层设计：

- **Stage 1-2 | 基础层（白銀之春 · 迷途之家の黒猫）**
  - CLI 客户端：基于 Cobra 框架的命令行工具
  - 服务端：基于 Go + Echo 的 RESTful API 服务
  - 数据库：支持 SQLite / MySQL / PostgreSQL

- **Stage 3-4 | 功能增强层（人偶裁判の夜 · 雪上の櫻花結界）**
  - 标签系统、别名管理、备注功能
  - 自动元数据抓取（标题、描述、关键词）
  - 智能监控系统（内容变更检测、邮件通知）

- **Stage 5-6 | 安全认证层（白玉樓階梯の幻闊 · 冥界大小姐の亡骸）**
  - JWT 双令牌认证（Access Token + Refresh Token）
  - 用户注册、登录、邮箱验证
  - 令牌安全存储（系统 Keyring）

## 核心特性

### Stage 6 | 冥界大小姐の亡骸 - 安全认证

- **双令牌机制**：JWT Access Token（短期） + Refresh Token（长期）
- **密码加密**：采用 bcrypt 算法安全存储
- **安全存储**：令牌存储在系统 keyring（libsecret / Secret Service）
- **自动刷新**：令牌过期自动刷新，无感知体验
- **邮箱验证**：完整的邮箱验证流程

### Stage 3 | 人偶裁判の夜 - 灵活的标签系统（客户端特色）

客户端通过标签系统实现了丰富的命令功能，这是本项目的核心特色：

- **多对多关联**：一个链接可以有多个标签，一个标签可以关联多个链接
- **标签组合查询**：支持通过多个标签组合条件搜索链接
- **批量标签操作**：支持通过 URL 或别名批量添加/删除标签
- **标签分类管理**：按标签分类管理和检索链接
- **用户隔离**：每个用户都有独立的标签命名空间
- **智能标签命令**：
  - `gate search -t "tag1,tag2"` - 按标签搜索
  - `gate add -l url -t "tag1,tag2"` - 添加链接时设置标签
  - `gate list -t tag` - 列出特定标签的所有链接

### Stage 2-3 | 迷途之家の黒猫 & 人偶裁判の夜 - 核心链接管理

- **多维度检索**：支持通过 URL、标签、别名快速查询
- **批量操作**：一次性添加或删除多个链接、标签、别名
- **别名系统**（Name）：为链接设置多个别名，避免重复输入长 URL
- **备注功能**：为每个链接添加个性化备注
- **关联管理**：链接、标签、别名之间的灵活关联

### Stage 4 | 雪上の櫻花結界 - 智能监控

- **元数据抓取**：自动获取网页标题、描述、关键词
- **定时监控**：后台定时检查链接状态和内容变化
- **自动检测**：自动发现网页内容更新
- **分级监控**：watching（高频）和 non-watching（低频）两种模式
- **变化通知**：内容变化时自动通知（支持邮件，SMTP 协议）

### Stage 5 | 白玉樓階梯の幻闊 - 便捷的 CLI 体验

- **交互式界面**：友好的命令行交互
- **智能提示**：清晰的错误提示和操作指引
- **自动认证**：智能令牌管理，无需频繁登录
- **快速打开**：`gate open` 直接在浏览器打开链接
- **跨平台支持**：Linux、macOS、Windows

## Stage 1 | 白銀之春 - 快速开始

### 前置要求

- Go 1.25.1+
- 数据库（三选一）：
  - **SQLite**（默认，无需额外安装）
  - **MySQL** 5.7+ / 8.0+
  - **PostgreSQL** 12.0+

### 安装步骤

```bash
# 1. 克隆项目
git clone https://github.com/sokx6/imperishable-gate.git
cd imperishable-gate

# 2. 安装依赖
go mod download

# 3. 配置环境变量（可选）
# 服务端配置（默认使用 SQLite，无需额外配置）
# 将项目根目录的 .env.example 复制到 cmd/gate-server/
cp .env.example cmd/gate-server/.env

# 客户端配置（可选，用于设置服务器地址）
# cmd/gate/ 目录下已有 .env.example
cp cmd/gate/.env.example cmd/gate/.env
# 或直接使用环境变量
export GATE_SERVER_ADDR=http://localhost:4514

# 4. 构建
go build -o bin/gate-server ./cmd/gate-server
go build -o bin/gate ./cmd/gate

# 5. 启动服务端
./bin/gate-server start
```

### 客户端使用

```bash
# Stage 6 | 用户系统
# 配置服务器地址（重要：务必包含 http:// 或 https:// 前缀）
export GATE_SERVER_ADDR=http://localhost:4514

# 注册用户
./bin/gate register
# 登录
./bin/gate login
# 查看当前登录状态
./bin/gate whoami

# Stage 2-3 | 链接管理
# 添加链接（自动抓取元数据）
./bin/gate add -l "https://thwiki.cc" -t "touhou,wiki" -N "thwiki"

# 为链接添加备注
./bin/gate add -l "https://thwiki.cc" --remark "东方 Project 中文维基"

# 查看所有链接
./bin/gate list

# 按标签搜索
./bin/gate search -t "touhou"

# Stage 5 | 打开链接
# 通过别名直接打开
./bin/gate open -n "thwiki"
# 一次打开多个
./bin/gate open -n "thwiki,pixiv"

# Stage 4 | 监控
# 为链接启用监控（高频检查）
./bin/gate watch -n "thwiki"

# 标签相关操作（客户端特色功能）
# 按标签组合搜索
./bin/gate search -t "touhou,wiki"
# 为已有链接添加标签
./bin/gate add -n "thwiki" -t "reference"
# 列出特定标签的所有链接
./bin/gate list -t "touhou"
```

**提示**：
- 设置 `GATE_SERVER_ADDR` 时，请务必包含协议前缀（`http://` 或 `https://`）
- 使用别名（Names）可以避免重复输入长 URL
- 令牌会自动存储在系统 Keyring 中，安全可靠
- **标签系统是客户端的核心特色**，充分利用标签可以实现高效的链接管理

详细使用说明请参考 [快速开始指南](docs/getting-started.md) 和 [完整客户端文档](docs/client-guide.md)。

## 文档

### 用户文档（新手指南）
- [快速开始](docs/getting-started.md) - Stage 1：白銀之春 - 安装和基本使用
- [配置说明](docs/configuration.md) - 数据库、服务端、客户端的完整配置
- [客户端指南](docs/client-guide.md) - 所有 CLI 命令的详细说明（重点介绍标签命令）
- [API 文档](docs/api.md) - 完整的 RESTful API 参考

### 开发文档（进阶内容）
- [架构设计](docs/architecture.md) - 从 Stage 1 到 Stage 6 的系统架构演进
- [贡献指南](docs/contributing.md) - 如何为项目贡献代码
- [部署指南](docs/deployment.md) - 生产环境部署最佳实践
- [安全特性](docs/security.md) - JWT、bcrypt、Keyring 安全机制详解

### 开发日志
- [开发日志](docs/devlog.md) - 开发过程记录与感悟

## 技术栈

### 后端服务
- **Go 1.25.1+**：编程语言
- **Echo v4**：轻量级 Web 框架，RESTful API 设计
- **GORM**：ORM 框架，优雅的数据库操作
- **数据库支持**：
  - **SQLite**（默认）- Stage 2 基础实现，零配置
  - **MySQL** - Stage 2 扩展支持
  - **PostgreSQL** - Stage 2 扩展支持
- **JWT (golang-jwt/jwt)**：Stage 6 认证机制
- **bcrypt**：Stage 6 密码加密
- **goquery**：Stage 4 网页元数据抓取
- **SMTP**：Stage 4 邮件通知功能

### CLI 客户端
- **Cobra**：CLI 框架，优雅的命令行设计
- **go-keyring**：Stage 6 凭证安全存储（libsecret / Secret Service）
- **跨平台支持**：Linux / macOS / Windows

## 开发进度

### 已完成 (v1.0.0)

#### Stage 1 | 白銀之春
- 客户端&服务端基础框架
- RESTful API 设计
- Ping 测试功能

#### Stage 2 | 迷途之家の黒猫
- 数据库集成（SQLite/MySQL/PostgreSQL）
- 链接 CRUD 操作
- 完整的 API 路由设计

#### Stage 3 | 人偶裁判の夜
- 标签管理系统（多对多关联）
- 别名管理（Name → Link 映射）
- 备注功能
- 通过标签/别名查询链接

#### Stage 4 | 雪上の櫻花結界
- 自动网页元数据抓取（标题、描述、关键词）
- 定时轮询机制（goroutine 实现）
- 链接监控系统（watching/non-watching）
- 邮件通知功能（SMTP 协议）

#### Stage 5 | 白玉樓階梯の幻闊
- 链接搜索功能（模糊搜索）
- `gate open` 命令（浏览器打开）
- 批量打开多个链接

#### Stage 6 | 冥界大小姐の亡骸
- 完整的用户系统（注册/登录/登出/注销）
- JWT 双令牌认证（Access Token + Refresh Token）
- 邮箱验证功能
- 令牌安全存储（系统 Keyring）
- 自动令牌刷新
- `whoami` 命令

### 计划中 (v2.0)
- ElasticSearch 集成（Stage 5 扩展）
- Collection/View 对象（Stage 5 扩展）
- 管理员系统（Stage 6 扩展）
- 组系统（Stage 6 扩展）
- 审计日志系统（Stage 6 扩展）
- 链接导入/导出功能
- 单元测试覆盖

## 贡献

欢迎贡献代码、报告问题或提出新功能建议！

### 贡献流程

1. **Fork** 本项目到你的账户
2. 创建特性分支：`git checkout -b feature/AmazingFeature`
3. 编写代码，遵循项目规范：
   - 提交信息使用 `feat:`、`fix:`、`docs:`、`refactor:` 等前缀
   - 注意代码组织，避免把所有代码写在一个文件里
   - 添加必要的错误处理
4. 提交更改：`git commit -m 'feat: Add some AmazingFeature'`
5. 推送到分支：`git push origin feature/AmazingFeature`
6. 创建 **Pull Request**

### 建议方向

- **Stage 5 扩展**：ElasticSearch 集成、Collection 系统
- **Stage 6 扩展**：管理员系统、组共享、审计日志
- **测试改进**：单元测试、集成测试
- **文档完善**：更多示例、最佳实践

详见 [贡献指南](docs/contributing.md)。

## 许可证

本项目采用 MIT 许可证。详见 [LICENSE](LICENSE) 文件。

## 致谢

- 感谢所有贡献者和给本项目 star 的朋友
- 特别感谢 **东方 Project（東方Project）** 游戏系列的灵感启发
  - 项目名称来源：**永夜抄（Imperishable Night）**
  - 架构设计灵感：**妖妖梦（Perfect Cherry Blossom）** 的 Stage 结构
  - 主题氛围：白玉楼、冥界、春雪樱花
- 感谢 THBWiki（东方 Project 中文维基）提供的丰富资源

## 联系方式

- **GitHub Issues**: [提交问题](https://github.com/sokx6/imperishable-gate/issues)
- **作者**: QQ 2841929072

---

<div align="center">

**[回到顶部](#imperishable-gate--不朽之门)**

Made with love and Go by [locxl](https://github.com/sokx6)

*Inspired by Touhou Project 上海アリス幻樂団 (Team Shanghai Alice)*

</div>
