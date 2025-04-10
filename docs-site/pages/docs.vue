<template>
  <div class="flex min-h-screen bg-gray-50">
    <!-- Sidebar -->
    <div class="w-64 bg-white shadow-lg fixed h-full overflow-y-auto">
      <div class="p-6 border-b">
        <img src="/images/logo.jpg" alt="EJECS Logo" class="h-12 w-auto mx-auto mb-4" />
        <div class="text-sm text-gray-500 text-center">Documentation v1.0</div>
      </div>
      
      <nav class="p-4">
        <!-- Getting Started -->
        <div class="mb-8">
          <h3 class="font-semibold text-gray-400 uppercase tracking-wide text-sm mb-3">Getting Started</h3>
          <ul class="space-y-2">
            <li>
              <NuxtLink to="/docs/introduction" class="flex items-center text-gray-600 hover:text-blue-600 hover:bg-blue-50 rounded-lg p-2 transition-all">
                <span class="mr-2">📚</span> Introduction
              </NuxtLink>
            </li>
            <li>
              <NuxtLink to="/docs/installation" class="flex items-center text-gray-600 hover:text-blue-600 hover:bg-blue-50 rounded-lg p-2 transition-all">
                <span class="mr-2">⚙️</span> Installation
              </NuxtLink>
            </li>
            <li>
              <NuxtLink to="/docs/quick-start" class="flex items-center text-gray-600 hover:text-blue-600 hover:bg-blue-50 rounded-lg p-2 transition-all">
                <span class="mr-2">🚀</span> Quick Start
              </NuxtLink>
            </li>
          </ul>
        </div>

        <!-- Core Concepts -->
        <div class="mb-8">
          <h3 class="font-semibold text-gray-400 uppercase tracking-wide text-sm mb-3">Core Concepts</h3>
          <ul class="space-y-2">
            <li>
              <NuxtLink to="/docs/components" class="flex items-center text-gray-600 hover:text-blue-600 hover:bg-blue-50 rounded-lg p-2 transition-all">
                <span class="mr-2">🧩</span> Components
              </NuxtLink>
            </li>
            <li>
              <NuxtLink to="/docs/systems" class="flex items-center text-gray-600 hover:text-blue-600 hover:bg-blue-50 rounded-lg p-2 transition-all">
                <span class="mr-2">⚡</span> Systems
              </NuxtLink>
            </li>
            <li>
              <NuxtLink to="/docs/entities" class="flex items-center text-gray-600 hover:text-blue-600 hover:bg-blue-50 rounded-lg p-2 transition-all">
                <span class="mr-2">🎮</span> Entities
              </NuxtLink>
            </li>
            <li>
              <NuxtLink to="/docs/relationships" class="flex items-center text-gray-600 hover:text-blue-600 hover:bg-blue-50 rounded-lg p-2 transition-all">
                <span class="mr-2">🔗</span> Relationships
              </NuxtLink>
            </li>
          </ul>
        </div>

        <!-- Advanced -->
        <div class="mb-8">
          <h3 class="font-semibold text-gray-400 uppercase tracking-wide text-sm mb-3">Advanced</h3>
          <ul class="space-y-2">
            <li>
              <NuxtLink to="/docs/networking" class="flex items-center text-gray-600 hover:text-blue-600 hover:bg-blue-50 rounded-lg p-2 transition-all">
                <span class="mr-2">🌐</span> Networking
              </NuxtLink>
            </li>
            <li>
              <NuxtLink to="/docs/events" class="flex items-center text-gray-600 hover:text-blue-600 hover:bg-blue-50 rounded-lg p-2 transition-all">
                <span class="mr-2">📡</span> Events
              </NuxtLink>
            </li>
            <li>
              <NuxtLink to="/docs/code-generation" class="flex items-center text-gray-600 hover:text-blue-600 hover:bg-blue-50 rounded-lg p-2 transition-all">
                <span class="mr-2">🔧</span> Code Generation
              </NuxtLink>
            </li>
          </ul>
        </div>

        <!-- Reference -->
        <div>
          <h3 class="font-semibold text-gray-400 uppercase tracking-wide text-sm mb-3">Reference</h3>
          <ul class="space-y-2">
            <li>
              <NuxtLink to="/docs/api-reference" class="flex items-center text-gray-600 hover:text-blue-600 hover:bg-blue-50 rounded-lg p-2 transition-all">
                <span class="mr-2">📖</span> API Reference
              </NuxtLink>
            </li>
            <li>
              <NuxtLink to="/docs/examples" class="flex items-center text-gray-600 hover:text-blue-600 hover:bg-blue-50 rounded-lg p-2 transition-all">
                <span class="mr-2">💡</span> Examples
              </NuxtLink>
            </li>
          </ul>
        </div>
      </nav>
    </div>

    <!-- Main Content -->
    <div class="ml-64 flex-1 p-8">
      <!-- Top Navigation Bar -->
      <div class="bg-white rounded-lg shadow-sm p-4 mb-8 flex items-center justify-between">
        <div class="flex items-center space-x-4">
          <button class="text-gray-500 hover:text-gray-700">
            <span class="sr-only">Search</span>
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
          </button>
          <div class="text-sm text-gray-500">Press <kbd class="px-2 py-1 bg-gray-100 rounded">⌘ K</kbd> to search</div>
        </div>
        <div class="flex items-center space-x-4">
          <a href="https://github.com/ejecs/ejecs" target="_blank" class="text-gray-500 hover:text-gray-700">
            <span class="sr-only">GitHub</span>
            <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
              <path fill-rule="evenodd" d="M12 2C6.477 2 2 6.484 2 12.017c0 4.425 2.865 8.18 6.839 9.504.5.092.682-.217.682-.483 0-.237-.008-.868-.013-1.703-2.782.605-3.369-1.343-3.369-1.343-.454-1.158-1.11-1.466-1.11-1.466-.908-.62.069-.608.069-.608 1.003.07 1.531 1.032 1.531 1.032.892 1.53 2.341 1.088 2.91.832.092-.647.35-1.088.636-1.338-2.22-.253-4.555-1.113-4.555-4.951 0-1.093.39-1.988 1.029-2.688-.103-.253-.446-1.272.098-2.65 0 0 .84-.27 2.75 1.026A9.564 9.564 0 0112 6.844c.85.004 1.705.115 2.504.337 1.909-1.296 2.747-1.027 2.747-1.027.546 1.379.202 2.398.1 2.651.64.7 1.028 1.595 1.028 2.688 0 3.848-2.339 4.695-4.566 4.943.359.309.678.92.678 1.855 0 1.338-.012 2.419-.012 2.747 0 .268.18.58.688.482A10.019 10.019 0 0022 12.017C22 6.484 17.522 2 12 2z" clip-rule="evenodd" />
            </svg>
          </a>
          <button class="text-gray-500 hover:text-gray-700">
            <span class="sr-only">Theme</span>
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z" />
            </svg>
          </button>
        </div>
      </div>

      <!-- Content Area -->
      <div class="bg-white rounded-lg shadow-sm p-8">
        <NuxtPage />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// Add keyboard shortcut for search
onMounted(() => {
  window.addEventListener('keydown', (e) => {
    if ((e.metaKey || e.ctrlKey) && e.key === 'k') {
      e.preventDefault()
      // TODO: Implement search
    }
  })
})
</script>

<style>
/* Custom scrollbar for sidebar */
.overflow-y-auto::-webkit-scrollbar {
  width: 4px;
}

.overflow-y-auto::-webkit-scrollbar-track {
  background: #f1f1f1;
}

.overflow-y-auto::-webkit-scrollbar-thumb {
  background: #cbd5e1;
  border-radius: 2px;
}

.overflow-y-auto::-webkit-scrollbar-thumb:hover {
  background: #94a3b8;
}

/* Smooth transitions */
.transition-all {
  transition-property: all;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  transition-duration: 150ms;
}
</style> 