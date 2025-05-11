<script setup lang="ts">
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
import type { Request } from '@/models/dto/Reservation'
import Caret from '@/assets/caret-right.svg'
import parseDate from '@/shared/parseDate'

defineProps<{ requests: Request[] }>()
</script>

<template>
  <DataTable :value="requests" striped-rows>
    <Column field="id" header="ID">
      <template #body="slotProps">
        <RouterLink class="table-link" :to="`/requests/${slotProps.data.id}`">
          <span>{{ slotProps.data.id }}</span>
          <img :src="Caret" :alt="Caret" />
        </RouterLink>
      </template>
    </Column>
    <Column field="author.name" header="Автор" />
    <Column field="created_at" header="Создана">
      <template #body="slotProps">
        <span>{{ parseDate(slotProps.data.created_at) }}</span>
      </template>
    </Column>
  </DataTable>
</template>

<style scoped>
.table-link {
  text-decoration: none;
  color: black;
  display: flex;
  gap: 4px;
}
</style>
