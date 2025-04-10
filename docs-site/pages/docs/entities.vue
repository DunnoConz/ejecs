<template>
  <div class="max-w-3xl mx-auto">
    <h1 class="text-4xl font-bold mb-8">Entities</h1>
    
    <div class="prose prose-lg">
      <p class="mb-6">
        Entities are the game objects in your EJECS world. They are lightweight identifiers that serve as containers for components,
        allowing you to compose complex game objects from simple data components.
      </p>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Creating Entities</h2>
      <p class="mb-4">Creating a new entity is simple:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- Spawn a new entity
local entityId = world:spawn()

-- Spawn and immediately add components
local playerId = world:spawnWith({
    Position = { x = 0, y = 0, z = 0 },
    Health = 100,
    Name = "Player"
})</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Working with Entities</h2>
      
      <h3 class="text-xl font-medium mt-6 mb-3">Adding Components</h3>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- Add individual components
world:add(entityId, "Position", { x = 0, y = 0, z = 0 })
world:add(entityId, "Health", 100)

-- Add multiple components at once
world:addMultiple(entityId, {
    Velocity = { x = 0, y = 0, z = 0 },
    Damage = 10,
    IsEnemy = true
})</code></pre>

      <h3 class="text-xl font-medium mt-6 mb-3">Checking Components</h3>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- Check if an entity has a specific component
if world:has(entityId, "Health") then
    -- Entity has health component
end

-- Check for multiple components
if world:hasAll(entityId, {"Position", "Velocity"}) then
    -- Entity has both position and velocity
end</code></pre>

      <h3 class="text-xl font-medium mt-6 mb-3">Removing Components</h3>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- Remove a single component
world:remove(entityId, "Damage")

-- Remove multiple components
world:removeMultiple(entityId, {"Velocity", "IsEnemy"})</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Destroying Entities</h2>
      <p class="mb-4">When an entity is no longer needed, it can be destroyed:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- Destroy an entity and all its components
world:destroy(entityId)</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Entity Patterns</h2>
      
      <h3 class="text-xl font-medium mt-6 mb-3">Entity Templates</h3>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- Create a template function for common entity types
local function createEnemy(world, x, y, z)
    return world:spawnWith({
        Position = { x = x, y = y, z = z },
        Health = 50,
        Damage = 5,
        IsEnemy = true,
        AI = { state = "patrol" }
    })
end

-- Use the template
local enemy1 = createEnemy(world, 0, 0, 0)
local enemy2 = createEnemy(world, 10, 0, 10)</code></pre>

      <h3 class="text-xl font-medium mt-6 mb-3">Entity Relationships</h3>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- Create parent-child relationships
local parent = world:spawn()
local child = world:spawn()

world:add(child, "Parent", parent)
world:add(parent, "Children", { child })</code></pre>

      <div class="mt-8 p-4 bg-blue-50 rounded-md">
        <h3 class="text-xl font-semibold mb-2">💡 Pro Tip</h3>
        <p>
          Keep your entities lightweight and focused. Instead of creating complex inheritance hierarchies,
          use composition through components to build up entity functionality.
        </p>
      </div>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Best Practices</h2>
      <ul class="list-disc pl-6 mb-6">
        <li>Use entity templates for consistent entity creation</li>
        <li>Clean up entities properly when they're no longer needed</li>
        <li>Consider using tags (boolean components) for simple categorization</li>
        <li>Keep track of important entities (like the player) using dedicated components</li>
        <li>Use meaningful naming conventions for entity templates and categories</li>
      </ul>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Next Steps</h2>
      <p class="mb-4">
        Learn how to organize and manage your entities effectively with 
        <NuxtLink to="/docs/components" class="text-blue-600 hover:text-blue-800">Components</NuxtLink> and
        <NuxtLink to="/docs/systems" class="text-blue-600 hover:text-blue-800">Systems</NuxtLink>.
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