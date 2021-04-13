<template>
  <La ref="lla" :lee="la">sadsdad</La>
  <Button label="bas" @click="degis"></Button>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import La from '@/views/La.vue'
import Button from 'primevue/button'

let la = ref('FFFFF')
let lla = ref(null)

onMounted(() => {
  console.log(lla.value)
})

const degis = () => (la.value = Date.now())
</script>
