<script setup lang="ts">
import Textarea from 'primevue/textarea'
import Select from 'primevue/select'
import FileUpload, { type FileUploadSelectEvent } from 'primevue/fileupload'
import Button from 'primevue/button'
import Message from 'primevue/message'
import { Form, type FormResolverOptions, type FormSubmitEvent } from '@primevue/forms'
import { reactive, ref } from 'vue'
import { type NewRequest, type NewRequestErrors } from '@/models/Request'
import { useToast } from 'primevue/usetoast'
import { Toast } from 'primevue'
import { parseFormResult } from '@/shared/parseFormResult'
import type Attachment from '@/models/Attachment'
import type { FileType } from '@/models/Attachment'
import Header from '@/components/Header.vue'

const toast = useToast()

const fileInput = ref<Attachment[]>([])
const initialValues = reactive<Omit<NewRequest, 'attachments'>>({
  event_id: '',
  description: '',
  status: 'new',
})

const handleFileLoad = (e: FileUploadSelectEvent) => {
  toast.add({
    severity: 'info',
    summary: 'Файл добавлен',
    life: 1400,
  })

  const files = (e.files as FileType[]).map<Attachment>((file) => ({
    filename: file.name,
    id: file.objectURL.split('/').pop() ?? '',
  }))

  fileInput.value = files
}

const resolver = ({ values }: FormResolverOptions) => {
  const errors: NewRequestErrors = {}

  if (!values.event_id) errors.event_id = [{ message: 'Не указан id мероприятия' }]

  return {
    errors,
  }
}

const onFormSubmit = ({ valid, states }: FormSubmitEvent) => {
  if (valid) {
    const result = parseFormResult<NewRequest>(states)
    const files = fileInput.value

    console.log({ ...result, attachments: files })

    toast.add({
      severity: 'success',
      summary: 'Заявка подана.',
      life: 3000,
    })
  }
}
</script>
<template>
  <Form
    v-slot="$form"
    :initialValues
    :resolver
    class="requests-new-container"
    @submit="onFormSubmit"
  >
    <Header back title="Новая заявка" />
    <div class="requests-new-input">
      <div class="requests-flex">
        <Select :options="['test']" name="event_id" placeholder="Выберите мероприятие" />
        <Message v-if="$form.event_id?.invalid" severity="error" size="small" variant="simple">{{
          $form.event_id.error?.message
        }}</Message>
      </div>
      <div class="requests-flex">
        <Textarea name="description" placeholder="Описание" style="resize: none" rows="12" />
        <Message v-if="$form.description?.invalid" severity="error" size="small" variant="simple">{{
          $form.description.error?.message
        }}</Message>
      </div>
    </div>
    <div class="requests-new-footer">
      <FileUpload auto multiple name="attachments" mode="basic" @select="handleFileLoad" />
      <Button :disabled="!fileInput?.length" type="submit" label="Отправить" />
    </div>
  </Form>
  <Toast />
</template>

<style scoped>
.requests-new-container {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.requests-flex {
  display: flex;
  flex-direction: column;
  width: 100%;
}

.requests-new-input {
  display: flex;
  width: 100%;
  flex-direction: column;
  gap: 36px;
}

.requests-new-footer {
  display: flex;
  gap: 16px;
  justify-content: end;
  align-items: center;
}
</style>
