<script setup lang="ts">
import { RouterLink } from 'vue-router'
import Back from '@/assets/back.png'
</script>
<template>
  <RouterLink class="back-button" to="/">
    <img :src="Back" :alt="Back" />
  </RouterLink>
</template>
