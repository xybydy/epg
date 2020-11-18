<template>
  <DataTable v-if="GetM3uDataLength > 0" />
</template>

<script>
import DataTable from '@/components/DataTable.vue'
import MainStore from '@/store/MainStore'
import { useRouter } from 'vue-router'
import { onBeforeMount } from 'vue'

export default {
  components: {
    DataTable,
  },
  setup() {
    let router = useRouter()
    const { GetM3uDataLength } = MainStore()

    onBeforeMount(() => {
      if (GetM3uDataLength.value <= 0) {
        router.push({ name: 'index' })
      }
    })
    return { GetM3uDataLength }
  },
}
</script>
