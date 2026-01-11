<script setup lang="ts">
import AuthService from '@/services/AuthService'
import { TOKEN_PROVIDE_KEY } from '@/api/instance'
import { InputText } from 'primevue'
import Button from 'primevue/button'
import { inject, useTemplateRef } from 'vue'
import { useRouter } from 'vue-router'

const authForm = useTemplateRef('authForm')

const router = useRouter()

const tokenProvider = inject<{ setToken: (token: string) => void }>(TOKEN_PROVIDE_KEY)

const handleLogin = async (ev: MouseEvent) => {
  ev.preventDefault()
  if (!authForm.value) return

  const values = authForm.value.elements as unknown as HTMLCollectionOf<HTMLInputElement>

  const token = await AuthService.login({ email: values[0].value, password: values[1].value })

  if (token) tokenProvider?.setToken(token)

  router.push('/')
}

const navigateToRegisterView = () => router.push('/register')
</script>
<template>
  <div class="auth">
    <div class="auth-form-container">
      <h1>Авторизация</h1>
      <form class="auth-form" ref="authForm">
        <InputText name="email" placeholder="Почта" type="email" />
        <InputText name="password" placeholder="Пароль" type="password" />
        <Button type="submit" label="Войти" @click="handleLogin" />
      </form>
      <Button variant="outlined" @click="navigateToRegisterView">Создать профиль</Button>
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
