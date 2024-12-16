<script setup lang="ts">
import Button from 'primevue/button'
</script>
<template>
  <div class="auth">
    <div class="auth-form">
      <h1>Авторизация</h1>
      <Button label="Войти" @click="$router.push('/')" />
    </div>
  </div>
</template>

<style lang="css" scoped>
.auth {
  display: flex;
  height: 80vh;
  align-items: center;
  justify-content: center;
}

.auth-form {
  padding: 20px 80px;
  border-radius: 16px;
  border: black 2px solid;
  display: flex;
  flex-direction: column;
  gap: 16px;
  align-items: center;
  justify-content: center;
}
</style>
