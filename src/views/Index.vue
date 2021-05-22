<template>
  <div class="container p-d-flex p-mx-auto">
    <BlockUI :blocked="isBlocked">
      <div class="login p-p-2">
        <TabView>
          <TabPanel header="Upload playlist"> <M3u /></TabPanel>
          <TabPanel header="Enter URL"> <URL /> </TabPanel>
          <TabPanel header="Xtream Login"> <Xtream /> </TabPanel>
        </TabView>
      </div>
    </BlockUI>
  </div>
  <Toast position="bottom-right" />
</template>

<script setup>
  import {ref} from 'vue'
  import TabView from 'primevue/tabview'
  import TabPanel from 'primevue/tabpanel'
  import BlockUI from 'primevue/blockui'
  import Toast from 'primevue/toast'

  import M3u from '@/views/login/M3u.vue'
  import URL from '@/views/login/Url.vue'
  import Xtream from '@/views/login/Xtream.vue'

  import eventBus from '@/emitter/eventBus.js'

  import {useToast} from 'primevue/usetoast'

  let toast = useToast()
  let isBlocked = ref(false)

  eventBus.on('onStart', () => (isBlocked.value = true))
  eventBus.on('onFinish', () => (isBlocked.value = false))
  eventBus.on('pushToast', (e) => {
    console.log(e)
    toast.add({
      severity: e.severity,
      summary: e.summary,
      detail: e.message,
      life: 3000,
    })
  })
</script>

<style lang="sass" scoped>
  .container
    min-height: 100vh
    justify-content: center
    align-items: center

  .login
    height: 400px
    border: 2px solid beige
</style>
