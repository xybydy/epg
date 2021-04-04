<template>
  <h1>{{ foo }}</h1>
</template>

<script setup>
import { defineProps } from 'vue'

const props = defineProps({
  foo: {
    type: String,
    required: false
  }
})


</script>
