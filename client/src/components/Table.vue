<script setup lang="ts">
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
import Button from 'primevue/button'
import type { ColumnTypeMap, TableDataMap } from '@/models/ui/Table'
import router from '@/router'

interface OwnProps {
  columns: ColumnTypeMap[]
  data: TableDataMap[]
}

const { columns, data } = defineProps<OwnProps>()

const onRowOpen = (data: Record<string, unknown>) => {
  router.push(router.currentRoute.value.path + '/' + data.id)
}
</script>

<template>
  <DataTable :value="data" striped-rows>
    <Column
      v-for="column in columns"
      :key="column.field"
      :field="column.field"
      :header="column.header"
    />
    <Column>
      <template #body="{ data }">
        <Button @click="() => onRowOpen(data)">Открыть</Button>
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
