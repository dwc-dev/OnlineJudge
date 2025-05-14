<template>
  <div class="w-full">
    <div class="container mx-auto px-4 py-4">
      <!-- 顶部卡片 -->
      <el-card shadow="never">
        <div class="relative flex justify-between">
          <div class="flex">
            <el-avatar
              :size="100"
              :src="userInfo.user_avatar_url"
              class="cursor-pointer border-4 border-white shadow-lg"
              @click="showPreview = true"
            />
            <el-image-viewer
              v-if="showPreview"
              :url-list="[userInfo.user_avatar_url]"
              :initial-index="0"
              @close="showPreview = false"
            />
            <span class="mt-4 ml-4 text-2xl font-bold text-gray-800">{{ userInfo.user_name }}</span>
          </div>
          <el-button
            type="primary"
            plain
            v-if="userInfo.user_id == currentUid"
            class="absolute right-0 bottom-0"
          >
            编辑信息
          </el-button>
        </div>
      </el-card>

      <!-- 导航标签 -->
      <el-tabs>
        <el-tab-pane label="主页"></el-tab-pane>
      </el-tabs>

      <!-- 主体内容 -->
      <div class="mt-4 grid grid-cols-3 gap-6">
        <!-- 左侧内容 -->
        <div class="col-span-2">
          <el-card shadow="never" class="mb-6">
            <template #header>
              <div class="font-bold">个人简介</div>
            </template>
            <p class="whitespace-pre-line text-gray-600">
              {{ userInfo.user_profile || '暂无简介' }}
            </p>
          </el-card>
        </div>

        <!-- 右侧信息 -->
        <div class="col-span-1">
          <!-- 数据统计
          <el-card shadow="never" class="mb-6">
            <div class="grid grid-cols-3 gap-4 text-center">
              <div class="p-3">
                <p class="text-xl font-bold text-gray-800">0</p>
                <p class="text-sm text-gray-500">提交数</p>
              </div>
              <div class="p-3">
                <p class="text-xl font-bold text-gray-800">0</p>
                <p class="text-sm text-gray-500">通过数</p>
              </div>
              <div class="p-3">
                <p class="text-xl font-bold text-gray-800">0</p>
                <p class="text-sm text-gray-500">尝试题数</p>
              </div>
            </div>
          </el-card> -->
          <!-- 用户信息 -->
          <el-card shadow="never">
            <template #header>
              <div class="font-bold">用户信息</div>
            </template>
            <div class="space-y-3">
              <div class="flex justify-between">
                <span class="text-gray-500">用户编号</span>
                <span>{{ userInfo.user_id }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-500">用户类型</span>
                <span>{{ role }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-500">注册时间</span>
                <span>{{ new Date(userInfo.create_at).toLocaleDateString() }}</span>
              </div>
            </div>
          </el-card>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router'
import { onMounted, ref, computed } from 'vue'
import { ElMessage } from 'element-plus'
import api from '@/api'
import { useUserStore } from '@/stores/user'

const route = useRoute()
const uid = route.params.id
const userStore = useUserStore()
const currentUid = computed(() => userStore.uid)
const userInfo = ref({
  user_id: 0,
  user_name: '',
  user_email: '',
  user_avatar_url: '',
  user_profile: '',
  user_role: '',
  create_at: '',
})
const role = ref('')
const showPreview = ref(false)

onMounted(async () => {
  try {
    const res = await api.user.getUserPublicInfoById(Number(uid))
    userInfo.value = res.data
    role.value = userInfo.value.user_role === 'admin' ? '管理员' : '普通用户'
    document.title = `OnlineJudge - ${userInfo.value.user_name}的个人空间`
  } catch {
    ElMessage.error('获取用户信息失败')
  }
})
</script>

<style scoped></style>
