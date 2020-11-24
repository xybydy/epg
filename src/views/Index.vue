<template>
  <EpgUploader @m3uLoaded="onM3uLoaded"></EpgUploader>
</template>

<script>
import { useRouter } from 'vue-router'
import EpgUploader from '@/components/EpgUploader'

export default {
  name: 'Index',
  components: {
    EpgUploader
  },
  setup() {
    const router = useRouter()
    let onM3uLoaded = () => router.push('/epg/view')
    return { EpgUploader, onM3uLoaded }
  }
}
</script>
