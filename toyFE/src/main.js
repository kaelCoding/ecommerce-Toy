import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { init_store } from './stores'
import loadFile from './plugins/loadFile'

const app = createApp(App)

const init_app = async () => {
    await init_store()
    app.use(router)
    app.use(loadFile)
    app.mount('#app')
}

init_app()


