<template>
  <FileUpload
    ref="uploadButton"
    name="upload"
    url="/"
    mode="basic"
    :auto="true"
    :maxFileSize="26214400"
    :fileLimit="1"
    chooseLabel="Yukle"
    :customUpload="true"
    @uploader="upload"
  />
</template>

<script>
import { ref, inject } from 'vue'

import FileUpload from 'primevue/fileupload'

import { parse } from 'iptv-playlist-parser'
import { SanitizeM3u } from '@/utils/m3u.js'

export default {
  name: 'EpgUploader',
  components: {
    FileUpload
  },
  emits: ['m3u-loaded'],
  setup(_, { emit }) {
    const uploadButton = ref(null)
    const AddM3uData = inject('AddM3uData')

    const upload = async e => {
      e.files[0].text().then(f => {
        AddM3uData(SanitizeM3u(parse(f).items))
        emit('m3u-loaded')
      })
      uploadButton.value.clear()
    }
    return { upload, uploadButton }
  }
}
</script>
