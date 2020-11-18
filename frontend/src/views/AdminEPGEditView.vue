<template>
  <DataTable />
</template>

<script>
import DataTable from '@/components/DataTable.vue'
export default {
  components: {
    DataTable,
  },
  setup() {
    fetch('http://localhost:8000/api/epg/get', {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
    })
      .then((resp) => resp.json())
      .then((data) => {
        if (data.status_code == 200) {
          console.log(data)
        }
      })
  },
}
</script>

<style lang="scss" scoped></style>
