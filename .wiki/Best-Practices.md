# Best Practices

This guide outlines recommended practices for using EJECS effectively in your projects.

## Component Design

### Keep Components Small and Focused

```ejecs
// Good: Small, focused components
component Position {
    x: number
    y: number
}

component Velocity {
    dx: number
    dy: number
}

// Bad: Large, unfocused component
component Physics {
    x: number
    y: number
    dx: number
    dy: number
    mass: number
    friction: number
    // ... many more fields
}
```

### Use Appropriate Types

```ejecs
// Good: Using appropriate types
component Player {
    name: string
    level: number
    inventory: {string: Item}
    skills: Skill[]
}

// Bad: Using strings for everything
component Player {
    name: string
    level: string  // Should be number
    inventory: string  // Should be map
    skills: string  // Should be array
}
```

### Leverage Default Values

```ejecs
// Good: Using default values
component Health {
    current: number = 100
    max: number = 100
    regenRate: number = 1
}

// Bad: No default values
component Health {
    current: number
    max: number
    regenRate: number
}
```

## System Design

### Single Responsibility Principle

```ejecs
// Good: Separate systems for different concerns
system Movement {
    query(Position, Velocity)
    {
        // Handle movement only
    }
}

system Collision {
    query(Position, Collider)
    {
        // Handle collisions only
    }
}

// Bad: One system doing everything
system Physics {
    query(Position, Velocity, Collider)
    {
        // Handle movement, collisions, and more
    }
}
```

### Use Parameters for Configuration

```ejecs
// Good: Configurable system
system Damage {
    query(Health)
    params {
        amount: number = 1
        type: string = "physical"
    }
    {
        // Use parameters
    }
}

// Bad: Hard-coded values
system Damage {
    query(Health)
    {
        // Hard-coded damage value
        local damage = 1
    }
}
```

### Consider System Frequency and Priority

```ejecs
// Good: Appropriate frequency and priority
system Physics {
    query(Position, Velocity)
    frequency(60)
    priority(1)
    {
        // Physics update
    }
}

system Render {
    query(Position, Sprite)
    frequency(60)
    priority(2)
    {
        // Rendering
    }
}

// Bad: No frequency or priority
system Update {
    query(Position, Velocity, Sprite)
    {
        // Everything in one system
    }
}
```

## Code Organization

### Group Related Components

```ejecs
// Good: Grouped components
// physics.ejecs
component Position { ... }
component Velocity { ... }
component Collider { ... }

// render.ejecs
component Sprite { ... }
component Animation { ... }

// Bad: All components in one file
// components.ejecs
component Position { ... }
component Sprite { ... }
component Velocity { ... }
component Animation { ... }
```

### Use Clear Naming Conventions

```ejecs
// Good: Clear, descriptive names
component PlayerInventory {
    items: {string: Item}
    gold: number
}

system UpdatePlayerPosition {
    query(Player, Position)
    {
        // Clear system purpose
    }
}

// Bad: Unclear names
component PI {
    i: {string: Item}
    g: number
}

system UPP {
    query(P, Pos)
    {
        // Unclear purpose
    }
}
```

## Performance Considerations

### Optimize Queries

```ejecs
// Good: Specific queries
system Movement {
    query(Position, Velocity)
    {
        // Only process entities with both components
    }
}

// Bad: Unnecessary components in query
system Movement {
    query(Position, Velocity, Sprite, Health)
    {
        // Processing entities that might not need movement
    }
}
```

### Use Appropriate Data Structures

```ejecs
// Good: Appropriate data structures
component Inventory {
    items: {string: Item}  // Map for quick lookup
    equipped: Equipment[]  // Array for ordered items
}

// Bad: Inappropriate data structures
component Inventory {
    items: Item[]  // Array for items that need lookup
    equipped: {string: Equipment}  // Map for ordered items
}
```

## Error Handling

### Validate Component Data

```ejecs
// Good: Data validation
system HealthUpdate {
    query(Health)
    {
        for _, entity in ipairs(entities) do
            local health = entity.Health
            health.current = math.max(0, math.min(health.max, health.current))
        end
    }
}

// Bad: No validation
system HealthUpdate {
    query(Health)
    {
        for _, entity in ipairs(entities) do
            local health = entity.Health
            health.current = health.current  // No bounds checking
        end
    }
}
```

### Handle Optional Fields

```ejecs
// Good: Proper optional field handling
system TransformUpdate {
    query(Transform)
    {
        for _, entity in ipairs(entities) do
            local transform = entity.Transform
            if transform.scale then
                -- Handle scale
            end
        end
    }
}

// Bad: Assuming optional fields exist
system TransformUpdate {
    query(Transform)
    {
        for _, entity in ipairs(entities) do
            local transform = entity.Transform
            transform.scale = transform.scale * 2  // Might be nil
        end
    }
}
```

## Testing

### Write Testable Systems

```ejecs
// Good: Testable system with clear inputs/outputs
system Damage {
    query(Health)
    params {
        amount: number = 1
    }
    {
        for _, entity in ipairs(entities) do
            local health = entity.Health
            health.current = health.current - amount
        end
    }
}

// Bad: System with side effects
system Damage {
    query(Health)
    {
        for _, entity in ipairs(entities) do
            local health = entity.Health
            health.current = health.current - 1
            love.audio.play("damage_sound")  // Side effect
        end
    }
}
```

### Document Complex Logic

```ejecs
// Good: Well-documented complex logic
system Pathfinding {
    query(Position, Target)
    {
        for _, entity in ipairs(entities) do
            // A* pathfinding algorithm
            // 1. Initialize open and closed sets
            // 2. Calculate heuristic for start node
            // 3. Process nodes until path found
            // ... detailed implementation
        end
    }
}

// Bad: Undocumented complex logic
system Pathfinding {
    query(Position, Target)
    {
        for _, entity in ipairs(entities) do
            -- Complex algorithm with no explanation
        end
    }
}
```

## Version Control

### Use Meaningful Commit Messages

```bash
# Good: Clear commit messages
git commit -m "Add damage system with type-based resistance"

# Bad: Unclear commit messages
git commit -m "Update code"
```

### Keep Generated Files Separate

```bash
# Good: Separate source and generated files
src/
  components.ejecs
  systems.ejecs
generated/
  components.lua
  systems.lua

# Bad: Mixed source and generated files
src/
  components.ejecs
  components.lua
  systems.ejecs
  systems.lua
``` 