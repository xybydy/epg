<template>
  <Dialog
    v-model:visible="visible"
    header="Grup Düzenle"
    modal
    closeOnEscape
    closable
    dismissableMask
  >
    <Dropdown
      v-model="selectedGroup"
      :options="Object.keys(props.groups)"
      placeholder="Grup Adı"
    ></Dropdown>
    <InputText v-model="newName" type="text" placeholder="Yeni isim" @keyup.enter="onYes" />
    <template #footer>
      <Button label="No" @click="visible = false" />
      <Button label="Yes" autofocus @click="onYes" />
    </template>
  </Dialog>
</template>

<script setup>
import { ref, defineProps } from 'vue'
import eventBus from '@/emitter/eventBus.js'

import Dialog from 'primevue/dialog'
import Dropdown from 'primevue/dropdown'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'

let selectedGroup = ref()
let newName = ref('')
let visible = ref(false)

let props = defineProps({
  groups: {
    type: Object,
  },
})

const onYes = () => {
  visible.value = false
  eventBus.emit('editGroup', { d_val: selectedGroup.value, text_val: newName.value })
}

eventBus.on('selectedEditGroupNameDialog', () => (visible.value = !visible.value))
</script>
