<template>
  <div class="flex max-h-fit items-center border-b-1 border-b-gray-200">
    <RouterLink to="/" class="mr-5 ml-5 text-xl font-bold text-gray-800">OnlineJudge</RouterLink>
    <el-menu mode="horizontal" :default-active="activeIndex" class="flex-1 gap-5" :router="true">
      <el-menu-item index="/">主页</el-menu-item>
      <el-menu-item index="/question">题库</el-menu-item>
      <el-menu-item index="/contest">比赛</el-menu-item>
    </el-menu>
    <div>
      <el-button plain class="mr-5" @click="handleLogin">登录/注册</el-button>
    </div>
  </div>
</template>

<style scoped>
.el-menu--horizontal.el-menu {
  border-bottom: none;
}
</style>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
const router = useRouter()
const activeIndex = ref(route.path)

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
</script>
