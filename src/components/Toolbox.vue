<template>
  <Menu style="flex-shrink: 0" :model="menuItems"></Menu>
</template>

<script setup>
  import {computed, defineProps, onMounted, ref} from 'vue'
  import Menu from 'primevue/menu'

  import eventBus from '@/emitter/eventBus.js'
  import {root_path} from '@/router/url'
  import {useToast} from 'primevue/usetoast'

  import {selectedItems} from '@/store/selectedItems.js'

  const emit = {
    selectedRemoveDialog: () => eventBus.emit('selectedRemoveDialog'),
    selectedEditTVGDialog: () => eventBus.emit('selectedEditTVGDialog'),
    selectedTvgRemove: () => eventBus.emit('selectedTvgRemove'),
    onTvgMatcher: () => eventBus.emit('onTvgMatcher'),
    clickUpdate: () => eventBus.emit('clickUpdate'),
    clickSave: () => eventBus.emit('clickSave'),
    ondownloadM3u: () => eventBus.emit('ondownloadM3u'),
    clickRemoveGroups: () => eventBus.emit('clickRemoveGroups'),
    selectedEditGroupNameDialog: () => eventBus.emit('selectedEditGroupNameDialog'),
    selectedEditChanPreTagDialog: () => eventBus.emit('selectedEditChanPreTagDialog'),
    selectedEditChanTagDialog: () => eventBus.emit('selectedEditChanTagDialog'),
    selectedEditGroupPreTagDialog: () => eventBus.emit('selectedEditGroupPreTagDialog'),
    selectedEditGroupTagDialog: () => eventBus.emit('selectedEditGroupTagDialog'),
  }

  let toast = useToast()

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

  const matchSelected = (country) => {
    let chan_list = []

    for (let item of selectedItems.value) {
      chan_list.push({name: item.chan_name, country: country})
    }

    fetch(root_path + '/api/match', {
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      body: JSON.stringify(chan_list),
    })
      .then((resp) => resp.json())
      .then((data) => {
        if (data.status_code === 200) {
          for (let ret of data.data) {
            for (let item of selectedItems.value) {
              if (ret.name === item.chan_name) {
                item.tvg_id = ret.tvg_id
                continue
              }
            }
          }
          toast.add({
            severity: 'success',
            summary: 'Success',
            detail: data.message,
            life: 3000,
          })
        } else {
          toast.add({
            severity: 'error',
            summary: data.status_code,
            detail: data.message,
            life: 3000,
          })
        }
      })
      .catch((error) => {
        toast.add({
          severity: 'error',
          summary: 'Error',
          detail: error,
          life: 3000,
        })
      })
  }

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
          command: () => emit.selectedEditTVGDialog(),
        },
        {
          label: 'Eşleştir',
          disabled: notSelectedItem,
          command: () => {
            eventBus.emit('openMatchDialog')
          },
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

  onMounted(() => {
    eventBus.on('selectedCountry', (e) => {
      matchSelected(e)
    })
  })
</script>
