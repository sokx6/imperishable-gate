# 贡献指南 | 参与开发

**[📖 简体中文](contributing.md) | [📘 English](contributing.en.md)**

> 🤝 *"幻想乡众人纷纷给你点了 star 并写了 issue！"*

感谢你考虑为 **Imperishable Gate（不朽之门）** 做出贡献！无论是报告 bug、提出新功能、改进文档，还是贡献代码，我们都非常欢迎！

## 🌸 项目规范说明

在上一世学习编程时，你就知道遵循规范的重要性。本项目遵循以下规范：

- **代码组织**：不要把所有代码写到一个文件里（MVC 模式）
- **Commit Message**：使用 `feat:`, `fix:`, `docs:`, `refactor:` 等前缀
- **错误处理**：添加必要的错误处理，让程序能应对常见意外
- **Git 工作流**：新功能在 `dev` 分支开发，每实现一点功能就 commit 一次

## 如何贡献

### 📝 报告 Bug

发现 bug？请创建 issue 并包含：

- 问题描述（发生在哪个 Stage 的功能？）
- 重现步骤
- 预期结果 vs 实际结果
- 环境信息（系统、Go 版本、数据库类型等）
- 相关日志输出

### 💡 提出新功能

有好的想法？（也许是 Stage 7 的内容？）

1. 先搜索是否已有类似的 issue
2. 创建新 issue 描述你的想法
3. 说明为什么需要这个功能
4. 如果可能，说明这个功能应该属于哪个 Stage

### 🔨 提交代码

#### 1. Fork 项目

```bash
# Fork 后克隆你的仓库
git clone https://github.com/your-username/imperishable-gate.git
cd imperishable-gate

# 添加上游仓库
git remote add upstream https://github.com/locxl/imperishable-gate.git
```

#### 2. 创建分支

```bash
# 功能分支
git checkout -b feature/your-feature-name

# 修复分支
git checkout -b fix/bug-description
```

#### 3. 编写代码

- 遵循现有代码风格
- 添加必要的注释
- 确保代码能够正常运行

#### 4. 提交更改

```bash
# 添加文件
git add .

# 提交（使用清晰的提交信息）
git commit -m "添加了某某功能"
# 或
git commit -m "修复了某某bug"
```

提交信息建议格式：
- `feat: 添加xxx功能`
- `fix: 修复xxx问题`
- `docs: 更新文档`
- `refactor: 重构xxx`
- `test: 添加测试`

#### 5. 推送和创建 PR

```bash
# 推送到你的仓库
git push origin feature/your-feature-name
```

然后在 GitHub 上创建 Pull Request。

## 代码规范

### Go 代码风格

```bash
# 格式化代码
gofmt -w .

# 或使用
go fmt ./...
```

### 基本规范

1. **命名**：使用有意义的变量名
   ```go
   // ✅ 好
   userID := 123
   userName := "Alice"
   
   // ❌ 不好
   u := 123
   n := "Alice"
   ```

2. **错误处理**：不要忽略错误
   ```go
   // ✅ 好
   if err != nil {
       return err
   }
   
   // ❌ 不好
   _ = someFunction()
   ```

3. **注释**：为导出的函数添加注释
   ```go
   // GetUser 根据ID获取用户信息
   func GetUser(id uint) (*User, error) {
       // ...
   }
   ```

## 测试

运行测试（如果有）：

```bash
go test ./...
```

## Pull Request 检查清单

提交 PR 前确认：

- [ ] 代码已格式化 (`gofmt -w .`)
- [ ] 代码可以正常编译运行
- [ ] 已添加必要的注释
- [ ] 提交信息清晰
- [ ] 已测试过功能

## 需要帮助？

- 查看现有 issues 寻找 `good first issue` 标签
- 在 issue 中提问
- 参考项目现有代码

## 行为准则

- 友好待人
- 尊重不同观点
- 接受建设性批评
- 专注于对项目最有利的事情

---

再次感谢你的贡献！每一个贡献都让项目变得更好 🎉
