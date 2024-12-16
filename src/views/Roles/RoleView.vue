<script setup lang="ts">
import BackButton from '@/components/BackButton.vue'
import { Role, type NewRole, type NewRoleErrors } from '@/models/Role'
import { parseFormResult } from '@/shared/parseFormResult'
import { Form, type FormResolverOptions, type FormSubmitEvent } from '@primevue/forms'
import { Button, InputText, useToast, Select, Message, Toast } from 'primevue'
import { reactive } from 'vue'

const toast = useToast()

const roles = [Role.USER, Role.ADMIN, Role.MODERATOR]

const initialValues = reactive<NewRole>({
  login: '',
  new_role: Role.USER,
})

const resolver = ({ values }: FormResolverOptions) => {
  const errors: NewRoleErrors = {}

  if (!values.login) errors.login = [{ message: 'Не указан логин' }]
  if (!values.new_role) errors.new_role = [{ message: 'Не указана роль' }]

  return {
    errors,
  }
}

const onFormSubmit = ({ valid, states }: FormSubmitEvent) => {
  console.log(valid)
  if (valid) {
    toast.add({
      severity: 'success',
      summary: 'Роль выдана.',
      life: 3000,
    })

    console.log(parseFormResult<NewRole>(states))
  }
}
</script>

<template>
  <div class="role-header">
    <BackButton />
    <h1 class="role-header-title">Выдать роль</h1>
  </div>
  <Form class="role-form" v-slot="$form" :initialValues :resolver @submit="onFormSubmit">
    <div class="role-form-full-width">
      <InputText name="login" placeholder="Логин" />
      <Message v-if="$form.login?.invalid" severity="error" size="small" variant="simple">
        {{ $form.login.error?.message }}
      </Message>
    </div>
    <div class="role-form-full-width">
      <Select :options="roles" name="new_role" placeholder="Роль" />
      <Message v-if="$form.role?.invalid" severity="error" size="small" variant="simple">
        {{ $form.role.error?.message }}
      </Message>
    </div>
    <div class="role-form-footer">
      <Button label="Добавить" type="submit" />
    </div>
  </Form>
  <Toast />
</template>

<style scoped>
.role-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 52px;
}

.role-header-title {
  width: 100%;
  text-align: center;
}

.role-form {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.role-form-footer {
  display: flex;
  justify-content: end;
}

.role-form-full-width {
  display: flex;
  flex-direction: column;
}
</style>
