# EJECS Language Reference

## Overview

EJECS is an embeddable Interface Definition Language (IDL) designed for defining Entity Component Systems. It focuses on generating clean, type-safe code with Luau (Roblox/Lune) as its primary target.

## Type System

### Basic Types
| EJECS Type | Luau Type | Default Value | Description |
|------------|-----------|---------------|-------------|
| number     | number    | 0            | Floating-point number |
| string     | string    | ""           | Text string |
| boolean    | boolean   | false        | True/false value |
| int        | number    | 0            | Integer value |
| Vector2    | Vector2   | Vector2.new()| 2D vector |
| Vector3    | Vector3   | Vector3.new()| 3D vector |
| CFrame     | CFrame    | CFrame.new() | Coordinate frame |
| Color3     | Color3    | Color3.new() | RGB color |
| ColorSequence | ColorSequence | ColorSequence.new() | Color gradient |
| NumberRange | NumberRange | NumberRange.new() | Number range |
| NumberSequence | NumberSequence | NumberSequence.new() | Number sequence |
| UDim      | UDim      | UDim.new()   | Dimension with scale and offset |
| UDim2     | UDim2     | UDim2.new()  | 2D dimension |
| Ray       | Ray       | Ray.new()    | 3D ray |
| Region3   | Region3   | Region3.new()| 3D region |
| Region3int16 | Region3int16 | Region3int16.new() | 3D region with 16-bit integers |
| Rect      | Rect      | Rect.new()   | 2D rectangle |
| Instance  | Instance  | nil          | Roblox instance reference |
| EnumItem  | EnumItem  | nil          | Enum value |
| BrickColor | BrickColor | BrickColor.new() | Brick color |

### Type Modifiers
- Optional: `?` suffix (e.g., `Vector3?`)
- Array: `[]` suffix (e.g., `Vector3[]`)
- Dictionary: `{[K]: V}` where K and V are types (e.g., `{[string]: number}`)
- Union: `|` operator (e.g., `string | number`)

### Roblox-Specific Types
```ejecs
component UIElement {
    UDim2 position;
    UDim2 size;
    Color3 backgroundColor;
    number transparency = 0;
    boolean visible = true;
}

component ParticleEmitter {
    ColorSequence color;
    NumberRange size;
    NumberRange speed;
    number lifetime = 1;
    number rate = 10;
    boolean enabled = true;
}

component Raycast {
    Ray ray;
    number maxDistance = 100;
    RaycastResult? result;  // Optional result
}

component Region {
    Region3 bounds;
    boolean isInside(Vector3 point) {
        return bounds:FindPointInRegion(point)
    }
}
```

## Components

Components are defined using the `component` keyword with type-before-name syntax:

```ejecs
component Transform {
    Vector3 position;
    CFrame orientation;
    Vector3? scale;  // Optional field
}

component Health {
    number maxHealth = 100;
    number currentHealth = 100;
    boolean isInvulnerable = false;
}

component Inventory {
    int capacity = 10;
    {[string]: int} items;  // Dictionary of item counts
    Instance[] equipped;    // Array of equipped items
}
```

## Systems

Systems are defined using the `system` keyword:

```ejecs
system Movement {
    query(Transform, Physics)
    params {
        number deltaTime;
        Vector3 gravity = Vector3.new(0, -9.81, 0);
    }
    frequency: fixed(60)
    priority: 100
    {
        // System implementation in Luau
        for _, entity in ipairs(entities) do
            local transform = entity.Transform
            local physics = entity.Physics
            
            transform.position = transform.position + physics.velocity * deltaTime
            physics.velocity = physics.velocity + gravity * deltaTime
        end
    }
}
```

## Embedding

EJECS can be embedded in Luau projects using the provided API:

```lua
local EJECS = require("ejecs")

-- Parse EJECS definitions
local world = EJECS.parse([[
    component Position {
        Vector3 value;
    }
    
    system UpdatePosition {
        query(Position)
        {
            -- System logic
        }
    }
]])

-- Generate Luau code
local code = EJECS.generate(world, {
    target = "luau",
    module = "game",
    namespace = "Components"
})
```

## Code Generation

The generated Luau code includes:
- Type definitions using Luau's type system
- Runtime component registration
- System implementations
- Type checking and validation

Example output:
```lua
--!strict
local Types = require(game.Types)

export type Position = {
    value: Vector3
}

local Components = {
    Position = {
        name = "Position",
        new = function(): Position
            return {
                value = Vector3.new(0, 0, 0)
            }
        end,
        validate = function(data: any): boolean
            return typeof(data.value) == "Vector3"
        end
    }
}

return Components
```

## Best Practices

1. **Type Safety**
   - Use specific types instead of generic ones
   - Leverage Luau's type system
   - Define clear component interfaces

2. **Performance**
   - Use appropriate data structures
   - Consider Roblox-specific optimizations
   - Profile generated code

3. **Integration**
   - Follow Roblox coding standards
   - Use consistent naming conventions
   - Document component dependencies

## Relationships

Relationships are defined using the `relationship` keyword:

```ejecs
relationship ChildOf {
    // Relationship definition
}
```

## Operators and Expressions

The following operators are supported:

- Arithmetic: `+`, `-`, `*`, `/`
- Assignment: `=`, `+=`
- Comparison: `==`, `!=`, `<`, `<=`, `>`, `>=`
- Logical: `&&`, `||`

## Comments

Single-line comments use `//`:

```ejecs
// This is a comment
component Example {
    // Field comment
    value: number;
}
```

## Code Generation

EJECS supports code generation for multiple target libraries:
- ECR (Entity Component Runtime)
- JECS (Just Enough Component System)

Use the command line tool to generate code:

```bash
ejecs -input input.ejecs -output output.lua -library ecr
``` 