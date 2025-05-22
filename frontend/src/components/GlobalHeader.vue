<template>
  <div class="flex items-center">
    <RouterLink to="/" class="mr-5 ml-5 text-xl font-bold text-gray-800">OnlineJudge</RouterLink>
    <el-menu mode="horizontal" :default-active="activeIndex" class="flex-1 gap-2" :router="true">
      <el-menu-item v-for="item in menuItems" :key="item.path" :index="item.path">{{
        item.name
      }}</el-menu-item>
    </el-menu>
    <div v-if="!isLogin">
      <el-button plain class="mr-5" @click="handleLogin">登录/注册</el-button>
    </div>
    <div v-else class="mr-5">
      <el-dropdown>
        <el-avatar :size="40" :src="userStore.avatar" class="cursor-pointer" />
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item @click="handleSpace">我的空间</el-dropdown-item>
            <el-dropdown-item @click="handleCenter">个人中心</el-dropdown-item>
            <el-dropdown-item @click="handleLogout">退出登录</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
</template>

<style scoped>
.el-menu--horizontal.el-menu {
  border-bottom: none;
}
</style>

<script setup lang="ts">
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { menuConfig } from '@/config/menuConfig'
import token from '@/api/http/token'
import api from '@/api'
import { ElMessage } from 'element-plus'
const userStore = useUserStore()
const route = useRoute()
const router = useRouter()
const isLogin = computed(() => userStore.isLogin)
const activeIndex = computed(() => {
  if (route.path.includes('/manage')) {
    return '/manage'
  }
  return route.path
})

const menuItems = computed(() => {
  return menuConfig.filter((item) => {
    if (item.requiresAuth) {
      return item.roles.includes(userStore.role)
    }
    return true
  })
})

const handleLogin = () => {
  router.push('/auth/login')
}

const handleLogout = async () => {
  await api.user.logout()
  token.clearAccessToken()
  token.clearRefreshToken()
  userStore.logout()
  router.push('/')
  ElMessage.success('退出成功')
}

const handleSpace = () => {
  router.push(`/space/${userStore.uid}`)
}

const handleCenter = () => {
  router.push('/center')
}
</script>
