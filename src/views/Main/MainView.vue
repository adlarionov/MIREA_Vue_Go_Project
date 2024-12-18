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

const inputForm = useTemplateRef<HTMLFormElement>('inputForm')

const requests = ref<Request[]>([])

// const requests: Request[] = [
//   {
//     id: '3fa85f64-5717-4562-b3fc-2c963f66afa6',
//     author: { login: 'Ярослав', id: 0, name: 'Ярослав', phone: '+797446465' },
//     event: 'Highload++',
//     created_at: '2024-12-18T20:40:24.063Z',
//   },
//   {
//     id: '3fa85f64-5717-4562-b3fc-2c963f66afa6',
//     author: { login: 'Ярослав', id: 0, name: 'Ярослав', phone: '+797446465' },
//     event: 'Highload++',
//     created_at: '2024-12-18T20:40:24.063Z',
//   },
//   {
//     id: '3fa85f64-5717-4562-b3fc-2c963f66afa6',
//     author: { login: 'Ярослав', id: 0, name: 'Ярослав', phone: '+797446465' },
//     event: 'Highload++',
//     created_at: '2024-12-18T20:40:24.063Z',
//   },
// ]

interface FilterFormItems {
  event_name: string
  author_login: string
  created_after: string
}

const handleSearchRequests = async (event: MouseEvent) => {
  event.preventDefault()
  if (!inputForm.value) return

  const filterForm = Array.from(inputForm.value.elements) as HTMLInputElement[]

  const payload: FilterFormItems = {
    event_name: filterForm[0].value,
    author_login: filterForm[1].value,
    created_after: new Date(filterForm[2].value).toISOString(),
  }

  console.log(payload)

  requests.value = await RequestsController.getRequests(
    undefined,
    payload.event_name,
    payload.author_login,
    payload.created_after,
  )
}

onMounted(async () => {
  requests.value = await RequestsController.getRequests()
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
