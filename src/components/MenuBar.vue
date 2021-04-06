<template>
  <Toolbar class="p-mb-4">
    <template #left>
      <EpgUploader class="p-mr-2"></EpgUploader>
      <Button
        :label="removeLabel"
        icon="pi pi-trash"
        class="p-button-danger p-mr-2"
        :disabled="!selectedItemsListLength"
        @click="$emit('selectedRemoveDialog')"
      />
      <Button
        label="Toplu TVG-ID Düzenle"
        icon="pi pi-file"
        class="p-button-danger p-mr-2"
        :disabled="!selectedItemsListLength"
        @click="$emit('selectedEditDialog')"
      />
      <Button
        label="TVG-ID Getir"
        icon="pi pi-cloud-download"
        class="p-button-danger"
        @click="$emit('onTvgMatcher')"
      />
    </template>

    <template #right>
      <Button
        v-if="props.new"
        label="Kaydet"
        icon="pi pi-save"
        class="p-button-help p-mr-2"
        @click="$emit('onSave')"
      />
      <Button
        v-else
        label="Guncelle"
        icon="pi pi-save"
        class="p-button-help p-mr-2"
        @click="$emit('onUpdate')"
      />
      <Button
        :disabled="downloadButtonLock"
        label="İndir"
        icon="pi pi-download"
        class="p-button-help"
        @click="downloadM3u"
      />
    </template>
  </Toolbar>
</template>

<script setup>
import { ref, computed, inject, defineEmit, defineProps } from 'vue'
import EpgUploader from '@/components/EpgUploader.vue'
import Buttom from 'primevue/button'
import Toolbar from 'primevue/toolbar'

const props = defineProps({
  new: {
    type: Boolean,
    required: true,
  },
})

const selectedItemsListLength = inject('LengthSelectedItems')

const emit = defineEmit([
  'selectedRemoveDialog',
  'selectedEditDialog',
  'onTvgMatcher',
  'onUpdate',
  'onSave',
])

let removeLabel = computed(() => {
  return selectedItemsListLength === 0 ? 'Sil' : `Sil (${selectedItemsListLength})`
})
</script>

<style scoped></style>
