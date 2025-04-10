# EJECS Embedding API

## Overview

The EJECS embedding API allows you to parse and generate EJECS definitions directly within your Luau code. This is particularly useful for:
- Runtime code generation
- Dynamic component creation
- Integration with Roblox Studio
- Custom tooling and plugins

## API Reference

### EJECS Module

```lua
local EJECS = require("ejecs")
```

### Core Functions

#### `parse(source: string): World`
Parses EJECS source code and returns a World object.

```lua
local world = EJECS.parse([[
    component Position {
        Vector3 value;
    }
    
    system Movement {
        query(Position)
        {
            -- System logic
        }
    }
]])
```

#### `generate(world: World, options: GenerateOptions): string`
Generates Luau code from a World object.

```lua
local code = EJECS.generate(world, {
    target = "luau",
    module = "game.Components",
    strict = true,
    namespace = "Components"
})
```

### GenerateOptions

| Option | Type | Description |
|--------|------|-------------|
| target | string | Output target ("luau") |
| module | string | Module path for imports |
| strict | boolean | Enable strict type checking |
| namespace | string | Namespace for generated types |

### World Object

The World object provides access to parsed definitions:

```lua
local world = EJECS.parse(source)

-- Access components
for name, component in pairs(world.components) do
    print(name, component)
end

-- Access systems
for name, system in pairs(world.systems) do
    print(name, system)
end
```

## Integration Examples

### Roblox Studio Plugin

```lua
local EJECS = require(script.Parent.EJECS)

local function generateFromSelection()
    local selection = game:GetService("Selection"):Get()
    local source = ""
    
    -- Collect EJECS definitions from selected objects
    for _, instance in ipairs(selection) do
        if instance:IsA("StringValue") and instance.Name:match("%.ejecs$") then
            source = source .. instance.Value .. "\n"
        end
    end
    
    -- Generate code
    local world = EJECS.parse(source)
    local code = EJECS.generate(world, {
        target = "luau",
        module = "game.Components"
    })
    
    -- Create output
    local output = Instance.new("ModuleScript")
    output.Name = "GeneratedComponents"
    output.Source = code
    output.Parent = game:GetService("ReplicatedStorage")
end
```

### Runtime Component Registration

```lua
local EJECS = require("ejecs")
local ECS = require("ecs")

local function registerDynamicComponents(source)
    local world = EJECS.parse(source)
    local code = EJECS.generate(world, {
        target = "luau",
        module = "game.Components"
    })
    
    -- Execute generated code
    local fn, err = loadstring(code)
    if not fn then
        error("Failed to load generated code: " .. err)
    end
    
    local components = fn()
    
    -- Register with ECS
    for name, component in pairs(components) do
        ECS.registerComponent(name, component)
    end
end
```

### Custom Type Validation

```lua
local EJECS = require("ejecs")

local function validateComponent(component, data)
    local world = EJECS.parse([[
        component Validation {
            ]] .. component .. [[
        }
    ]])
    
    local code = EJECS.generate(world, {
        target = "luau",
        module = "game.Validation"
    })
    
    local fn, err = loadstring(code)
    if not fn then
        error("Failed to load validation code: " .. err)
    end
    
    local validation = fn()
    return validation.Validation.validate(data)
end
```

## Best Practices

1. **Error Handling**
   - Always validate parse results
   - Handle generation errors gracefully
   - Provide meaningful error messages

2. **Performance**
   - Cache parsed worlds when possible
   - Minimize runtime code generation
   - Use appropriate validation levels

3. **Security**
   - Validate input sources
   - Sanitize generated code
   - Use appropriate sandboxing

4. **Maintenance**
   - Version control EJECS definitions
   - Document custom types
   - Keep generated code up to date 