<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import type { BookingResponseDto } from '@/models/dto/Booking'
import { BookingService } from '@/services/BookingService'
import { Card, Panel, Tag } from 'primevue'
import { formImageUrl } from '@/shared/formImageUrl'
import Header from '@/components/Header.vue'

const route = useRoute()
const pageId = route.params.id

const data = ref<BookingResponseDto>()

const requestInitData = async () => {
  const booking = await BookingService.getBookingById(Number(pageId))

  data.value = booking
}

const severityByStatus: Record<string, string> = {
  active: 'success',
  unactive: 'danger',
}

onMounted(() => requestInitData())
</script>
<template>
  <Header back to="/bookings" title="Заявка" style="margin: 12px 0px" />
  <Card style="height: 600px">
    <template #title>Заявка №{{ data?.id }}</template>
    <template #subtitle>
      <span> Цена: {{ data?.price }} RUB </span>
      <span>
        Статус: <Tag :severity="severityByStatus[data?.status || 'active']">{{ data?.status }}</Tag>
      </span>
    </template>
    <template #content>
      <Panel toggleable :header="`Событие - ${data?.Event.name}`">
        <p>{{ data?.Event.description }}</p>
      </Panel>
      <Panel toggleable :header="`Локация - ${data?.Venue.name}`">
        <p>{{ data?.Venue.description }}</p>
        <p>Вместимость - {{ data?.Venue.capacity }}</p>
        <p>Адрес - {{ data?.Venue.address }}</p>
        <img
          class="image"
          :src="formImageUrl(data?.Venue.Image.image_url || '')"
          alt="venue image"
        />
      </Panel>
    </template>
  </Card>
</template>

<style scoped>
.image {
  width: 100%;
  height: 100%;
}
</style>
