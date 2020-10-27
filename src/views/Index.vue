<template>
  <div class="p-d-flex p-flex-column p-mt-3">
    <FileUpload
      name="upload"
      url="/"
      ref="uploadButton"
      mode="basic"
      :auto="true"
      :maxFileSize="26214400"
      :fileLimit="1"
      :customUpload="true"
      @uploader="upload"
      class="p-mb-4"
    />
    <DataTable v-if="m3uData.isLoaded" :m3u="m3uData.m3uList"></DataTable>
  </div>
</template>

<script>
import { ref, reactive } from 'vue'

import FileUpload from 'primevue/fileupload'

import { parse } from 'iptv-playlist-parser'

import DataTable from '../components/DataTable.vue'

export default {
  name: 'Index',
  setup() {
    let m3uData = reactive({
      m3uList: [],
      isLoaded: false
    })
    const uploadButton = ref(null)

    const upload = async e => {
      e.files[0].text().then(f => {
        m3uData.m3uList = parse(f).items
        m3uData.isLoaded = true
      })
      uploadButton.value.clear()
    }

    return { upload, uploadButton, m3uData }
  },
  components: {
    FileUpload,
    DataTable
  }
}
</script>
