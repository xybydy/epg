<template>
  <Dialog v-model:visible="visible" :style="{width: '450px'}" header="Sil" modal dismissableMask>
    <div class="p-d-flex p-ai-center confirmation-content">
      <i class="pi pi-exclamation-triangle p-mr-3" style="font-size: 2rem" />
      <span>Seçtiğin kanallar silinsin mi?</span>
    </div>
    <template #footer>
      <Button label="No" icon="pi pi-times" class="p-button-text" @click="visible = false" />
      <Button label="Yes" icon="pi pi-check" @click="onYes" />
    </template>
  </Dialog>
</template>

<script setup>
  import {ref, onMounted} from 'vue'

  import eventBus from '@/emitter/eventBus.js'

  import Dialog from 'primevue/dialog'
  import Button from 'primevue/button'

  let visible = ref(false)

  const onYes = () => {
    visible.value = false
    eventBus.emit('removeSelectedItems')
  }

  onMounted(() => {
    eventBus.on('selectedRemoveDialog', () => (visible.value = true))
  })
</script>
