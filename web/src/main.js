import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import ToastService from 'primevue/toastservice'

import 'primevue/resources/themes/saga-blue/theme.css'

import 'primeflex/primeflex.css'
import 'primevue/resources/primevue.min.css'

import 'primeicons/primeicons.css'
import '@/style/main.scss'

import Button from 'primevue/button'

const app = createApp(App)
app.use(ToastService)
app.use(router)

app.component('Button', Button)

app.mount('#app')
