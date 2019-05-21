<template>
  <Toast />

  <La ref="lla" :lee="la">sadsdad</La>
  <Button label="bas" @click="degis"></Button>
  <Button label="LA" @click="diger"></Button>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import La from '@/views/La.vue'
import Button from 'primevue/button'
import Toast from 'primevue/toast'
import { useToast } from 'primevue/usetoast'

let toast = useToast()

let la = ref('FFFFF')
let lla = ref(null)

onMounted(() => {
  console.log(lla.value)
})

const degis = () => (la.value = Date.now())

const diger = () => {
  toast.add({
    severity: 'error',
    summary: 'Error',
    detail: 'Invalid file type',
    life: 3000,
  })
}
</script>
