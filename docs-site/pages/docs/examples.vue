<template>
  <div class="max-w-3xl mx-auto">
    <h1 class="text-4xl font-bold mb-8">Examples</h1>
    
    <div class="prose prose-lg">
      <p class="mb-6">
        Learn from these practical examples that demonstrate how to use EJECS in real game scenarios.
        Each example showcases different aspects of the ECS architecture.
      </p>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Basic Character Controller</h2>
      <p class="mb-4">A simple character controller with movement and input handling:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- Components
world:registerComponent("Position", "table")
world:registerComponent("Velocity", "table")
world:registerComponent("Input", "table")
world:registerComponent("Player", "boolean")

-- Create player entity
local player = world:spawnWith({
    Position = { x = 0, y = 0, z = 0 },
    Velocity = { x = 0, y = 0, z = 0 },
    Input = { forward = 0, right = 0 },
    Player = true
})

-- Input system
local function InputSystem(world)
    for id, input in world:query({"Input"}) do
        -- Get input state
        local userInput = game:GetService("UserInputService")
        
        -- Update input component
        input.forward = 0
        input.right = 0
        
        if userInput:IsKeyDown(Enum.KeyCode.W) then
            input.forward = 1
        elseif userInput:IsKeyDown(Enum.KeyCode.S) then
            input.forward = -1
        end
        
        if userInput:IsKeyDown(Enum.KeyCode.D) then
            input.right = 1
        elseif userInput:IsKeyDown(Enum.KeyCode.A) then
            input.right = -1
        end
    end
end

-- Movement system
local function MovementSystem(world)
    for id, pos, vel, input in world:query({"Position", "Velocity", "Input"}) do
        -- Update velocity based on input
        vel.x = input.right * 5
        vel.z = input.forward * 5
        
        -- Update position
        pos.x = pos.x + vel.x
        pos.z = pos.z + vel.z
    end
end

world:addSystem(InputSystem)
world:addSystem(MovementSystem)</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Combat System</h2>
      <p class="mb-4">A combat system with health, damage, and death handling:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- Components
world:registerComponent("Health", "number")
world:registerComponent("MaxHealth", "number")
world:registerComponent("Damage", "number")
world:registerComponent("Dead", "boolean")
world:registerComponent("Team", "string")

-- Create entities
local player = world:spawnWith({
    Health = 100,
    MaxHealth = 100,
    Team = "Player"
})

local enemy = world:spawnWith({
    Health = 50,
    MaxHealth = 50,
    Team = "Enemy",
    Damage = 10
})

-- Damage system
local function DamageSystem(world)
    for id, health, damage in world:query({"Health", "Damage"}) do
        -- Apply damage
        health = health - damage
        
        -- Check for death
        if health <= 0 then
            world:add(id, "Dead", true)
            world:remove(id, "Health")
        else
            world:set(id, "Health", health)
        end
        
        -- Remove processed damage
        world:remove(id, "Damage")
    end
end

-- Death system
local function DeathSystem(world)
    for id, _ in world:query({"Dead"}) do
        -- Handle death effects
        local position = world:get(id, "Position")
        if position then
            -- Spawn death particles
            createDeathEffect(position)
        end
        
        -- Clean up entity
        world:destroy(id)
    end
end

world:addSystem(DamageSystem)
world:addSystem(DeathSystem)</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Inventory System</h2>
      <p class="mb-4">An inventory system with items and equipment:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- Components
world:registerComponent("Inventory", "table")
world:registerComponent("Equipment", "table")
world:registerComponent("Item", "table")
world:registerComponent("Pickupable", "boolean")

-- Create player with inventory
local player = world:spawnWith({
    Inventory = {
        items = {},
        capacity = 10
    },
    Equipment = {
        weapon = nil,
        armor = nil
    }
})

-- Create an item in the world
local sword = world:spawnWith({
    Item = {
        name = "Iron Sword",
        type = "weapon",
        damage = 15
    },
    Position = { x = 10, y = 0, z = 10 },
    Pickupable = true
})

-- Pickup system
local function PickupSystem(world)
    for playerId, inventory, playerPos in world:query({"Inventory", "Position"}) do
        for itemId, item, itemPos in world:query({"Item", "Position", "Pickupable"}) do
            -- Check if player is near item
            local distance = calculateDistance(playerPos, itemPos)
            if distance < 2 then
                -- Add item to inventory
                if #inventory.items < inventory.capacity then
                    table.insert(inventory.items, item)
                    world:destroy(itemId)
                end
            end
        end
    end
end

-- Equipment system
local function EquipmentSystem(world)
    for id, equipment, inventory in world:query({"Equipment", "Inventory"}) do
        -- Handle equipment changes
        if equipment.pendingEquip then
            local item = equipment.pendingEquip
            if item.type == "weapon" then
                equipment.weapon = item
            elseif item.type == "armor" then
                equipment.armor = item
            end
            equipment.pendingEquip = nil
        end
    end
end

world:addSystem(PickupSystem)
world:addSystem(EquipmentSystem)</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">AI Behavior System</h2>
      <p class="mb-4">A simple AI system with different behavior states:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- Components
world:registerComponent("AI", "table")
world:registerComponent("Target", "number")
world:registerComponent("PatrolPath", "table")

-- Create an AI entity
local enemy = world:spawnWith({
    Position = { x = 0, y = 0, z = 0 },
    AI = {
        state = "patrol",
        alertness = 0,
        lastSeenTarget = nil
    },
    PatrolPath = {
        points = {
            { x = 0, y = 0, z = 0 },
            { x = 10, y = 0, z = 0 },
            { x = 10, y = 0, z = 10 },
            { x = 0, y = 0, z = 10 }
        },
        currentPoint = 1
    }
})

-- AI State system
local function AISystem(world)
    for id, ai, pos in world:query({"AI", "Position"}) do
        -- Update AI based on current state
        if ai.state == "patrol" then
            local path = world:get(id, "PatrolPath")
            if path then
                -- Move to next patrol point
                local target = path.points[path.currentPoint]
                moveTowards(pos, target)
                
                -- Check if reached point
                if reachedPosition(pos, target) then
                    path.currentPoint = (path.currentPoint % #path.points) + 1
                end
            end
        elseif ai.state == "chase" then
            if ai.lastSeenTarget then
                local targetPos = world:get(ai.lastSeenTarget, "Position")
                if targetPos then
                    -- Chase target
                    moveTowards(pos, targetPos)
                else
                    -- Lost target, return to patrol
                    ai.state = "patrol"
                end
            end
        end
        
        -- Check for targets
        for targetId, targetPos in world:query({"Position", "Player"}) do
            if canSeeTarget(pos, targetPos) then
                ai.state = "chase"
                ai.lastSeenTarget = targetId
                break
            end
        end
    end
end

world:addSystem(AISystem)</code></pre>

      <div class="mt-8 p-4 bg-blue-50 rounded-md">
        <h3 class="text-xl font-semibold mb-2">💡 Pro Tip</h3>
        <p>
          These examples can be combined and extended to create more complex game mechanics.
          Try mixing and matching different systems to create your own unique gameplay features.
        </p>
      </div>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Next Steps</h2>
      <p class="mb-4">
        Now that you've seen some practical examples, learn about 
        <NuxtLink to="/docs/advanced/optimization" class="text-blue-600 hover:text-blue-800">optimizing your ECS implementation</NuxtLink>
        for better performance.
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