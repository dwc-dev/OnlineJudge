<template>
  <div class="flex flex-col items-center justify-center space-y-6 px-4 py-12">
    <h2
      class="bg-gradient-to-r from-blue-300 to-blue-500 bg-clip-text text-3xl leading-relaxed font-extrabold text-transparent"
    >
      {{ activeTab === 'login' ? loginSentence : registerSentence }}
    </h2>
    <el-tabs v-model="activeTab" class="w-80" @tab-change="handleTabChange">
      <el-tab-pane label="登录" name="login">
        <el-form :model="loginForm" class="mt-4 space-y-6">
          <el-form-item>
            <el-input
              v-model="loginForm.email"
              placeholder="请输入邮箱"
              type="email"
              :prefix-icon="Message"
            />
          </el-form-item>
          <el-form-item>
            <el-input
              v-model="loginForm.password"
              type="password"
              show-password
              placeholder="请输入密码"
              :prefix-icon="Lock"
            />
          </el-form-item>
          <el-button type="primary" class="w-full" @click="handleLogin">登录</el-button>
        </el-form>
      </el-tab-pane>

      <el-tab-pane label="注册" name="register">
        <el-form :model="registerForm" class="mt-4 space-y-6">
          <el-form-item>
            <el-input
              v-model="registerForm.username"
              placeholder="请输入用户名"
              :prefix-icon="User"
            />
          </el-form-item>
          <el-form-item>
            <el-input
              v-model="registerForm.email"
              placeholder="请输入邮箱"
              type="email"
              :prefix-icon="Message"
            />
          </el-form-item>
          <el-form-item>
            <el-input
              v-model="registerForm.password"
              placeholder="请输入密码"
              type="password"
              show-password
              :prefix-icon="Lock"
            />
          </el-form-item>
          <el-button type="primary" class="w-full" @click="handleRegister">注册</el-button>
        </el-form>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script setup lang="ts">
import { ref, watchEffect } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { Message, Lock, User } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import api from '@/api'
const route = useRoute()
const router = useRouter()
const activeTab = ref(route.params.type || 'login')

const handleTabChange = (tabName: string) => {
  router.push(`/auth/${tabName}`)
}

watchEffect(() => {
  activeTab.value = route.params.type || 'login'
})

const loginSentence = 'Your journey continues here.'
const registerSentence = 'Sign up. Level up.'

const loginForm = ref({
  email: '',
  password: '',
})

const registerForm = ref({
  username: '',
  email: '',
  password: '',
})

const handleLogin = async () => {
  api.user.login(loginForm.value.email, loginForm.value.password).then(() => {
    ElMessage.success('登录成功！')
  })
}

const handleRegister = async () => {
  await api.user.register(
    registerForm.value.username,
    registerForm.value.email,
    registerForm.value.password,
  )
  ElMessage.success('注册成功！')
}
</script>

<style scoped></style>
