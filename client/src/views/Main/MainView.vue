<script setup lang="ts">
// Components
import Button from 'primevue/button'
import { inject, onMounted, ref } from 'vue'
import { Menubar } from 'primevue'
import router from '@/router'
import type { MenuItem } from 'primevue/menuitem'
import { RouterView } from 'vue-router'
import { TOKEN_PROVIDE_KEY } from '@/api/instance'

onMounted(async () => {})

const menuItems = ref<MenuItem[]>([
  {
    label: 'Заявки',
    command() {
      router.push('/')
    },
  },
  {
    label: 'События',
    command() {
      router.push('/events')
    },
  },
  {
    label: 'Площадки',
    command() {
      router.push('/venues')
    },
  },
])

const tokenProvider = inject<{ clearToken: () => void }>(TOKEN_PROVIDE_KEY)

const handleLogout = () => {
  tokenProvider?.clearToken()
  router.push('/auth')
}
</script>
<template>
  <div>
    <Menubar :model="menuItems">
      <template #end>
        <Button label="Выйти" @click="handleLogout" severity="danger" variant="text" />
      </template>
    </Menubar>
    <RouterView />
  </div>
</template>

<style scoped>
.main-table {
  margin-bottom: 56px;
}
</style>
