<template>
  <Dialog v-model:visible="dialogVisible" :header="header" modal closeOnEscape>
    <div class="p-field">
      <Dropdown :options="choices" :editable="true"> </Dropdown>
      <p class="p-text-light" style="font-size: 0.8rem">
        5 karakterden uzun etiketleri desteklemiyor.
      </p>
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
