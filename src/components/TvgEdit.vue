<template>
  <Dialog
    v-model:visible="visible"
    :style="{width: '350px'}"
    header="TVG-ID Düzenle"
    modal
    dismissableMask
    closable
  >
    <div class="p-d-flex p-jc-between p-ai-center p-input-filled p-p-2">
      <span>Yeni TVG-ID: </span>
      <InputText v-model="newVal" autofocus />
    </div>
    <template #footer>
      <Button label="Vazgeç" icon="pi pi-times" class="p-button-text" @click="visible = false" />
      <Button label="Onayla" icon="pi pi-check" @click="onYes" />
    </template>
  </Dialog>
</template>

<script setup>
  import {ref, onMounted} from 'vue'
  import eventBus from '@/emitter/eventBus.js'

  import Dialog from 'primevue/dialog'
  import InputText from 'primevue/inputtext'
  import Button from 'primevue/button'

  let visible = ref(false)
  let newVal = ref('')

  const onYes = () => {
    visible.value = false
    eventBus.emit('editTvg', newVal.value)
    newVal.value = ''
  }

  onMounted(() => {
    eventBus.on('selectedEditTVGDialog', () => (visible.value = true))
  })
</script>
