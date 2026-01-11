<script setup lang="ts">
import { RouterView } from 'vue-router'
import { Toast } from 'primevue'
import { provide } from 'vue'
import { TOKEN_PROVIDE_KEY, TokenMiddleware } from '@/api/instance'

const { clearToken, setToken } = TokenMiddleware()

provide(TOKEN_PROVIDE_KEY, { setToken, clearToken })
</script>

<template>
  <div class="root">
    <Toast />
    <RouterView />
  </div>
</template>

<style scoped>
.root {
  padding: 24px 16px;
}
</style>
