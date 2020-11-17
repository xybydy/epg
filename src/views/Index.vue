<template>
  <div>
    <FileUpload
      name="upload"
      url="/"
      mode="basic"
      :auto="true"
      :maxFileSize="26214400"
      :fileLimit="1"
      :customUpload="true"
      @uploader="upload"
      class="p-mb-4"
    />
  </div>
</template>

<script>
import FileUpload from 'primevue/fileupload'
import { parse } from 'iptv-playlist-parser'
import { SanitizeM3u } from '@/utils/m3u.js'
import { useRouter } from 'vue-router'
import MainStore from '@/store/MainStore'

export default {
  name: 'Index',
  setup() {
    const router = useRouter()
    const { AddM3uData } = MainStore()

    const upload = async e => {
      e.files[0].text().then(f => {
        AddM3uData(SanitizeM3u(parse(f).items))
        router.push({ name: 'userEpgView' })
      })
    }
    return { upload }
  },
  components: {
    FileUpload
  }
}
</script>
