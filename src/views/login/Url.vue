<template>
  <div class="p-flex-column">
    <div class="p-field p-fluid">
      <InputText
        id="url"
        v-model="url"
        :class="validateUrl"
        type="text"
        placeholder="Enter URL"
      ></InputText>
    </div>
    <div class="p-field" style="float: right">
      <Button class="p-mr-3" :disabled="sendDisabled" @click="clickSend">Gonder</Button>
      <Button @click="clickReset">Reset</Button>
    </div>
  </div>
</template>

<script setup>
import eventBus from '@/emitter/eventBus.js'

import { ref, computed, inject, defineEmit } from 'vue'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'

import { parse } from 'iptv-playlist-parser'
import { SanitizeM3u } from '@/utils/m3u.js'
import { useRouter } from 'vue-router'

const urlValidateRe = /^http(s)?:\/\/[\w!#$&'()*+,./:;=?@[\]~\-]+$/

let url = ref('')
let sendDisabled = ref(false)

const AddM3uData = inject('AddM3uData')
const router = useRouter()

let validateUrl = computed(() => {
  let urlClass = 'p-inputtext-sm'

  if (url.value.length > 0 && !urlValidateRe.test(url.value)) {
    urlClass = 'p-inputtext-sm p-invalid'
  }

  return urlClass
})

const clickSend = () => {
  eventBus.emit('onStart')
  if (!urlValidateRe.test(url.value) || url.value.length === 0) {
    console.error('invalid url')
  } else {
    sendDisabled.value = true
    fetch(url.value)
      .then((resp) => resp.text())
      .then((data) => {
        AddM3uData(SanitizeM3u(parse(data).items))
        router.push('/epg/view')
      })
      .catch(() => {
        eventBus.emit('pushToast', {
          severity: 'error',
          summary: 'Error',
          message: 'Error on getting data',
        })
      })
      .finally(() => {
        sendDisabled.value = false
        eventBus.emit('onFinish')
      })
  }
}

const clickReset = () => {
  url.value = ''
}
</script>

<style lang="sass" scoped></style>
