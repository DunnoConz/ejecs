# Examples

This page contains practical examples of EJECS usage, from basic to advanced scenarios.

## Basic Examples

### Simple 2D Game

A basic example showing position and velocity components with movement and rendering systems.

```ejecs
// Basic components for a 2D game
component Position {
    x: number
    y: number
}

component Velocity {
    dx: number
    dy: number
}

component Sprite {
    image: string
    width: number = 32
    height: number = 32
}

// Movement system
system Movement {
    query(Position, Velocity)
    frequency(60)
    {
        for _, entity in ipairs(entities) do
            local pos = entity.Position
            local vel = entity.Velocity
            pos.x = pos.x + vel.dx
            pos.y = pos.y + vel.dy
        end
    }
}

// Render system
system Render {
    query(Position, Sprite)
    {
        for _, entity in ipairs(entities) do
            local pos = entity.Position
            local sprite = entity.Sprite
            love.graphics.draw(sprite.image, pos.x, pos.y)
        end
    }
}
```

### Health and Damage System

A simple health system with damage and healing capabilities.

```ejecs
component Health {
    current: number = 100
    max: number = 100
}

component Damageable {
    armor: number = 0
    resistance: number = 0
}

system TakeDamage {
    query(Health, Damageable)
    params {
        amount: number = 1
        type: string = "physical"
    }
    {
        for _, entity in ipairs(entities) do
            local health = entity.Health
            local damageable = entity.Damageable
            
            local damage = amount
            if type == "physical" then
                damage = damage - damageable.armor
            end
            damage = damage * (1 - damageable.resistance)
            
            health.current = math.max(0, health.current - damage)
        end
    }
}

system Heal {
    query(Health)
    params {
        amount: number = 10
    }
    {
        for _, entity in ipairs(entities) do
            local health = entity.Health
            health.current = math.min(health.max, health.current + amount)
        end
    }
}
```

## Intermediate Examples

### Inventory System

A more complex example showing an inventory system with items and equipment.

```ejecs
component Item {
    name: string
    value: number
    weight: number
    stackable: boolean = false
    quantity: number = 1
}

component Inventory {
    items: {string: Item}
    maxSize: number = 20
    maxWeight: number = 100
    gold: number = 0
}

component Equipment {
    slots: {string: Item}
    maxSlots: number = 10
}

system AddItem {
    query(Inventory)
    params {
        item: Item
    }
    {
        for _, entity in ipairs(entities) do
            local inv = entity.Inventory
            if inv.items[item.name] then
                if item.stackable then
                    inv.items[item.name].quantity = inv.items[item.name].quantity + item.quantity
                else
                    -- Handle non-stackable items
                end
            else
                inv.items[item.name] = item
            end
        end
    }
}

system EquipItem {
    query(Equipment, Inventory)
    params {
        itemName: string
        slot: string
    }
    {
        for _, entity in ipairs(entities) do
            local equip = entity.Equipment
            local inv = entity.Inventory
            
            if inv.items[itemName] and not equip.slots[slot] then
                equip.slots[slot] = inv.items[itemName]
                inv.items[itemName] = nil
            end
        end
    }
}
```

### State Machine

An example showing how to implement a state machine using components.

```ejecs
component State {
    current: string = "idle"
    previous: string?
    transitionTime: number = 0
}

component Animation {
    frames: string[]
    currentFrame: number = 1
    frameTime: number = 0.1
    loop: boolean = true
}

system StateTransition {
    query(State)
    params {
        newState: string
    }
    {
        for _, entity in ipairs(entities) do
            local state = entity.State
            if state.current ~= newState then
                state.previous = state.current
                state.current = newState
                state.transitionTime = 0
            end
        end
    }
}

system UpdateAnimation {
    query(State, Animation)
    {
        for _, entity in ipairs(entities) do
            local state = entity.State
            local anim = entity.Animation
            
            anim.frameTime = anim.frameTime - dt
            if anim.frameTime <= 0 then
                anim.currentFrame = anim.currentFrame + 1
                if anim.currentFrame > #anim.frames then
                    if anim.loop then
                        anim.currentFrame = 1
                    else
                        anim.currentFrame = #anim.frames
                    end
                end
                anim.frameTime = 0.1
            end
        end
    }
}
```

## Advanced Examples

### Networked Game State

An example showing how to handle networked game state with components.

```ejecs
component Networked {
    id: string
    lastUpdate: number = 0
    interpolationTime: number = 0.1
}

component Position {
    x: number
    y: number
    targetX: number?
    targetY: number?
}

component Velocity {
    dx: number
    dy: number
}

system NetworkUpdate {
    query(Networked, Position)
    params {
        updates: {string: {x: number, y: number}}
    }
    {
        for _, entity in ipairs(entities) do
            local net = entity.Networked
            local pos = entity.Position
            
            if updates[net.id] then
                local update = updates[net.id]
                pos.targetX = update.x
                pos.targetY = update.y
                net.lastUpdate = love.timer.getTime()
            end
        end
    }
}

system InterpolatePosition {
    query(Networked, Position)
    {
        for _, entity in ipairs(entities) do
            local net = entity.Networked
            local pos = entity.Position
            
            if pos.targetX and pos.targetY then
                local alpha = (love.timer.getTime() - net.lastUpdate) / net.interpolationTime
                alpha = math.min(1, alpha)
                
                pos.x = pos.x + (pos.targetX - pos.x) * alpha
                pos.y = pos.y + (pos.targetY - pos.y) * alpha
            end
        end
    }
}
```

### Particle System

An advanced example showing a particle system implementation.

```ejecs
component Particle {
    position: Position
    velocity: Velocity
    lifetime: number
    age: number = 0
    color: {r: number, g: number, b: number, a: number}
    size: number
}

component ParticleEmitter {
    rate: number = 10
    lifetime: number = 1
    color: {r: number, g: number, b: number, a: number}
    size: number = 1
    spread: number = 1
    speed: number = 1
}

system EmitParticles {
    query(ParticleEmitter, Position)
    {
        for _, entity in ipairs(entities) do
            local emitter = entity.ParticleEmitter
            local pos = entity.Position
            
            local particlesToEmit = emitter.rate * dt
            for i = 1, particlesToEmit do
                local angle = love.math.random() * math.pi * 2
                local distance = love.math.random() * emitter.spread
                
                local particle = {
                    position = {
                        x = pos.x + math.cos(angle) * distance,
                        y = pos.y + math.sin(angle) * distance
                    },
                    velocity = {
                        dx = math.cos(angle) * emitter.speed,
                        dy = math.sin(angle) * emitter.speed
                    },
                    lifetime = emitter.lifetime,
                    color = emitter.color,
                    size = emitter.size
                }
                
                ecr:add(particle)
            end
        end
    }
}

system UpdateParticles {
    query(Particle)
    {
        for _, entity in ipairs(entities) do
            local particle = entity.Particle
            particle.age = particle.age + dt
            
            if particle.age >= particle.lifetime then
                ecr:remove(entity)
            else
                local alpha = 1 - (particle.age / particle.lifetime)
                particle.color.a = alpha
                
                particle.position.x = particle.position.x + particle.velocity.dx * dt
                particle.position.y = particle.position.y + particle.velocity.dy * dt
            end
        end
    }
}
```

## Running the Examples

To run these examples:

1. Save the code to a `.ejecs` file
2. Generate the code for your target platform:
```bash
ejecs generate -f ecr example.ejecs > example.lua
```
3. Integrate the generated code into your game

## Contributing Examples

We welcome contributions of new examples! Please follow these guidelines:

1. Keep examples focused and clear
2. Include comments explaining complex parts
3. Test the examples before submitting
4. Follow the EJECS style guide
5. Include both the EJECS source and expected output 