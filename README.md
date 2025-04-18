# EJECS - Entity Component System IDL

EJECS (Entity Component System Interface Definition Language) is a domain-specific language for defining Entity Component Systems with a focus on clean syntax and code generation. It provides a declarative way to define components, systems, and relationships, with support for generating code for multiple ECS libraries.

## Features

- **Component Definitions**: Define data structures with typed fields and default values
- **System Specifications**: Declare systems with queries, parameters, and code blocks
- **Relationship Modeling**: Define relationships between entities with type safety
- **Code Generation**: Generate code for multiple ECS libraries (ECR and JECS)
- **Type Safety**: Strong typing with support for basic and complex types
- **System Configuration**: Configure systems with parameters, frequency, and priority
- **Extensible**: Support for custom types and code generation targets

## Installation

```bash
go install github.com/your-org/ejecs/cmd/ejecs@latest
```

## Quick Start

1. Create a new EJECS file (`game.ejecs`):

```ejecs
component Position {
    x: number;
    y: number;
}

component Velocity {
    dx: number;
    dy: number;
}

system Movement {
    query(Position, Velocity)
    {
        for _, entity in ipairs(entities) do
            local pos = entity.Position
            local vel = entity.Velocity
            pos.x = pos.x + vel.dx
            pos.y = pos.y + vel.dy
        end
    }
}
```

2. Generate code:

```bash
ejecs -input game.ejecs -output game.lua -library ecr
```

## Documentation

- [Language Reference](.wiki/Language-Reference.md): Complete language syntax and features
- [Component System](.wiki/Component-System.md): Component definition and usage
- [Code Generation](.wiki/Code-Generation.md): Code generation options and targets
- [Embedding API](.wiki/Embedding-API.md): How to embed EJECS in your application
- [Best Practices](.wiki/Best-Practices.md): Guidelines for writing EJECS code
- [Examples](.wiki/Examples.md): Example EJECS files and usage
- [Tutorials](.wiki/Tutorials.md): Step-by-step guides
- [FAQ](.wiki/FAQ.md): Frequently asked questions
- [Troubleshooting](.wiki/Troubleshooting.md): Common issues and solutions

## Project Structure

```
ejecs/
├── cmd/
│   └── ejecs/          # Command line tool
├── internal/
│   ├── ast/           # Abstract Syntax Tree
│   ├── lexer/         # Lexical analysis
│   ├── parser/        # Syntax parsing
│   └── generator/     # Code generation
├── examples/          # Example EJECS files
└── .wiki/            # Documentation
```

## Contributing

We welcome contributions! Please see our [Contributing Guide](.wiki/Contributing.md) for details on how to get started.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. 