<template>
  <div class="p-4">
    <h1 class="text-2xl font-bold">主页</h1>
    <el-button type="primary" class="mt-2" @click="handleTestButton">TEST</el-button>
  </div>
</template>

<script setup lang="ts">
import api from '@/api'
const handleTestButton = async () => {
  const data = await api.user.getUserInfo()
  alert(JSON.stringify(data?.data))
}
</script>
