# Component System

## Overview

The EJECS component system provides a way to define data structures that can be attached to entities. Components are pure data containers with no behavior.

## Component Definition

Components are defined using the `component` keyword followed by a name and a block of fields:

```ejecs
component Position {
    x: number;
    y: number;
}
```

## Field Types

### Basic Types
- `number`: Floating-point numbers
- `string`: Text strings
- `boolean`: True/false values
- `int`: Integer values

### Field Modifiers
- Optional fields: Use `?` suffix
  ```ejecs
  component Transform {
      scale: number?;  // Optional field
  }
  ```
- Arrays: Use `[]` suffix
  ```ejecs
  component Inventory {
      items: string[];  // Array of strings
  }
  ```

### Default Values
Fields can have default values:
```ejecs
component Health {
    current: number = 100;
    max: number = 100;
}
```

## Field Syntax

Each field declaration must end with a semicolon (;) and follows this pattern:
```ejecs
fieldName: type [= defaultValue];
```

## Component Usage in Systems

Components are referenced in system queries:
```ejecs
system Movement {
    query(Position, Velocity)  // References Position and Velocity components
    {
        // System implementation
    }
}
```

## Best Practices

1. **Naming**
   - Use PascalCase for component names
   - Use camelCase for field names
   - Names should be descriptive and clear

2. **Organization**
   - Keep components focused on a single aspect
   - Group related data in the same component
   - Split large components into smaller, more focused ones

3. **Types**
   - Use the most appropriate type for each field
   - Consider using optional fields for non-required data
   - Use arrays when dealing with collections

4. **Documentation**
   - Comment complex fields or non-obvious usage
   - Document any assumptions or constraints
   - Explain relationships with other components 