<script setup lang="ts">
import { reactive, ref } from 'vue'
import Button from 'primevue/button'
import { Form, type FormResolverOptions, type FormSubmitEvent } from '@primevue/forms'
import { FileUpload, InputNumber, InputText, Textarea, type FileUploadUploadEvent } from 'primevue'
import { VenueService } from '@/services/VenueService'
import type { VenueRequestDto } from '@/models/dto/Venue'
import URLs, { API_CONTEXT } from '@/api/URLs'

const initialValues = reactive<Partial<VenueRequestDto>>({
  address: '',
  capacity: 0,
  description: '',
  name: '',
})

const venueId = ref<number>()

const fileupload = ref()

const upload = () => {
  fileupload.value.upload()
}

const resolver = ({ values }: FormResolverOptions) => {
  return { values }
}

const handleCreateVenue = async ({ values }: FormSubmitEvent) => {
  const newVenueId = await VenueService.createVenue(values as VenueRequestDto)

  venueId.value = newVenueId

  return values
}

const onFileUpload = (event: FileUploadUploadEvent) => {
  console.log(event)
}
</script>

<template>
  <div class="form-wrapper">
    <h1>Создать локацию</h1>
    <Form class="form" @submit="handleCreateVenue" :initialValues :resolver form>
      <InputText name="name" placeholder="Название" />
      <InputNumber name="price" placeholder="Введите максимальное кол-во людей" />
      <InputText name="address" placeholder="Адрес" />
      <Textarea name="description" placeholder="Введите описание локации" width="100%" autoResize />

      <Button type="submit" label="Отправить" style="width: 120px" s />
    </Form>
    <!-- <div v-if="Boolean(venueId)"> -->
    <FileUpload
      ref="fileupload"
      mode="basic"
      name="image"
      :url="venueId ? `${API_CONTEXT}${URLs.addImageToVenueById(venueId)}` : ''"
      accept="image/*"
      :maxFileSize="1000000"
      @upload="onFileUpload"
    />
    <Button label="Upload" @click="upload" severity="secondary" />
    <!-- </div> -->
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
