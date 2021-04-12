<template>
  <La :lee="la"></La>
  <Button label="bas" @click="degis"></Button>
  <GroupRemove></GroupRemove>
</template>

<script setup>
import { ref } from 'vue'
import GroupRemove from '@/components/GroupRemove.vue'
import La from '@/views/La.vue'
import Button from 'primevue/button'

let la = ref('FFFFF')
const degis = () => (la.value = Date.now())
</script>
