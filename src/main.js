import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import ToastService from 'primevue/toastservice'

import 'primevue/resources/themes/saga-blue/theme.css'
import 'primeflex/primeflex.css'
import 'primevue/resources/primevue.min.css'
import 'primeicons/primeicons.css'

import '@/style/main.scss'

import { GetM3uData, LengthM3uData, AddM3uData } from '@/store/MainStore'

const app = createApp(App)

app
  .provide('GetM3uData', GetM3uData)
  .provide('LengthM3uData', LengthM3uData)
  .provide('AddM3uData', AddM3uData)
  .use(ToastService)
  .use(router)

app.mount('#app')
