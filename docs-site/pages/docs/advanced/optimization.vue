<template>
  <div class="max-w-3xl mx-auto">
    <h1 class="text-4xl font-bold mb-8">Optimization Guide</h1>
    
    <div class="prose prose-lg">
      <p class="mb-6">
        Performance is crucial in game development. This guide covers best practices and techniques for optimizing your EJECS implementation.
      </p>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Query Optimization</h2>
      <p class="mb-4">Queries are one of the most performance-critical aspects of an ECS system. Here's how to optimize them:</p>

      <h3 class="text-xl font-medium mt-6 mb-3">Component Order</h3>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- Order components from least to most common
for id, rare, common in world:query({"RareComponent", "CommonComponent"}) do
    -- This is more efficient than querying CommonComponent first
end</code></pre>

      <h3 class="text-xl font-medium mt-6 mb-3">Query Caching</h3>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- Cache query results when appropriate
local function createCachedSystem()
    local queryCache = {}
    
    return function(world)
        -- Only update cache periodically or on specific events
        if shouldUpdateCache then
            queryCache = {}
            for id, comp in world:query({"ExpensiveComponent"}) do
                queryCache[id] = comp
            end
        end
        
        -- Use cached results
        for id, comp in pairs(queryCache) do
            -- Process cached data
        end
    end
end</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Component Design</h2>
      
      <h3 class="text-xl font-medium mt-6 mb-3">Data Layout</h3>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- Prefer flat structures over nested ones
-- Good
type Position = {
    x: number,
    y: number,
    z: number
}

-- Avoid
type Transform = {
    position: {
        x: number,
        y: number,
        z: number
    },
    rotation: {
        x: number,
        y: number,
        z: number
    }
}</code></pre>

      <h3 class="text-xl font-medium mt-6 mb-3">Component Granularity</h3>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- Split large components into smaller, focused ones
-- Instead of
type Character = {
    health: number,
    mana: number,
    position: Vector3,
    inventory: { [string]: number }
}

-- Use
type Health = number
type Mana = number
type Position = Vector3
type Inventory = { [string]: number }</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">System Optimization</h2>
      
      <h3 class="text-xl font-medium mt-6 mb-3">System Scheduling</h3>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- Group systems by update frequency
local function createTickSystem(interval)
    local accumulator = 0
    return function(world, dt)
        accumulator = accumulator + dt
        if accumulator >= interval then
            accumulator = accumulator - interval
            -- Run less frequent updates
        end
    end
end

world:addSystem(MovementSystem) -- Every frame
world:addSystem(createTickSystem(0.1)) -- Every 0.1 seconds
world:addSystem(createTickSystem(1.0)) -- Every second</code></pre>

      <h3 class="text-xl font-medium mt-6 mb-3">Batch Processing</h3>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>local function BatchProcessingSystem(world)
    local batch = {}
    local batchSize = 0
    
    -- Collect entities for processing
    for id, data in world:query({"ProcessData"}) do
        batch[batchSize + 1] = { id = id, data = data }
        batchSize = batchSize + 1
        
        -- Process batch when full
        if batchSize >= 100 then
            processBatch(world, batch)
            batch = {}
            batchSize = 0
        end
    end
    
    -- Process remaining entities
    if batchSize > 0 then
        processBatch(world, batch)
    end
end</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Memory Management</h2>
      <ul class="list-disc pl-6 mb-6">
        <li>Destroy entities and remove components when they're no longer needed</li>
        <li>Reuse entities through object pooling for frequently created/destroyed objects</li>
        <li>Use appropriate data structures for your component storage needs</li>
        <li>Consider using flags or bitfields for simple state tracking</li>
      </ul>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Profiling and Monitoring</h2>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>local function createProfilingSystem(name)
    return function(world)
        local startTime = os.clock()
        
        -- System logic here
        
        local endTime = os.clock()
        print(string.format("%s took %.3f ms", name, (endTime - startTime) * 1000))
    end
end

world:addSystem(createProfilingSystem("MovementSystem"))</code></pre>

      <div class="mt-8 p-4 bg-blue-50 rounded-md">
        <h3 class="text-xl font-semibold mb-2">💡 Pro Tip</h3>
        <p>
          Profile your systems regularly and focus optimization efforts on the most performance-critical parts of your game.
          Remember that premature optimization can lead to more complex, harder-to-maintain code.
        </p>
      </div>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Next Steps</h2>
      <p class="mb-4">
        Learn about other advanced topics like 
        <NuxtLink to="/docs/advanced/networking" class="text-blue-600 hover:text-blue-800">Networking</NuxtLink> and
        <NuxtLink to="/docs/advanced/testing" class="text-blue-600 hover:text-blue-800">Testing</NuxtLink>.
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