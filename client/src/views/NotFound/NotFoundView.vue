<script setup lang="ts">
import { useRouter } from 'vue-router'
import Button from 'primevue/button'
import Card from 'primevue/card'

const router = useRouter()

const goHome = () => {
  router.push('/')
}
</script>

<template>
  <div class="not-found">
    <Card class="card">
      <template #title>
        <span class="code">404</span>
      </template>

      <template #content>
        <p class="title">Страница не найдена</p>
        <p class="description">
          Возможно, вы перешли по неверной ссылке или страница была удалена.
        </p>

        <Button label="На главную" icon="pi pi-home" @click="goHome" />
      </template>
    </Card>
  </div>
</template>

<style scoped>
.not-found {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--surface-ground);
}

.card {
  width: 420px;
  text-align: center;
}

.code {
  font-size: 72px;
  font-weight: 700;
  color: var(--primary-color);
}

.title {
  font-size: 20px;
  font-weight: 600;
  margin-bottom: 8px;
}

.description {
  font-size: 14px;
  color: var(--text-color-secondary);
  margin-bottom: 24px;
}
</style>
