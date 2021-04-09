<template>
  <Dialog visible :style="{ width: '200px' }" :modal="true" :closable="false" header="Please Wait">
    <ProgressSpinner style="display: flex" />
  </Dialog>
  <La :lee="la"></La>
  <Button label="bas" @click="degis"></Button>
</template>

<script setup>
import { ref } from 'vue'
import La from '@/views/La.vue'
import Button from 'primevue/button'
import Dialog from 'primevue/dialog'
import ProgressSpinner from 'primevue/progressspinner'

let la = ref('FFFFF')
const degis = () => (la.value = Date.now())
</script>
