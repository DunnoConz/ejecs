<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Navigation -->
    <nav class="bg-white shadow-sm">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <div class="flex">
            <NuxtLink to="/" class="flex items-center">
              <span class="text-xl font-bold text-gray-900">EJECS</span>
            </NuxtLink>
          </div>
          <div class="flex space-x-8">
            <NuxtLink to="/docs" class="inline-flex items-center px-1 pt-1 text-gray-900">
              Documentation
            </NuxtLink>
            <NuxtLink to="/examples" class="inline-flex items-center px-1 pt-1 text-gray-900">
              Examples
            </NuxtLink>
            <NuxtLink to="/download" class="inline-flex items-center px-1 pt-1 text-gray-900">
              Download
            </NuxtLink>
            <a href="https://github.com/ejecs/ejecs" target="_blank" class="inline-flex items-center px-1 pt-1 text-gray-900">
              GitHub
            </a>
          </div>
        </div>
      </div>
    </nav>

    <!-- Main content -->
    <main>
      <NuxtPage />
    </main>

    <!-- Footer -->
    <footer class="bg-white mt-auto">
      <div class="max-w-7xl mx-auto py-12 px-4 sm:px-6 md:flex md:items-center md:justify-between lg:px-8">
        <div class="mt-8 md:mt-0">
          <p class="text-center text-base text-gray-400">
            &copy; {{ new Date().getFullYear() }} EJECS. All rights reserved.
          </p>
        </div>
      </div>
    </footer>
  </div>
</template>

<style>
.page-enter-active,
.page-leave-active {
  transition: all 0.4s;
}
.page-enter-from,
.page-leave-to {
  opacity: 0;
  filter: blur(1rem);
}
</style> 