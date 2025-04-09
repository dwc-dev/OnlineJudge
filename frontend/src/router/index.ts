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
      path: '/question',
      name: 'questionList',
      component: () => import('@/views/QuestionView.vue'),
    },
    {
      path: '/question/:id',
      name: 'questionDetail',
      component: () => import('@/views/CodingView.vue'),
    },
    {
      path: '/contest',
      name: 'contest',
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
