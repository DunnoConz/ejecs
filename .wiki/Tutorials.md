# Tutorials

This page contains step-by-step tutorials to help you learn EJECS.

## Getting Started

### Tutorial 1: Your First EJECS Project

Learn how to create a simple 2D game with EJECS.

1. **Setup**
   ```bash
   # Install EJECS
   go install github.com/your-org/ejecs/cmd/ejecs@latest

   # Create project directory
   mkdir my-game
   cd my-game
   ```

2. **Create Basic Components**
   ```ejecs
   // components.ejecs
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
   ```

3. **Add Movement System**
   ```ejecs
   // systems.ejecs
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
   ```

4. **Generate Code**
   ```bash
   ejecs generate -f ecr components.ejecs systems.ejecs > game.lua
   ```

5. **Run the Game**
   ```bash
   lua game.lua
   ```

### Tutorial 2: Adding Health and Combat

Learn how to implement a health and combat system.

1. **Add Health Component**
   ```ejecs
   component Health {
       current: number = 100
       max: number = 100
   }

   component Damageable {
       armor: number = 0
       resistance: number = 0
   }
   ```

2. **Create Combat Systems**
   ```ejecs
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
   ```

3. **Add Healing System**
   ```ejecs
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

### Tutorial 3: Inventory System

Learn how to create an inventory system.

1. **Define Item Components**
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
   ```

2. **Add Inventory Management Systems**
   ```ejecs
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
   ```

### Tutorial 4: State Machine

Learn how to implement a state machine.

1. **Create State Components**
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
   ```

2. **Add State Management Systems**
   ```ejecs
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
   ```

## Advanced Tutorials

### Tutorial 5: Networked Game State

Learn how to handle networked game state.

1. **Add Network Components**
   ```ejecs
   component Networked {
       id: string
       lastUpdate: number = 0
       interpolationTime: number = 0.1
   }
   ```

2. **Create Network Systems**
   ```ejecs
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
   ```

### Tutorial 6: Particle System

Learn how to create a particle system.

1. **Define Particle Components**
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
   ```

2. **Add Particle Systems**
   ```ejecs
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
   ```

## Best Practices

1. **Component Design**
   - Keep components small and focused
   - Use appropriate types
   - Leverage default values

2. **System Design**
   - Follow single responsibility principle
   - Use parameters for configuration
   - Consider frequency and priority

3. **Code Organization**
   - Group related components
   - Use clear naming conventions
   - Document complex logic

## Next Steps

1. Explore the [Examples](Examples) page for more complex examples
2. Read the [Best Practices](Best-Practices) guide
3. Check out the [Language Reference](Language-Reference) for detailed syntax
4. Join the community and share your projects 