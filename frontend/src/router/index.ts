import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      redirect: '/question',
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
      path: '/competition/:id/question/:qid',
      name: 'competitionQuestion',
      component: () => import('@/views/CodingView.vue'),
    },
    {
      path: '/competition',
      name: 'competition',
      component: () => import('@/views/CompetitionView.vue'),
    },
    {
      path: '/competition/:id',
      name: 'competitionDetail',
      component: () => import('@/views/CompetitionDetailView.vue'),
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
    {
      path: '/space/:id',
      name: 'space',
      component: () => import('@/views/SpaceView.vue'),
    },
    {
      path: '/center',
      name: 'center',
      component: () => import('@/views/CenterView.vue'),
    },
    {
      path: '/judgelist',
      name: 'judgelist',
      component: () => import('@/views/JudgeListView.vue'),
    },
    {
      path: '/manage',
      name: 'manage',
      redirect: '/manage/user',
      component: () => import('@/views/ManageView.vue'),
      children: [
        {
          path: 'user',
          name: 'userManage',
          component: () => import('@/views/manage/UserManageView.vue'),
        },
        {
          path: 'question',
          name: 'questionManage',
          component: () => import('@/views/manage/QuestionManageView.vue'),
        },
        {
          path: 'competition',
          name: 'competitionManage',
          component: () => import('@/views/manage/CompetitionManageView.vue'),
        },
      ],
    },
  ],
})

export default router
