<template>
  <div class="max-w-3xl mx-auto">
    <!-- Hero Section -->
    <div class="text-center mb-16">
      <div class="inline-block p-2 px-4 bg-blue-50 text-blue-600 rounded-full text-sm font-medium mb-4">
        Welcome to EJECS
      </div>
      <h1 class="text-5xl font-bold mb-4 bg-gradient-to-r from-blue-600 to-purple-600 text-transparent bg-clip-text">
        Build Better Games with EJECS
      </h1>
      <p class="text-xl text-gray-600 mb-8">
        A powerful scripting language for defining game systems that compiles to highly optimized code
      </p>
      <div class="flex justify-center space-x-4">
        <NuxtLink to="/docs/quick-start" class="inline-flex items-center px-6 py-3 border border-transparent text-base font-medium rounded-lg shadow-sm text-white bg-blue-600 hover:bg-blue-700 transition-colors">
          Quick Start Guide
          <svg class="ml-2 -mr-1 w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7l5 5m0 0l-5 5m5-5H6" />
          </svg>
        </NuxtLink>
        <a href="https://github.com/ejecs/ejecs" target="_blank" class="inline-flex items-center px-6 py-3 border border-gray-300 text-base font-medium rounded-lg text-gray-700 bg-white hover:bg-gray-50 transition-colors">
          View on GitHub
          <svg class="ml-2 -mr-1 w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
            <path fill-rule="evenodd" d="M12 2C6.477 2 2 6.484 2 12.017c0 4.425 2.865 8.18 6.839 9.504.5.092.682-.217.682-.483 0-.237-.008-.868-.013-1.703-2.782.605-3.369-1.343-3.369-1.343-.454-1.158-1.11-1.466-1.11-1.466-.908-.62.069-.608.069-.608 1.003.07 1.531 1.032 1.531 1.032.892 1.53 2.341 1.088 2.91.832.092-.647.35-1.088.636-1.338-2.22-.253-4.555-1.113-4.555-4.951 0-1.093.39-1.988 1.029-2.688-.103-.253-.446-1.272.098-2.65 0 0 .84-.27 2.75 1.026A9.564 9.564 0 0112 6.844c.85.004 1.705.115 2.504.337 1.909-1.296 2.747-1.027 2.747-1.027.546 1.379.202 2.398.1 2.651.64.7 1.028 1.595 1.028 2.688 0 3.848-2.339 4.695-4.566 4.943.359.309.678.92.678 1.855 0 1.338-.012 2.419-.012 2.747 0 .268.18.58.688.482A10.019 10.019 0 0022 12.017C22 6.484 17.522 2 12 2z" clip-rule="evenodd" />
          </svg>
        </a>
      </div>
    </div>

    <!-- Features Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-16">
      <div class="bg-white p-6 rounded-xl shadow-sm border border-gray-100 hover:shadow-md transition-shadow">
        <div class="text-blue-600 mb-4">
          <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 3v2m6-2v2M9 19v2m6-2v2M5 9H3m2 6H3m18-6h-2m2 6h-2M7 19h10a2 2 0 002-2V7a2 2 0 00-2-2H7a2 2 0 00-2 2v10a2 2 0 002 2zM9 9h6v6H9V9z" />
          </svg>
        </div>
        <h3 class="text-lg font-semibold mb-2">Type-Safe Components</h3>
        <p class="text-gray-600">Define game components with strong typing and automatic validation.</p>
      </div>

      <div class="bg-white p-6 rounded-xl shadow-sm border border-gray-100 hover:shadow-md transition-shadow">
        <div class="text-purple-600 mb-4">
          <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
          </svg>
        </div>
        <h3 class="text-lg font-semibold mb-2">Optimized Systems</h3>
        <p class="text-gray-600">Write high-level system logic that compiles to highly optimized code.</p>
      </div>

      <div class="bg-white p-6 rounded-xl shadow-sm border border-gray-100 hover:shadow-md transition-shadow">
        <div class="text-green-600 mb-4">
          <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z" />
          </svg>
        </div>
        <h3 class="text-lg font-semibold mb-2">Entity Relationships</h3>
        <p class="text-gray-600">Model complex game relationships with built-in hierarchy support.</p>
      </div>

      <div class="bg-white p-6 rounded-xl shadow-sm border border-gray-100 hover:shadow-md transition-shadow">
        <div class="text-red-600 mb-4">
          <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6V4m0 2a2 2 0 100 4m0-4a2 2 0 110 4m-6 8a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4m6 6v10m6-2a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4" />
          </svg>
        </div>
        <h3 class="text-lg font-semibold mb-2">Event Handling</h3>
        <p class="text-gray-600">Define and handle game events with type-safe interfaces.</p>
      </div>
    </div>

    <!-- Code Example Section -->
    <div class="mb-16">
      <h2 class="text-2xl font-bold mb-6">Quick Example</h2>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="bg-gray-900 rounded-xl p-6 text-white">
          <div class="flex items-center justify-between mb-4">
            <span class="text-sm text-gray-400">game.jecs</span>
            <button class="text-gray-400 hover:text-white transition-colors">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z" />
              </svg>
            </button>
          </div>
          <pre class="text-sm font-mono"><code>component Position {
    x: number
    y: number
    z: number
}

@replicated
component Health {
    current: number
    max: number
}

system Movement {
    query: (Position, Velocity)
    frequency: 60Hz
    priority: 1
}</code></pre>
        </div>

        <div class="bg-gray-900 rounded-xl p-6 text-white">
          <div class="flex items-center justify-between mb-4">
            <span class="text-sm text-gray-400">Generated Code</span>
            <button class="text-gray-400 hover:text-white transition-colors">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z" />
              </svg>
            </button>
          </div>
          <pre class="text-sm font-mono"><code>// Optimized output
struct Position {
    float x, y, z;
};

struct Health {
    float current, max;
};

void movement_system(
    Position* positions,
    Velocity* velocities,
    size_t count
) {
    #pragma omp parallel for
    for(size_t i = 0; i < count; i++) {
        positions[i].x += velocities[i].x;
        positions[i].y += velocities[i].y;
        positions[i].z += velocities[i].z;
    }
}</code></pre>
        </div>
      </div>
    </div>

    <!-- Getting Started Section -->
    <div class="bg-gradient-to-r from-blue-500 to-purple-600 rounded-xl p-8 text-white">
      <h2 class="text-2xl font-bold mb-4">Ready to Get Started?</h2>
      <p class="mb-6">Learn how to write EJECS scripts and generate optimized code for your game.</p>
      <NuxtLink to="/docs/installation" class="inline-flex items-center px-6 py-3 border border-transparent text-base font-medium rounded-lg text-blue-600 bg-white hover:bg-gray-50 transition-colors">
        Installation Guide
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
.bg-clip-text {
  -webkit-background-clip: text;
}
</style> 