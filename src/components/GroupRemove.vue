<template>
  <Dialog header="Grup Kaldır" modal closeOnEscape closable style="width: 400px" @hide="onHide">
    <div class="p-d-flex-column">
      <div v-for="group of props.groups" :key="group" class="p-field-checkbox">
        <Checkbox :id="group" v-model="selectedGroups" name="group" :value="group"></Checkbox>
        <label :for="group">{{ group }}</label>
      </div>
    </div>
    <template #footer>
      <Button label="No" @click="eventBus.emit('clickNoRemoveGroups')" />
      <Button
        label="Yes"
        autofocus
        @click="eventBus.emit('clickYesRemoveGroups', selectedGroups)"
      />
    </template>
  </Dialog>
</template>

<script setup>
import { defineProps, ref } from 'vue'
import eventBus from '@/emitter/eventBus.js'
import Dialog from 'primevue/dialog'
import Checkbox from 'primevue/checkbox'
import Button from 'primevue/button'

let selectedGroups = ref([])

const onHide = () => {
  selectedGroups.value = []
}

let props = defineProps({
  groups: {
    type: Array,
  },
})
</script>
