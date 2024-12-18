<script setup lang="ts">
import Textarea from 'primevue/textarea'
import Header from '@/components/Header.vue'
import { Button, Tag, useToast } from 'primevue'
import File from '@/assets/file-icon.png'
import { ref, onMounted } from 'vue'
import RequestsController from '@/controllers/requestsController'
import { useRoute, useRouter } from 'vue-router'
import type { RequestDto } from '@/models/Request'

const route = useRoute()

const router = useRouter()
const toast = useToast()

const data = ref<RequestDto>()

onMounted(async () => {
  try {
    if (!Array.isArray(route.params.request_id))
      data.value = await RequestsController.getRequestById(route.params.request_id)
  } catch {
    toast.add({
      severity: 'error',
      closable: false,
      summary: 'Такой заявки не существует',
      life: 3000,
    })
    router.push('/')
  }
})
</script>
<template>
  <div>
    <Header back title="Заявка" />
    <div class="request-info">
      <div class="request-info-files">
        <Textarea disabled style="resize: none" rows="8" />
        <div class="request-info-files-container">
          <div class="file">
            <img width="60" :src="File" :alt="File" />
            <span>file 1</span>
          </div>
          <div class="file">
            <img width="60" :src="File" :alt="File" />
            <span>file 1</span>
          </div>
        </div>
      </div>
      <div class="request-info-table">
        <h4>id: <span style="margin-left: 8px"></span></h4>
        <h4>Автор: <span style="margin-left: 8px"></span></h4>
        <h4>Мероприятие: <span style="margin-left: 8px"></span></h4>
        <h4>Создана: <span style="margin-left: 8px"></span></h4>
        <h4>Обновлена: <span style="margin-left: 8px"></span></h4>
        <h4>
          Статус: <Tag style="margin-left: 8px"><span>test</span></Tag>
        </h4>
      </div>
    </div>
    <div class="request-comments">
      <h3 class="request-comments-title">Комментарии</h3>
      <Textarea rows="10" disabled style="resize: none" />
      <div class="request-comments-description">
        <p>date</p>
        <h4>author</h4>
        <Button label="Новый комментарий" />
      </div>
    </div>
  </div>
</template>

<style scoped>
.request-info {
  display: flex;
  justify-content: space-between;
  margin-bottom: 36px;
}

.request-info-files {
  display: flex;
  flex-direction: column;
  flex-basis: 60%;
  gap: 20px;
}

.request-info-files-container {
  display: flex;
  gap: 16px;
}

.file {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
}

.request-info-table {
  display: flex;
  flex-direction: column;
  flex-basis: 30%;
  gap: 8px;
}

.request-comments {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.request-comments-title {
  text-align: center;
  margin-bottom: 28px;
}

.request-comments-description {
  display: flex;
  flex-direction: column;
  align-items: end;
  gap: 8px;
}
</style>
