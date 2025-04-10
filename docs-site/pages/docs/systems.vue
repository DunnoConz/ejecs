<template>
  <div class="max-w-3xl mx-auto">
    <!-- Header -->
    <div class="mb-12">
      <div class="inline-block p-2 px-4 bg-purple-50 text-purple-600 rounded-full text-sm font-medium mb-4">
        Core Concepts
      </div>
      <h1 class="text-4xl font-bold mb-4">System Definitions</h1>
      <p class="text-xl text-gray-600">
        Systems in EJECS are interfaces that specify what components they operate on and how they
        should be executed. The actual implementation is provided by your chosen ECS framework.
      </p>
    </div>

    <!-- Basic System Definition -->
    <div class="mb-12">
      <h2 class="text-2xl font-bold mb-6 flex items-center">
        <span class="text-purple-600 mr-3">
          <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
          </svg>
        </span>
        Basic System Definition
      </h2>
      <p class="text-gray-600 mb-6">Define a system with its required components:</p>
      <div class="bg-gray-900 rounded-xl p-6 text-white mb-6">
        <div class="flex items-center justify-between mb-4">
          <span class="text-sm text-gray-400">movement_system.jecs</span>
          <button class="text-gray-400 hover:text-white transition-colors">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z" />
            </svg>
          </button>
        </div>
        <pre class="text-sm font-mono"><code>system Movement {
    -- Components this system operates on
    query: (Position, Velocity)

    -- Optional system configuration
    frequency: 60Hz
    priority: 1
}</code></pre>
      </div>
    </div>

    <!-- System Attributes -->
    <div class="mb-12">
      <h2 class="text-2xl font-bold mb-6 flex items-center">
        <span class="text-blue-600 mr-3">
          <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z" />
          </svg>
        </span>
        System Attributes
      </h2>
      <p class="text-gray-600 mb-6">Configure system behavior with attributes:</p>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
        <div class="bg-gray-900 rounded-xl p-6 text-white">
          <div class="flex items-center justify-between mb-4">
            <span class="text-sm text-gray-400">Server Systems</span>
          </div>
          <pre class="text-sm font-mono"><code>@server
system Combat {
    query: (Health, Damage)
    priority: 2
}

@client
system Animation {
    query: (Model, AnimState)
    frequency: 60Hz
}</code></pre>
        </div>
        <div class="bg-gray-900 rounded-xl p-6 text-white">
          <div class="flex items-center justify-between mb-4">
            <span class="text-sm text-gray-400">Parallel Systems</span>
          </div>
          <pre class="text-sm font-mono"><code>@parallel
system ParticleUpdate {
    query: (Position, Emitter)
    batch_size: 100
    frequency: 30Hz
}</code></pre>
        </div>
      </div>
    </div>

    <!-- System Dependencies -->
    <div class="mb-12">
      <h2 class="text-2xl font-bold mb-6 flex items-center">
        <span class="text-green-600 mr-3">
          <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 11.5V14m0-2.5v-6a2.5 2.5 0 015 0v6m-5 0h5m-5 0a2.5 2.5 0 002.5 2.5H15m-4 0a2.5 2.5 0 002.5-2.5M15 11.5V14" />
          </svg>
        </span>
        System Dependencies
      </h2>
      <p class="text-gray-600 mb-6">Specify dependencies between systems:</p>
      <div class="bg-gray-900 rounded-xl p-6 text-white mb-6">
        <pre class="text-sm font-mono"><code>system Physics {
    query: (Position, Velocity)
    priority: 1
}

system Combat {
    query: (Health, Damage)
    after: Physics  -- Run after physics
}

system Rendering {
    query: (Position, Model)
    after: [Physics, Combat]  -- Multiple deps
}</code></pre>
      </div>
    </div>

    <!-- Query Specifications -->
    <div class="mb-12">
      <h2 class="text-2xl font-bold mb-6 flex items-center">
        <span class="text-red-600 mr-3">
          <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
        </span>
        Query Specifications
      </h2>
      <p class="text-gray-600 mb-6">Define how systems query for entities:</p>
      <div class="bg-gray-900 rounded-xl p-6 text-white mb-6">
        <pre class="text-sm font-mono"><code>system AI {
    -- Multiple component queries
    queries {
        idle: {
            required: [AI, Position],
            optional: [Path],
            without: [Combat]
        },
        combat: {
            required: [AI, Position, Combat],
            optional: [Weapon]
        }
    }
}</code></pre>
      </div>
    </div>

    <!-- Pro Tips -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-12">
      <div class="bg-blue-50 rounded-xl p-6">
        <div class="flex items-center text-blue-600 mb-4">
          <svg class="w-8 h-8 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z" />
          </svg>
          <h3 class="text-lg font-semibold">Performance Tip</h3>
        </div>
        <p class="text-gray-700">
          Use system dependencies and priorities to create a predictable update order.
          This makes your game's behavior more deterministic and easier to debug.
        </p>
      </div>
      <div class="bg-purple-50 rounded-xl p-6">
        <div class="flex items-center text-purple-600 mb-4">
          <svg class="w-8 h-8 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          <h3 class="text-lg font-semibold">Best Practice</h3>
        </div>
        <p class="text-gray-700">
          Keep your systems focused on a single responsibility. If a system is doing too much,
          consider splitting it into multiple systems with clear dependencies.
        </p>
      </div>
    </div>

    <!-- Next Steps -->
    <div class="bg-gradient-to-r from-purple-500 to-blue-600 rounded-xl p-8 text-white">
      <h2 class="text-2xl font-bold mb-4">Ready to Learn More?</h2>
      <p class="mb-6">Learn how to define entity relationships in the next guide.</p>
      <NuxtLink to="/docs/relationships" class="inline-flex items-center px-6 py-3 border border-transparent text-base font-medium rounded-lg text-purple-600 bg-white hover:bg-gray-50 transition-colors">
        Relationships Guide
        <svg class="ml-2 -mr-1 w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7l5 5m0 0l-5 5m5-5H6" />
        </svg>
      </NuxtLink>
    </div>
  </div>
</template>

<script setup lang="ts">
// Add any required script logic here
</script>

<style scoped>
/* Add any scoped styles here */
</style> 