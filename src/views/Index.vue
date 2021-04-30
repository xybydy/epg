<template>
  <div class="container p-d-flex p-mx-auto">
    <div class="login p-p-2">
      <TabView>
        <TabPanel header="Upload playlist"> <M3u /></TabPanel>
        <TabPanel header="Enter URL"> <URL /> </TabPanel>
        <TabPanel header="Xtream Login"> <Xtream /> </TabPanel>
      </TabView>
    </div>
  </div>
</template>

<script setup>
import TabView from 'primevue/tabview'
import TabPanel from 'primevue/tabpanel'

import M3u from '@/views/login/M3u.vue'
import URL from '@/views/login/Url.vue'
import Xtream from '@/views/login/Xtream.vue'
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
