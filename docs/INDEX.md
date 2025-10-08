# Imperishable Gate 文档中心 | 不朽之门

**[📖 简体中文](INDEX.md) | [📘 English](INDEX.en.md)**

> 🌸 *"想要深入了解这个系统？让我们翻开白玉楼的典籍..."* 🌸

欢迎来到 **Imperishable Gate（不朽之门）** 的文档中心！本项目是一个受东方 Project 启发的现代化命令行链接管理系统，完整实现了从 **Stage 1「白銀之春」** 到 **Stage 6「冥界大小姐的亡骸」** 的所有功能。

---

## 🌸 Stage 1-2 | 白銀之春 & 迷途之家の黒猫 - 快速开始

### 新手必读

- **[📘 快速开始指南](getting-started.md)**  
  *Stage 1 基础：从安装到运行你的第一个命令*
  - 环境准备（Go、数据库选择：SQLite/MySQL/PostgreSQL）
  - 服务端和客户端安装
  - 基本命令使用示例（add / list / ping）
  
- **[⚙️ 配置说明](configuration.md)**  
  *Stage 2 数据库：多数据库支持配置*
  - 服务端环境变量配置（DB_TYPE, DSN, JWT_SECRET）
  - 客户端配置（GATE_SERVER_ADDR）
  - 数据库连接字符串示例

---

## 🎭 Stage 3-5 | 人偶裁判の夜 & 雪上の櫻花結界 & 白玉樓階梯の幻闊 - 功能使用

### 客户端使用指南

- **[💻 客户端完整手册](client-guide.md)**  
  *所有 CLI 命令的详细说明*
  - **Stage 2-3**：链接、标签、别名、备注管理
  - **Stage 4**：监控功能（watch）、元数据抓取
  - **Stage 5**：搜索（search）和快速打开（open）
  - **Stage 6**：用户系统（register / login / whoami）

### API 开发文档

- **[🔌 RESTful API 文档](api.md)**  
  *完整的 HTTP API 接口说明*
  - 认证相关 API（JWT Token）
  - 链接管理 API（CRUD）
  - 标签和别名 API  
  - 监控和通知 API

---

## 🏗️ Stage 1-6 | 架构与开发 - 技术深入

### 开发者文档

- **[🏯 架构设计](architecture.md)**  
  *从 Stage 1 到 Stage 6 的系统演进*
  - 前后端分离架构
  - 数据库设计（ER 图）
  - 技术栈选型理由
  - 各 Stage 功能模块划分

- **[🤝 贡献指南](contributing.md)**  
  *如何为项目贡献代码*
  - Fork 和 Pull Request 流程
  - Commit Message 规范
  - 代码审查标准

---

## 🚀 部署与安全 - 生产环境

### 运维文档

- **[🚀 部署指南](deployment.md)**  
  *生产环境部署最佳实践*
  - 服务器配置
  - systemd 服务设置
  - Nginx 反向代理
  - 备份策略

- **[🔒 安全特性详解](security.md)**  
  *Stage 6：深入理解安全机制*
  - JWT 双令牌认证原理（Access Token + Refresh Token）
  - bcrypt 密码加密（不会明文存密码吧？）
  - Keyring 安全存储（libsecret / Secret Service）
  - 数据隔离和权限控制

---

## 📝 开发日志与记录

### 项目历程

- **[📖 Devlog](devlog.md)**  
  *开发过程记录和心路历程*
  - 各 Stage 的实现过程
  - 遇到的问题和解决方案
  - 技术选型的思考

---

## 🎮 东方 Project 元素

本项目深受东方 Project 系列启发：

### 🌸 命名来源
- **项目名称**：永夜抄（Imperishable Night）
- **Stage 架构**：妖妖梦（Perfect Cherry Blossom）

### 📜 Stage 标题（来自妖妖梦）
1. **Stage 1**: 白銀之春（White & Pink Spring）
2. **Stage 2**: 迷途之家の黒猫（Black Cat of the Lost Home）
3. **Stage 3**: 人偶裁判の夜（Night of the Doll's Judgment）
4. **Stage 4**: 雪上の櫻花結界（Cherry Blossom Barrier on Snow）
5. **Stage 5**: 白玉樓階梯の幻闊（Phantom Expanse of Hakugyokurou Stairs）
6. **Stage 6**: 冥界大小姐の亡骸（The Corpse of the Netherworld Mistress）

### 🏯 世界观设定
- **地点**：白玉楼（西行寺家的庭园）
- **身份**：庭师（打扫院子、给主子做饭）
- **目标**：管理互联网冲浪时收集的链接

---

## 🌸 常见问题 FAQ

### Q: 我应该从哪里开始？
**A**: 新手请从 **[快速开始指南](getting-started.md)** 开始！它会带你完成 Stage 1 的基础搭建。

### Q: 如何切换数据库？
**A**: 查看 **[配置说明](configuration.md)** 中的数据库配置部分。支持 SQLite（默认）/ MySQL / PostgreSQL。

### Q: 客户端无法连接服务器？
**A**: 确保设置 `GATE_SERVER_ADDR` 时包含了 `http://` 或 `https://` 前缀！

### Q: Token 存储安全吗？
**A**: 非常安全！查看 **[安全特性](security.md)** 了解 Keyring 机制（Stage 6 实现）。

### Q: 如何贡献代码？
**A**: 阅读 **[贡献指南](contributing.md)**，然后提交你的 Pull Request！遵循 `feat:`, `fix:`, `docs:` 等 commit 规范。

### Q: 支持哪些功能？
**A**: 
- **Stage 1-2**: 基础客户端/服务端架构、数据库集成
- **Stage 3**: 标签、别名（Name）、备注系统
- **Stage 4**: 元数据抓取、智能监控、邮件通知
- **Stage 5**: 搜索、快速打开（open）
- **Stage 6**: 用户系统、JWT 认证、邮箱验证

---

<div align="center">

## 🎯 快速链接

| 我想... | 查看文档 |
|--------|---------|
| 快速安装并运行 | [快速开始](getting-started.md) |
| 配置数据库和服务器 | [配置说明](configuration.md) |
| 学习所有命令 | [客户端指南](client-guide.md) |
| 了解 API 接口 | [API 文档](api.md) |
| 理解系统架构 | [架构设计](architecture.md) |
| 参与开发 | [贡献指南](contributing.md) |
| 部署到生产环境 | [部署指南](deployment.md) |
| 了解安全机制 | [安全特性](security.md) |

---

### 🌸 *"经过七天的紧张开发，你的神秘妙妙软件终于发布了 1.0 版本"* 🌸

**[⬆ 返回主页](../README.md)** | **[GitHub 仓库](https://github.com/sokx6/imperishable-gate)**

*Made with ❤️ and 🌸 | Inspired by Touhou Project © 上海アリス幻樂団*

</div>
