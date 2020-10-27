import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import 'primeflex/primeflex.css'

import 'primevue/resources/themes/saga-blue/theme.css'
import 'primeicons/primeicons.css'

import Button from 'primevue/button'

const app = createApp(App)

app.component('Button', Button)
app.use(router)
app.mount('#app')
