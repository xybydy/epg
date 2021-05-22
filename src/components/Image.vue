<template>
  <img :src="imgUrl" :alt="alt" style="height: 48px" @error="onError" />
</template>

<script setup>
  import {defineProps, ref, computed, watch} from 'vue'
  import no_image from '@/assets/no-image.png'

  let err = ref(false)

  const onError = () => {
    if (!err.value) {
      err.value = true
    }
  }

  const imgUrl = computed(() => {
    return err.value ? no_image : props.source
  })

  const props = defineProps(['source', 'alt'])

  watch(
    () => props.source,
    () => {
      err.value = false
    }
  )
</script>
