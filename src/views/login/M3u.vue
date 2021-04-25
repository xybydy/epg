<template>
  <div class="p-d-flex p-ai-center">
    <label class="p-mr-3" for="uploader">Upload M3u: </label
    ><EpgUploader id="uploader" @m3uLoaded="onM3uLoaded"></EpgUploader>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'

import eventBus from '@/emitter/eventBus.js'

import EpgUploader from '@/components/index/EpgUploader.vue'

const router = useRouter()

eventBus.on('upload_file_size')
eventBus.on('invalid_type')

let onM3uLoaded = () => router.push('/epg/view')
</script>

<style lang="sass" scoped></style>
