<template>
  <VMenuBar class="p-mr-6" />
  <DataTable :m3u="m3uData" :new="true"></DataTable>
  <!-- <MenuBar props.new downloadButtonLock></MenuBar> -->
</template>

<script setup>
import { inject } from 'vue'
import { useRouter } from 'vue-router'
import DataTable from '../components/DataTable.vue'
import VMenuBar from '../components/VMenuBar.vue'

const router = useRouter()
let m3uData = inject('GetM3uData')
let m3uDataLength = inject('LengthM3uData')

if (m3uDataLength.value === 0) {
  router.replace({ name: 'Index' })
}
</script>
