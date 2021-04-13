<template>
  <Dialog v-model:visible="dialogVisible" :header="header" modal closeOnEscape>
    <div class="p-field">
      <Dropdown
        v-model="selectedTag"
        :options="choices"
        :editable="true"
        @keyup.enter="eventBus.emit('editPreTag', $event.target.value)"
      >
      </Dropdown>
    </div>
  </Dialog>
</template>

<script setup>
import eventBus from '@/emitter/eventBus.js'

import { ref, defineProps, onMounted } from 'vue'

import Dialog from 'primevue/dialog'
import Dropdown from 'primevue/dropdown'
import preTags from '@/utils/preTags.js'

const props = defineProps({
  visible: {
    type: Boolean,
  },
})

let dialogVisible = ref(props.visible)
let choices = ref(preTags)
let selectedTag = ref('HEVC')

onMounted(() => {
  eventBus.on('selectedEditChanPreTagDialog', () => {
    dialogVisible.value = !dialogVisible.value
    editType.value = 'chan'
    header.value = 'Kanal Etiket Duzenle'
  })
  eventBus.on('selectedEditGroupPreTagDialog', () => {
    dialogVisible.value = !dialogVisible.value
    editType.value = 'group'
    header.value = 'Grup Etiket Duzenle'
  })
})

let editType = ref()

let header = ref()
</script>
