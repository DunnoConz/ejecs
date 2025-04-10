# EJECS Best Practices

This guide outlines recommended practices when using EJECS in your projects.

## Component Design

### Keep Components Small and Focused

✅ Good:
```lua
-- Position only handles position
local Position = world:component()
type PositionData = {
    x: number,
    y: number,
    z: number,
}

-- Velocity only handles velocity
local Velocity = world:component()
type VelocityData = {
    x: number,
    y: number,
    z: number,
}
```

❌ Bad:
```lua
-- Mixing unrelated data
local MovementData = world:component()
type MovementData = {
    x: number,
    y: number,
    z: number,
    speed: number,
    health: number,  -- Unrelated to movement
    name: string,    -- Unrelated to movement
}
```

### Use Type Definitions

✅ Good:
```lua
type PositionData = {
    x: number,
    y: number,
    z: number,
}

local position = world:get(entity, Position) :: PositionData
```

❌ Bad:
```lua
local position = world:get(entity, Position)  -- Type inference only
```

## System Design

### Single Responsibility

✅ Good:
```lua
-- Movement system only handles movement
world:system({
    name = "Movement",
    query = {Position, Velocity},
    callback = function(entity, components)
        local position = components[Position]
        local velocity = components[Velocity]
        position.x += velocity.x
    end
})

-- Health system only handles health
world:system({
    name = "Health",
    query = {Health},
    callback = function(entity, components)
        local health = components[Health]
        health.value = math.min(health.value + health.regen, health.max)
    end
})
```

❌ Bad:
```lua
-- Mixing different responsibilities
world:system({
    name = "MovementAndHealth",
    query = {Position, Velocity, Health},
    callback = function(entity, components)
        -- Movement logic
        components[Position].x += components[Velocity].x
        
        -- Health logic (should be separate)
        components[Health].value += components[Health].regen
    end
})
```

### Query Only Required Components

✅ Good:
```lua
world:system({
    name = "Movement",
    query = {Position, Velocity},  -- Only what's needed
    callback = function(entity, components)
        local position = components[Position]
        local velocity = components[Velocity]
        position.x += velocity.x
    end
})
```

❌ Bad:
```lua
world:system({
    name = "Movement",
    query = {Position, Velocity, Health, Name},  -- Unnecessary components
    callback = function(entity, components)
        local position = components[Position]
        local velocity = components[Velocity]
        position.x += velocity.x
        -- Health and Name are never used
    end
})
```

## Relationship Design

### Use Meaningful Relationship Names

✅ Good:
```lua
local ChildOf = world:component()
local OwnedBy = world:component()
local AttachedTo = world:component()
```

❌ Bad:
```lua
local Rel1 = world:component()
local Rel2 = world:component()
```

### Check Relationship Validity

✅ Good:
```lua
world:system({
    name = "PrintChildren",
    query = {Name, ChildOf},
    callback = function(entity, components)
        local target = world:getTarget(entity, ChildOf)
        if not target then
            warn("Invalid child relationship for entity:", entity)
            return
        end
        -- Process valid relationship
    end
})
```

❌ Bad:
```lua
world:system({
    name = "PrintChildren",
    query = {Name, ChildOf},
    callback = function(entity, components)
        local target = world:getTarget(entity, ChildOf)
        print(world:get(target, Name).value)  -- Might crash
    end
})
```

## Performance Optimization

### Minimize System Updates

✅ Good:
```lua
-- Update only when needed
if gameState.isPlaying then
    world:update()
end
```

❌ Bad:
```lua
-- Updating every frame regardless
while true do
    world:update()
    task.wait()
end
```

### Efficient Component Access

✅ Good:
```lua
world:system({
    name = "Movement",
    callback = function(entity, components)
        local position = components[Position]  -- Cache component
        local velocity = components[Velocity]  -- Cache component
        
        position.x += velocity.x
        position.y += velocity.y
        position.z += velocity.z
    end
})
```

❌ Bad:
```lua
world:system({
    name = "Movement",
    callback = function(entity, components)
        -- Accessing components multiple times
        components[Position].x += components[Velocity].x
        components[Position].y += components[Velocity].y
        components[Position].z += components[Velocity].z
    end
})
```

## Error Handling

### Validate Component Data

✅ Good:
```lua
world:system({
    name = "Movement",
    callback = function(entity, components)
        local position = components[Position]
        local velocity = components[Velocity]
        
        if not position or not velocity then
            warn("Missing required components for entity:", entity)
            return
        end
        
        position.x += velocity.x
    end
})
```

❌ Bad:
```lua
world:system({
    name = "Movement",
    callback = function(entity, components)
        -- No validation
        components[Position].x += components[Velocity].x
    end
})
```

## Code Organization

### Group Related Components

✅ Good:
```lua
-- Movement components
local Position = world:component()
local Velocity = world:component()
local Acceleration = world:component()

-- Combat components
local Health = world:component()
local Damage = world:component()
local Defense = world:component()
```

❌ Bad:
```lua
-- Mixed/unorganized components
local Position = world:component()
local Health = world:component()
local Velocity = world:component()
local Damage = world:component()
```

### Document Component Types

✅ Good:
```lua
-- @type Position - Represents an entity's position in 3D space
type PositionData = {
    x: number,  -- X coordinate
    y: number,  -- Y coordinate
    z: number,  -- Z coordinate
}

-- @type Health - Represents an entity's health status
type HealthData = {
    current: number,  -- Current health value
    max: number,      -- Maximum health value
    regen: number,    -- Health regeneration per update
}
```

❌ Bad:
```lua
type PositionData = {x: number, y: number, z: number}
type HealthData = {current: number, max: number, regen: number}
``` 