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
import { onMounted } from 'vue'
import EventsController from '@/controllers/eventsController'
import type { Event } from '@/models/Event'
import { AxiosError } from 'axios'
import { useRouter } from 'vue-router'
import FilesController from '@/controllers/filesController'
import RequestsController from '@/controllers/requestsController'

const toast = useToast()
const router = useRouter()

const fileInput = ref<Attachment | undefined>()
const selectOptions = ref<Event[]>([])
const initialValues = reactive<Omit<NewRequest, 'attachments'>>({
  event_id: '',
  description: '',
  status: 'new',
})

onMounted(async () => {
  try {
    const response = await EventsController.getEvents()

    selectOptions.value = response
  } catch (e) {
    if (e instanceof AxiosError) {
      toast.add({
        severity: 'error',
        closable: false,
        summary: e.message,
        life: 3000,
      })
    }
    return
  }
})

const handleFileLoad = async (e: FileUploadSelectEvent) => {
  toast.add({
    severity: 'info',
    summary: 'Файл добавлен',
    life: 1400,
  })

  const files = e.files as FileType[]

  console.log(files)

  // const files = new Blob()
  if (!files.length) {
    fileInput.value = undefined
    return
  }

  const blob = new Blob([files[0]], { type: 'text/plain' })

  const response = await FilesController.loadFile(blob)

  fileInput.value = {
    filename: files[0].name,
    id: response.file_id,
  }
}

const resolver = ({ values }: FormResolverOptions) => {
  const errors: NewRequestErrors = {}

  if (!values.event_id) errors.event_id = [{ message: 'Не указан id мероприятия' }]

  return {
    errors,
  }
}

const onFormSubmit = async ({ valid, states, reset }: FormSubmitEvent) => {
  console.log('here')
  if (valid) {
    try {
      const result = parseFormResult<NewRequest>(states)
      const files = fileInput.value

      const response = await RequestsController.createRequest({
        ...result,
        status: 'new',
        attachments: files ? [files] : [],
      })

      if (response) {
        toast.add({
          severity: 'success',
          summary: 'Заявка подана',
          life: 3000,
        })
        router.push('/')
        reset()
      }
    } catch (e) {
      if (e instanceof AxiosError) {
        toast.add({
          severity: 'error',
          summary: e.message,
          life: 3000,
        })
      }
      return
    }
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
        <Select
          :loading="!selectOptions.length"
          :options="selectOptions"
          optionLabel="name"
          optionValue="id"
          name="event_id"
          placeholder="Выберите мероприятие"
        />
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
      <FileUpload auto name="attachments" @select="handleFileLoad" />
      <Button type="submit" label="Отправить" />
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
