<template>
  <DataTable
    v-if="!loadingDialog"
    v-model:selection="selectedItems"
    class="p-datatable-sm p-datatable-striped editable-cells-table"
    editMode="cell"
    :filters="filters"
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
    <Column :rowReorder="true" headerStyle="width: 3rem" :reorderableColumn="false" />

    <Column selectionMode="multiple" headerStyle="width: 3em"></Column>

    <Column field="chan_name" header="NAME" autoLayout="true">
      <template #editor="slotProps">
        <InputText v-model="slotProps.data[slotProps.column.props.field]" @focus="onSelectInput" />
      </template>
      <template #filter>
        <InputText v-model="filters.chan_name" type="text" placeholder="İsim ara" />
      </template>
    </Column>
    <Column field="group_title" header="GROUP" autoLayout="true">
      <template #editor="slotProps">
        <InputText v-model="slotProps.data[slotProps.column.props.field]" @focus="onSelectInput" />
      </template>
      <template #filter>
        <InputText v-model="filters.group_title" type="text" placeholder="Group ara" />
      </template>
    </Column>
    <Column field="tvg_id" header="TVG-ID" autoLayout="true">
      <template #editor="slotProps">
        <InputText v-model="slotProps.data[slotProps.column.props.field]" @focus="onSelectInput" />
      </template>
      <template #filter>
        <InputText v-model="filters.tvg_id" type="text" placeholder="TVG-ID ara" />
      </template>
    </Column>
  </DataTable>

  <Dialog
    v-model:visible="deleteItemsDialog"
    :style="{ width: '450px' }"
    header="Onayla"
    :modal="true"
  >
    <div class="confirmation-content">
      <i class="pi pi-exclamation-triangle p-mr-3" style="font-size: 2rem" />
      <span v-if="selectedItems">Seçtiğin kanallar silinsin mi?</span>
    </div>
    <template #footer>
      <Button
        label="No"
        icon="pi pi-times"
        class="p-button-text"
        @click="deleteItemsDialog = false"
      />
      <Button label="Yes" icon="pi pi-check" class="p-button-text" @click="removeSelectedItems" />
    </template>
  </Dialog>

  <Dialog
    v-model:visible="BulkTvgIdDialog"
    :style="{ width: '350px' }"
    header="TVG-ID Düzenle"
    :modal="true"
  >
    <div class="p-input-filled p-p-2">
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
      <Button label="Onayla" icon="pi pi-check" class="p-button-text" @click="editBulkItems" />
    </template>
  </Dialog>

  <Dialog
    v-model:visible="loadingDialog"
    :style="{ width: '200px' }"
    :modal="true"
    :closable="false"
    header="Please Wait"
  >
    <ProgressSpinner style="display: flex" />
  </Dialog>

  <Toast position="bottom-right" />
</template>

<script setup>
import { ref, watch, defineProps, onMounted } from 'vue'
import eventBus from '@/emitter/eventBus.js'

import ProgressSpinner from 'primevue/progressspinner'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
import InputText from 'primevue/inputtext'
import Dialog from 'primevue/dialog'
import Button from 'primevue/button'

import { useToast } from 'primevue/usetoast'
import Toast from 'primevue/toast'

import { ExportM3u } from '@/utils/m3u.js'
import { deDupe } from '@/utils'

import EpgUploader from '@/components/EpgUploader.vue'
import MenuBar from '@/components/MenuBar.vue'

import { selectedItems } from '@/store/selectedItems.js'

import { root_path } from '@/router/url'

const pushItems = (toArray, fromArray) => {
  for (let item of fromArray.value) {
    let pushItem = { [item._id]: item }
    Object.assign(toArray, pushItem)
  }
}

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

let filters = ref({})

let deleteItemsDialog = ref(false)

let BulkTvgIdDialog = ref(false)

let loadingDialog = props.m3u.length === 0 ? ref(true) : ref(false)

let reOrderedList = ref(props.m3u)
let itemsToSavedList = []

let downloadButtonLock = ref(true)
let newTvgId = ref('')

let changedItems = {}

let onCellComplete = (e) => Object.assign(changedItems, { [e.data._id]: e.data })

let onRowReorder = (e) => (reOrderedList.value = e.value)
let onSelectInput = (e) => e.target.select()

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
      selectedItems.value = []
      loadingDialog.value = false
    })
}

const postUpdate = () => {
  loadingDialog.value = true

  itemsToSavedList = Object.entries(changedItems).map(([k, v]) => {
    return { _id: v._id, group_title: v.group_title, chan_name: v.chan_name, tvg_id: v.tvg_id }
  })

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
      selectedItems.value = []
      changedItems = {}
      loadingDialog.value = false
    })
}

const removeSelectedItems = () => {
  reOrderedList.value = reOrderedList.value.filter((val) => !selectedItems.value.includes(val))
  deleteItemsDialog.value = false
  selectedItems.value = []
}

const editBulkItems = () => {
  for (let item of selectedItems.value) {
    for (let i = 0; i < reOrderedList.value.length; i++) {
      if (reOrderedList.value[i]._id === item._id) {
        reOrderedList.value[i].tvg_id = newTvgId.value
        break
      }
    }
  }
  pushItems(changedItems, selectedItems)
  selectedItems.value = []
  newTvgId.value = ''
  BulkTvgIdDialog.value = false
}

const downloadM3u = () => {
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
      loadingDialog.value = false
      selectedItems.value = []
    })
}

onMounted(() => {
  eventBus.on('selectedRemoveDialog', () => (deleteItemsDialog.value = true))
  eventBus.on('selectedEditDialog', () => (BulkTvgIdDialog.value = true))
  eventBus.on('onTvgMatcher', onTvgMatcher)
  eventBus.on('onUpdate', postUpdate)
  eventBus.on('onSave', postSave)
  eventBus.on('ondownloadM3u', downloadM3u)
})

watch(
  () => props.m3u,
  (f) => {
    loadingDialog.value = false
    reOrderedList.value = f
  }
)
</script>
