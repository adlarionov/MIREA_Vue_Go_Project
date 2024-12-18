<script setup lang="ts">
import AuthApi from '@/api/AuthApi'
import { InputText } from 'primevue'
import Button from 'primevue/button'
import { useTemplateRef } from 'vue'
import { useRouter } from 'vue-router'

const authForm = useTemplateRef('authForm')

const router = useRouter()

const handleLogin = (ev: MouseEvent) => {
  ev.preventDefault()
  if (!authForm.value) return

  const values = authForm.value.elements as unknown as HTMLCollectionOf<HTMLInputElement>

  AuthApi.login(values[0].value, values[1].value)

  router.push('/')
}
</script>
<template>
  <div class="auth">
    <div class="auth-form-container">
      <h1>Авторизация</h1>
      <form class="auth-form" ref="authForm">
        <InputText name="username" placeholder="Логин" type="text" />
        <InputText name="password" placeholder="Пароль" type="password" />
        <Button type="submit" label="Войти" @click="handleLogin" />
      </form>
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

.auth-form-container {
  padding: 20px 80px;
  border-radius: 16px;
  border: black 2px solid;
  display: flex;
  flex-direction: column;
  gap: 16px;
  align-items: center;
  justify-content: center;
}

.auth-form {
  display: flex;
  flex-direction: column;
  gap: 12px;
}
</style>
