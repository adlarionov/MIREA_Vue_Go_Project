<script setup lang="ts">
import Header from '@/components/Header.vue'
import Button from 'primevue/button'
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
      <Button label="Выйти" @click="$router.push('/auth')" severity="danger" />
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
