<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import type { BookingResponseDto } from '@/models/dto/Booking'
import { BookingService } from '@/services/BookingService'

const route = useRoute()

const data = ref<BookingResponseDto>()

const requestInitData = async () => {
  const booking = await BookingService.getBookingById(Number(route.params.id))

  data.value = booking
}

onMounted(() => requestInitData())
</script>
<template>
  <div></div>
</template>

<style scoped></style>
