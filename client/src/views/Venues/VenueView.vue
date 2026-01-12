<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { Card, Panel } from 'primevue'
import { formImageUrl } from '@/shared/formImageUrl'
import Header from '@/components/Header.vue'
import type { VenueResponseDto } from '@/models/dto/Venue'
import { VenueService } from '@/services/VenueService'

const route = useRoute()
const pageId = route.params.id

const data = ref<VenueResponseDto>()

const requestInitData = async () => {
  const venue = await VenueService.getVenueById(Number(pageId))

  data.value = venue
}

onMounted(() => requestInitData())
</script>
<template>
  <Header back to="/venues" title="Локации" style="margin: 12px 0px" />
  <Card style="height: 600px">
    <template #title>Локация №{{ data?.id }}</template>
    <template #content>
      <Panel toggleable :header="`Локация - ${data?.name}`">
        <p>{{ data?.description }}</p>
        <p>Вместимость - {{ data?.capacity }}</p>
        <p>Адрес - {{ data?.address }}</p>
        <img class="image" :src="formImageUrl(data?.Image.image_url || '')" alt="venue image" />
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
