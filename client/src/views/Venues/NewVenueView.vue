<script setup lang="ts">
import { reactive, ref } from 'vue'
import Button from 'primevue/button'
import { Form, type FormResolverOptions, type FormSubmitEvent } from '@primevue/forms'
import { FileUpload, InputNumber, InputText, Textarea, type FileUploadSelectEvent } from 'primevue'
import { VenueService } from '@/services/VenueService'
import type { VenueRequestDto } from '@/models/dto/Venue'

const initialValues = reactive<Partial<VenueRequestDto['venue']>>({
  address: '',
  capacity: 0,
  description: '',
  name: '',
})

const venueId = ref<number>()

const onFileSelect = async (event: FileUploadSelectEvent) => {
  const file = event.files[0]

  const formData = new FormData()
  formData.append('Image', file)

  console.log(formData)

  if (venueId.value) {
    await VenueService.addImageToVenue(venueId.value, formData)
  }
}

const resolver = ({ values }: FormResolverOptions) => {
  return { values }
}

const handleCreateVenue = async ({ values }: FormSubmitEvent) => {
  const newVenueId = await VenueService.createVenue(values as VenueRequestDto['venue'])

  venueId.value = newVenueId

  return values
}
</script>

<template>
  <div class="form-wrapper">
    <h1>Создать локацию</h1>
    <Form class="form" @submit="handleCreateVenue" :initialValues :resolver form>
      <InputText name="name" placeholder="Название" />
      <InputNumber name="capacity" placeholder="Введите максимальное кол-во людей" />
      <InputText name="address" placeholder="Адрес" />
      <Textarea name="description" placeholder="Введите описание локации" width="100%" autoResize />

      <Button type="submit" label="Отправить" style="width: 120px" s />
    </Form>
    <div v-if="Boolean(venueId)">
      <FileUpload
        mode="basic"
        name="image"
        customUpload
        accept="image/*"
        :maxFileSize="1000000"
        @select="onFileSelect"
      />
    </div>
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
