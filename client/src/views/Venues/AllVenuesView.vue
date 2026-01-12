<script setup lang="ts">
import Table from '@/components/Table.vue'
import Button from 'primevue/button'
import { venueTableSchema } from '@/models/ui/Table'
import { onMounted, ref } from 'vue'
import router from '@/router'
import { flatVenueSchema } from '@/views/Bookings/utils/tableSchemaBuilder'
import type { VenueTableItem } from '@/models/dto/Venue'
import { VenueService } from '@/services/VenueService'

const data = ref<VenueTableItem[]>([])

const requestInitData = async () => {
  const venues = await VenueService.getVenues()

  if (venues && venues.length > 0) data.value = flatVenueSchema(venues)
}

onMounted(() => requestInitData())

const navigateToCreateVenue = () => router.push('/venues/new')
</script>

<template>
  <div class="all-venues-table-wrapper">
    <h1>Локации</h1>
    <Table :columns="venueTableSchema" :data="data" style="max-height: 800px" />
    <Button style="margin-top: 24px" @click="navigateToCreateVenue">Создать локацию</Button>
  </div>
</template>

<style scoped>
.all-venues-table-wrapper {
  margin: 36px 0px;
}
</style>
