# Contributing Guide

**[简体中文](contributing.md) | [English](contributing.en.md)**

> *"Everyone in Gensokyo starred your repo and opened issues!"*

Thank you for considering contributing to **Imperishable Gate**! Whether it's reporting bugs, proposing new features, improving documentation, or contributing code, we welcome all contributions!

## Project Standards

From my past life of learning programming, I know the importance of following standards. This project follows these conventions:

- **Code Organization**: Don't put all code in one file (MVC pattern)
- **Commit Messages**: Use prefixes like `feat:`, `fix:`, `docs:`, `refactor:`, etc.
- **Error Handling**: Add necessary error handling to cope with common exceptions
- **Git Workflow**: Develop new features on the `dev` branch, commit frequently as you make progress

## How to Contribute

### Reporting Bugs

Found a bug? Please create an issue including:

- Description of the problem (which Stage's functionality?)
- Steps to reproduce
- Expected vs actual results
- Environment information (OS, Go version, database type, etc.)
- Relevant log output

### Proposing New Features

Have a great idea? (Maybe for Stage 7?)

1. Search for similar existing issues first
2. Create a new issue describing your idea
3. Explain why this feature is needed
4. If possible, indicate which Stage this feature should belong to

### Submitting Code

#### 1. Fork the Project

```bash
# Clone your forked repository
git clone https://github.com/your-username/imperishable-gate.git
cd imperishable-gate

# Add upstream repository
git remote add upstream https://github.com/sokx6/imperishable-gate.git
```

#### 2. Create a Branch

```bash
# Feature branch
git checkout -b feature/your-feature-name

# Bug fix branch
git checkout -b fix/bug-description
```

#### 3. Write Code

- Follow existing code style
- Add necessary comments
- Ensure code runs properly

#### 4. Commit Changes

```bash
# Add files
git add .

# Commit with clear commit message
git commit -m "Add some feature"
# or
git commit -m "Fix some bug"
```

Recommended commit message format:
- `feat: Add xxx feature`
- `fix: Fix xxx issue`
- `docs: Update documentation`
- `refactor: Refactor xxx`
- `test: Add tests`

#### 5. Push and Create PR

```bash
# Push to your repository
git push origin feature/your-feature-name
```

Then create a Pull Request on GitHub.

## Code Standards

### Go Code Style

```bash
# Format code
gofmt -w .

# Or use
go fmt ./...
```

### Basic Standards

1. **Naming**: Use meaningful variable names
   ```go
   // Good
   userID := 123
   userName := "Alice"
   
   // Bad
   u := 123
   n := "Alice"
   ```

2. **Error Handling**: Don't ignore errors
   ```go
   // Good
   if err != nil {
       return err
   }
   
   // Bad
   _ = someFunction()
   ```

3. **Comments**: Add comments for exported functions
   ```go
   // GetUser retrieves user information by ID
   func GetUser(id uint) (*User, error) {
       // ...
   }
   ```

## Testing

Run tests (if available):

```bash
go test ./...
```

## Pull Request Checklist

Before submitting a PR, confirm:

- [ ] Code is formatted (`gofmt -w .`)
- [ ] Code compiles and runs properly
- [ ] Necessary comments are added
- [ ] Commit messages are clear
- [ ] Functionality has been tested

## Need Help?

- Check existing issues for `good first issue` tags
- Ask questions in issues
- Refer to existing project code

## Code of Conduct

- Be friendly to others
- Respect different viewpoints
- Accept constructive criticism
- Focus on what's best for the project

---

Thank you again for your contribution! Every contribution makes the project better
