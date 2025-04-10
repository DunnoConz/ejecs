<template>
  <div class="max-w-3xl mx-auto">
    <h1 class="text-4xl font-bold mb-8">API Reference</h1>
    
    <div class="prose prose-lg">
      <p class="mb-6">
        Complete documentation of all EJECS functions, methods, and types.
      </p>

      <h2 class="text-2xl font-semibold mt-8 mb-4">World</h2>
      <p class="mb-4">The main container for your ECS game state.</p>

      <h3 class="text-xl font-medium mt-6 mb-3">Constructor</h3>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>local world = EJECS.World.new()</code></pre>

      <h3 class="text-xl font-medium mt-6 mb-3">Methods</h3>
      
      <h4 class="text-lg font-medium mt-4 mb-2">Component Registration</h4>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- Register a new component type
world:registerComponent(name: string, type: string, options?: table)

-- Supported types:
-- "string" - String values
-- "number" - Numeric values
-- "boolean" - Boolean values
-- "table" - Complex data structures
-- "Instance" - Roblox instances

-- Options:
-- replicate: boolean - Whether to sync across network
-- reliable: boolean - Use reliable network transport</code></pre>

      <h4 class="text-lg font-medium mt-4 mb-2">Entity Management</h4>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- Create a new entity
local entityId = world:spawn()

-- Create entity with components
local entityId = world:spawnWith(components: table)

-- Check if entity exists
local exists = world:exists(entityId: number)

-- Destroy an entity and all its components
world:destroy(entityId: number)</code></pre>

      <h4 class="text-lg font-medium mt-4 mb-2">Component Operations</h4>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- Add a component to an entity
world:add(entityId: number, componentName: string, value: any)

-- Add multiple components
world:addMultiple(entityId: number, components: table)

-- Get a component value
local value = world:get(entityId: number, componentName: string)

-- Set a component value
world:set(entityId: number, componentName: string, value: any)

-- Remove a component
world:remove(entityId: number, componentName: string)

-- Remove multiple components
world:removeMultiple(entityId: number, componentNames: table)

-- Check if entity has component
local has = world:has(entityId: number, componentName: string)

-- Check if entity has all components
local hasAll = world:hasAll(entityId: number, componentNames: table)

-- Check if component value has changed
local changed = world:hasChanged(entityId: number, componentName: string)</code></pre>

      <h4 class="text-lg font-medium mt-4 mb-2">Querying</h4>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- Query entities with specific components
for entityId, componentA, componentB in world:query(componentNames: table) do
    -- Process entities
end

-- Example:
for id, position, velocity in world:query({"Position", "Velocity"}) do
    -- Update position based on velocity
end</code></pre>

      <h4 class="text-lg font-medium mt-4 mb-2">Systems</h4>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- Add a system
world:addSystem(system: function)

-- Remove a system
world:removeSystem(system: function)

-- Update all systems
world:update(deltaTime: number)</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Component Types</h2>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- String component
world:registerComponent("Name", "string")
world:add(entity, "Name", "Player")

-- Number component
world:registerComponent("Health", "number")
world:add(entity, "Health", 100)

-- Boolean component
world:registerComponent("IsEnemy", "boolean")
world:add(entity, "IsEnemy", true)

-- Table component
world:registerComponent("Position", "table")
world:add(entity, "Position", { x = 0, y = 0, z = 0 })

-- Instance component
world:registerComponent("Model", "Instance")
world:add(entity, "Model", workspace.PlayerModel)</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">System Function</h2>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- Basic system structure
local function ExampleSystem(world)
    -- System receives world as parameter
    for entityId, componentA, componentB in world:query({"ComponentA", "ComponentB"}) do
        -- Process components
    end
end

-- System with parameters
local function createParameterizedSystem(param)
    return function(world)
        -- Access param in closure
        for entityId, component in world:query({"Component"}) do
            -- Use param in processing
        end
    end
end</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Events</h2>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- Component change events
world.onComponentAdded(entityId: number, componentName: string)
world.onComponentRemoved(entityId: number, componentName: string)
world.onComponentChanged(entityId: number, componentName: string)

-- Entity events
world.onEntityCreated(entityId: number)
world.onEntityDestroyed(entityId: number)</code></pre>

      <div class="mt-8 p-4 bg-blue-50 rounded-md">
        <h3 class="text-xl font-semibold mb-2">💡 Pro Tip</h3>
        <p>
          Use the TypeScript-style type annotations in your code comments to make your code more maintainable
          and to help catch potential type errors early in development.
        </p>
      </div>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Next Steps</h2>
      <p class="mb-4">
        Now that you understand the API, check out the 
        <NuxtLink to="/docs/examples" class="text-blue-600 hover:text-blue-800">Examples</NuxtLink>
        to see how these functions are used in practice.
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