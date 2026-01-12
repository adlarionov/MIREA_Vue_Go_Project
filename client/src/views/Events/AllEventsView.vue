<script setup lang="ts">
import Table from '@/components/Table.vue'
import Button from 'primevue/button'
import { eventTableSchema } from '@/models/ui/Table'
import { onMounted, ref } from 'vue'
import router from '@/router'
import { flatEventSchema } from '@/views/Bookings/utils/tableSchemaBuilder'
import { EventService } from '@/services/EventService'
import type { EventTableItem } from '@/models/dto/Event'

const data = ref<EventTableItem[]>([])

const requestInitData = async () => {
  const events = await EventService.getEvents()

  if (events && events.length > 0) data.value = flatEventSchema(events)
}

onMounted(() => requestInitData())

const navigateToCreateEvent = () => router.push('/events/new')
</script>

<template>
  <div class="all-events-table-wrapper">
    <h1>События</h1>
    <Table :columns="eventTableSchema" :data="data" style="max-height: 800px" />
    <Button style="margin-top: 24px" @click="navigateToCreateEvent">Создать событие</Button>
  </div>
</template>

<style scoped>
.all-events-table-wrapper {
  margin: 36px 0px;
}
</style>
