<template>
  <EpgUploader @m3uLoaded="onM3uLoaded"></EpgUploader>
</template>

<script setup>
import { useRouter } from 'vue-router'
import EpgUploader from '@/components/EpgUploader.vue'

const router = useRouter()
let onM3uLoaded = () => router.push('/epg/view')

</script>
