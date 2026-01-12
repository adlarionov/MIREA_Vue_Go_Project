<script setup lang="ts">
import { reactive } from 'vue'
import Button from 'primevue/button'
import { Form, type FormResolverOptions, type FormSubmitEvent } from '@primevue/forms'
import { InputText, Textarea } from 'primevue'
import type { EventRequestDto } from '@/models/dto/Event'
import { EventService } from '@/services/EventService'

const initialValues = reactive<Partial<EventRequestDto>>({
  description: '',
  name: '',
})

const resolver = ({ values }: FormResolverOptions) => {
  return { values }
}

const handleCreateEvent = async ({ values }: FormSubmitEvent) => {
  EventService.createEvent(values as EventRequestDto)

  return values
}
</script>

<template>
  <div class="form-wrapper">
    <h1>Создать событие</h1>
    <Form class="form" @submit="handleCreateEvent" :initialValues :resolver form>
      <InputText name="name" placeholder="Введите название события" />
      <Textarea name="description" placeholder="Введите описание события" width="100%" autoResize />
      <Button type="submit" label="Отправить" style="width: 120px" />
    </Form>
  </div>
</template>

<style scoped>
.form-wrapper {
  margin-top: 56px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 24px;
}
.form {
  display: flex;
  flex-direction: column;
  gap: 16px;
  width: 80vw;
}
</style>
