<template>
  <div class="h-full w-full" v-loading="loading">
    <el-tabs v-model="activeTab" class="mx-auto mt-10 w-1/2">
      <el-tab-pane label="个人信息" name="info">
        <div class="h-full w-full">
          <el-form label-width="100px">
            <div class="mb-4 flex items-center justify-center">
              <el-avatar
                :size="100"
                :src="userInfo.avatar"
                class="cursor-pointer border-4 border-white shadow-lg"
                @click="handleClickAvatar"
                title="点击更换头像"
              />
            </div>
            <el-form-item label="用户ID">
              <span>{{ userInfo.uid }}</span>
            </el-form-item>
            <el-form-item label="用户名">
              <el-input v-model="userInfo.username" />
            </el-form-item>
            <el-form-item label="邮箱">
              <span>{{ userInfo.email }}</span>
            </el-form-item>
            <el-form-item label="角色">
              <span>{{ roleToText(userInfo.role) }}</span>
            </el-form-item>
            <el-form-item label="个人简介">
              <el-input v-model="userInfo.profile" type="textarea" />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="handleUpdatePersonalInfo">更新个人信息</el-button>
            </el-form-item>
          </el-form>
        </div>
      </el-tab-pane>
      <el-tab-pane label="修改密码" name="password">
        <div class="h-full w-full">
          <el-form label-width="100px">
            <el-form-item label="旧密码">
              <el-input
                v-model="passwordInfo.oldPassword"
                placeholder="请输入旧密码"
                type="password"
                maxlength="30"
                show-password
              />
            </el-form-item>
            <el-form-item label="新密码">
              <el-tooltip placement="right" effect="light">
                <template #content>
                  <p>密码格式要求：</p>
                  <p>1. 密码长度为8-30位</p>
                  <p>2. 密码只能包含数字、大小写字母和特殊字符~!@#$%^&*()_+</p>
                  <p>3. 新密码不能和旧密码相同</p>
                </template>
                <el-input
                  v-model="passwordInfo.newPassword"
                  placeholder="请输入新密码"
                  type="password"
                  maxlength="30"
                  show-password
                />
              </el-tooltip>
            </el-form-item>
            <el-form-item label="确认密码">
              <el-input
                v-model="passwordInfo.confirmPassword"
                placeholder="请再次输入新密码"
                type="password"
                maxlength="30"
                show-password
              />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="handleUpdatePassword">更新密码</el-button>
            </el-form-item>
          </el-form>
        </div>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useUserStore } from '@/stores/user'
import api from '@/api'
import { ElMessage } from 'element-plus'

const userStore = useUserStore()
const loading = ref(false)
const userInfo = ref({
  uid: userStore.uid,
  username: userStore.username,
  avatar: userStore.avatar,
  email: userStore.email,
  role: userStore.role,
  profile: userStore.profile,
  avatar_base64: '',
})

const passwordInfo = ref({
  oldPassword: '',
  newPassword: '',
  confirmPassword: '',
})

const activeTab = ref('info')

const roleToText = (role: string) => {
  return role === 'admin' ? '管理员' : '普通用户'
}

const handleClickAvatar = () => {
  // 选择文件，并转为base64
  const file = document.createElement('input')
  file.type = 'file'
  file.accept = 'image/jpeg, image/png'
  file.onchange = () => {
    userInfo.value.avatar = URL.createObjectURL(file.files![0])
    const reader = new FileReader()
    reader.readAsDataURL(file.files![0])
    reader.onload = (e) => {
      userInfo.value.avatar_base64 = e.target?.result as string
      userInfo.value.avatar_base64 = userInfo.value.avatar_base64.split(',')[1]
    }
  }
  file.click()
}

const handleUpdatePersonalInfo = () => {
  loading.value = true
  if (userInfo.value.username === '') {
    ElMessage.error('请输入用户名')
    loading.value = false
    return
  }
  if (userInfo.value.profile === '') {
    ElMessage.error('请输入个人简介')
    loading.value = false
    return
  }
  const personalInfo = {
    user_name: userInfo.value.username,
    user_profile: userInfo.value.profile,
    avatar_base64: userInfo.value.avatar_base64,
  }
  api.user
    .updatePersonalInfo(personalInfo)
    .then(() => {
      ElMessage.success('更新个人信息成功')
      userStore.username = userInfo.value.username
      userStore.profile = userInfo.value.profile
      userStore.avatar = userInfo.value.avatar
    })
    .catch(() => {
      ElMessage.error('更新个人信息失败')
    })
    .finally(() => {
      loading.value = false
    })
}

const handleUpdatePassword = () => {
  loading.value = true
  if (
    passwordInfo.value.oldPassword === '' ||
    passwordInfo.value.newPassword === '' ||
    passwordInfo.value.confirmPassword === ''
  ) {
    ElMessage.error('请输入完整信息')
    loading.value = false
    return
  }
  if (passwordInfo.value.newPassword !== passwordInfo.value.confirmPassword) {
    ElMessage.error('新密码与确认密码不一致')
    loading.value = false
    return
  }
  api.user
    .updatePassword(passwordInfo.value.oldPassword, passwordInfo.value.newPassword)
    .then(() => {
      ElMessage.success('更新密码成功')
      passwordInfo.value.oldPassword = ''
      passwordInfo.value.newPassword = ''
      passwordInfo.value.confirmPassword = ''
    })
    .catch((error) => {
      ElMessage.error(error.message)
    })
    .finally(() => {
      loading.value = false
    })
}
</script>
