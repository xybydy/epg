<template>
  <DataTable
    v-model:filters="filters"
    class="p-datatable-sm p-datatable-striped editable-cells-table"
    dataKey="_id"
    :rowsPerPageOptions="[25, 50, 100, 250]"
    :rows="25"
    editMode="cell"
    paginator
    currentPageReportTemplate="Showing {first} to {last} of {totalRecords}"
  >
    <template #empty> Sonuc bulunamadi.</template>
    <template #header>
      <div class="p-d-flex p-jc-end">
        <span class="p-input-icon-left">
          <i class="pi pi-search" />
          <InputText placeholder="Ara" @keyup.enter="onSearch($event.target.value)" />
        </span>
      </div>
    </template>
  </DataTable>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import DataTable from 'primevue/datatable'
import InputText from 'primevue/inputtext'
import { FilterMatchMode } from 'primevue/api'
import { useToast } from 'primevue/usetoast'
import Toast from 'primevue/toast'

import { root_path } from '@/router/url'

const toast = useToast()

const onSearch = (e) => {
  filters.value.global.value = e
}

onMounted(() => {
  fetch(root_path + '/api/tvg')
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
})

let filters = ref({ global: { value: null, matchMode: FilterMatchMode.CONTAINS } })
</script>
