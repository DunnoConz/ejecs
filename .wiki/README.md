# EJECS: Entity Component System IDL

**EJECS** is an **Interface Definition Language (IDL)** designed specifically for defining **Entity Component System (ECS)** structures like components, systems, and relationships. It aims to provide a clear, concise, and type-safe way to describe your ECS architecture, primarily targeting **Luau** code generation for game engines like Roblox.

## Features

*   **Component Definition:** Define data components with typed fields.
*   **System Definition:** Describe systems with queries, parameters (basic parsing), frequency/priority (basic parsing), and embedded code blocks.
*   **Relationship Modeling:** Define parent-child or other relationships between components/entities (basic parsing).
*   **Type Safety:** Enforces type checking within the IDL.
*   **Luau Generation:** Generates Luau code structure based on the definitions (Generator implementation is ongoing).
*   **Table Type:** Supports map-like structures using `table<KeyType, ValueType>`.

## Installation

Ensure you have Go installed (version 1.18 or later recommended).

1.  Clone the repository:
    ```bash
    git clone https://github.com/your-org/ejecs.git
    cd ejecs
    ```
2.  Build the compiler:
    ```bash
    go build -o bin/ejecs ./cmd/ejecs
    ```
    This will create the `ejecs` executable in the `bin` directory.

## Quick Start / Usage

1.  **Create an EJECS file** (e.g., `definitions.ejecs`):

    ```coffeescript
    // definitions.ejecs

    component Position {
        number x;
        number y;
    }

    component Velocity {
        number dx = 0;
        number dy = 0;
    }

    component PlayerInfo {
        string name;
        boolean? admin; // Optional field
        table<string, number> stats;
    }

    // Example System
    system Movement {
        query(Position, Velocity) // Which components the system operates on
        // params { number deltaTime; } // Params block parsing TBD
        // frequency: 60 // Basic parsing, complex values skipped
        // priority: 1   // Basic parsing
        {
            // Luau code block (treated as raw string for now)
            for entityId, pos, vel in world:query(Position, Velocity) do
                pos.x = pos.x + vel.dx -- * deltaTime -- Assuming deltaTime available
                pos.y = pos.y + vel.dy -- * deltaTime
            end
        }
    }

    // Example Relationship (basic parsing)
    @parent
    relationship ChildOf {
        child: EntityA
        parent: EntityB
    }
    ```

2.  **Generate Luau Code:**

    Run the compiler, specifying your input `.ejecs` file and the desired output `.luau` file:

    ```bash
    ./bin/ejecs -input definitions.ejecs -output generated_ecs.luau
    ```

3.  **Use in your Project:** Integrate the generated Luau code (`generated_ecs.luau`) into your game engine environment.

## Language Features

### Component Definition

Components define the data associated with entities.

```coffeescript
component Health {
    number current = 100;
    number max = 100;
}

component Tags {
    table<string, boolean> list; // Example using table
}
```

### System Definition

Systems contain the logic that operates on entities with specific components.

```coffeescript
system Regeneration {
    query(Health)       // Target entities with Health component
    frequency: 10       // How often to run (basic value parsing)
    priority: 5         // Execution order (basic value parsing)
    {
        // Luau code to execute
        for id, health in world:query(Health) do
            if health.current < health.max then
                health.current = health.current + 1
            end
        end
    }
}
```
*(Note: Parsing for `params`, complex `frequency`, and code block integration is still under development)*

### Relationship Definition

Define relations between component types (basic parsing).

```coffeescript
@many_to_one // Optional type annotation
relationship BelongsTo {
    child: Item
    parent: Inventory
}
```

## Development Status

EJECS is currently under active development.

*   **Parser:** Mostly complete for core syntax (components, systems, relationships, table types). Default value and complex frequency/priority parsing is basic.
*   **Generator:** Basic structure exists, but Luau code generation logic needs implementation.
*   **Error Handling:** Basic error reporting is in place.

## Contributing

Contributions are welcome! Please refer to `CONTRIBUTING.md` for guidelines.

## License

This project is licensed under the [MIT License](LICENSE). 