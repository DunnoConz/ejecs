--!strict
-- EJECS Relationships Example
local ejecs = require("../src/init")
local world = ejecs.World.new()

-- Components
local Position = world:component()
local Velocity = world:component()
local Name = world:component()
local ChildOf = world:component()

-- Component Types
type PositionData = {
    x: number,
    y: number,
    z: number,
}

type VelocityData = {
    x: number,
    y: number,
    z: number,
}

type NameData = {
    value: string,
}

-- Helper Functions
local function getName(entity: number): string
    local nameComponent = world:get(entity, Name)
    if nameComponent then
        return nameComponent.value
    end
    return tostring(entity)
end

-- Entity Creation
local function createEntity(name: string, x: number, y: number, z: number, vx: number, vy: number, vz: number): number
    local entity = world:createEntity()
    world:addComponent(entity, Position, {x = x, y = y, z = z})
    world:addComponent(entity, Velocity, {x = vx, y = vy, z = vz})
    world:addComponent(entity, Name, {value = name})
    return entity
end

-- Create example entities
local parent = createEntity("Parent", 0, 0, 0, 1, 1, 1)
local child = createEntity("Child", 10, 10, 10, -1, -1, -1)
world:addRelation(child, ChildOf, parent)

-- System: Movement
world:system({
    name = "Movement",
    query = {
        Position,
        Velocity,
    },
    callback = function(entity: number, components: {[any]: any})
        local position = components[Position]
        local velocity = components[Velocity]
        
        if not position or not velocity then
            return
        end

        position.x += velocity.x
        position.y += velocity.y
        position.z += velocity.z

        print(string.format("Entity %d moved to (%.1f, %.1f, %.1f)", 
            entity, position.x, position.y, position.z))
    end
})

-- System: PrintChildren
world:system({
    name = "PrintChildren",
    query = {
        Name,
        ChildOf,
    },
    callback = function(entity: number, components: {[any]: any})
        local name = components[Name]
        if not name then return end

        -- Get relationship targets
        local childofTarget = world:getTarget(entity, ChildOf)
        
        if not childofTarget then
            print(string.format("Entity %s has ChildOf component but no target", getName(entity)))
            return
        end

        print(string.format("Relationship: %s is a child of %s", name.value, getName(childofTarget)))
    end
})

-- Main loop for testing
print("Starting simulation...")

print("\nInitial state:")
world:update() -- This will run both systems once

print("\nRunning 5 frames:")
for i = 1, 5 do
    print(string.format("\nFrame %d:", i))
    world:update() -- This will run both systems once
end 