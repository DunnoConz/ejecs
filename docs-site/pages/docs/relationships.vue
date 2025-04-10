<template>
  <div class="max-w-3xl mx-auto">
    <h1 class="text-4xl font-bold mb-8">Entity Relationships</h1>
    
    <div class="prose prose-lg">
      <p class="mb-6">
        EJECS allows you to define explicit relationships between entities. These relationships are first-class
        citizens in the IDL and can be used to model parent-child hierarchies, ownership, or any other type of
        entity association.
      </p>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Basic Relationship Definition</h2>
      <p class="mb-4">Define a simple one-to-one relationship:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>relationship OwnedBy {
    owner: Entity
    item: Entity
}

relationship ControlledBy {
    controller: Entity
    controlled: Entity
}</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Relationship Types</h2>
      <p class="mb-4">EJECS supports different types of relationships:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>@one_to_one  -- Each entity can only have one relationship
relationship Marriage {
    spouse1: Entity
    spouse2: Entity
}

@one_to_many  -- One entity can relate to many others
relationship ChildOf {
    parent: Entity
    @index  -- Index for efficient child lookup
    child: Entity
}

@many_to_many  -- Many-to-many relationship
relationship TeamMembership {
    team: Entity
    player: Entity
}</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Relationship Attributes</h2>
      <p class="mb-4">Add metadata and constraints to relationships:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>@replicated  -- Relationship syncs between server/client
relationship Friendship {
    friend1: Entity
    friend2: Entity
    @range(0, 100)
    strength: number
}

@server  -- Server-only relationship
relationship AITarget {
    seeker: Entity
    target: Entity
    @timestamp
    lastSeen: number
}</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Component Requirements</h2>
      <p class="mb-4">Specify required components for relationship participants:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>relationship ChildOf {
    @requires(Transform)  -- Parent must have Transform
    parent: Entity
    
    @requires(Transform, Model)  -- Child needs both
    child: Entity
}

relationship Equipped {
    @requires(Inventory)
    owner: Entity
    
    @requires(Item)
    item: Entity
}</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Querying Relationships</h2>
      <p class="mb-4">Example of how to use relationships in systems:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>system ParentTransform {
    using Transform
    
    queries {
        children: {
            relationship: ChildOf,
            select: [parent, child]
        }
    }
}

system InventoryUpdate {
    using Inventory, Item
    
    queries {
        equipment: {
            relationship: Equipped,
            select: [owner, item]
        }
    }
}</code></pre>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Generated Code</h2>
      <p class="mb-4">EJECS generates relationship management code:</p>
      <pre class="bg-gray-100 p-4 rounded-md mb-6"><code>-- Generated from EJECS relationship definition
export type ChildOfRelation = {
    parent: Entity,
    child: Entity
}

export type ChildOfManager = {
    create: (parent: Entity, child: Entity) -> (),
    remove: (parent: Entity, child: Entity) -> (),
    getParent: (child: Entity) -> Entity?,
    getChildren: (parent: Entity) -> {Entity}
}</code></pre>

      <div class="mt-8 p-4 bg-blue-50 rounded-md">
        <h3 class="text-xl font-semibold mb-2">💡 Pro Tip</h3>
        <p>
          Use relationship indices to optimize queries when you frequently need to look up related entities
          in a specific direction (e.g., finding all children of a parent).
        </p>
      </div>

      <h2 class="text-2xl font-semibold mt-8 mb-4">Next Steps</h2>
      <p class="mb-4">
        Learn about advanced features in the 
        <NuxtLink to="/docs/events" class="text-blue-600 hover:text-blue-800">Events Guide</NuxtLink>.
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