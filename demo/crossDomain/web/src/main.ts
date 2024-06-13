import { createApp } from 'vue'

import router from './router'
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'

import './styles/tailwind.css'

import App from './App.vue'

const store = createPinia()
store.use(piniaPluginPersistedstate)

createApp(App).use(store).use(router).mount('#app')
