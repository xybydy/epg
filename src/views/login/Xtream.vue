<template>
  <div class="p-flex-column">
    <div class="p-field p-d-flex p-jc-between">
      <InputText
        id="username"
        v-model="userRef"
        class="p-inputtext-sm"
        placeholder="Username"
      ></InputText>
      <Password
        v-model="passRef"
        class="p-inputtext-sm"
        placeholder="Password"
        toggleMask
        :feedback="false"
      ></Password>
    </div>
    <div class="p-field p-fluid">
      <InputText
        id="xtream_ "
        v-model="urlRef"
        class="p-inputtext-sm"
        placeholder="Server URL"
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

import { ref, inject } from 'vue'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import Password from 'primevue/password'

import { BuildUrl } from '@/utils/m3u'

import { useRouter } from 'vue-router'
import { parse } from 'iptv-playlist-parser'
import { SanitizeM3u } from '@/utils/m3u.js'

const AddM3uData = inject('AddM3uData')

const urlValidateRe = /^http(s)?:\/\/[\w!#$&'()*+,./:;=?@[\]~\-]+$/

const router = useRouter()

const clickReset = () => {
  passRef.value = ''
  userRef.value = ''
  urlRef.value = ''
}

const clickSend = async () => {
  eventBus.emit('onStart')

  if (!urlValidateRe.test(urlRef.value) || urlRef.value.length === 0) {
    console.error('invalid url')
  } else {
    sendDisabled.value = true

    let url = BuildUrl(userRef.value, passRef.value, urlRef.value)

    fetch(url)
      .then((resp) => {
        return resp.text()
      })
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

let passRef = ref('V8XitVMusx')
let userRef = ref('Cem@lkarsli')
let urlRef = ref('http://tpaneltv.com:8080/')
let sendDisabled = ref(false)
</script>
