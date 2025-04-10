<template>
  <div class="max-w-3xl mx-auto">
    <h1 class="text-4xl font-bold mb-8">Code Generation</h1>
    
    <div class="prose prose-lg">
      <p class="mb-6">
        EJECS generates type-safe, optimized code from your IDL definitions. The generated code includes
        type definitions, validation logic, network replication, and system templates that you can implement
        in your game code.
      </p>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Generation Process</h2>
      <p class="mb-4">EJECS follows these steps when generating code:</p>
      <ol class="list-decimal pl-6 mb-6">
        <li>Parse and validate IDL files</li>
        <li>Generate type definitions</li>
        <li>Create component storage and access code</li>
        <li>Build system templates</li>
        <li>Set up network replication</li>
        <li>Generate event handling code</li>
      </ol>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Generated Files</h2>
      <p class="mb-4">The compiler produces several files:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>generated/
  ├── Types.lua           -- Type definitions
  ├── Components/         -- Component implementations
  │   ├── Position.lua
  │   ├── Velocity.lua
  │   └── ...
  ├── Systems/           -- System templates
  │   ├── Movement.lua
  │   ├── Combat.lua
  │   └── ...
  ├── Events/           -- Event definitions and handlers
  │   ├── Damage.lua
  │   ├── Collision.lua
  │   └── ...
  └── Network/          -- Network replication code
      ├── Replication.lua
      └── Authority.lua</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Type Definitions</h2>
      <p class="mb-4">Generated type definitions include:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- Generated from EJECS definitions
export type Entity = number

export type Position = {
    x: number,
    y: number,
    z: number
}

export type ComponentStorage<T> = {
    get: (entity: Entity) -> T?,
    set: (entity: Entity, value: T) -> (),
    remove: (entity: Entity) -> (),
    has: (entity: Entity) -> boolean
}

export type SystemContext = {
    world: World,
    storage: {[string]: ComponentStorage<any>},
    events: EventDispatcher,
    network: NetworkManager
}</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Component Implementation</h2>
      <p class="mb-4">Generated component code includes storage and validation:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- Generated Position component
local Position = {}

function Position.new()
    return {
        storage = {},
        validate = function(value)
            return type(value) == "table"
                and type(value.x) == "number"
                and type(value.y) == "number"
                and type(value.z) == "number"
        end,
        
        get = function(self, entity)
            return self.storage[entity]
        end,
        
        set = function(self, entity, value)
            assert(self:validate(value), "Invalid Position data")
            self.storage[entity] = value
        end
    }
end

return Position</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">System Templates</h2>
      <p class="mb-4">Generated system templates ready for implementation:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- Generated Movement system template
local Types = require(script.Parent.Parent.Types)

local MovementSystem = {}

function MovementSystem.new()
    return {
        name = "Movement",
        priority = 1,
        frequency = 60,
        
        components = {
            "Position",
            "Velocity"
        },
        
        -- Implement this function
        update = function(context)
            local world = context.world
            local pos = context.storage.Position
            local vel = context.storage.Velocity
            
            -- Your implementation here
        end
    }
end

return MovementSystem</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Compiler Options</h2>
      <p class="mb-4">Configure code generation with compiler flags:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code># Basic compilation
ejecs compile *.jecs

# Generate with specific options
ejecs compile *.jecs \
    --out src/generated \
    --format luau \
    --optimize \
    --strict-types \
    --generate-tests</code></pre>

      <div class="mt-8 p-4 bg-blue-50 rounded-md">
        <h3 class="text-xl font-semibold mb-2">💡 Pro Tip</h3>
        <p>
          Use the <code>--watch</code> flag during development to automatically regenerate code when your
          IDL files change. This keeps your implementation in sync with your definitions.
        </p>
      </div>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Next Steps</h2>
      <p class="mb-4">
        Now that you understand how EJECS works, check out our 
        <NuxtLink to="/docs/examples" class="text-blue-600 hover:text-blue-800">Example Projects</NuxtLink>
        to see it in action.
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
// Add any required script logic here
</script>

<style scoped>
/* Add any scoped styles here */
</style> 