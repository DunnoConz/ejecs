# EJECS API Documentation

## Table of Contents
- [World](#world)
- [Components](#components)
- [Systems](#systems)
- [Relationships](#relationships)
- [Types](#types)

## World

The World is the main container for your ECS. It manages all entities, components, and systems.

### Creation

```lua
local ejecs = require("ejecs")
local world = ejecs.World.new()
```

### Methods

#### `world:component()`
Creates a new component type.

```lua
local Position = world:component()
```

#### `world:createEntity()`
Creates a new entity and returns its ID.

```lua
local entityId = world:createEntity()
```

#### `world:addComponent(entity, component, data)`
Adds component data to an entity.

```lua
world:addComponent(entity, Position, {x = 0, y = 0, z = 0})
```

#### `world:get(entity, component)`
Gets component data for an entity. Returns nil if the component doesn't exist.

```lua
local positionData = world:get(entity, Position)
if positionData then
    print(positionData.x, positionData.y, positionData.z)
end
```

#### `world:addRelation(source, relationship, target)`
Creates a relationship between two entities.

```lua
world:addRelation(childEntity, ChildOf, parentEntity)
```

#### `world:getTarget(entity, relationship)`
Gets the target entity of a relationship. Returns nil if no relationship exists.

```lua
local parentId = world:getTarget(childEntity, ChildOf)
```

#### `world:system(config)`
Registers a new system.

```lua
world:system({
    name = "Movement",
    query = {Position, Velocity},
    callback = function(entity, components)
        -- System logic
    end
})
```

#### `world:update()`
Updates all systems once.

```lua
world:update()
```

## Components

Components are pure data containers. They define the properties that entities can have.

### Creating Components

```lua
-- Define the component
local Position = world:component()

-- Define its type (optional but recommended)
type PositionData = {
    x: number,
    y: number,
    z: number,
}
```

### Using Components

```lua
-- Add component to entity
world:addComponent(entity, Position, {
    x = 0,
    y = 0,
    z = 0,
})

-- Get component data
local position = world:get(entity, Position)
```

## Systems

Systems contain the logic that operates on entities with specific components.

### System Configuration

```lua
type SystemConfig = {
    name: string,
    query: {Component},
    callback: (entity: number, components: {[Component]: ComponentData}) -> ()
}
```

### Creating Systems

```lua
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
```

## Relationships

Relationships allow entities to be connected to other entities.

### Creating Relationships

```lua
-- Define a relationship type
local ChildOf = world:component()

-- Create a relationship
world:addRelation(childEntity, ChildOf, parentEntity)
```

### Querying Relationships

```lua
-- Get the target of a relationship
local parentId = world:getTarget(childEntity, ChildOf)

-- Use in systems
world:system({
    name = "PrintChildren",
    query = {Name, ChildOf},
    callback = function(entity, components)
        local target = world:getTarget(entity, ChildOf)
        print("Child:", entity, "Parent:", target)
    end
})
```

## Types

EJECS is fully type-safe when used with Luau's strict type checking.

### Component Data Types

```lua
type PositionData = {
    x: number,
    y: number,
    z: number,
}

type VelocityData = {
    x: number,
    y: number,
    z: number,
}

type NameData = {
    value: string,
}
```

### System Types

```lua
type SystemConfig = {
    name: string,
    query: {Component},
    callback: (entity: number, components: {[Component]: ComponentData}) -> ()
}
```

### Type Safety

```lua
-- Type-safe component access
local position = world:get(entity, Position) :: PositionData
local velocity = world:get(entity, Velocity) :: VelocityData

-- Type-safe system callbacks
world:system({
    name = "Movement",
    query = {Position, Velocity},
    callback = function(entity: number, components: {[any]: any})
        local position = components[Position] :: PositionData
        local velocity = components[Velocity] :: VelocityData
        -- Type-safe operations
        position.x += velocity.x
    end
}) 