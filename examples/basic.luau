--!strict
-- EJECS Basic Example
local ejecs = require("../src/init")
local world = ejecs.World.new()

-- Define components
local Position = world:component()
local Velocity = world:component()

-- Define component types
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

-- Create an entity
local entity = world:createEntity()
world:addComponent(entity, Position, {x = 0, y = 0, z = 0})
world:addComponent(entity, Velocity, {x = 1, y = 2, z = 3})

-- Define a movement system
world:system({
    name = "Movement",
    query = {Position, Velocity},
    callback = function(entity: number, components: {[any]: any})
        local position = components[Position]
        local velocity = components[Velocity]
        
        position.x += velocity.x
        position.y += velocity.y
        position.z += velocity.z

        print(string.format("Entity %d moved to (%.1f, %.1f, %.1f)", 
            entity, position.x, position.y, position.z))
    end
})

-- Run simulation
print("Starting basic simulation...")

print("\nRunning 3 frames:")
for i = 1, 3 do
    print(string.format("\nFrame %d:", i))
    world:update()
end 