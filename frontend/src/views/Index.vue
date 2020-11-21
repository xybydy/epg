<template>
  <div class="p-d-flex p-flex-column p-mt-3">
    <FileUpload
      v-if="!m3uData.isLoaded"
      ref="uploadButton"
      name="upload"
      url="/"
      mode="basic"
      :auto="true"
      :max-file-size="26214400"
      :file-limit="1"
      :custom-upload="true"
      class="p-mb-4"
      @uploader="upload"
    />
    <DataTable v-if="m3uData.isLoaded" :m3u="m3uData.m3uList"></DataTable>
  </div>
</template>

<script>
import { ref, reactive } from 'vue'

import FileUpload from 'primevue/fileupload'

import { parse } from 'iptv-playlist-parser'
import { SanitizeM3u } from '@/utils/m3u.js'

import DataTable from '../components/DataTable.vue'

export default {
  name: 'Index',
  components: {
    FileUpload,
    DataTable
  },
  setup() {
    let m3uData = reactive({
      m3uList: [],
      isLoaded: false
    })
    const uploadButton = ref(null)

    const upload = async e => {
      e.files[0].text().then(f => {
        m3uData.m3uList = SanitizeM3u(parse(f).items)
        m3uData.isLoaded = true
      })
      uploadButton.value.clear()
    }
    return { upload, uploadButton, m3uData }
  }
}
</script>
