# EJECS - Embedded Just Enough Component System

A lightweight, type-safe Entity Component System (ECS) implementation for Luau/Roblox.

## Features

- Fully type-safe with Luau strict type checking
- Component-based architecture
- System-based logic processing
- Relationship support
- Frame-based system updates
- Zero external dependencies

## Documentation

- [API Documentation](docs/API.md) - Detailed API reference
- [Best Practices](docs/BEST_PRACTICES.md) - Recommended practices and patterns

## Project Structure

```
ejecs/
├── src/
│   ├── init.luau       # Main entry point
│   ├── types.luau      # Type definitions
│   └── world.luau      # World implementation
├── examples/
│   ├── basic.luau      # Basic usage example
│   └── relationships.luau  # Relationship example
├── docs/
│   ├── API.md          # API documentation
│   └── BEST_PRACTICES.md  # Best practices guide
└── README.md
```

## Installation

1. Copy the `src` directory into your project
2. Require the module:

```lua
local ejecs = require("path/to/ejecs")
```

## Quick Start

```lua
local ejecs = require("ejecs")
local world = ejecs.World.new()

-- Define components
local Position = world:component()
local Velocity = world:component()

-- Create an entity
local entity = world:createEntity()
world:addComponent(entity, Position, {x = 0, y = 0, z = 0})
world:addComponent(entity, Velocity, {x = 1, y = 1, z = 1})

-- Define a system
world:system({
    name = "Movement",
    query = {Position, Velocity},
    callback = function(entity, components)
        local position = components[Position]
        local velocity = components[Velocity]
        
        position.x += velocity.x
        position.y += velocity.y
        position.z += velocity.z
    end
})

-- Update systems
world:update()
```

## Examples

Check the `examples` directory for detailed examples:

1. `basic.luau` - Basic component and system usage
2. `relationships.luau` - Entity relationships and complex queries

## Type Safety

EJECS is designed to be fully type-safe with Luau's strict type checking:

```lua
-- Define component types
type PositionData = {
    x: number,
    y: number,
    z: number,
}

-- Type-safe component access
local position = world:get(entity, Position) :: PositionData
```

See the [Best Practices](docs/BEST_PRACTICES.md) guide for more type safety tips.

## Contributing

Contributions are welcome! Please read our contributing guidelines and submit pull requests to our repository.

## License

MIT 