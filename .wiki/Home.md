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

## Quick Start

1. Install EJECS:
```bash
go install github.com/your-org/ejecs/cmd/ejecs@latest
```

2. Create your first EJECS file:
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

3. Generate code:
```bash
ejecs -input game.ejecs -output game.lua -library ecr
```

## Documentation

- [Language Reference](Language-Reference.md): Complete language syntax and features
- [Component System](Component-System.md): Component definition and usage
- [Code Generation](Code-Generation.md): Code generation options and targets
- [Embedding API](Embedding-API.md): How to embed EJECS in your application
- [Best Practices](Best-Practices.md): Guidelines for writing EJECS code
- [Examples](Examples.md): Example EJECS files and usage
- [Tutorials](Tutorials.md): Step-by-step guides
- [FAQ](FAQ.md): Frequently asked questions
- [Troubleshooting](Troubleshooting.md): Common issues and solutions

## Contributing

We welcome contributions! Please see our [Contributing Guide](Contributing.md) for details on how to get started.

## License

EJECS is licensed under the MIT License. See [LICENSE](LICENSE) for details. 