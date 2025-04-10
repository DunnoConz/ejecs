<template>
  <div class="max-w-3xl mx-auto">
    <h1 class="text-4xl font-bold mb-8">Quick Start Guide</h1>
    
    <div class="prose prose-lg">
      <p class="mb-6">
        Let's create a simple game using EJECS to demonstrate the basics of the Entity Component System IDL.
        We'll build a basic player movement system with position and velocity components.
      </p>

      <h2 class="text-2xl font-semibold mt-8 mb-4">1. Create a New Project</h2>
      <p class="mb-4">Initialize a new EJECS project:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>ejecs init my-game
cd my-game</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">2. Define Components</h2>
      <p class="mb-4">Create a file named <code>components.jecs</code>:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>@replicated
component Position {
    x: number
    y: number
    z: number
}

@replicated
component Velocity {
    x: number
    y: number
    z: number
}

@replicated
component PlayerInput {
    moveDirection: Vector3
    isJumping: boolean
}</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">3. Define Systems</h2>
      <p class="mb-4">Create a file named <code>systems.jecs</code>:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>@server
system Movement {
    using Position, Velocity
    frequency: 60Hz
    priority: 1
}

@client
system InputHandler {
    using PlayerInput
    frequency: 60Hz
    priority: 0
}</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">4. Generate Code</h2>
      <p class="mb-4">Compile your EJECS files to generate Luau code:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>ejecs compile *.jecs --out src/generated</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">5. Implement System Logic</h2>
      <p class="mb-4">Create the movement system implementation in <code>src/systems/Movement.lua</code>:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>local Types = require(game:GetService("ReplicatedStorage").Generated.Types)

local MovementSystem = {}

function MovementSystem.new()
    return {
        update = function(world, dt)
            for id, pos, vel in world:query({"Position", "Velocity"}) do
                pos.x += vel.x * dt
                pos.y += vel.y * dt
                pos.z += vel.z * dt
            end
        end
    }
end

return MovementSystem</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">6. Set Up the Game</h2>
      <p class="mb-4">Create your main game script:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>local ReplicatedStorage = game:GetService("ReplicatedStorage")
local World = require(ReplicatedStorage.Packages.EJECS).World

local world = World.new()

-- Register systems
world:addSystem(require(script.Systems.Movement).new())
world:addSystem(require(script.Systems.InputHandler).new())

-- Create a player entity
local player = world:createEntity({
    Position = { x = 0, y = 0, z = 0 },
    Velocity = { x = 0, y = 0, z = 0 },
    PlayerInput = { moveDirection = Vector3.new(), isJumping = false }
})</code></pre>

      <div class="mt-8 p-4 bg-blue-50 rounded-md">
        <h3 class="text-xl font-semibold mb-2">💡 Pro Tip</h3>
        <p>
          Use the <code>ejecs watch</code> command during development to automatically recompile your EJECS files
          when they change.
        </p>
      </div>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Next Steps</h2>
      <p class="mb-4">
        Learn more about component definitions in the 
        <NuxtLink to="/docs/components" class="text-blue-600 hover:text-blue-800">Components Guide</NuxtLink>.
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