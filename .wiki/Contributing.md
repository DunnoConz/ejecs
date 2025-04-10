# Contributing to EJECS

Thank you for your interest in contributing to EJECS! This guide will help you get started with contributing to the project.

## Getting Started

1. Fork the repository
2. Clone your fork:
```bash
git clone https://github.com/your-username/ejecs.git
cd ejecs
```

3. Set up the development environment:
```bash
# Install dependencies
go mod download

# Build the project
go build ./cmd/ejecs
```

## Development Workflow

1. Create a new branch for your feature:
```bash
git checkout -b feature/your-feature-name
```

2. Make your changes following the [code style guidelines](#code-style)

3. Run tests:
```bash
go test ./...
```

4. Commit your changes:
```bash
git commit -m "Description of your changes"
```

5. Push to your fork:
```bash
git push origin feature/your-feature-name
```

6. Create a pull request

## Code Style

### Go Code

Follow the standard Go formatting guidelines:
- Use `gofmt` or `goimports` for formatting
- Follow effective Go practices
- Write clear, documented code

Example:
```go
// Package parser implements the EJECS parser.
package parser

// Parser represents an EJECS parser.
type Parser struct {
    lexer *lexer.Lexer
    // ...
}

// ParseComponent parses a component definition.
func (p *Parser) ParseComponent() (*ast.Component, error) {
    // ...
}
```

### EJECS Code

When contributing examples or documentation:
- Use clear, descriptive names
- Include comments for complex logic
- Follow the [best practices](Best-Practices)

Example:
```ejecs
// Component for tracking entity position
component Position {
    x: number
    y: number
}

// System for updating entity position
system Movement {
    query(Position, Velocity)
    frequency(60)
    {
        // Update position based on velocity
        for _, entity in ipairs(entities) do
            local pos = entity.Position
            local vel = entity.Velocity
            pos.x = pos.x + vel.dx
            pos.y = pos.y + vel.dy
        end
    }
}
```

## Testing

### Writing Tests

1. Create test files with `_test.go` suffix
2. Use table-driven tests where appropriate
3. Include both success and failure cases

Example:
```go
func TestParseComponent(t *testing.T) {
    tests := []struct {
        input    string
        expected *ast.Component
        err      bool
    }{
        {
            input: "component Position { x: number }",
            expected: &ast.Component{
                Name: "Position",
                Fields: []*ast.Field{
                    {Name: "x", Type: "number"},
                },
            },
            err: false,
        },
        // ... more test cases
    }

    for _, test := range tests {
        t.Run(test.input, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test -run TestParseComponent
```

## Documentation

### Updating Documentation

1. Update relevant wiki pages
2. Keep examples up to date
3. Document new features

### Writing Documentation

- Use clear, concise language
- Include code examples
- Follow the existing documentation style
- Update the [FAQ](FAQ) if adding new features

## Pull Requests

### Creating a Pull Request

1. Ensure your branch is up to date:
```bash
git fetch origin
git rebase origin/main
```

2. Push your changes:
```bash
git push origin feature/your-feature-name
```

3. Create a pull request on GitHub

### Pull Request Guidelines

- Provide a clear description of changes
- Reference related issues
- Include tests for new features
- Update documentation as needed
- Ensure all tests pass

## Issue Tracking

### Creating Issues

When creating issues:
- Use clear, descriptive titles
- Provide detailed descriptions
- Include reproduction steps
- Add relevant code examples
- Specify expected vs actual behavior

### Issue Labels

- `bug`: Something isn't working
- `enhancement`: New feature or request
- `documentation`: Documentation updates
- `question`: Further information needed
- `help wanted`: Extra attention needed

## Community Guidelines

### Code of Conduct

- Be respectful and inclusive
- Focus on constructive feedback
- Help others when possible
- Follow the project's code of conduct

### Communication

- Use GitHub issues for bug reports
- Use GitHub discussions for questions
- Join the Discord community for real-time chat

## Release Process

### Versioning

EJECS follows semantic versioning:
- Major version for incompatible changes
- Minor version for new features
- Patch version for bug fixes

### Creating a Release

1. Update version numbers
2. Update changelog
3. Create release tag
4. Build and test release
5. Publish release notes

## Getting Help

If you need help:
1. Check the [documentation](Home)
2. Search existing issues
3. Ask in the Discord community
4. Create a new issue

## Thank You!

Thank you for contributing to EJECS! Your contributions help make the project better for everyone. 