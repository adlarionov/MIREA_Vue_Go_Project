import './main.module.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { PrimeVue } from '@primevue/core'
import Aura from '@primevue/themes/aura'
import ToastService from 'primevue/toastservice'

const app = createApp(App)

app
  .use(ToastService)
  .use(PrimeVue, {
    theme: {
      preset: Aura,
    },
  })
  .use(router)
  .mount('#app')
