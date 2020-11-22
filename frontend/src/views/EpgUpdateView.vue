<template>
  <DataTable :m3u="m3uData" :new="false"></DataTable>
</template>

<script>
import { inject } from 'vue'
import DataTable from '../components/DataTable.vue'

export default {
  components: {
    DataTable
  },
  setup() {
    let m3uData = inject('GetM3uData')
    return { m3uData }
  }
}
</script>
