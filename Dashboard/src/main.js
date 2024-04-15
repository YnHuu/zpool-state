import './assets/main.css'

import { createApp } from 'vue'

import App from './App.vue'
import router from './router'
import stores from './stores';

import 'virtual:svg-icons-register'
import Sicon from '@/components/Sicon'

import Particles from '@tsparticles/vue3'
import { loadSlim } from '@tsparticles/slim'

const app = createApp(App)

app.use(stores)
app.use(router)
app.use(Particles, {
    init: async engine => {
        await loadSlim(engine);
    },
})
app.component('Sicon', Sicon)
// app.use(Alert)
app.mount('#app')
