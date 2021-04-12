<template>
  <Dialog v-model:visible="dialogVisible" :header="header" modal closeOnEscape>
    <div class="p-field">
      <InputText
        id="tag-text"
        type="text"
        placeholder="Ayraç"
        @keyup.enter="
          editTag($event.target.value, $event.target.nextElementSibling.children[0].value)
        "
      />
      <InputNumber
        id="max-char"
        v-model="editTagSliderVal"
        inputStyle="width:2rem"
        mode="decimal"
        :max="5"
        :min="1"
        showButtons
      ></InputNumber>
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
import InputNumber from 'primevue/inputnumber'
import InputText from 'primevue/inputtext'

const props = defineProps({
  visible: {
    type: Boolean,
  },
})

let dialogVisible = ref(props.visible)

onMounted(() => {
  eventBus.on('selectedEditChanTagDialog', () => {
    dialogVisible.value = !dialogVisible.value
    editType.value = 'chan'
    header.value = 'Kanal Etiket Duzenle'
  })
  eventBus.on('selectedEditGroupTagDialog', () => {
    dialogVisible.value = !dialogVisible.value
    editType.value = 'group'
    header.value = 'Grup Etiket Duzenle'
  })
})

let editType = ref()

let header = ref()
let editTagSliderVal = ref(1)

const editTag = (ayrac, num) => {
  let obj = { val: { ayrac: ayrac, num: num }, type: '' }

  if (editType.value === 'group') {
    obj = { val: { ayrac: ayrac, num: num }, type: 'group' }
  } else if (editType.value === 'chan') {
    obj = { val: { ayrac: ayrac, num: num }, type: 'chan' }
  } else {
    console.log('editTag failed.', props.editType)
  }

  eventBus.emit('editChanTag', obj)
}
</script>
