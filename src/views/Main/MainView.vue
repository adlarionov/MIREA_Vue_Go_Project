<script setup lang="ts">
// Components
import Table from '@/components/Table.vue'

import InputText from 'primevue/inputtext'
import MainHeader from './components/MainHeader.vue'
import MainFooter from './components/MainFooter.vue'
import type { Request } from '@/models/Request'

const requests: Request[] = [
  {
    id: '11111111-2222-4444-2222',
    author: { login: 'Ярослав', id: 0, name: 'Ярослав', phone: '+797446465' },
    event: 'Highload++',
    createdAt: '11.01.2002 12:12',
  },
  {
    id: '11111111-2222-4444-2222',
    author: { login: 'Ярослав', id: 0, name: 'Ярослав', phone: '+797446465' },
    event: 'Highload++',
    createdAt: '11.01.2002 12:12',
  },
  {
    id: '11111111-2222-4444-2222',
    author: { login: 'Ярослав', id: 0, name: 'Ярослав', phone: '+797446465' },
    event: 'Highload++',
    createdAt: '11.01.2002 12:12',
  },
]
</script>
<template>
  <div>
    <MainHeader />
    <div class="main-filters">
      <InputText placeholder="Мероприятие" />
      <InputText placeholder="Автор" />
      <InputText placeholder="Создана после" />
    </div>
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
