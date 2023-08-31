import { createApp } from 'vue'
import App from './StudentApp.vue'
import { createPinia } from 'pinia';
import router from './router/router.ts';

const pinia = createPinia()
const app = createApp(App)

app.use(router).use(pinia).mount('#app')
