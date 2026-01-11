<script setup lang="ts">
import Textarea from 'primevue/textarea'
import Button from 'primevue/button'
import Message from 'primevue/message'
import { Form, type FormResolverOptions, type FormSubmitEvent } from '@primevue/forms'
import { reactive } from 'vue'
import { useToast } from 'primevue/usetoast'
import { InputText, Toast } from 'primevue'
import { parseFormResult } from '@/shared/parseFormResult'
import Header from '@/components/Header.vue'
import { AxiosError } from 'axios'

const initialValues = reactive({
  description: '',
  name: '',
})

const resolver = ({ values }: FormResolverOptions) => {
  const errors = {}

  return {
    errors,
  }
}

const onFormSubmit = async ({ valid, states, reset }: FormSubmitEvent) => {
  if (valid) {
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
