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
    @before-send="fatih"
    @upload="fatih"
    @error="fatih"
  />
  <Toast position="bottom-right" />
</template>

<script setup>
import { ref, inject, defineEmit } from 'vue'
import FileUpload from 'primevue/fileupload'
import { useToast } from 'primevue/usetoast'
import Toast from 'primevue/toast'

import { parse } from 'iptv-playlist-parser'
import { SanitizeM3u } from '@/utils/m3u.js'

const emit = defineEmit(['m3u-loaded'])
const toast = useToast()

const FILE_SIZE_LIMIT = 52428800

let getExtension = (e) => {
  return e.split('.').pop()
}

const uploadButton = ref(null)
const AddM3uData = inject('AddM3uData')

function fatih(e) {
  console.log('qweeq')
}

const upload = async (e) => {
  let file = e.files[0]

  let file_ext = getExtension(file.name)
  let file_size = file.size

  if (file_size > FILE_SIZE_LIMIT) {
    uploadButton.value.clear()
    uploadButton.value.uploadedFileCount = 0
    toast.add({
      severity: 'error',
      summary: 'Error',
      detail: 'File size should not be higher than 50MB.',
      life: 3000,
    })
    return
  }

  if (!['m3u', 'm3u8'].includes(file_ext)) {
    uploadButton.value.clear()
    uploadButton.value.uploadedFileCount = 0
    toast.add({
      severity: 'error',
      summary: 'Error',
      detail: 'Invalid file type',
      life: 3000,
    })
    return
  }

  file.text().then((f) => {
    AddM3uData(SanitizeM3u(parse(f).items))
    emit('m3u-loaded')
  })
}
</script>
