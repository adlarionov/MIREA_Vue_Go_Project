<script setup lang="ts">
import { computed, onMounted, reactive, ref } from 'vue'
import Select from 'primevue/select'
import Button from 'primevue/button'
import { Form, type FormResolverOptions, type FormSubmitEvent } from '@primevue/forms'
import type { BookingRequestDto } from '@/models/dto/Booking'
import { InputNumber } from 'primevue'
import { VenueService } from '@/services/VenueService'
import type { VenueResponseDto } from '@/models/dto/Venue'
import type { EventResponseDto } from '@/models/dto/Event'
import { EventService } from '@/services/EventService'
import { getOptionsFromData } from '@/shared/getOptionsFromData'
import { BookingService } from '@/services/BookingService'

const initialValues = reactive<Partial<BookingRequestDto>>({
  venue_id: 0,
  event_id: 0,
  price: 0,
  status: 'active',
})

const venueRef = ref<VenueResponseDto[]>([])
const eventRef = ref<EventResponseDto[]>([])

const availableEvents = computed(() => getOptionsFromData(eventRef.value, 'id', 'name'))
const availableVenues = computed(() => getOptionsFromData(venueRef.value, 'id', 'name'))

const statuses = [
  { label: 'Активная', value: 'active' },
  { label: 'Неактивная', value: 'unactive' },
]

onMounted(async () => {
  const [venues, events] = await Promise.all([VenueService.getVenues(), EventService.getEvents()])

  eventRef.value = events
  venueRef.value = venues
})

const resolver = ({ values }: FormResolverOptions) => {
  return { values }
}

const handleCreateBooking = async ({ values }: FormSubmitEvent) => {
  BookingService.createBooking(values as BookingRequestDto)

  return values
}
</script>

<template>
  <div class="form-wrapper">
    <h1>Создать заявку</h1>
    <Form class="form" @submit="handleCreateBooking" :initialValues :resolver form>
      <Select
        name="event_id"
        :options="availableEvents"
        optionLabel="label"
        optionValue="value"
        placeholder="Выберите мероприятие"
      />

      <Select
        name="venue_id"
        :options="availableVenues"
        optionLabel="label"
        optionValue="value"
        placeholder="Выберите локацию"
      />
      <InputNumber name="price" placeholder="Введите цену" />
      <Select
        name="status"
        :options="statuses"
        optionLabel="label"
        optionValue="value"
        placeholder="Выберите статус заявки"
      />
      <Button type="submit" label="Отправить" />
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
  max-width: 400px;
}
</style>
