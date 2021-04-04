<template>
  <FileUpload
    ref='uploadButton'
    name='upload'
    url='/'
    mode='basic'
    :auto='true'
    :maxFileSize='26214400'
    :fileLimit='1'
    chooseLabel='Yukle'
    :customUpload='true'
    @uploader='upload'
  />
</template>

<script setup>
import { ref, inject, defineEmit } from 'vue'
import FileUpload from 'primevue/fileupload'

import { parse } from 'iptv-playlist-parser'
import { SanitizeM3u } from '@/utils/m3u.js'

const emit = defineEmit(['m3u-loaded'])

const uploadButton = ref(null)
const AddM3uData = inject('AddM3uData')

const upload = async e => {
  e.files[0].text().then(f => {
    AddM3uData(SanitizeM3u(parse(f).items))
    emit('m3u-loaded')
  })
  uploadButton.value.clear()
}

</script>
