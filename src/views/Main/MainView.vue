<script setup lang="ts">
// Components
import Table from '@/components/Table.vue'

import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import MainHeader from './components/MainHeader.vue'
import MainFooter from './components/MainFooter.vue'
import type { Request } from '@/models/Request'
import { onMounted, ref, useTemplateRef } from 'vue'
import RequestsController from '@/controllers/requestsController'
import { AxiosError } from 'axios'
import { useToast } from 'primevue'
import EventsController from '@/controllers/eventsController'

const inputForm = useTemplateRef<HTMLFormElement>('inputForm')

const toast = useToast()
const requests = ref<Request[]>([])

interface FilterFormItems {
  event_name: string
  author_login: string
  created_after: string
}

const handleSearchRequests = async (event: MouseEvent) => {
  event.preventDefault()
  if (!inputForm.value) return

  try {
    const filterForm = Array.from(inputForm.value.elements) as HTMLInputElement[]

    const payload: FilterFormItems = {
      event_name: filterForm[0].value,
      author_login: filterForm[1].value,
      created_after: new Date(filterForm[2].value).toISOString(),
    }

    console.log(payload)

    let event_id = undefined
    if (payload.event_name) {
      const commonEvent = await EventsController.getEvents(payload.event_name)
      event_id = commonEvent[0].event_id
    }

    requests.value = await RequestsController.getRequests(
      undefined,
      event_id,
      payload.author_login,
      payload.created_after,
    )
  } catch (e) {
    if (e instanceof AxiosError) {
      if (e.status === 404)
        toast.add({
          severity: 'info',
          closable: false,
          summary: 'Нет заявок по заданным фильтрам',
          life: 1500,
        })
      else
        toast.add({
          severity: 'error',
          closable: false,
          summary: 'Ошибка при выводе заявок',
          life: 3000,
        })
    }
    return
  }
}

onMounted(async () => {
  try {
    requests.value = await RequestsController.getRequests()
  } catch {
    toast.add({
      severity: 'error',
      closable: false,
      summary: 'Ошибка при выводе заявок',
      life: 3000,
    })
  }
})
</script>
<template>
  <div>
    <MainHeader />
    <form ref="inputForm" class="main-filters">
      <InputText name="event_name" placeholder="Мероприятие" />
      <InputText name="author_login" placeholder="Автор" />
      <InputText type="datetime-local" name="created_after" placeholder="Создана после" />
      <Button label="Поиск" variant="outlined" type="submit" @click="handleSearchRequests" />
    </form>
    <div class="main-table">
      <Table showGridlines :requests />
    </div>
    <MainFooter />
  </div>
</template>

<style scoped>
.main-filters {
  display: flex;
  align-items: center;
  gap: 20px;
  margin-bottom: 36px;
}
.main-table {
  margin-bottom: 56px;
}
</style>
