<template>
  <Menu style="flex-shrink: 0" :model="menuItems"></Menu>
</template>

<script setup>
import { computed, defineProps, ref } from 'vue'
import EpgUploader from '@/components/index/EpgUploader.vue'
import Menu from 'primevue/menu'
import Button from 'primevue/button'
import Toolbar from 'primevue/toolbar'

import eventBus from '@/emitter/eventBus.js'

import { selectedItems } from '@/store/selectedItems.js'

const emit = {
  selectedRemoveDialog: () => eventBus.emit('selectedRemoveDialog'),
  selectedEditDialog: () => eventBus.emit('selectedEditDialog'),
  selectedTvgRemove: () => eventBus.emit('selectedTvgRemove'),
  onTvgMatcher: () => eventBus.emit('onTvgMatcher'),
  clickUpdate: () => eventBus.emit('clickUpdate'),
  clickSave: () => eventBus.emit('clickSave'),
  ondownloadM3u: () => eventBus.emit('ondownloadM3u'),
  clickRemoveGroups: () => eventBus.emit('clickNoRemoveGroups'),
  selectedEditGroupNameDialog: () => eventBus.emit('selectedEditGroupNameDialog'),
  selectedEditChanPreTagDialog: () => eventBus.emit('selectedEditChanPreTagDialog'),
  selectedEditChanTagDialog: () => eventBus.emit('selectedEditChanTagDialog'),
  selectedEditGroupPreTagDialog: () => eventBus.emit('selectedEditGroupPreTagDialog'),
  selectedEditGroupTagDialog: () => eventBus.emit('selectedEditGroupTagDialog'),
}

const props = defineProps({
  new: {
    type: Boolean,
  },
  downloadButtonLock: {
    type: Boolean,
  },
})

let removeLabel = computed(() => {
  return selectedItems.value.length === 0 ? 'Sil' : `Sil (${selectedItems.value.length})`
})

let notSelectedItem = computed(() => {
  return selectedItems.value.length === 0
})

const menuItems = ref([
  {
    label: 'Kanal',
    items: [
      {
        label: removeLabel,
        command: () => emit.selectedRemoveDialog(),
        disabled: notSelectedItem,
      },
      {
        label: 'Etiket Kaldır',
        command: () => emit.selectedEditChanTagDialog(),
      },
      {
        label: 'Ön Etiket Kaldır',
        command: () => emit.selectedEditChanPreTagDialog(),
      },
    ],
  },
  {
    separator: true,
  },
  {
    label: 'TVG-ID',
    items: [
      {
        label: 'Düzenle',
        disabled: notSelectedItem,
        command: () => emit.selectedEditDialog(),
      },
      {
        label: 'Eşleştir',
        disabled: notSelectedItem,
        command: () => {},
      },
      {
        label: 'Kaldır',
        disabled: notSelectedItem,
        command: () => emit.selectedTvgRemove(),
      },
    ],
  },
  {
    separator: true,
  },
  {
    label: 'Gruplar',
    items: [
      {
        label: 'Düzenle',
        command: () => emit.selectedEditGroupNameDialog(),
      },
      {
        label: 'Kaldır',
        command: () => emit.clickRemoveGroups(),
      },
      {
        label: 'Etiket Kaldır',
        command: () => emit.selectedEditGroupTagDialog(),
      },
      {
        label: 'Ön Etiket Kaldır',
        command: () => emit.selectedEditGroupPreTagDialog(),
      },
    ],
  },
  {
    separator: true,
  },
  {
    label: 'Kaydet',
    command: () => emit.clickSave(),
  },
  {
    label: 'Güncelle',
    command: () => emit.clickUpdate(),
  },
])
</script>

<style scoped></style>
