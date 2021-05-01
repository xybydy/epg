<template>
  <Dialog
    v-model:visible="visible"
    header="Grup Kaldır"
    modal
    closeOnEscape
    closable
    dismissableMask
    style="width: 400px"
    @hide="onHide"
  >
    <div class="p-d-flex-column">
      <div v-for="(num, group) in props.groups" :key="group" class="p-field-checkbox">
        <Checkbox :id="group" v-model="selectedGroups" name="group" :value="group"></Checkbox>
        <label :for="group"
          >{{ group }} <b>({{ num }})</b></label
        >
      </div>
    </div>
    <template #footer>
      <Button label="No" @click="onHide" />
      <Button label="Yes" autofocus @click="onYes" />
    </template>
  </Dialog>
</template>

<script setup>
import { defineProps, ref, onMounted } from 'vue'
import eventBus from '@/emitter/eventBus.js'
import Dialog from 'primevue/dialog'
import Checkbox from 'primevue/checkbox'
import Button from 'primevue/button'

let selectedGroups = ref([])
let visible = ref(false)

const onHide = () => {
  if (visible.value) {
    visible.value = false
  }
  selectedGroups.value = []
}

const onYes = () => {
  visible.value = false
  eventBus.emit('clickYesRemoveGroups', selectedGroups.value)
}

let props = defineProps({
  groups: {
    type: Object,
  },
})

onMounted(() => {
  eventBus.on('clickRemoveGroups', () => {
    visible.value = !visible.value
  })
})
</script>
