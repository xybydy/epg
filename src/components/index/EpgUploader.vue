g
<template>
  <FileUpload
    ref="uploadButton"
    name="fileUploader"
    url="/"
    mode="basic"
    :fileLimit="1"
    chooseLabel="Yukle"
    customUpload
    auto
    @uploader="upload"
  />
</template>

<script setup>
import { ref, inject, defineEmit } from 'vue'
import FileUpload from 'primevue/fileupload'

import { parse } from 'iptv-playlist-parser'
import { SanitizeM3u } from '@/utils/m3u.js'
import eventBus from '../../emitter/eventBus'

const emit = defineEmit(['m3u-loaded'])

const FILE_SIZE_LIMIT = 52428800

let getExtension = (e) => {
  return e.split('.').pop()
}

const uploadButton = ref(null)
const AddM3uData = inject('AddM3uData')

const upload = async (e) => {
  let file = e.files[0]

  let file_ext = getExtension(file.name)
  let file_size = file.size

  if (file_size > FILE_SIZE_LIMIT) {
    uploadButton.value.clear()
    uploadButton.value.uploadedFileCount = 0
    eventBus.emit('upload_file_size')
    return
  }

  if (!['m3u', 'm3u8'].includes(file_ext)) {
    uploadButton.value.clear()
    uploadButton.value.uploadedFileCount = 0
    eventBus.emit('invalid_type')
    return
  }

  file.text().then((f) => {
    AddM3uData(SanitizeM3u(parse(f).items))
    emit('m3u-loaded')
  })
}
</script>
