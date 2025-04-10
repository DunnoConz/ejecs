<template>
  <div class="max-w-3xl mx-auto">
    <h1 class="text-4xl font-bold mb-8">Installation</h1>
    
    <div class="prose prose-lg">
      <p class="mb-6">
        EJECS can be installed through Wally, the package manager for Roblox. You can also build it from source
        if you need the latest features or want to contribute to development.
      </p>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Prerequisites</h2>
      <ul class="list-disc pl-6 mb-6">
        <li>Roblox Studio</li>
        <li>Wally (Package Manager)</li>
        <li>Rojo (for source installation)</li>
      </ul>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Installing via Wally</h2>
      <p class="mb-4">Add EJECS to your <code>wally.toml</code>:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>[dependencies]
ejecs = "yourusername/ejecs@1.0.0"</code></pre>
      
      <p class="mb-4">Then run Wally to install:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>wally install</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Building from Source</h2>
      <p class="mb-4">Clone the repository and build using Rojo:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code># Clone the repository
git clone https://github.com/yourusername/ejecs.git
cd ejecs

# Build with Rojo
rojo build -o ejecs.rbxm</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">IDE Integration</h2>
      <p class="mb-4">
        EJECS provides language server support for VS Code and other editors that support the Language Server Protocol.
        Install the EJECS extension for your editor to get:
      </p>
      <ul class="list-disc pl-6 mb-6">
        <li>Syntax highlighting</li>
        <li>Code completion</li>
        <li>Error checking</li>
        <li>Go to definition</li>
      </ul>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Verifying Installation</h2>
      <p class="mb-4">Create a simple EJECS file to verify your installation:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- test.jecs
component Test {
    value: string
}

system TestSystem {
    using Test
}</code></pre>

      <p class="mb-4">Then run the EJECS compiler:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>ejecs compile test.jecs</code></pre>

      <div class="mt-8 p-4 bg-blue-50 rounded-md">
        <h3 class="text-xl font-semibold mb-2">💡 Pro Tip</h3>
        <p>
          Use the <code>ejecs init</code> command to create a new project with a basic configuration and example files.
        </p>
      </div>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Next Steps</h2>
      <p class="mb-4">
        Follow the <NuxtLink to="/docs/quick-start" class="text-blue-600 hover:text-blue-800">Quick Start Guide</NuxtLink>
        to create your first EJECS project.
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