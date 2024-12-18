<script setup lang="ts">
import { onMounted } from 'vue'
import { RouterView, useRouter } from 'vue-router'
import AuthApi from './api/AuthApi'
import { Toast } from 'primevue'

const router = useRouter()

onMounted(() => {
  const user = AuthApi.auth()
  if (!user) router.push('/auth')
})
</script>

<template>
  <div class="root">
    <Toast />
    <RouterView />
  </div>
</template>

<style scoped>
.root {
  padding: 60px 40px;
}
</style>
