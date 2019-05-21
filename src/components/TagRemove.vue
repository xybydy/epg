<template>
  <Dialog v-model:visible="visible" :header="header" modal closeOnEscape dismissableMask>
    <div class="p-field">
      <Dropdown v-model="selectedTag" :options="choices" :editable="true" @keyup.enter="editTag">
      </Dropdown>
    </div>
    <Button style="float: right" @click="editTag">OK</Button>
  </Dialog>
</template>

<script setup>
import eventBus from '@/emitter/eventBus.js'

import { ref, defineProps, onMounted } from 'vue'

import Dialog from 'primevue/dialog'
import Dropdown from 'primevue/dropdown'
import Button from 'primevue/button'
import preTags from '@/utils/preTags.js'
import { selectedItems } from '../store/selectedItems'

let visible = ref(false)
let choices = ref(preTags)
let selectedTag = ref('HEVC')

let editType = ref()
let header = ref()

const editTag = () => {
  visible.value = false
  eventBus.emit('startLoading')
  let obj = { val: '', type: '' }

  if (editType.value === 'group') {
    obj = { val: selectedTag.value, type: 'group' }
  } else if (editType.value === 'chan') {
    obj = { val: selectedTag.value, type: 'chan' }
  } else {
    console.log('editTag failed.', editType)
  }

  eventBus.emit('editTag', obj)
  eventBus.emit('stopLoading')
}

onMounted(() => {
  eventBus.on('selectedEditChanTagDialog', () => {
    visible.value = !visible.value
    editType.value = 'chan'
    header.value = 'Kanal Etiket Duzenle'
  })
  eventBus.on('selectedEditGroupTagDialog', () => {
    visible.value = !visible.value
    editType.value = 'group'
    header.value = 'Grup Etiket Duzenle'
  })
})
</script>
