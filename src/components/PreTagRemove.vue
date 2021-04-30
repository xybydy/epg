<template>
  <Dialog v-model:visible="dialogVisible" :header="header" modal closeOnEscape>
    <div class="p-field">
      <InputText
        id="tag-text"
        type="text"
        placeholder="Ayraç"
        @keyup.enter="
          editPreTag($event.target.value, $event.target.nextElementSibling.children[0].value)
        "
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

let dialogVisible = ref(props.visible)

onMounted(() => {
  eventBus.on('selectedEditChanPreTagDialog', () => {
    dialogVisible.value = !dialogVisible.value
    editType.value = 'chan'
    header.value = 'Kanal Ön Etiket Duzenle'
  })
  eventBus.on('selectedEditGroupPreTagDialog', () => {
    dialogVisible.value = !dialogVisible.value
    editType.value = 'group'
    header.value = 'Grup Ön Etiket Duzenle'
  })
})

let editType = ref()

let header = ref()
let editPreTagSliderVal = ref(1)

const editPreTag = (ayrac, num) => {
  let obj = { val: { ayrac: ayrac, num: num }, type: '' }

  if (editType.value === 'group') {
    obj = { val: { ayrac: ayrac, num: num }, type: 'group' }
  } else if (editType.value === 'chan') {
    obj = { val: { ayrac: ayrac, num: num }, type: 'chan' }
  } else {
    console.log('editPreTag failed.', props.editType)
  }

  eventBus.emit('editPreTag', obj)
}
</script>
