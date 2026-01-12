<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { Card, Panel } from 'primevue'
import Header from '@/components/Header.vue'
import type { EventResponseDto } from '@/models/dto/Event'
import { EventService } from '@/services/EventService'

const route = useRoute()
const pageId = route.params.id

const data = ref<EventResponseDto>()

const requestInitData = async () => {
  const event = await EventService.getEventById(Number(pageId))

  data.value = event
}

onMounted(() => requestInitData())
</script>
<template>
  <Header back to="/events" title="События" style="margin: 12px 0px" />
  <Card style="height: 600px">
    <template #title>Событие №{{ data?.id }}</template>
    <template #subtitle>
      <span> Организация: {{ data?.Organization.name }} </span>
    </template>
    <template #content>
      <Panel toggleable>
        <p>{{ data?.description }}</p>
      </Panel>
    </template>
  </Card>
</template>
