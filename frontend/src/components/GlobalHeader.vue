<template>
  <div class="flex max-h-fit items-center border-b-1 border-b-gray-200">
    <RouterLink to="/" class="mr-5 ml-5 text-xl font-bold text-gray-800">OnlineJudge</RouterLink>
    <el-menu mode="horizontal" :default-active="activeIndex" class="flex-1 gap-5" :router="true">
      <el-menu-item index="/">主页</el-menu-item>
      <el-menu-item index="/question">题库</el-menu-item>
      <el-menu-item index="/contest">比赛</el-menu-item>
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
import { ref, computed, watch, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'
import jwt from '@/api/http/jwt'
import api from '@/api'

const userStore = useUserStore()
const route = useRoute()
const router = useRouter()
const activeIndex = ref(route.path)
const isLogin = computed(() => userStore.isLogin)

onMounted(() => {
  if (jwt.getToken()) {
    api.user
      .getUserInfo()
      .then((res) => {
        userStore.setUserInfo({
          uid: res.data.user_id,
          username: res.data.user_name,
          avatar: res.data.user_avatar_url,
          email: res.data.user_email,
          role: res.data.user_role,
          profile: res.data.user_profile,
        })
      })
      .catch((err) => {
        ElMessage.error(err.message)
      })
  }
})

// 监听路由变化，更新激活菜单项
watch(
  () => route.path,
  (newPath) => {
    activeIndex.value = newPath
  },
)

const handleLogin = () => {
  router.push('/auth/login')
}

const handleLogout = () => {
  jwt.clearToken()
  userStore.logout()
  router.push('/')
}

const handleSpace = () => {
  router.push(`/space/${userStore.uid}`)
}

const handleCenter = () => {
  router.push(`/center/${userStore.uid}`)
}
</script>
