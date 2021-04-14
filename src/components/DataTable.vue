<template>
  <DataTable
    v-if="!loadingDialog"
    v-model:selection="selectedItems"
    v-model:filters="filters"
    class="p-datatable-sm p-datatable-striped editable-cells-table"
    editMode="cell"
    dataKey="_id"
    paginatorTemplate="CurrentPageReport FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink RowsPerPageDropdown"
    :value="reOrderedList"
    :paginator="true"
    :rowsPerPageOptions="[25, 50, 100, 250]"
    :rows="25"
    currentPageReportTemplate="Showing {first} to {last} of {totalRecords}"
    @row-reorder="onRowReorder"
    @cell-edit-complete="onCellComplete"
  >
    <template #empty> Sonuc bulunamadi. </template>
    <template #header>
      <div class="p-d-flex p-jc-end">
        <span class="p-input-icon-left">
          <i class="pi pi-search" />
          <InputText placeholder="Ara" @keyup.enter="onSearch($event.target.value)" />
        </span>
      </div>
    </template>
    <Column :rowReorder="true" headerStyle="width: 3rem" :reorderableColumn="false" />

    <Column selectionMode="multiple" headerStyle="width: 3em"></Column>

    <Column field="chan_name" header="NAME" autoLayout="true">
      <template #editor="slotProps">
        <InputText v-model="slotProps.data[slotProps.column.props.field]" @focus="onSelectInput" />
      </template>
    </Column>
    <Column field="group_title" header="GROUP" autoLayout="true">
      <template #editor="slotProps">
        <InputText v-model="slotProps.data[slotProps.column.props.field]" @focus="onSelectInput" />
      </template>
    </Column>
    <Column field="tvg_id" header="TVG-ID" autoLayout="true">
      <template #editor="slotProps">
        <InputText v-model="slotProps.data[slotProps.column.props.field]" @focus="onSelectInput" />
      </template>
    </Column>
    <Column
      field="matcher"
      headerStyle="width: 4rem; text-align: center"
      bodyStyle="text-align: center; overflow: visible"
    >
      <template #body="slotProps">
        <MatchButton :data="slotProps"></MatchButton>
      </template>
    </Column>
  </DataTable>

  <Dialog v-model:visible="deleteItemsDialog" :style="{ width: '450px' }" header="Onayla" modal>
    <div class="p-d-flex p-ai-center confirmation-content">
      <i class="pi pi-exclamation-triangle p-mr-3" style="font-size: 2rem" />
      <span>Seçtiğin kanallar silinsin mi?</span>
    </div>
    <template #footer>
      <Button
        label="No"
        icon="pi pi-times"
        class="p-button-text"
        @click="deleteItemsDialog = false"
      />
      <Button label="Yes" icon="pi pi-check" @click="removeSelectedItems" />
    </template>
  </Dialog>

  <Dialog
    v-model:visible="BulkTvgIdDialog"
    :style="{ width: '350px' }"
    header="TVG-ID Düzenle"
    modal
  >
    <div class="p-d-flex p-jc-between p-ai-center p-input-filled p-p-2">
      <span>Yeni TVG-ID: </span>
      <InputText v-model="newTvgId" autofocus />
    </div>
    <template #footer>
      <Button
        label="Vazgeç"
        icon="pi pi-times"
        class="p-button-text"
        @click="BulkTvgIdDialog = false"
      />
      <Button label="Onayla" icon="pi pi-check" @click="editBulkItems" />
    </template>
  </Dialog>

  <Dialog
    v-model:visible="loadingDialog"
    :style="{ width: '200px' }"
    modal
    :closable="false"
    header="Please Wait"
  >
    <ProgressSpinner style="display: flex" />
  </Dialog>

  <GroupRemove v-model:visible="groupRemovalVisible" :groups="getGroups"></GroupRemove>

  <Dialog v-model:visible="editGroupNameDialog" header="Grup Düzenle" modal closeOnEscape>
    <Dropdown v-model="selectedGroup" :options="getGroups" placeholder="Grup Adı"></Dropdown>
    <InputText
      type="text"
      placeholder="Yeni isim"
      @keyup.enter="
        editGroupName($event.target.previousElementSibling.outerText, $event.target.value)
      "
    />
  </Dialog>

  <Dialog v-model:visible="editChanPreTagDialog" header="Grup Düzenle" modal closeOnEscape>
    <Dropdown v-model="selectedGroup" :options="getGroups" placeholder="Grup Adı"></Dropdown>
    <InputText
      type="text"
      placeholder="Yeni isim"
      @keyup.enter="
        editGroupName($event.target.previousElementSibling.outerText, $event.target.value)
      "
    />
  </Dialog>

  <PreTagRemove :visible="editChanPreTagDialog"></PreTagRemove>
  <TagRemove :visible="editChanTagDialog"></TagRemove>
  <Toast position="bottom-right" />
</template>

<script setup>
import { ref, watch, defineProps, onMounted, computed } from 'vue'
import eventBus from '@/emitter/eventBus.js'

import ProgressSpinner from 'primevue/progressspinner'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
import InputText from 'primevue/inputtext'
import Dialog from 'primevue/dialog'
import Button from 'primevue/button'
import Dropdown from 'primevue/dropdown'
import InputNumber from 'primevue/inputnumber'

import { FilterMatchMode } from 'primevue/api'
import { useToast } from 'primevue/usetoast'
import Toast from 'primevue/toast'

import { ExportM3u } from '@/utils/m3u.js'
import { deDupe } from '@/utils'

import EpgUploader from '@/components/EpgUploader.vue'
import MenuBar from '@/components/MenuBar.vue'
import GroupRemove from '@/components/GroupRemove.vue'
import TagRemove from '@/components/TagRemove.vue'
import PreTagRemove from '@/components/PreTagRemove.vue'
import MatchButton from '@/components/MatchButton.vue'

import { selectedItems } from '@/store/selectedItems.js'

import { root_path } from '@/router/url'

const props = defineProps({
  m3u: {
    type: Object,
    required: true,
  },
  new: {
    type: Boolean,
    required: true,
  },
})

const toast = useToast()

let filters = ref({ global: { value: null, matchMode: FilterMatchMode.CONTAINS } })

let cleanup = () => {
  newTvgId.value = ''
  selectedItems.value = []
  loadingDialog.value = false
  filters.value.global.value = ''
}

let pushItems = (toArray, fromArray) => {
  for (let item of fromArray.value) {
    let pushItem = { [item._id]: item }
    Object.assign(toArray, pushItem)
  }
}

let deleteItemsDialog = ref(false)
let BulkTvgIdDialog = ref(false)
let loadingDialog = props.m3u.length === 0 ? ref(true) : ref(false)
let reOrderedList = ref(props.m3u)
let itemsToSavedList = []
let downloadButtonLock = ref(true)
let newTvgId = ref('')
let changedItems = {}
let selectedGroup = ref()
let editGroupNameDialog = ref(false)
let editChanTagDialog = ref(false)
let editChanPreTagDialog = ref(false)
let groupRemovalVisible = ref(false)
let matchingProgress = ref(false)

let onCellComplete = (e) => Object.assign(changedItems, { [e.data._id]: e.data })

let onRowReorder = (e) => (reOrderedList.value = e.value)
let onSelectInput = (e) => e.target.select()

let editNames = async () => {
  loadingDialog.value = true
  for (let item of selectedItems.value) {
    item.tvg_id = newTvgId.value
  }
  pushItems(changedItems, selectedItems)
  cleanup()
}

const postSave = () => {
  loadingDialog.value = true

  itemsToSavedList = deDupe(reOrderedList.value).map((item) => {
    return {
      group_title: item.group_title.trim(),
      chan_name: item.chan_name.trim(),
      tvg_id: item.tvg_id.trim(),
    }
  })
  downloadButtonLock.value = false
  fetch(root_path + '/api/save', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(itemsToSavedList),
  })
    .then((resp) => resp.json())
    .then((data) => {
      if (data.status_code == 200) {
        toast.add({
          severity: 'success',
          summary: 'Success',
          detail: data.message,
          life: 3000,
        })
      } else {
        toast.add({
          severity: 'error',
          summary: 'Error',
          detail: `${data.status_code} - ${data.message}`,
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
    .finally(() => {
      cleanup()
    })
}

const postUpdate = () => {
  loadingDialog.value = true

  itemsToSavedList = Object.entries(changedItems).map(([k, v]) => {
    return { _id: v._id, group_title: v.group_title, chan_name: v.chan_name, tvg_id: v.tvg_id }
  })

  if (itemsToSavedList.length === 0) {
    loadingDialog.value = false
    toast.add({
      severity: 'warn',
      summary: 'Warning',
      detail: `No items has changed!`,
      life: 3000,
    })
    return
  }

  downloadButtonLock.value = false
  fetch(root_path + '/api/update', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(itemsToSavedList),
  })
    .then((resp) => resp.json())
    .then((data) => {
      if (data.status_code == 200) {
        toast.add({
          severity: 'success',
          summary: 'Success',
          detail: data.message,
          life: 3000,
        })
      } else {
        toast.add({
          severity: 'error',
          summary: 'Error',
          detail: `${data.status_code} - ${data.message}`,
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
    .finally(() => {
      cleanup()
    })
}

const sleepNow = (delay) => new Promise((resolve) => setTimeout(resolve, delay))

const removeSelectedItems = async () => {
  deleteItemsDialog.value = false
  loadingDialog.value = true
  await sleepNow(30)
  reOrderedList.value = reOrderedList.value.filter((val) => !selectedItems.value.includes(val))
  cleanup()
}

const editBulkItems = async () => {
  BulkTvgIdDialog.value = false
  loadingDialog.value = true
  await sleepNow(30)
  editNames()
  cleanup()
}

const downloadM3u = () => {
  loadingDialog.value = true

  if (reOrderedList.value.length > 0) {
    ExportM3u(reOrderedList.value)
  } else {
    toast.add({
      severity: 'warn',
      summary: 'Uyarı',
      detail: 'İndirilecek bir değişiklik olmadı.',
      life: 3000,
    })
  }
  loadingDialog.value = false
}

const onTvgMatcher = () => {
  loadingDialog.value = true
  itemsToSavedList = selectedItems.value.map((item) => {
    return {
      _id: item._id,
      chan_name: item.chan_name,
    }
  })

  fetch(root_path + '/api/match', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(itemsToSavedList),
  })
    .then((resp) => resp.json())
    .then((data) => {
      if (data.status_code == 200) {
        console.log(data.data)
        for (let outerItem of data.data) {
          for (let innerItem of reOrderedList.value) {
            if (innerItem._id === outerItem._id) {
              innerItem.tvg_id = outerItem.tvg_id
              break
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
          summary: 'Error',
          detail: `${data.status_code} - ${data.message}`,
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
    .finally(() => {
      cleanup()
    })
}

const selectedTvgRemove = () => {
  loadingDialog.value = true

  for (let item of selectedItems.value) {
    item.tvg_id = ''
  }
  pushItems(changedItems, selectedItems)
  cleanup()
}

const onSearch = (e) => {
  filters.value.global.value = e
}

const getGroups = computed(() => {
  let arr = [...new Set([...reOrderedList.value].map((x) => x.group_title))].sort()
  return arr
})

const editTag = (e) => {
  editChanPreTagDialog.value = false
  loadingDialog.value = true
  let separator = e.val.ayrac.trim()

  if (e.type === 'chan') {
    if (selectedItems.value > 0) {
      for (let item of selectedItems.value) {
        let sep_index = item.chan_name.indexOf(separator)
        if (sep_index === -1 || sep_index > e.val.num) {
          console.log('No work:', item.chan_name)
        } else {
          item.chan_name = item.chan_name
            .slice(sep_index + separator.length, item.chan_name.length)
            .trim()
        }
      }
    } else {
      for (let item of reOrderedList.value) {
        let sep_index = item.chan_name.indexOf(separator)
        if (sep_index === -1 || sep_index > e.val.num) {
          console.log('No work:', item.chan_name)
        } else {
          item.chan_name = item.chan_name
            .slice(sep_index + separator.length, item.chan_name.length)
            .trim()
        }
      }
    }
  } else if (e.type === 'group') {
    for (let item of reOrderedList.value) {
      let sep_index = item.group_title.indexOf(separator)
      if (sep_index === -1 || sep_index > e.val.num) {
        console.log('No work:', item.group_title)
      } else {
        item.chan_name = item.group_title
          .slice(sep_index + separator.length, item.group_title.length)
          .trim()
      }
    }
  }

  cleanup()
}

const editPreTag = (e) => {
  loadingDialog.value = true
  let val = e.val.trim()
  if (selectedItems.value.length > 0) {
    for (let i = 0; i < selectedItems.value.length; i++) {
      if (e.type === 'chan') {
        let stripped = selectedItems.value[i].chan_name.replace(' ' + val, '').trim()
        selectedItems.value[i].chan_name = stripped
      } else if (e.type === 'group') {
        let stripped = selectedItems.value[i].group_title.replace(' ' + val, '').trim()
        selectedItems.value[i].group_title = stripped
      }
    }
  } else {
    if (e.type === 'chan') {
      for (let i = 0; i < reOrderedList.value.length; i++) {
        let stripped = reOrderedList.value[i].chan_name.replace(' ' + val, '').trim()
        reOrderedList.value[i].chan_name = stripped
      }
    } else if (e.type === 'group') {
      for (let i = 0; i < reOrderedList.value.length; i++) {
        let stripped = reOrderedList.value[i].group_title.replace(' ' + val, '').trim()
        reOrderedList.value[i].group_title = stripped
      }
    }
  }
  cleanup()
}

const editGroupName = (d_val, text_val) => {
  loadingDialog.value = true
  editGroupNameDialog.value = false

  reOrderedList.value
    .filter((val) => val.group_title === d_val)
    .map((val) => (val.group_title = text_val))

  cleanup()
}

onMounted(() => {
  console.log('mounted')
  eventBus.on('selectedRemoveDialog', () => (deleteItemsDialog.value = true))
  eventBus.on('selectedEditDialog', () => (BulkTvgIdDialog.value = true))
  eventBus.on('onTvgMatcher', onTvgMatcher)
  eventBus.on('clickUpdate', postUpdate)
  eventBus.on('clickSave', postSave)
  eventBus.on('ondownloadM3u', downloadM3u)
  eventBus.on('selectedTvgRemove', selectedTvgRemove)
  eventBus.on('clickNoRemoveGroups', () => (groupRemovalVisible.value = !groupRemovalVisible.value))
  eventBus.on('clickYesRemoveGroups', async (e) => {
    groupRemovalVisible.value = !groupRemovalVisible.value
    loadingDialog.value = true
    await sleepNow(30)
    e.map((group_title) => {
      reOrderedList.value = reOrderedList.value.filter((val) => val.group_title !== group_title)
    })
    loadingDialog.value = false
    selectedItems.value = []
  })
  eventBus.on(
    'selectedEditGroupNameDialog',
    () => (editGroupNameDialog.value = !editGroupNameDialog.value)
  )
  eventBus.on('editTag', (e) => editTag(e))
  eventBus.on('editPreTag', (e) => editPreTag(e))
})

watch(
  () => props.m3u,
  (f) => {
    loadingDialog.value = false
    reOrderedList.value = f
  }
)
</script>
