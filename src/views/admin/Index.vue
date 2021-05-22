<template>
  <DataTable
    v-model:filters="filters"
    v-model:selection="selectedItems"
    dataKey="_id"
    :value="values"
    :loading="loading"
    :paginator="true"
    :rows="50"
    :rowsPerPageOptions="[25, 50, 100, 250]"
  >
    <template #header>
      <div class="p-d-flex p-jc-between">
        <div class="p-d-flex">
          <Button label="Add" icon="pi pi-plus" class="p-mr-2" @click="addNew" />
          <Button
            label="Delete"
            class="p-button-danger"
            icon="pi pi-trash"
            :disabled="!selectedItems || !selectedItems.length"
            @click="delSelected"
          />
        </div>
        <span class="p-input-icon-left">
          <i class="pi pi-search" />
          <InputText placeholder="Ara" @keyup.enter="onSearch($event.target.value)" />
        </span></div
    ></template>
    <Column selectionMode="multiple" headerStyle="width: 4em" />

    <Column field="logo" header="" headerStyle="width: 6em">
      <template #body="slotProps">
        <Image :source="slotProps.data.logo" :alt="slotProps.data.name" />
      </template>
    </Column>
    <Column field="name" header="Name" />
    <Column field="tvg_id" header="Tvg-ID" />
    <Column field="country" header="Country" />
  </DataTable>

  <NewChannel />
  <Toast />
</template>

<script setup>
  import {onBeforeMount, ref, onMounted} from 'vue'
  import eventBus from '@/emitter/eventBus.js'

  import ToastService from 'primevue/toastservice'
  import DataTable from 'primevue/datatable'
  import InputText from 'primevue/inputtext'
  import Column from 'primevue/column'
  import Button from 'primevue/button'
  import Toast from 'primevue/toast'

  import {FilterMatchMode} from 'primevue/api'

  import {root_path} from '../../router/url'

  import Image from '@/components/Image.vue'
  import NewChannel from '@/components/admin/AddNewChannel.vue'
  import {useToast} from 'primevue/usetoast'

  let toast = useToast()

  let values = ref([])
  let loading = ref(true)
  let filters = ref({global: {value: null, matchMode: FilterMatchMode.CONTAINS}})
  let selectedItems = ref([])

  const onSearch = (e) => {
    filters.value.global.value = e
  }

  const addNew = () => {
    eventBus.emit('adminNewChanDialog')
  }

  const delSelected = () => {
    let ids = selectedItems.value.map((x) => x._id)
    fetch(root_path + '/api/chan/delete', {
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      body: JSON.stringify(ids),
    })
      .then((resp) => resp.json())
      .then((data) => {
        values.value = values.value.filter((val) => !selectedItems.value.includes(val))
        toast.add({
          severity: 'success',
          summary: 'Bilgi',
          detail: 'Channel removed successfully',
          life: 3000,
        })
        console.log(data)
      })
      .catch((error) => {
        toast.add({
          severity: 'error',
          summary: 'Hata',
          detail: 'Error: please check console logs',
          life: 3000,
        })
        console.error(error)
      })
  }

  onMounted(() => {
    eventBus.on('adminNewChanSend', (e) => {
      fetch(root_path + '/api/chan/add', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(e),
      })
        .then((resp) => resp.json())
        .then((data) => {
          if (data.status_code === 200 && data.data.length > 0) {
            let obj = {
              _id: data.data[0],
              country: e.country,
              logo: e.logo,
              name: e.name,
              tvg_id: e.tvg_id,
            }
            values.value.push(obj)
            toast.add({
              severity: 'success',
              summary: 'Bilgi',
              detail: 'Channel added successfully',
              life: 3000,
            })
          } else {
            toast.add({
              severity: 'error',
              summary: 'Hata',
              detail: 'Error: please check console logs',
              life: 3000,
            })
            console.error(data)
          }
        })
        .catch((error) => {
          toast.add({
            severity: 'error',
            summary: 'Hata',
            detail: 'Error: please check console logs',
            life: 3000,
          })
          console.error(error)
        })
    })
  })

  onBeforeMount(() => {
    fetch(root_path + '/api/all', {method: 'GET', headers: {'Content-Type': 'application/json'}})
      .then((resp) => resp.json())
      .then((data) => {
        if (data.status_code === 200) {
          console.log(data)
          values.value = data.data
          loading.value = false
        }
      })
  })
</script>
