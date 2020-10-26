<template>
  <p>qweqeeqeq</p>
  <FileUpload
    name="m3ufile"
    url="/"
    mode="basic"
    :auto="true"
    :maxFileSize="26214400"
    :fileLimit="1"
    :customUpload="true"
    @uploader="upload"
  />
</template>

<script>
import FileUpload from 'primevue/fileupload'

export default {
  setup() {
    const upload = e => {
      if (e.files ===

    }
    return { FileUpload, upload }
  }
}
</script>

<style lang="scss" scoped></style>
