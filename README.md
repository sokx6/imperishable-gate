# Imperishable Gate

<div align="center">

![Go Version](https://img.shields.io/bvim cmd/gate-server/.env

# 4. 构建
go build -o bin/gate-server ./cmd/gate-server
go build -o bin/gate ./cmd/gate

# 5. 启动服务端（默认使用 SQLite，无需额外配置数据库）-1.25.1+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-green.svg)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Database-336791?style=flat&logo=postgresql)
![Echo](https://img.shields.io/badge/Echo-Web_Framework-00C7B7?style=flat)

**现代化的命令行链接管理系统**

[快速开始](#-快速开始) • [功能特性](#-核心特性) • [文档](#-文档) • [贡献](#-贡献)

</div>

---

**Imperishable Gate（不朽之门）** 是一个受东方 Project 启发的现代化命令行链接管理系统。它提供了完整的链接生命周期管理功能，包括添加、删除、查询、标签分类、别名管理、备注、元数据自动抓取和智能链接监控等功能。

## 📖 项目简介

本项目采用**前后端分离架构**：
- **后端服务**：基于 Go + Echo + PostgreSQL + GORM 构建的高性能 RESTful API 服务
- **CLI 客户端**：基于 Cobra 框架的强大命令行工具，支持跨平台使用

## ✨ 核心特性

### 🔐 安全认证
- **双令牌机制**：JWT Access Token（短期） + Refresh Token（长期）
- **密码加密**：采用 bcrypt 算法安全存储
- **安全存储**：令牌存储在系统 keyring（Linux/macOS/Windows）
- **自动刷新**：令牌过期自动刷新，无感知体验
- **邮箱验证**：支持邮箱验证功能

### 🔗 强大的链接管理
- **多维度检索**：支持通过 URL、标签、别名快速查询
- **批量操作**：一次性添加或删除多个链接、标签、别名
- **智能抓取**：自动获取网页元数据（标题、描述、关键词）
- **状态追踪**：实时记录 HTTP 状态码，监控链接健康
- **关联管理**：链接、标签、别名之间的灵活关联

### 🏷️ 灵活的标签系统
- **多对多关联**：一个链接可以有多个标签
- **标签分类**：按标签分类管理和检索链接
- **批量标签操作**：支持通过 URL 或别名批量添加/删除标签
- **用户隔离**：标签在用户范围内唯一

### 👀 智能监控系统
- **定时监控**：后台定时检查链接状态和内容变化
- **自动检测**：自动发现网页内容更新
- **分级监控**：watching（高频）和 non-watching（低频）两种模式
- **变化通知**：内容变化时自动通知（支持邮件）

### 🖥️ 便捷的 CLI 体验
- **交互式界面**：友好的命令行交互
- **智能提示**：清晰的错误提示和操作指引
- **自动认证**：智能令牌管理，无需频繁登录
- **跨平台支持**：支持 Linux、macOS、Windows

## 🚀 快速开始

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
cp .env.example cmd/gate-server/.env

# 客户端配置（可选，用于设置服务器地址）
cp cmd/gate/.env.example cmd/gate/.env
# 或直接使用环境变量
export GATE_SERVER_ADDR=http://localhost:4514

# 4. 构建
go build -o bin/gate-server ./cmd/gate-server
go build -o bin/gate ./cmd/gate

# 6. 启动服务端
./bin/gate-server start
```

### 客户端使用

```bash
# 配置服务器地址（重要：务必包含 http:// 或 https:// 前缀）
export GATE_SERVER_ADDR=http://localhost:4514

# 注册用户
./bin/gate register

# 登录
./bin/gate login

# 添加链接
./bin/gate add -l "https://example.com" -t "tech,blog" -N "mysite"

# 查看链接
./bin/gate list

# 打开链接
./bin/gate open -n "mysite"
```

> **重要提示**：设置 `GATE_SERVER_ADDR` 时，请务必包含协议前缀（`http://` 或 `https://`），否则默认会使用 `https://`，可能导致本地开发时连接失败。

详细使用说明请参考 [快速开始指南](docs/getting-started.md)。

## 📚 文档

### 用户文档
- [📘 快速开始](docs/getting-started.md) - 安装和基本使用
- [⚙️ 配置说明](docs/configuration.md) - 环境变量和配置详解
- [💻 客户端指南](docs/client-guide.md) - 完整的 CLI 命令文档
- [🔌 API 文档](docs/api.md) - RESTful API 接口说明

### 开发文档
- [🏗️ 架构设计](docs/architecture.md) - 系统架构和技术栈
- [👨‍💻 开发指南](docs/development.md) - 开发环境搭建和规范
- [🤝 贡献指南](docs/contributing.md) - 如何参与项目开发
- [🚀 部署指南](docs/deployment.md) - 生产环境部署
- [🔒 安全特性](docs/security.md) - 安全机制详解

## 🏗️ 技术栈

### 后端
- **Go 1.25.1+**：编程语言
- **Echo v4**：Web 框架
- **GORM**：ORM 框架  
- **数据库支持**：SQLite（默认）/ MySQL / PostgreSQL
- **JWT**：认证机制
- **goquery**：网页抓取

### 客户端
- **Cobra**：CLI 框架
- **go-keyring**：凭证安全存储

## 📋 功能清单

### 已完成
- ✅ 用户注册、登录、登出
- ✅ 邮箱验证功能
- ✅ JWT 双令牌认证
- ✅ 链接 CRUD 操作
- ✅ 标签管理
- ✅ 别名管理
- ✅ 备注功能
- ✅ 元数据自动抓取
- ✅ 链接监控和变化通知
- ✅ CLI 客户端
- ✅ 自动令牌刷新

### 计划中
- 📝 响应格式统一化
- 📝 完善的错误处理
- 📝 单元测试
- 📝 链接导入/导出功能
- 📝 搜索功能增强

## 🤝 贡献

欢迎贡献代码、报告问题或提出新功能建议！

1. Fork 项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建 Pull Request

详见 [贡献指南](docs/contributing.md)。

## 📄 许可证

本项目采用 MIT 许可证。详见 [LICENSE](LICENSE) 文件。

## 🙏 致谢

- 感谢所有贡献者
- 灵感来源于东方 Project

## 📞 联系方式

- **GitHub Issues**: [提交问题](https://github.com/sokx6/imperishable-gate/issues)
- **GitHub Discussions**: [参与讨论](https://github.com/sokx6/imperishable-gate/discussions)

---

<div align="center">

**[⬆ 回到顶部](#imperishable-gate)**

Made with ❤️ by [sokx6](https://github.com/sokx6)

</div>
