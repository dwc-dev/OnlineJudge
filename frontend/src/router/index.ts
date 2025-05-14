import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from '@/stores/user'
import token from '@/api/http/token'
import { ElMessage } from 'element-plus'
import type { RouteLocationNormalized, NavigationGuardNext } from 'vue-router'

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
      path: '/center/:id',
      name: 'center',
      component: () => import('@/views/CenterView.vue'),
    },
    {
      path: '/judgelist',
      name: 'judgelist',
      component: () => import('@/views/JudgeListView.vue'),
      meta: {
        requiresAuth: true,
      },
    },
    {
      path: '/manage',
      name: 'manage',
      redirect: '/manage/user',
      component: () => import('@/views/ManageView.vue'),
      meta: {
        requiresAuth: true,
        roles: ['admin'],
      },
      children: [
        {
          path: 'user',
          name: 'userManage',
          component: () => import('@/views/manage/UserManageView.vue'),
          meta: {
            requiresAuth: true,
            roles: ['admin'],
          },
        },
        {
          path: 'question',
          name: 'questionManage',
          component: () => import('@/views/manage/QuestionManageView.vue'),
          meta: {
            requiresAuth: true,
            roles: ['admin'],
          },
        },
        {
          path: 'competition',
          name: 'competitionManage',
          component: () => import('@/views/manage/CompetitionManageView.vue'),
          meta: {
            requiresAuth: true,
            roles: ['admin'],
          },
        },
      ],
    },
  ],
})

// 定义路由元数据接口
declare module 'vue-router' {
  interface RouteMeta {
    requiresAuth?: boolean
    roles?: string[]
  }
}

// 全局前置守卫
router.beforeEach(
  (to: RouteLocationNormalized, from: RouteLocationNormalized, next: NavigationGuardNext) => {
    const userStore = useUserStore()

    // 判断是否需要登录权限
    const requiresAuth = to.matched.some((record) => record.meta.requiresAuth)

    // 判断用户是否已登录
    const isLoggedIn = userStore.isLogin || !!token.getAccessToken()

    // 判断路由是否需要特定角色
    const requiredRoles: string[] = to.matched.some(
      (record) => record.meta.roles && record.meta.roles.length > 0,
    )
      ? to.matched.find((record) => record.meta.roles)?.meta.roles || []
      : []

    // 如果路由需要认证但用户未登录
    if (requiresAuth && !isLoggedIn) {
      ElMessage.warning('请先登录')
      next({ path: '/auth/login', query: { redirect: to.fullPath } })
      return
    }

    // 如果路由需要特定角色权限
    if (requiredRoles.length > 0 && !requiredRoles.includes(userStore.role)) {
      ElMessage.error('您没有权限访问此页面')
      next({ path: '/' })
      return
    }

    // 如果用户已登录，尝试访问登录/注册页面，重定向到首页
    if (isLoggedIn && to.path.startsWith('/auth')) {
      next({ path: '/' })
      return
    }

    next()
  },
)

export default router
