<script setup lang="ts">
import Textarea from 'primevue/textarea'
import Button from 'primevue/button'
import Message from 'primevue/message'
import { Form, type FormResolverOptions, type FormSubmitEvent } from '@primevue/forms'
import { reactive } from 'vue'
import { useToast } from 'primevue/usetoast'
import { InputText, Toast } from 'primevue'
import { parseFormResult } from '@/shared/parseFormResult'
import type { NewEvent, NewEventError } from '@/models/Event'
import Header from '@/components/Header.vue'
import EventsController from '@/controllers/eventsController'
import { AxiosError } from 'axios'

const toast = useToast()

const initialValues = reactive<NewEvent>({
  description: '',
  name: '',
})

const resolver = ({ values }: FormResolverOptions) => {
  const errors: NewEventError = {}

  if (!values.name) errors.name = [{ message: 'Мероприятие не названо' }]
  if (!values.description) errors.description = [{ message: 'Мероприятие без описания' }]

  return {
    errors,
  }
}

const onFormSubmit = async ({ valid, states, reset }: FormSubmitEvent) => {
  if (valid) {
    const values = parseFormResult<NewEvent>(states)

    try {
      const response = await EventsController.createEvent(values)

      if (response)
        toast.add({
          severity: 'success',
          summary: 'Мероприятие успешно создано',
          life: 3000,
        })

      reset()
    } catch (e) {
      if (e instanceof AxiosError) {
        toast.add({
          severity: 'error',
          life: 3000,
          summary: e.message,
          closable: false,
        })
      }
      return
    }
  }
}
</script>
<template>
  <Form v-slot="$form" :initialValues :resolver class="events-new-container" @submit="onFormSubmit">
    <Header back title="Новое мероприятие" />
    <div class="event-new-input">
      <div class="event-flex">
        <InputText name="name" placeholder="Название" />
        <Message v-if="$form.name?.invalid" severity="error" size="small" variant="simple">{{
          $form.name.error?.message
        }}</Message>
      </div>
      <div class="event-flex">
        <Textarea name="description" placeholder="Описание" style="resize: none" rows="12" />
        <Message v-if="$form.description?.invalid" severity="error" size="small" variant="simple">{{
          $form.description.error?.message
        }}</Message>
      </div>
    </div>
    <div class="event-new-footer">
      <Button type="submit" label="Добавить" />
    </div>
  </Form>
  <Toast />
</template>

<style scoped>
.event-new-container {
  display: flex;
  flex-direction: column;
  gap: 16px;
}
.event-new-title {
  text-align: center;
  width: 100%;
}

.event-flex {
  display: flex;
  flex-direction: column;
  width: 100%;
}

.event-new-input {
  display: flex;
  width: 100%;
  flex-direction: column;
  gap: 36px;
  margin-bottom: 24px;
}

.event-new-footer {
  display: flex;
  gap: 16px;
  justify-content: end;
  align-items: center;
}
</style>
