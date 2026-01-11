<script setup lang="ts">
import AuthApi from '@/services/AuthService'
import type { RegisterRequestDto } from '@/models/dto/Auth'
import router from '@/router'
import { Form, type FormResolverOptions, type FormSubmitEvent } from '@primevue/forms'
import { InputText } from 'primevue'
import Button from 'primevue/button'
import { reactive, ref } from 'vue'

const isButtonDisabled = ref(true)

const initialValues = reactive({
  name: '',
  phone: 799999999,
  email: '',
  password: '',
})

const resolver = ({ values }: FormResolverOptions) => {
  validateForm(values)

  return { values }
}

const validateForm = (values: Record<string, unknown>) => {
  const valueArr = Object.values(values)

  const isAllValuesNotEmpty = valueArr.filter(Boolean).length === valueArr.length

  if (isAllValuesNotEmpty) isButtonDisabled.value = false
  else isButtonDisabled.value = true
}

const handleRegister = async ({ values }: FormSubmitEvent) => {
  const registerStatusMessage = await AuthApi.register(values as RegisterRequestDto)
  if (registerStatusMessage) navigateToAuthPage()
}

const navigateToAuthPage = () => router.push('/auth')
</script>
<template>
  <div class="register">
    <h1>Создать профиль</h1>
    <hr class="register-ruler" />
    <div class="register-form-container">
      <Form class="register-form" @submit="handleRegister" :initialValues :resolver>
        <InputText name="name" placeholder="Название" type="text" required />
        <InputText name="phone" placeholder="Телефон" type="tel" required />
        <InputText name="email" placeholder="Почта" type="email" required />
        <InputText name="password" placeholder="Пароль" type="password" required />
        <Button type="submit" label="Зарегистрироваться" :disabled="isButtonDisabled" />
      </Form>
      <Button
        label="Вернуться"
        variant="text"
        class="register-back-button"
        @click="navigateToAuthPage"
      />
    </div>
  </div>
</template>

<style scoped>
.register {
  display: flex;
  height: 80vh;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  gap: 8px;
}

.register-ruler {
  height: 1px;
  width: 35vw;
  color: white;
}

.register-form-container {
  padding: 20px 80px;
  display: flex;
  flex-direction: column;
  gap: 16px;
  align-items: center;
  justify-content: center;
}

.register-form {
  width: 250px;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.register-back-button {
  width: 100%;
}
</style>
