<template>
  <Dialog v-model:visible="visible" :header="header" modal closeOnEscape dismissableMask closable>
    <div class="p-field">
      <InputText
        id="tag-text"
        v-model="editInput"
        type="text"
        placeholder="Ayraç"
        @keyup.enter="onYes"
      />
      <InputNumber
        id="max-char"
        v-model="editPreTagSliderVal"
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
    <Button style="float: right" @click="onYes">OK</Button>
  </Dialog>
</template>

<script setup>
import eventBus from '@/emitter/eventBus.js'

import { ref, defineProps, onMounted } from 'vue'

import Dialog from 'primevue/dialog'
import InputNumber from 'primevue/inputnumber'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'

const props = defineProps({
  visible: {
    type: Boolean,
  },
})

let visible = ref(props.visible)

onMounted(() => {
  eventBus.on('selectedEditChanPreTagDialog', () => {
    visible.value = !visible.value
    editType.value = 'chan'
    header.value = 'Kanal Ön Etiket Duzenle'
  })
  eventBus.on('selectedEditGroupPreTagDialog', () => {
    visible.value = !visible.value
    editType.value = 'group'
    header.value = 'Grup Ön Etiket Duzenle'
  })
})

let editType = ref()
let header = ref()

let editInput = ref('')
let editPreTagSliderVal = ref(1)

const onYes = () => {
  visible.value = false
  eventBus.emit('startLoading')

  let obj = { val: { ayrac: editInput.value, num: editPreTagSliderVal.value }, type: '' }

  if (editType.value === 'group') {
    obj = { val: { ayrac: editInput.value, num: editPreTagSliderVal.value }, type: 'group' }
  } else if (editType.value === 'chan') {
    obj = { val: { ayrac: editInput.value, num: editPreTagSliderVal.value }, type: 'chan' }
  } else {
    console.log('editPreTag failed.', props.editType)
  }

  eventBus.emit('editPreTag', obj)
  eventBus.emit('stopLoading')
}
</script>
