# Troubleshooting

This guide helps you diagnose and fix common issues with EJECS.

## Common Errors

### Parser Errors

#### "Unexpected token: X"
This error occurs when the parser encounters an unexpected token.

Possible causes:
1. Syntax error in your EJECS file
2. Missing or extra braces/brackets
3. Invalid character in identifier

Example fix:
```ejecs
// Error: Unexpected token ':'
component Player {
    health: number  // Correct
    health : number // Error: Extra space before colon
}
```

#### "Expected '}' but found X"
This error occurs when a component or system block is not properly closed.

Example fix:
```ejecs
// Error: Missing closing brace
component Position {
    x: number
    y: number
// }  // Add missing brace
```

### Type Errors

#### "Unknown type: X"
This error occurs when you reference a type that hasn't been defined.

Example fix:
```ejecs
// Error: Unknown type "Vector2"
component Position {
    pos: Vector2  // Error: Vector2 not defined
}

// Fix: Define the type first
component Vector2 {
    x: number
    y: number
}

component Position {
    pos: Vector2  // Now valid
}
```

#### "Type mismatch: expected X, got Y"
This error occurs when you assign a value of the wrong type.

Example fix:
```ejecs
// Error: Type mismatch
component Player {
    health: number = "100"  // Error: string instead of number
}

// Fix: Use correct type
component Player {
    health: number = 100  // Correct
}
```

### System Errors

#### "Invalid query: component X not found"
This error occurs when a system queries a component that doesn't exist.

Example fix:
```ejecs
// Error: Health component not defined
system Damage {
    query(Health)  // Error: Health not defined
    {
        // ...
    }
}

// Fix: Define the component first
component Health {
    current: number
    max: number
}

system Damage {
    query(Health)  // Now valid
    {
        // ...
    }
}
```

#### "Duplicate system definition"
This error occurs when you define the same system twice.

Example fix:
```ejecs
// Error: Duplicate system
system Movement { ... }
system Movement { ... }  // Error: Duplicate

// Fix: Use unique names
system Movement { ... }
system AdvancedMovement { ... }  // Correct
```

## Debugging Tips

### Using the Debug Flag

Enable debug output when generating code:
```bash
ejecs generate -f ecr --debug my_game.ejecs
```

This will show:
- Parsing steps
- Type checking results
- Code generation details

### Adding Debug Systems

Create systems specifically for debugging:
```ejecs
system DebugPosition {
    query(Position)
    {
        for _, entity in ipairs(entities) do
            print(string.format(
                "Entity %d position: %f, %f",
                entity.id,
                entity.Position.x,
                entity.Position.y
            ))
        end
    }
}

system DebugHealth {
    query(Health)
    {
        for _, entity in ipairs(entities) do
            if entity.Health.current < entity.Health.max * 0.5 then
                print(string.format(
                    "Entity %d health low: %d/%d",
                    entity.id,
                    entity.Health.current,
                    entity.Health.max
                ))
            end
        end
    }
}
```

### Validation Systems

Add systems to validate component data:
```ejecs
system ValidateHealth {
    query(Health)
    {
        for _, entity in ipairs(entities) do
            assert(
                entity.Health.current >= 0,
                string.format("Health cannot be negative: %d", entity.Health.current)
            )
            assert(
                entity.Health.current <= entity.Health.max,
                string.format("Health exceeds maximum: %d > %d", 
                    entity.Health.current, 
                    entity.Health.max
                )
            )
        end
    }
}

system ValidatePosition {
    query(Position)
    {
        for _, entity in ipairs(entities) do
            assert(
                not (math.isnan(entity.Position.x) or math.isnan(entity.Position.y)),
                "Position contains NaN values"
            )
        end
    }
}
```

## Performance Issues

### Slow Systems

If a system is running slowly:
1. Check the query complexity
2. Look for unnecessary component access
3. Consider system frequency

Example optimization:
```ejecs
// Before: Complex query with unnecessary components
system Update {
    query(Position, Velocity, Sprite, Health)
    {
        for _, entity in ipairs(entities) do
            // Only using Position and Velocity
            entity.Position.x = entity.Position.x + entity.Velocity.dx
            entity.Position.y = entity.Position.y + entity.Velocity.dy
        end
    }
}

// After: Optimized query
system Movement {
    query(Position, Velocity)
    {
        for _, entity in ipairs(entities) do
            entity.Position.x = entity.Position.x + entity.Velocity.dx
            entity.Position.y = entity.Position.y + entity.Velocity.dy
        end
    }
}
```

### Memory Usage

High memory usage can be caused by:
1. Too many entities
2. Large component data
3. Unnecessary component creation

Example fix:
```ejecs
// Before: Creating components for all entities
system Spawn {
    query(SpawnPoint)
    {
        for _, entity in ipairs(entities) do
            local newEntity = ecr:add({
                Position = {x = 0, y = 0},
                Velocity = {dx = 0, dy = 0},
                Sprite = {image = "default.png"},
                Health = {current = 100, max = 100}
            })
        end
    }
}

// After: Only create needed components
system Spawn {
    query(SpawnPoint)
    {
        for _, entity in ipairs(entities) do
            local newEntity = ecr:add({
                Position = {x = 0, y = 0},
                Velocity = {dx = 0, dy = 0}
            })
            // Add other components only when needed
        end
    }
}
```

## Common Problems and Solutions

### Entity Relationships Not Working

If entity relationships aren't working:
1. Check that both entities exist
2. Verify the relationship components are properly set
3. Ensure systems are querying the correct components

Example fix:
```ejecs
system UpdateParentChild {
    query(Parent, Position)
    {
        for _, entity in ipairs(entities) do
            local parent = entity.Parent
            if parent.child then
                local child = parent.child
                if child.Position then  // Check if child exists and has Position
                    child.Position.x = entity.Position.x
                    child.Position.y = entity.Position.y
                end
            end
        end
    }
}
```

### Systems Not Running

If systems aren't running:
1. Check system frequency
2. Verify component queries
3. Look for syntax errors

Example fix:
```ejecs
// Before: System might not run
system Update {
    query(Position)
    {
        // ...
    }
}

// After: Explicit frequency
system Update {
    query(Position)
    frequency(60)
    {
        // ...
    }
}
```

### Code Generation Issues

If generated code has problems:
1. Check the EJECS syntax
2. Verify type definitions
3. Look for platform-specific issues

Example fix:
```ejecs
// Before: Platform-specific code without checks
system Render {
    query(Position, Sprite)
    {
        for _, entity in ipairs(entities) do
            love.graphics.draw(entity.Sprite.image, entity.Position.x, entity.Position.y)
        end
    }
}

// After: Platform-specific code with checks
system Render {
    query(Position, Sprite)
    {
        for _, entity in ipairs(entities) do
            if PLATFORM == "love2d" then
                love.graphics.draw(entity.Sprite.image, entity.Position.x, entity.Position.y)
            else
                -- Other platform rendering code
            end
        end
    }
}
```

## Getting Help

If you're still having issues:
1. Check the [FAQ](FAQ)
2. Search existing issues
3. Create a new issue with:
   - EJECS code that reproduces the problem
   - Expected vs actual behavior
   - Error messages
   - Platform information 