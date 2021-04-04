<template>
  <DataTable :m3u='m3uData' :new='true'></DataTable>
</template>

<script setup>
import { inject } from 'vue'
import { useRouter } from 'vue-router'
import DataTable from '../components/DataTable.vue'

const router = useRouter()
let m3uData = inject('GetM3uData')
let m3uDataLength = inject('GetM3uDataLength')

if (m3uDataLength.value === 0) {
  router.replace({ name: 'Index' })
}

</script>
