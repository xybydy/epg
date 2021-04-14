<template>
  <Button v-if="!matchingProgress" type="button" icon="pi pi-cog" @click="matchTvg()"></Button>
  <i v-else class="pi pi-spin pi-spinner" style="font-size: 2.4rem; color: blue"></i>
</template>

<script setup>
import { ref, defineProps } from 'vue'
import Button from 'primevue/button'

let props = defineProps(['data'])

let matchingProgress = ref(false)

const matchTvg = () => {
  matchingProgress.value = true
}
</script>
