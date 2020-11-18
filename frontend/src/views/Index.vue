<template>
  <div>
    <FileUpload
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
  components: {
    FileUpload,
  },
  setup() {
    const router = useRouter()
    const { AddM3uData } = MainStore()

    const upload = async (value) => {
      value.files[0].text().then((f) => {
        AddM3uData(SanitizeM3u(parse(f).items))
        router.push({ name: 'userEpgView' })
      })
    }
    return { upload }
  },
}
</script>
