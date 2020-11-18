<template>
  <Toolbar class="p-mb-4">
    <template #left>
      <FileUpload
        mode="basic"
        label="Ekle"
        :max-file-size="26214400"
        :file-limit="1"
        choose-label="Ekle"
        class="p-mr-2"
      />
      <Button
        :label="removeLabel"
        icon="pi pi-trash"
        class="p-button-danger p-mr-2"
        :disabled="!selectedItemsList || !selectedItemsList.length"
        @click="showRemoveSelected"
      />
      <Button
        label="Toplu TVG-ID Düzenle"
        icon="pi pi-file"
        class="p-button-danger p-mr-2"
        :disabled="!selectedItemsList || !selectedItemsList.length"
        @click="showEditSelected"
      />
      <!-- <Button
        label="TVG-ID Getir"
        icon="pi pi-cloud-download"
        class="p-button-danger"
        @click="TvgMatcherDialog"
      /> -->
    </template>

    <template #right>
      <Button label="Kaydet" icon="pi pi-save" class="p-button-help p-mr-2" @click="onSave" />
      <Button
        :disabled="downloadButtonLock"
        label="İndir"
        icon="pi pi-download"
        class="p-button-help"
        @click="downloadM3u"
      />
    </template>
  </Toolbar>
</template>

<script>
import Toolbar from 'primevue/toolbar'
import FileUpload from 'primevue/fileupload'

export default {
  components: {
    Toolbar,
    FileUpload,
  },
  setup() {
    return {}
  },
}
</script>
