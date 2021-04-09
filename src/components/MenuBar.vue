<template>
  <Toolbar class="p-mb-4">
    <template #left>
      <Button
        :label="removeLabel"
        icon="pi pi-trash"
        class="p-button-danger p-mr-2"
        :disabled="!selectedItems || !selectedItems.length"
        @click="emit.selectedRemoveDialog"
      />
      <Button
        label="Toplu TVG-ID Düzenle"
        icon="pi pi-file"
        class="p-button-danger p-mr-2"
        :disabled="!selectedItems || !selectedItems.length"
        @click="emit.selectedEditDialog"
      />
      <Button
        label="TVG-ID Getir"
        icon="pi pi-cloud-download"
        class="p-button-danger"
        :disabled="!selectedItems || !selectedItems.length"
        @click="emit.onTvgMatcher"
      />
    </template>
    <template #right>
      <Button
        v-if="props.new"
        label="Kaydet"
        icon="pi pi-save"
        class="p-button-help p-mr-2"
        @click="emit.onSave"
      />
      <Button
        v-else
        label="Guncelle"
        icon="pi pi-save"
        class="p-button-help p-mr-2"
        @click="emit.clickUpdate"
      />
      <Button
        :disabled="props.downloadButtonLock"
        label="İndir"
        icon="pi pi-download"
        class="p-button-help"
        @click="emit.ondownloadM3u"
      />
    </template>
  </Toolbar>
</template>

<script setup>
import { computed, defineProps } from 'vue'
import EpgUploader from '@/components/EpgUploader.vue'
import Button from 'primevue/button'
import Toolbar from 'primevue/toolbar'

import eventBus from '@/emitter/eventBus.js'

import { selectedItems } from '@/store/selectedItems.js'

const props = defineProps({
  new: {
    type: Boolean,
  },
  downloadButtonLock: {
    type: Boolean,
  },
})

const emit = {
  selectedRemoveDialog: () => eventBus.emit('selectedRemoveDialog'),
  selectedEditDialog: () => eventBus.emit('selectedEditDialog'),
  onTvgMatcher: () => eventBus.emit('onTvgMatcher'),
  clickUpdate: () => eventBus.emit('clickUpdate'),
  clickSave: () => eventBus.emit('clickSave'),
  ondownloadM3u: () => eventBus.emit('ondownloadM3u'),
}

let removeLabel = computed(() => {
  return selectedItems.value.length === 0 ? 'Sil' : `Sil (${selectedItems.value.length})`
})
</script>

<style scoped></style>
