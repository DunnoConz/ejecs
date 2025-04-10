<template>
  <div class="max-w-3xl mx-auto">
    <h1 class="text-4xl font-bold mb-8">Networking Guide</h1>
    
    <div class="prose prose-lg">
      <p class="mb-6">
        This guide covers how to use EJECS in multiplayer games, including state synchronization and client-server architecture.
      </p>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Client-Server Architecture</h2>
      <p class="mb-4">In a networked game, you typically have:</p>
      <ul class="list-disc pl-6 mb-6">
        <li>Server: Authoritative world state, runs game logic</li>
        <li>Clients: Local world state, prediction, and interpolation</li>
      </ul>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Component Replication</h2>
      <p class="mb-4">Mark components that need to be synchronized across the network:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- Server-side component registration
world:registerComponent("Position", "table", {
    replicate = true,  -- This component will be synced
    reliable = false   -- Use unreliable transport for position updates
})

world:registerComponent("Health", "number", {
    replicate = true,
    reliable = true    -- Use reliable transport for health updates
})</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Server Implementation</h2>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>local ReplicatedStorage = game:GetService("ReplicatedStorage")
local RemoteEvent = Instance.new("RemoteEvent")
RemoteEvent.Name = "WorldState"
RemoteEvent.Parent = ReplicatedStorage

local function NetworkSystem(world)
    -- Collect replicated component changes
    local changes = {}
    for id, pos, health in world:query({"Position", "Health"}) do
        if world:hasChanged(id, "Position") or world:hasChanged(id, "Health") then
            changes[id] = {
                position = pos,
                health = health
            }
        end
    end
    
    -- Send changes to all clients
    if next(changes) then
        RemoteEvent:FireAllClients(changes)
    end
end

world:addSystem(NetworkSystem)</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Client Implementation</h2>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>local ReplicatedStorage = game:GetService("ReplicatedStorage")
local RemoteEvent = ReplicatedStorage:WaitForChild("WorldState")

local function ClientNetworkSystem(world)
    -- Handle incoming state updates
    RemoteEvent.OnClientEvent:Connect(function(changes)
        for entityId, data in pairs(changes) do
            -- Create entity if it doesn't exist
            if not world:exists(entityId) then
                world:spawn(entityId)  -- Use same ID as server
            end
            
            -- Update components
            world:set(entityId, "Position", data.position)
            world:set(entityId, "Health", data.health)
        end
    end)
end

-- Add client-side prediction
local function PredictionSystem(world)
    for id, pos, vel in world:query({"Position", "Velocity"}) do
        -- Predict next position
        local predicted = {
            x = pos.x + vel.x,
            y = pos.y + vel.y,
            z = pos.z + vel.z
        }
        
        -- Store prediction
        world:add(id, "PredictedPosition", predicted)
    end
end</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">State Reconciliation</h2>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>local function ReconciliationSystem(world)
    for id, serverPos, predictedPos in world:query({"Position", "PredictedPosition"}) do
        -- Calculate difference between predicted and actual
        local error = {
            x = math.abs(serverPos.x - predictedPos.x),
            y = math.abs(serverPos.y - predictedPos.y),
            z = math.abs(serverPos.z - predictedPos.z)
        }
        
        -- If error is too large, correct the position
        if error.x > ERROR_THRESHOLD then
            world:set(id, "Position", serverPos)
            -- Optionally interpolate to new position
        end
    end
end</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Best Practices</h2>
      <ul class="list-disc pl-6 mb-6">
        <li>Minimize network traffic by only sending changed components</li>
        <li>Use client-side prediction for smooth movement</li>
        <li>Implement state reconciliation to handle network latency</li>
        <li>Consider bandwidth and choose appropriate update rates</li>
        <li>Use reliable networking for critical game state</li>
      </ul>

      <div class="mt-8 p-4 bg-blue-50 rounded-md">
        <h3 class="text-xl font-semibold mb-2">💡 Pro Tip</h3>
        <p>
          Consider using delta compression when sending updates to reduce bandwidth usage.
          Only send the differences between the current and previous state.
        </p>
      </div>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Next Steps</h2>
      <p class="mb-4">
        Learn how to test your networked ECS implementation in the 
        <NuxtLink to="/docs/advanced/testing" class="text-blue-600 hover:text-blue-800">Testing Guide</NuxtLink>.
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