export const menuConfig = [
  {
    path: '/',
    name: '主页',
    roles: [],
    requiresAuth: true,
  },
  {
    path: '/question',
    name: '题库',
    roles: ['user', 'admin'],
    requiresAuth: false,
  },
  {
    path: '/competition',
    name: '比赛',
    roles: ['user', 'admin'],
    requiresAuth: false,
  },
  {
    path: '/judgelist',
    name: '评测记录',
    roles: ['user', 'admin'],
    requiresAuth: true,
  },
  {
    path: '/manage',
    name: '系统管理',
    roles: ['admin'],
    requiresAuth: true,
  },
]
