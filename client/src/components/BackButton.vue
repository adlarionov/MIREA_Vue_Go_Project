<script setup lang="ts">
import { RouterLink } from 'vue-router'
import Back from '@/assets/back.png'

defineProps<{ to: string }>()
</script>
<template>
  <RouterLink class="back-button" :to="to">
    <img :src="Back" :alt="Back" />
  </RouterLink>
</template>
