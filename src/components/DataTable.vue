<template>
  <Toolbar class="p-mb-4">
    <template v-slot:left>
      <FileUpload mode="basic" label="Ekle" :maxFileSize="26214400" :fileLimit="1" chooseLabel="Ekle" class="p-mr-2" />
      <Button
        :label="removeLabel"
        icon="pi pi-trash"
        class="p-button-danger p-mr-2"
        @click="showRemoveSelected"
        :disabled="!selectedItemsList || !selectedItemsList.length"
      />
      <Button
        label="Toplu TVG-ID Düzenle"
        icon="pi pi-file"
        class="p-button-danger p-mr-2"
        @click="showEditSelected"
        :disabled="!selectedItemsList || !selectedItemsList.length"
      />
      <Button label="TVG-ID Getir" icon="pi pi-cloud-download" class="p-button-danger" @click="TvgMatcherDialog" />
    </template>

    <template v-slot:right>
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
    class="p-datatable-sm p-datatable-striped editable-cells-table"
    editMode="cell"
    :filters="filters"
    dataKey="id"
    v-model:selection="selectedItemsList"
    paginatorTemplate="CurrentPageReport FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink RowsPerPageDropdown"
    :value="reOrderedList"
    :paginator="true"
    :rowsPerPageOptions="[25, 50, 100, 250]"
    :rows="25"
    currentPageReportTemplate="Showing {first} to {last} of {totalRecords}"
    @row-reorder="onRowReorder"
  >
    <Column :rowReorder="true" headerStyle="width: 3rem" :reorderableColumn="false" />

    <Column selectionMode="multiple" headerStyle="width: 3em"></Column>

    <Column field="name" header="NAME" autoLayout="true">
      <template #editor="slotProps">
        <InputText v-model="slotProps.data[slotProps.column.props.field]" @focus="onSelectInput" />
      </template>
      <template #filter>
        <InputText type="text" v-model="filters.name" placeholder="İsim ara" />
      </template>
    </Column>
    <Column field="group_title" header="GROUP" autoLayout="true">
      <template #editor="slotProps">
        <InputText v-model="slotProps.data[slotProps.column.props.field]" @focus="onSelectInput" />
      </template>
      <template #filter>
        <InputText type="text" v-model="filters.group_title" placeholder="Group ara" />
      </template>
    </Column>
    <Column field="tvg_id" header="TVG-ID" autoLayout="true">
      <template #editor="slotProps">
        <InputText v-model="slotProps.data[slotProps.column.props.field]" @focus="onSelectInput" /> </template
      ><template #filter> <InputText type="text" v-model="filters.tvg_id" placeholder="TVG-ID ara" /> </template
    ></Column>
    <Column field="tvg_name" header="TVG-NAME" autoLayout="true">
      <template #editor="slotProps">
        <InputText v-model="slotProps.data[slotProps.column.props.field]" @focus="onSelectInput" /> </template
      ><template #filter> <InputText type="text" v-model="filters.tvg_name" placeholder="TVG-NAME ara" /> </template
    ></Column>
    <Column field="tvg_logo" header="TVG-LOGO" autoLayout="true">
      <template #editor="slotProps">
        <InputText v-model="slotProps.data[slotProps.column.props.field]" @focus="onSelectInput" /> </template
    ></Column>
    <Column field="url" header="URL" autoLayout="true">
      <template #editor="slotProps">
        <InputText v-model="slotProps.data[slotProps.column.props.field]" @focus="onSelectInput" /> </template
    ></Column>
  </DataTable>

  <Dialog v-model:visible="renameItemsDialog" :style="{ width: '450px' }" header="Onayla" :modal="true">
    <div class="confirmation-content">
      <i class="pi pi-exclamation-triangle p-mr-3" style="font-size: 2rem" />
      <span v-if="selectedItemsList">Seçtiğin kanallar silinsin mi?</span>
    </div>
    <template #footer>
      <Button label="No" icon="pi pi-times" class="p-button-text" @click="renameItemsDialog = false" />
      <Button label="Yes" icon="pi pi-check" class="p-button-text" @click="removeSelectedItemsList" />
    </template>
  </Dialog>

  <Dialog v-model:visible="BulkTvgIdDialog" :style="{ width: '350px' }" header="TVG-ID Düzenle" :modal="true">
    <div class="p-input-filled p-p-2"><span>Yeni TVG-ID: </span> <InputText autofocus v-model="newTvgId" /></div>
    <template #footer>
      <Button label="Vazgeç" icon="pi pi-times" class="p-button-text" @click="BulkTvgIdDialog = false" />
      <Button label="Onayla" icon="pi pi-check" class="p-button-text" @click="editBulkItems" />
    </template>
  </Dialog>

  <Toast position="bottom-right" />
</template>

<script>
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
import InputText from 'primevue/inputtext'
import Toolbar from 'primevue/toolbar'
import FileUpload from 'primevue/fileupload'
import Dialog from 'primevue/dialog'
import { useToast } from 'primevue/usetoast'
import Toast from 'primevue/toast'

import { ExportM3u } from '@/utils/m3u.js'
import { ref, computed } from 'vue'

export default {
  props: {
    m3u: Object
  },
  setup(props) {
    const toast = useToast()

    let filters = ref({})

    let renameItemsDialog = ref(false)
    const showRemoveSelected = () => (renameItemsDialog.value = true)

    let BulkTvgIdDialog = ref(false)
    const showEditSelected = () => (BulkTvgIdDialog.value = true)

    //const showTvgMatcher = () => ()

    let selectedItemsList = ref([])
    let reOrderedList = ref(props.m3u)
    let itemsToSavedList = ref([])

    let downloadButtonLock = ref(true)
    let newTvgId = ref('')
    let removeLabel = computed(() => {
      if (selectedItemsList.value.length < 1) {
        return 'Sil'
      } else {
        return `Sil (${selectedItemsList.value.length})`
      }
    })

    let onRowReorder = e => (reOrderedList.value = e.value)
    let onSelectInput = e => e.target.select()
    let onSave = () => {
      itemsToSavedList.value = reOrderedList.value
      downloadButtonLock.value = false
      console.log(itemsToSavedList)
      console.log(JSON.stringify(itemsToSavedList))
      fetch('/api/save', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(itemsToSavedList.value)
      })
        .then()
        .then(data => {
          if (data.status == 200) {
            toast.add({
              severity: 'success',
              summary: 'Success',
              detail: data,
              life: 3000
            })
          } else {
            toast.add({
              severity: 'error',
              summary: 'Error',
              detail: `${data.status} - ${data.statusText}`,
              life: 3000
            })
          }
        })
        .catch(data => {
          toast.add({
            severity: 'error',
            summary: 'Error',
            detail: data,
            life: 3000
          })
        })
    }

    const removeSelectedItemsList = () => {
      reOrderedList.value = reOrderedList.value.filter(val => !selectedItemsList.value.includes(val))
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
      if (itemsToSavedList.value.length > 0) {
        ExportM3u(itemsToSavedList.value)
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
      filters
      //showTvgMatcher
    }
  },
  components: {
    DataTable,
    Column,
    InputText,
    Toolbar,
    Dialog,
    FileUpload,
    Toast
  }
}
</script>

<style lang="scss" scoped></style>
