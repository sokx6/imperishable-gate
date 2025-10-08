# 文档索引

欢迎查看 Imperishable Gate 的文档！

## 📚 快速导航

### 🎯 新手入门

1. **[快速开始](getting-started.md)** - 安装和基本使用
2. **[配置说明](configuration.md)** - 环境变量和配置
3. **[客户端指南](client-guide.md)** - CLI 命令参考

### 📖 使用文档

- **[API 文档](api.md)** - RESTful API 接口说明
- **[架构设计](architecture.md)** - 系统架构和技术栈
- **[安全特性](security.md)** - 安全机制说明

### 👨‍💻 开发相关

- **[开发指南](development.md)** - 开发环境和代码规范
- **[贡献指南](contributing.md)** - 如何为项目做贡献
- **[部署指南](deployment.md)** - 服务器部署方法

## 🔍 常见问题

### 安装配置

- [环境要求](configuration.md#环境要求)
- [数据库配置](configuration.md#数据库配置)
- [客户端配置](configuration.md#客户端配置)

### 使用问题

- [客户端无法连接](getting-started.md#1-客户端无法连接服务器)
- [Linux keyring 错误](getting-started.md#2-linux-下-keyring-错误)
- [数据库连接失败](getting-started.md#3-数据库连接失败)

### 开发问题

- [项目结构](development.md#项目结构)
- [添加新功能](development.md#添加新功能)
- [调试方法](development.md#调试)

## 📝 重要提示

### 客户端配置

配置服务器地址时**务必包含协议前缀**：

```bash
# ✅ 正确
export GATE_SERVER_ADDR=http://localhost:4514

# ❌ 错误（会默认使用 https://）
export GATE_SERVER_ADDR=localhost:4514
```

详见：[客户端配置](configuration.md#客户端配置)

## 📋 文档列表

### 用户文档
- [快速开始](getting-started.md)
- [配置说明](configuration.md)
- [客户端指南](client-guide.md)
- [API 文档](api.md)

### 技术文档
- [架构设计](architecture.md)
- [安全特性](security.md)

### 开发文档
- [开发指南](development.md)
- [贡献指南](contributing.md)
- [部署指南](deployment.md)

## 🔗 外部资源

### Go 学习
- [Go 官方文档](https://go.dev/doc/)
- [Go by Example](https://gobyexample.com/)

### 框架文档
- [Echo 文档](https://echo.labstack.com/)
- [GORM 文档](https://gorm.io/docs/)
- [Cobra 文档](https://github.com/spf13/cobra)

## 💡 建议阅读顺序

**新用户**：
1. 快速开始
2. 配置说明
3. 客户端指南

**开发者**：
1. 架构设计
2. 开发指南
3. 贡献指南

**部署运维**：
1. 配置说明
2. 安全特性
3. 部署指南

---

有问题？查看[常见问题](#-常见问题)或在 [GitHub Issues](https://github.com/sokx6/imperishable-gate/issues) 提问。
