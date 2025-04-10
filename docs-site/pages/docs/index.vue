<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()

// Redirect to introduction page
router.replace('/docs/introduction')
</script>

<template>
  <div>Redirecting to introduction...</div>
</template> 