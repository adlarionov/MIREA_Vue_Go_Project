<script setup lang="ts">
import BackButton from '@/components/BackButton.vue'

defineProps<{ back?: boolean; title: string }>()
</script>
<template>
  <header class="header">
    <BackButton v-if="back" />
    <h1
      class="header-new-title"
      :style="{ marginLeft: !!$slots.buttons ? '210px' : 0, marginRight: back ? '30px' : 0 }"
    >
      {{ title }}
    </h1>
    <div class="slot">
      <slot name="buttons"></slot>
    </div>
  </header>
</template>

<style scoped>
.header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 48px;
}

.slot {
  display: flex;
  gap: 16px;
}

.header-new-title {
  text-align: center;
  width: 100%;
}
</style>
