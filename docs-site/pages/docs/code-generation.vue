<template>
  <div class="max-w-3xl mx-auto">
    <h1 class="text-4xl font-bold mb-8">Code Generation</h1>
    
    <div class="prose prose-lg">
      <p class="mb-6">
        EJECS generates highly optimized C++ code from your component and system definitions. The generated code includes
        type definitions, efficient component storage, parallel system execution, and automatic memory management.
      </p>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Generation Process</h2>
      <p class="mb-4">EJECS follows these steps when generating code:</p>
      <ol class="list-decimal pl-6 mb-6">
        <li>Parse and validate EJECS files</li>
        <li>Generate C++ type definitions</li>
        <li>Create optimized component storage</li>
        <li>Build system implementations with parallel execution</li>
        <li>Generate memory management code</li>
        <li>Add debug instrumentation</li>
      </ol>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Generated Files</h2>
      <p class="mb-4">The compiler produces several files:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>generated/
  ├── types.hpp           // Type definitions
  ├── components/         // Component implementations
  │   ├── position.hpp
  │   ├── velocity.hpp
  │   └── ...
  ├── systems/           // System implementations
  │   ├── movement.hpp
  │   ├── combat.hpp
  │   └── ...
  ├── storage/          // Component storage
  │   ├── archetype.hpp
  │   ├── sparse_set.hpp
  │   └── ...
  └── debug/            // Debug utilities
      ├── profiler.hpp
      └── inspector.hpp</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Type Definitions</h2>
      <p class="mb-4">Generated type definitions include:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>// Generated from EJECS definitions
using Entity = std::uint32_t;

struct Position {
    float x;
    float y;
    float z;
};

template&lt;typename T&gt;
class ComponentStorage {
public:
    T* get(Entity entity);
    void set(Entity entity, const T& value);
    void remove(Entity entity);
    bool has(Entity entity) const;
private:
    SparseSet&lt;T&gt; storage;
};</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Component Implementation</h2>
      <p class="mb-4">Generated component code includes optimized storage and SIMD operations:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>// Generated Position component
class PositionStorage : public ComponentStorage&lt;Position&gt; {
public:
    void update_batch(Entity* entities, size_t count) {
        #pragma omp simd
        for (size_t i = 0; i < count; i++) {
            auto pos = get(entities[i]);
            // SIMD optimized operations
        }
    }
    
    // Memory-aligned allocation
    Position* allocate(size_t count) {
        return static_cast&lt;Position*&gt;(
            std::aligned_alloc(32, count * sizeof(Position))
        );
    }
};</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">System Implementation</h2>
      <p class="mb-4">Generated systems with parallel execution:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>// Generated Movement system
class MovementSystem {
public:
    static constexpr int Priority = 1;
    static constexpr int Frequency = 60;
    
    void update(World& world) {
        auto& positions = world.get_storage&lt;Position&gt;();
        auto& velocities = world.get_storage&lt;Velocity&gt;();
        
        auto view = world.view&lt;Position, Velocity&gt;();
        
        #pragma omp parallel for
        for (const auto entity : view) {
            auto pos = positions.get(entity);
            auto vel = velocities.get(entity);
            
            pos->x += vel->x;
            pos->y += vel->y;
            pos->z += vel->z;
        }
    }
};</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Compiler Options</h2>
      <p class="mb-4">Configure code generation with compiler flags:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code># Basic compilation
ejecs -input game.jecs -output game.cpp

# Generate with specific options
ejecs -input game.jecs -output game.cpp \
    --optimize \
    --simd \
    --parallel \
    --debug-info</code></pre>

      <div class="mt-8 p-4 bg-blue-50 rounded-md">
        <h3 class="text-xl font-semibold mb-2">💡 Pro Tip</h3>
        <p>
          Use the <code>--debug-info</code> flag during development to include profiling and inspection tools
          in the generated code. This helps identify performance bottlenecks and debug issues.
        </p>
      </div>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Next Steps</h2>
      <p class="mb-4">
        Now that you understand how EJECS generates optimized code, check out our 
        <NuxtLink to="/docs/examples" class="text-blue-600 hover:text-blue-800">Example Projects</NuxtLink>
        to see it in action.
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