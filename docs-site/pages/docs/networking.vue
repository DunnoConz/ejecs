<template>
  <div class="max-w-3xl mx-auto">
    <h1 class="text-4xl font-bold mb-8">Networking</h1>
    
    <div class="prose prose-lg">
      <p class="mb-6">
        EJECS provides built-in support for network replication in Roblox games. The IDL allows you to specify
        which components, systems, and events should be replicated between server and client, ensuring type-safe
        network communication.
      </p>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Component Replication</h2>
      <p class="mb-4">Mark components for network replication:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>@replicated  -- Syncs to all clients
component Position {
    x: number
    y: number
    z: number
}

@replicated(owner)  -- Only syncs to owning client
component Inventory {
    items: {string}
    gold: number
}

@replicated(reliable)  -- Uses reliable network transport
component CriticalState {
    health: number
    status: string
}</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">System Network Context</h2>
      <p class="mb-4">Define where systems should run:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>@server  -- Runs only on server
system Physics {
    using Position, Velocity
    frequency: 60Hz
}

@client  -- Runs only on client
system Animation {
    using Model, AnimationState
}

@shared  -- Runs on both server and client
system ParticleEffect {
    using Position, ParticleEmitter
}</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Network Events</h2>
      <p class="mb-4">Configure event network behavior:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>@replicated  -- Server to client event
event PlayerSpawned {
    player: Entity
    position: Vector3
}

@server_to_client(owner)  -- Only sent to owning client
event InventoryUpdated {
    player: Entity
    item: string
    quantity: number
}

@client_to_server  -- Client to server event
event PlayerInput {
    player: Entity
    moveDirection: Vector3
    actions: {string}
}</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Network Optimization</h2>
      <p class="mb-4">Control replication frequency and precision:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>@replicated
@rate(10Hz)  -- Update 10 times per second
component Position {
    @precision(0.1)  -- Round to nearest 0.1
    x: number
    y: number
    z: number
}

@replicated
@delta  -- Only send changed fields
component PlayerState {
    health: number
    stamina: number
    status: string
}</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Network Authority</h2>
      <p class="mb-4">Specify data ownership and authority:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>@replicated
@server_owned  -- Server has final authority
component Position {
    x: number
    y: number
    z: number
}

@replicated
@client_owned  -- Client can modify directly
component Input {
    moveDirection: Vector3
    actions: {string}
}

@replicated
@shared_owned  -- Both can modify with reconciliation
component Prediction {
    position: Vector3
    velocity: Vector3
}</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Generated Network Code</h2>
      <p class="mb-4">EJECS generates network replication code:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- Generated from EJECS definitions
export type NetworkedComponent = {
    id: string,
    owner: Player?,
    replicationType: "all" | "owner" | "none",
    transportType: "reliable" | "unreliable",
    lastUpdate: number
}

export type NetworkManager = {
    replicate: (entity: Entity, component: string) -> (),
    isReplicated: (entity: Entity, component: string) -> boolean,
    getOwner: (entity: Entity) -> Player?
}</code></pre>

      <div class="mt-8 p-4 bg-blue-50 rounded-md">
        <h3 class="text-xl font-semibold mb-2">💡 Pro Tip</h3>
        <p>
          Use delta compression and rate limiting on frequently changing components to reduce network traffic.
          Only replicate what clients need to know.
        </p>
      </div>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Next Steps</h2>
      <p class="mb-4">
        Learn about code generation in the 
        <NuxtLink to="/docs/code-generation" class="text-blue-600 hover:text-blue-800">Code Generation Guide</NuxtLink>.
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