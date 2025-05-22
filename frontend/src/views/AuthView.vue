<template>
  <div class="flex flex-col items-center justify-center space-y-6 px-4 py-12" v-loading="loading">
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
            <el-tooltip placement="right" effect="light">
              <template #content>
                <p>密码格式要求：</p>
                <p>1. 密码长度为8-30位</p>
                <p>2. 密码只能包含数字、大小写字母和特殊字符~!@#$%^&*()_+</p>
                <p>3. 新密码不能和旧密码相同</p>
              </template>
              <el-input
                v-model="registerForm.password"
                placeholder="请输入密码"
                type="password"
                show-password
                :prefix-icon="Lock"
              />
            </el-tooltip>
          </el-form-item>
          <el-form-item>
            <el-input
              v-model="registerForm.confirmPassword"
              placeholder="请再次输入密码"
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
import { ref, watchEffect, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { Message, Lock, User } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/stores/user'
import api from '@/api'
const route = useRoute()
const router = useRouter()
const activeTab = ref(route.params.type || 'login')
const userStore = useUserStore()
const loading = ref(false)
const handleTabChange = (tabName: string) => {
  router.push(`/auth/${tabName}`)
}

watchEffect(() => {
  activeTab.value = route.params.type || 'login'
})

onMounted(() => {
  document.title = 'OnlineJudge - 登录/注册'
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
  confirmPassword: '',
})

const handleLogin = async () => {
  loading.value = true
  api.user
    .login(loginForm.value.email, loginForm.value.password)
    .then(() => {
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
          ElMessage.success('登录成功！')
          router.push('/')
        })
        .catch((err) => {
          ElMessage.error(err.message)
        })
    })
    .catch((err) => {
      ElMessage.error(err.message)
    })
    .finally(() => {
      loading.value = false
    })
}

const handleRegister = async () => {
  loading.value = true
  if (
    registerForm.value.password === '' ||
    registerForm.value.confirmPassword === '' ||
    registerForm.value.username === '' ||
    registerForm.value.email === ''
  ) {
    ElMessage.error('请填写完整信息')
    loading.value = false
    return
  }
  if (registerForm.value.password !== registerForm.value.confirmPassword) {
    ElMessage.error('两次输入的密码不一致')
    loading.value = false
    return
  }
  try {
    await api.user.register(
      registerForm.value.username,
      registerForm.value.email,
      registerForm.value.password,
    )
    activeTab.value = 'login'
    loginForm.value.email = registerForm.value.email
    loginForm.value.password = registerForm.value.password
    registerForm.value.username = ''
    registerForm.value.email = ''
    registerForm.value.password = ''
    registerForm.value.confirmPassword = ''
    loading.value = false
    ElMessage.success('注册成功！')
  } catch (error: unknown) {
    if (error instanceof Error) {
      ElMessage.error(error.message)
    }
    loading.value = false
  }
}
</script>

<style scoped></style>
