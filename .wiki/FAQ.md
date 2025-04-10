# Frequently Asked Questions

## General Questions

### What is EJECS?
EJECS (Entity Component System IDL) is a declarative language for defining Entity Component Systems. It provides a clean, type-safe syntax for describing components, systems, and their relationships, with support for code generation to various ECS implementations.

### Why use EJECS instead of writing ECS code directly?
EJECS offers several advantages:
- Clean, declarative syntax for ECS definitions
- Type safety and validation
- Code generation for multiple platforms
- Easy maintenance and refactoring
- Documentation as code
- Consistent architecture across projects

### What platforms does EJECS support?
Currently, EJECS supports generating code for:
- ECR (Lua)
- JECS (Lua)
- More platforms planned for future releases

## Language Features

### What types are supported in EJECS?
EJECS supports:
- Basic types: `number`, `string`, `boolean`, `entity`
- Complex types: arrays (`[]`), maps (`{string: Type}`)
- Optional fields (`?`)
- Custom types (component references)

### How do I define relationships between entities?
You can define relationships using the `entity` type:

```ejecs
component Parent {
    child: entity
}

component Child {
    parent: entity
}
```

### Can I have optional fields in components?
Yes, use the `?` suffix to mark a field as optional:

```ejecs
component Transform {
    position: Position
    rotation: number = 0
    scale: number?  // Optional field
}
```

## System Design

### How do I configure system behavior?
Use the `params` block to define configurable parameters:

```ejecs
system Damage {
    query(Health)
    params {
        amount: number = 1
        type: string = "physical"
    }
    {
        // Use parameters in system logic
    }
}
```

### What's the difference between frequency and priority?
- `frequency`: How often the system runs (in updates per second)
- `priority`: Order in which systems run (lower numbers run first)

```ejecs
system Physics {
    query(Position, Velocity)
    frequency(60)  // Run 60 times per second
    priority(1)    // Run before systems with higher priority
    {
        // Physics update
    }
}
```

### Can systems have side effects?
While possible, it's generally recommended to keep systems focused on data transformation. Side effects like audio playback or rendering should be handled by dedicated systems.

## Code Generation

### How do I generate code from EJECS files?
Use the `ejecs generate` command:

```bash
# Generate ECR code
ejecs generate -f ecr my_game.ejecs > my_game_ecr.lua

# Generate JECS code
ejecs generate -f jecs my_game.ejecs > my_game_jecs.lua
```

### Can I customize the generated code?
Yes, you can use code blocks in systems to insert custom code:

```ejecs
system CustomLogic {
    query(Position)
    {
        -- Custom Lua code here
        for _, entity in ipairs(entities) do
            -- Your custom logic
        end
    }
}
```

### How do I handle platform-specific code?
Use conditional code blocks or separate systems for platform-specific logic:

```ejecs
system PlatformSpecific {
    query(Position)
    {
        if PLATFORM == "mobile" then
            -- Mobile-specific code
        else
            -- Desktop-specific code
        end
    }
}
```

## Best Practices

### How should I organize my EJECS files?
Group related components and systems in separate files:

```
src/
  physics/
    position.ejecs
    velocity.ejecs
    collision.ejecs
  render/
    sprite.ejecs
    animation.ejecs
  game/
    player.ejecs
    inventory.ejecs
```

### How do I handle complex game logic?
Break down complex logic into smaller, focused systems:

```ejecs
// Instead of one large system
system GameLogic {
    // ... complex logic
}

// Use multiple focused systems
system Movement { ... }
system Combat { ... }
system Inventory { ... }
```

### What's the best way to handle entity creation and destruction?
Use dedicated systems for entity lifecycle management:

```ejecs
system SpawnEnemy {
    query(SpawnPoint)
    params {
        enemyType: string
    }
    {
        for _, entity in ipairs(entities) do
            local enemy = ecr:add({
                Position = {x = 0, y = 0},
                Health = {current = 100, max = 100}
            })
        end
    }
}

system Cleanup {
    query(Health)
    {
        for _, entity in ipairs(entities) do
            if entity.Health.current <= 0 then
                ecr:remove(entity)
            end
        end
    }
}
```

## Troubleshooting

### Common Error Messages

#### "Unknown type: X"
This error occurs when you reference a type that hasn't been defined. Make sure all component types are defined before use.

#### "Duplicate component definition"
You've defined the same component twice. Check for duplicate component definitions in your files.

#### "Invalid query"
The system is trying to query components that don't exist. Verify that all components in the query are defined.

### Debugging Tips

1. Use the `--debug` flag when generating code:
```bash
ejecs generate -f ecr --debug my_game.ejecs
```

2. Add validation systems:
```ejecs
system ValidateHealth {
    query(Health)
    {
        for _, entity in ipairs(entities) do
            assert(entity.Health.current >= 0, "Health cannot be negative")
            assert(entity.Health.current <= entity.Health.max, "Health exceeds maximum")
        end
    }
}
```

3. Use logging in systems:
```ejecs
system DebugPosition {
    query(Position)
    {
        for _, entity in ipairs(entities) do
            print(string.format("Entity position: %f, %f", 
                entity.Position.x, entity.Position.y))
        end
    }
}
```

## Contributing

### How can I contribute to EJECS?
1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Submit a pull request

### What should I include in my pull request?
- Clear description of changes
- Tests for new features
- Updated documentation
- Example usage if applicable

### Where can I get help?
- GitHub Issues
- Discord community
- Documentation
- Example projects 