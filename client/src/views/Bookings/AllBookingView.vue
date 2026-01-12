<script setup lang="ts">
import Table from '@/components/Table.vue'
import Button from 'primevue/button'
import type { BookingTableItem } from '@/models/dto/Booking'
import { bookingTableSchema } from '@/models/ui/Table'
import { BookingService } from '@/services/BookingService'
import { onMounted, ref } from 'vue'
import router from '@/router'
import { flatBookingSchema } from '@/views/Bookings/utils/tableSchemaBuilder'

const data = ref<BookingTableItem[]>([])

const requestInitData = async () => {
  const bookings = await BookingService.getBookings()

  if (bookings && bookings.length > 0) data.value = flatBookingSchema(bookings)
}

onMounted(() => requestInitData())

const navigateToCreateBooking = () => router.push('/bookings/new')
</script>

<template>
  <div class="all-bookings-table-wrapper">
    <h1>Заявки</h1>
    <Table :columns="bookingTableSchema" :data="data" style="max-height: 800px" />
    <Button style="margin-top: 24px" @click="navigateToCreateBooking">Создать заявку</Button>
  </div>
</template>

<style scoped>
.all-bookings-table-wrapper {
  margin: 36px 0px;
}
</style>
