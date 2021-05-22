<template>
  <Button v-if="!matchingProgress" type="button" icon="pi pi-cog" @click="matchTvg()"></Button>
  <i v-else class="pi pi-spin pi-spinner" style="font-size: 2.4rem; color: blue"></i>
</template>

<script setup>
  import {ref, defineProps} from 'vue'
  import {useToast} from 'primevue/usetoast'
  import eventBus from '@/emitter/eventBus.js'

  import Button from 'primevue/button'

  import {root_path} from '@/router/url'

  let props = defineProps(['data'])
  let toast = useToast()

  let matchingProgress = ref(false)

  const matchTvg = () => {
    matchingProgress.value = true

    fetch(root_path + '/api/match', {
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      body: JSON.stringify([{name: props.data.data.chan_name}]),
    })
      .then((resp) => resp.json())
      .then((data) => {
        if (data.status_code === 200) {
          data['index'] = props.data.index
          eventBus.emit('matchedTVG', {data: data})
          toast.add({
            severity: 'success',
            summary: 'Success',
            detail: data.message,
            life: 3000,
          })
        } else {
          toast.add({
            severity: 'error',
            summary: data.status_code,
            detail: data.message,
            life: 3000,
          })
        }
      })
      .catch((error) => {
        toast.add({
          severity: 'error',
          summary: 'Error',
          detail: error,
          life: 3000,
        })
      })

    matchingProgress.value = false
  }
</script>
