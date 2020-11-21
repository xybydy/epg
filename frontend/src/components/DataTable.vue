<template>
  <Toolbar class="p-mb-4">
    <template #left>
      <FileUpload
        mode="basic"
        label="Ekle"
        :max-file-size="26214400"
        :file-limit="1"
        choose-label="Ekle"
        class="p-mr-2"
      />
      <Button
        :label="removeLabel"
        icon="pi pi-trash"
        class="p-button-danger p-mr-2"
        :disabled="!selectedItemsList || !selectedItemsList.length"
        @click="showRemoveSelected"
      />
      <Button
        label="Toplu TVG-ID Düzenle"
        icon="pi pi-file"
        class="p-button-danger p-mr-2"
        :disabled="!selectedItemsList || !selectedItemsList.length"
        @click="showEditSelected"
      />
      <Button
        label="TVG-ID Getir"
        icon="pi pi-cloud-download"
        class="p-button-danger"
        @click="TvgMatcherDialog"
      />
    </template>

    <template #right>
      <Button label="Kaydet" icon="pi pi-save" class="p-button-help p-mr-2" @click="onSave" />
      <Button
        :disabled="downloadButtonLock"
        label="İndir"
        icon="pi pi-download"
        class="p-button-help"
        @click="downloadM3u"
      />
    </template>
  </Toolbar>

  <DataTable
    v-model:selection="selectedItemsList"
    class="p-datatable-sm p-datatable-striped editable-cells-table"
    edit-mode="cell"
    :filters="filters"
    data-key="id"
    paginator-template="CurrentPageReport FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink RowsPerPageDropdown"
    :value="reOrderedList"
    :paginator="true"
    :rows-per-page-options="[25, 50, 100, 250]"
    :rows="25"
    current-page-report-template="Showing {first} to {last} of {totalRecords}"
    @row-reorder="onRowReorder"
  >
    <Column :row-reorder="true" header-style="width: 3rem" :reorderable-column="false" />

    <Column selection-mode="multiple" header-style="width: 3em"></Column>

    <Column field="name" header="NAME" auto-layout="true">
      <template #editor="slotProps">
        <InputText v-model="slotProps.data[slotProps.column.props.field]" @focus="onSelectInput" />
      </template>
      <template #filter>
        <InputText v-model="filters.name" type="text" placeholder="İsim ara" />
      </template>
    </Column>
    <Column field="group_title" header="GROUP" auto-layout="true">
      <template #editor="slotProps">
        <InputText v-model="slotProps.data[slotProps.column.props.field]" @focus="onSelectInput" />
      </template>
      <template #filter>
        <InputText v-model="filters.group_title" type="text" placeholder="Group ara" />
      </template>
    </Column>
    <Column field="tvg_id" header="TVG-ID" auto-layout="true">
      <template #editor="slotProps">
        <InputText
          v-model="slotProps.data[slotProps.column.props.field]"
          @focus="onSelectInput"
        /> </template
      ><template #filter>
        <InputText v-model="filters.tvg_id" type="text" placeholder="TVG-ID ara" /> </template
    ></Column>
    <!-- <Column field="tvg_name" header="TVG-NAME" autoLayout="true">
      <template #editor="slotProps">
        <InputText
          v-model="slotProps.data[slotProps.column.props.field]"
          @focus="onSelectInput"
        /> </template
      ><template #filter>
        <InputText type="text" v-model="filters.tvg_name" placeholder="TVG-NAME ara" /> </template
    ></Column> -->
    <Column field="tvg_logo" header="TVG-LOGO" auto-layout="true">
      <template #editor="slotProps">
        <InputText
          v-model="slotProps.data[slotProps.column.props.field]"
          @focus="onSelectInput"
        /> </template
    ></Column>
    <Column field="url" header="URL" auto-layout="true">
      <template #editor="slotProps">
        <InputText
          v-model="slotProps.data[slotProps.column.props.field]"
          @focus="onSelectInput"
        /> </template
    ></Column>
  </DataTable>

  <Dialog
    v-model:visible="renameItemsDialog"
    :style="{ width: '450px' }"
    header="Onayla"
    :modal="true"
  >
    <div class="confirmation-content">
      <i class="pi pi-exclamation-triangle p-mr-3" style="font-size: 2rem" />
      <span v-if="selectedItemsList">Seçtiğin kanallar silinsin mi?</span>
    </div>
    <template #footer>
      <Button
        label="No"
        icon="pi pi-times"
        class="p-button-text"
        @click="renameItemsDialog = false"
      />
      <Button
        label="Yes"
        icon="pi pi-check"
        class="p-button-text"
        @click="removeSelectedItemsList"
      />
    </template>
  </Dialog>

  <Dialog
    v-model:visible="BulkTvgIdDialog"
    :style="{ width: '350px' }"
    header="TVG-ID Düzenle"
    :modal="true"
  >
    <div class="p-input-filled p-p-2">
      <span>Yeni TVG-ID: </span> <InputText v-model="newTvgId" autofocus />
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
    header="Saving"
  >
    <ProgressSpinner style="display:flex" />
  </Dialog>

  <Toast position="bottom-right" />
</template>

<script>
import ProgressSpinner from 'primevue/progressspinner'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
import InputText from 'primevue/inputtext'
import Toolbar from 'primevue/toolbar'
import FileUpload from 'primevue/fileupload'
import Dialog from 'primevue/dialog'
import { useToast } from 'primevue/usetoast'
import Toast from 'primevue/toast'

import { ExportM3u } from '@/utils/m3u.js'
import { deDupe } from '@/utils'
import { ref, computed } from 'vue'
import { root_path } from '@/router/urls'

export default {
  components: {
    DataTable,
    Column,
    InputText,
    Toolbar,
    Dialog,
    FileUpload,
    Toast
  },
  props: {
    m3u: {
      type: Object,
      required: true
    }
  },

  setup(props) {
    const toast = useToast()

    let filters = ref({})

    let renameItemsDialog = ref(false)
    const showRemoveSelected = () => (renameItemsDialog.value = true)

    let BulkTvgIdDialog = ref(false)
    const showEditSelected = () => (BulkTvgIdDialog.value = true)

    let loadingDialog = ref(false)
    let selectedItemsList = ref([])
    let reOrderedList = ref(props.m3u)
    let itemsToSavedList = []

    let downloadButtonLock = ref(true)
    let newTvgId = ref('')
    let removeLabel = computed(() => {
      return selectedItemsList.value.length === 0
        ? 'Sil'
        : `Sil (${selectedItemsList.value.length})`
    })

    let onRowReorder = e => (reOrderedList.value = e.value)
    let onSelectInput = e => e.target.select()
    let onSave = () => {
      loadingDialog.value = true

      itemsToSavedList = deDupe(reOrderedList.value).map(item => {
        return { group_title: item.group_title, chan_name: item.name, tvg_id: item.tvg_id }
      })
      downloadButtonLock.value = false
      fetch(root_path + '/api/save', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(itemsToSavedList)
      })
        .then(resp => resp.json())
        .then(data => {
          if (data.status_code == 200) {
            toast.add({
              severity: 'success',
              summary: 'Success',
              detail: data.message,
              life: 3000
            })
          } else {
            toast.add({
              severity: 'error',
              summary: 'Error',
              detail: `${data.status_code} - ${data.message}`,
              life: 3000
            })
          }
        })
        .catch(error => {
          toast.add({
            severity: 'error',
            summary: 'Error',
            detail: error,
            life: 3000
          })
        })
        .finally(() => (loadingDialog.value = false))
    }

    const removeSelectedItemsList = () => {
      reOrderedList.value = reOrderedList.value.filter(
        val => !selectedItemsList.value.includes(val)
      )
      renameItemsDialog.value = false
      selectedItemsList.value = []
    }
    const editBulkItems = () => {
      for (let item of selectedItemsList.value) {
        for (let i = 0; i < reOrderedList.value.length; i++) {
          if (reOrderedList.value[i].id === item.id) {
            reOrderedList.value[i].tvg_id = newTvgId.value
            break
          }
        }
      }
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
          life: 3000
        })
      }
    }

    return {
      renameItemsDialog,
      selectedItemsList,
      onSelectInput,
      showRemoveSelected,
      onRowReorder,
      reOrderedList,
      downloadButtonLock,
      downloadM3u,
      removeSelectedItemsList,
      removeLabel,
      onSave,
      BulkTvgIdDialog,
      showEditSelected,
      editBulkItems,
      newTvgId,
      filters,
      ProgressSpinner,
      loadingDialog
    }
  }
}
</script>

<style lang="scss" scoped></style>
