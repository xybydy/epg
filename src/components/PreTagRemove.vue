<template>
  <Dialog v-model:visible="dialogVisible" :header="header" modal closeOnEscape>
    <div class="p-field">
      <Dropdown
        v-model="selectedTag"
        :options="choices"
        :editable="true"
        @keyup.enter="editPreTag()"
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
import { selectedItems } from '../store/selectedItems'

const props = defineProps({
  visible: {
    type: Boolean,
  },
})

let dialogVisible = ref(props.visible)
let choices = ref(preTags)
let selectedTag = ref('HEVC')

let editType = ref()
let header = ref()

const editPreTag = () => {
  let obj = { val: '', type: '' }

  if (editType.value === 'group') {
    obj = { val: selectedTag.value, type: 'group' }
  } else if (editType.value === 'chan') {
    obj = { val: selectedTag.value, type: 'chan' }
  } else {
    console.log('editTag failed.', props.editType)
  }

  eventBus.emit('editPreTag', obj)
}

onMounted(() => {
  eventBus.on('selectedEditChanPreTagDialog', () => {
    dialogVisible.value = !dialogVisible.value
    editType.value = 'chan'
    header.value = 'Kanal Etiket Duzenle'
  })
  eventBus.on('selectedEditGroupPreTagDialog', () => {
    dialogVisible.value = !dialogVisible.value
    editType.value = 'group'
    header.value = 'Grup Ön Etiket Duzenle'
  })
})
</script>
