<script setup lang="ts">
import Button from 'primevue/button'
</script>

<template>
  <div class="main-footer">
    <Button label="Создать мероприятие" @click="$router.push('/events/new')" />
    <Button label="Создать заявку" @click="$router.push('/requests/new')" />
  </div>
</template>

<style scoped>
.main-footer {
  display: flex;
  align-items: center;
  gap: 16px;
  justify-content: end;
}
</style>
