import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import './style.css'

import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import ws from './assets/js/websocket'
import request from './util/request'


const app = createApp(App)
app.use(store).use(router).use(ElementPlus).mount('#app')
app.config.globalProperties.$axios = request
app.config.globalProperties.$ws = ws
//全局引入所有图标
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}