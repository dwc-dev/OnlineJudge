import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import token from './api/http/token'
import api from './api'
import { useUserStore } from './stores/user'

const app = createApp(App)

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

const pinia = createPinia()
app.use(pinia)
app.use(router)

// 初始化用户数据
const initUserData = async () => {
  if (token.getAccessToken()) {
    try {
      const userStore = useUserStore()
      const res = await api.user.getUserInfo()
      userStore.setUserInfo({
        uid: res.data.user_id,
        username: res.data.user_name,
        avatar: res.data.user_avatar_url,
        email: res.data.user_email,
        role: res.data.user_role,
        profile: res.data.user_profile,
      })
    } catch (error) {
      console.error('获取用户信息失败', error)
    }
  }
}

// 确保用户数据加载完成后再挂载应用
initUserData().finally(() => {
  app.mount('#app')
})
