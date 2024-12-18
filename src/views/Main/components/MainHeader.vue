<script setup lang="ts">
import AuthApi from '@/api/AuthApi'
import Header from '@/components/Header.vue'
import Button from 'primevue/button'
import { useRouter } from 'vue-router'

const router = useRouter()

const handleLogout = () => {
  AuthApi.logout()
  router.push('/auth')
}
</script>
<template>
  <Header title="Заявки">
    <template #buttons>
      <Button
        label="Роли"
        @click="$router.push('/roles')"
        variant="outlined"
        severity="secondary"
      />
      <Button label="Выйти" @click="handleLogout" severity="danger" />
    </template>
  </Header>
</template>
<style scoped>
.main-title {
  text-align: center;
  margin-left: 216px;
  width: 100%;
}
</style>
