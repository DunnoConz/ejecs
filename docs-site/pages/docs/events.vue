<template>
  <div class="max-w-3xl mx-auto">
    <h1 class="text-4xl font-bold mb-8">Events</h1>
    
    <div class="prose prose-lg">
      <p class="mb-6">
        EJECS provides a type-safe event system that allows components and systems to communicate through
        well-defined events. Events can be used to trigger state changes, notify systems of important
        occurrences, or coordinate behavior across different parts of your game.
      </p>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Event Definition</h2>
      <p class="mb-4">Define events with their payload types:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>event Damage {
    target: Entity
    amount: number
    type: string
}

event CollisionOccurred {
    entity1: Entity
    entity2: Entity
    point: Vector3
    normal: Vector3
}

event StateChanged {
    entity: Entity
    from: string
    to: string
}</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Event Attributes</h2>
      <p class="mb-4">Configure event behavior with attributes:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>@replicated  -- Event will be sent to clients
event PlayerJoined {
    player: Entity
    name: string
}

@server  -- Event only processed on server
event AdminCommand {
    issuer: Entity
    command: string
    args: {string}
}

@client  -- Event only processed on client
event UIStateChanged {
    widget: string
    state: string
}</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Event Handlers</h2>
      <p class="mb-4">Define how systems handle events:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>system DamageHandler {
    using Health
    
    events {
        onDamage: Damage  -- Handle damage events
    }
}

system CollisionSystem {
    using Position, Collider
    
    events {
        onCollision: CollisionOccurred,
        onTrigger: TriggerEntered
    }
}</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Event Validation</h2>
      <p class="mb-4">Add validation rules to event payloads:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>event ItemPickup {
    player: Entity
    item: Entity
    
    @range(1, 100)
    quantity: number
    
    @validate("validatePickupConditions")
    conditions: {string}
}

event Achievement {
    player: Entity
    
    @pattern("^[A-Z][a-z]+$")
    name: string
    
    @range(0, 1000)
    score: number
}</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Event Priorities</h2>
      <p class="mb-4">Control event handling order:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>system InventorySystem {
    using Inventory
    
    events {
        @priority(1)  -- Handle before other systems
        onItemPickup: ItemPickup
    }
}

system AchievementSystem {
    events {
        @priority(2)  -- Handle after inventory update
        onItemPickup: ItemPickup
    }
}</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Generated Code</h2>
      <p class="mb-4">EJECS generates type-safe event handling code:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- Generated from EJECS event definition
export type DamageEvent = {
    target: Entity,
    amount: number,
    type: string
}

export type EventDispatcher = {
    emit: (event: string, payload: any) -> (),
    on: (event: string, callback: (payload: any) -> ()) -> Connection,
    once: (event: string, callback: (payload: any) -> ()) -> Connection
}</code></pre>

      <div class="mt-8 p-4 bg-blue-50 rounded-md">
        <h3 class="text-xl font-semibold mb-2">💡 Pro Tip</h3>
        <p>
          Use event validation to ensure that event payloads are correct before they're processed.
          This is especially important for replicated events that cross network boundaries.
        </p>
      </div>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Next Steps</h2>
      <p class="mb-4">
        Learn about networking in the 
        <NuxtLink to="/docs/networking" class="text-blue-600 hover:text-blue-800">Networking Guide</NuxtLink>.
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