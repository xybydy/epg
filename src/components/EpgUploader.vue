<template>
  <FileUpload
    ref="uploadButton"
    name="upload"
    url="/"
    mode="basic"
    :maxFileSize="52428800"
    :fileLimit="1"
    chooseLabel="Yukle"
    customUpload
    auto
    @uploader="upload"
    @error="onUpload"
  />
</template>

<script setup>
import { ref, inject, defineEmit } from 'vue'
import FileUpload from 'primevue/fileupload'

import { parse } from 'iptv-playlist-parser'
import { SanitizeM3u } from '@/utils/m3u.js'

const emit = defineEmit(['m3u-loaded'])

let getExtension = (e) => {
  return e.split('.').pop()
}

const uploadButton = ref(null)
const AddM3uData = inject('AddM3uData')

const onUpload = () => {
  console.log('qweeq')
}

const upload = async (e) => {
  let file = e.files[0]
  let file_ext = getExtension(file.name)

  if (['m3u', 'm3u8'].includes(file_ext)) {
    file.text().then((f) => {
      AddM3uData(SanitizeM3u(parse(f).items))
      emit('m3u-loaded')
    })
  } else {
    return new Error('unsupported file type')
  }

  uploadButton.value.clear()
  uploadButton.value.uploadedFileCount = 0
}
</script>
