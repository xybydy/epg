<template>
  <DataTable :m3u="m3uData" :new="false"></DataTable>
</template>

<script setup>
  import {inject} from 'vue'
  import DataTable from '../components/DataTable.vue'
  import {root_path} from '@/router/url'

  let m3uData = inject('GetM3uData')
  let AddM3uData = inject('AddM3uData')

  fetch(root_path + '/api/get', {
    method: 'GET',
    headers: {'Content-Type': 'application/json'},
  })
    .then((resp) => resp.json())
    .then((data) => AddM3uData(data.data))
    .catch((error) => console.log(error))
</script>
