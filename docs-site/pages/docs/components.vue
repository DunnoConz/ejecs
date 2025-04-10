<template>
  <div class="max-w-3xl mx-auto">
    <h1 class="text-4xl font-bold mb-8">Component Definitions</h1>
    
    <div class="prose prose-lg">
      <p class="mb-6">
        EJECS allows you to define components using a TypeScript-like type system. These definitions describe
        the structure and constraints of data that entities can hold, which will be used to generate type-safe
        implementations in your target language.
      </p>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Basic Type Definitions</h2>
      <p class="mb-4">Define components with primitive types:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>component Position {
    x: number
    y: number
    z: number
}

component Health {
    current: number
    max: number
}

component Name {
    value: string
}

component IsActive {
    value: boolean
}</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Structured Types</h2>
      <p class="mb-4">Use complex types and nested structures:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>component Inventory {
    items: Item[]
    capacity: number
    equipped: {
        weapon: Item?
        armor: Item?
    }
}

component Transform {
    position: Vector3
    rotation: CFrame
    scale: Vector3
}</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Component Attributes</h2>
      <p class="mb-4">Add metadata and constraints to components:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>@replicated  -- Component will sync between server and client
component PlayerState {
    health: number
    stamina: number
}

@server  -- Component exists only on server
component AIState {
    @range(0, 100)
    awareness: number
    
    @default("patrol")
    state: string
}

@client  -- Component exists only on client
component Visual {
    model: Instance
    @optional
    effects: Instance[]
}</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Component Relationships</h2>
      <p class="mb-4">Define relationships between entities:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>relationship ChildOf {
    parent: Entity
    @index
    child: Entity
}

relationship OwnedBy {
    owner: Entity
    item: Entity
}</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Component Validation</h2>
      <p class="mb-4">Add validation rules to component fields:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>component Damage {
    @range(0, 1000)
    amount: number

    @pattern("^[a-zA-Z]+$")
    type: string

    @validate("validateMultiplier")
    multiplier: number
}</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Generated Code</h2>
      <p class="mb-4">EJECS generates type definitions and validation code:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- Generated from EJECS component definition
export type Position = {
    x: number,
    y: number,
    z: number
}

export type PositionValidator = {
    validate: (component: Position) -> boolean,
    getErrors: (component: Position) -> {string}
}</code></pre>

      <div class="mt-8 p-4 bg-blue-50 rounded-md">
        <h3 class="text-xl font-semibold mb-2">💡 Pro Tip</h3>
        <p>
          Choose meaningful names for your components and their fields. The generated code will use these names,
          so clear and descriptive names will make your codebase more maintainable.
        </p>
      </div>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Next Steps</h2>
      <p class="mb-4">
        Learn how to define system interfaces in the 
        <NuxtLink to="/docs/systems" class="text-blue-600 hover:text-blue-800">Systems Guide</NuxtLink>.
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