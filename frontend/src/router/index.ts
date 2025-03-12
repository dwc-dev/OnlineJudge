import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('@/views/HomeView.vue'),
    },
    {
      path: '/problems',
      name: 'problems',
      component: () => import('@/views/ProblemView.vue'),
    },
    {
      path: '/contests',
      name: 'contests',
      component: () => import('@/views/ContestView.vue'),
    },
    {
      path: '/auth/:type(login|register)', // 动态路由
      name: 'auth',
      component: () => import('@/views/AuthView.vue'),
    },
    {
      path: '/auth',
      redirect: '/auth/login', // 默认重定向到登录页
    },
  ],
})

export default router
