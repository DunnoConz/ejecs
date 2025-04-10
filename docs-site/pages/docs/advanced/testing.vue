<template>
  <div class="max-w-3xl mx-auto">
    <h1 class="text-4xl font-bold mb-8">Testing Guide</h1>
    
    <div class="prose prose-lg">
      <p class="mb-6">
        Testing is crucial for maintaining a reliable ECS implementation. This guide covers unit testing, integration testing,
        and debugging strategies for EJECS systems and components.
      </p>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Unit Testing Components</h2>
      <p class="mb-4">Test individual components and their interactions:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>local function testHealthComponent()
    local world = EJECS.World.new()
    world:registerComponent("Health", "number")
    
    local entity = world:spawn()
    world:add(entity, "Health", 100)
    
    assert(world:has(entity, "Health"), "Entity should have Health component")
    assert(world:get(entity, "Health") == 100, "Health should be 100")
    
    world:set(entity, "Health", 50)
    assert(world:get(entity, "Health") == 50, "Health should be updated to 50")
end</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Testing Systems</h2>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>local function testMovementSystem()
    local world = EJECS.World.new()
    world:registerComponent("Position", "table")
    world:registerComponent("Velocity", "table")
    
    -- Create test entity
    local entity = world:spawnWith({
        Position = { x = 0, y = 0, z = 0 },
        Velocity = { x = 1, y = 0, z = 0 }
    })
    
    -- Define system
    local function MovementSystem(world)
        for id, pos, vel in world:query({"Position", "Velocity"}) do
            pos.x = pos.x + vel.x
            pos.y = pos.y + vel.y
            pos.z = pos.z + vel.z
        end
    end
    
    -- Run system and verify results
    world:addSystem(MovementSystem)
    world:update(1)
    
    local pos = world:get(entity, "Position")
    assert(pos.x == 1, "Entity should move 1 unit in x direction")
    assert(pos.y == 0, "Entity should not move in y direction")
    assert(pos.z == 0, "Entity should not move in z direction")
end</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Integration Testing</h2>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>local function testCombatSystem()
    local world = EJECS.World.new()
    
    -- Register components
    world:registerComponent("Health", "number")
    world:registerComponent("Damage", "number")
    world:registerComponent("Dead", "boolean")
    
    -- Create test entities
    local player = world:spawnWith({
        Health = 100
    })
    
    local enemy = world:spawnWith({
        Health = 50,
        Damage = 20
    })
    
    -- Define systems
    local function DamageSystem(world)
        for id, health, damage in world:query({"Health", "Damage"}) do
            health = health - damage
            if health <= 0 then
                world:add(id, "Dead", true)
            end
            world:set(id, "Health", health)
        end
    end
    
    -- Run systems and check results
    world:addSystem(DamageSystem)
    world:update(1)
    
    assert(world:get(enemy, "Health") == 30, "Enemy health should be reduced")
    assert(not world:has(enemy, "Dead"), "Enemy should not be dead")
end</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Debugging Tools</h2>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>local function createDebugSystem(name)
    return function(world)
        print(string.format("=== %s Debug Info ===", name))
        
        -- Log entity counts
        local count = 0
        for _ in world:query({}) do
            count = count + 1
        end
        print(string.format("Total entities: %d", count))
        
        -- Log component states
        for id, health in world:query({"Health"}) do
            print(string.format("Entity %d: Health = %d", id, health))
        end
    end
end

world:addSystem(createDebugSystem("Combat"))</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Performance Testing</h2>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>local function benchmarkSystem(world, systemFunc, iterations)
    local times = {}
    
    for i = 1, iterations do
        local start = os.clock()
        systemFunc(world)
        local elapsed = os.clock() - start
        times[i] = elapsed
    end
    
    -- Calculate statistics
    local total = 0
    local max = times[1]
    local min = times[1]
    
    for _, time in ipairs(times) do
        total = total + time
        max = math.max(max, time)
        min = math.min(min, time)
    end
    
    return {
        average = total / iterations,
        max = max,
        min = min
    }
end</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Best Practices</h2>
      <ul class="list-disc pl-6 mb-6">
        <li>Write tests for all critical systems and components</li>
        <li>Use test fixtures to set up common test scenarios</li>
        <li>Test edge cases and error conditions</li>
        <li>Profile performance in realistic scenarios</li>
        <li>Use debugging tools to track system behavior</li>
      </ul>

      <div class="mt-8 p-4 bg-blue-50 rounded-md">
        <h3 class="text-xl font-semibold mb-2">💡 Pro Tip</h3>
        <p>
          Create helper functions for common test setup and assertions to make your tests more readable and maintainable.
          Consider using a test framework like TestEZ for more structured testing.
        </p>
      </div>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Next Steps</h2>
      <p class="mb-4">
        Check out the <NuxtLink to="/docs/examples" class="text-blue-600 hover:text-blue-800">Examples</NuxtLink> section
        for complete game implementations using EJECS.
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