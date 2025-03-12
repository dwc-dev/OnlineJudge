<template>
  <el-container class="min-h-screen">
    <el-header class="!p-0">
      <GlobalHeader />
    </el-header>
    <el-main>
      <RouterView />
    </el-main>
    <el-footer class="flex h-14 items-center justify-center bg-gray-100">
      <span class="text-gray-600">OnlineJudge Version 1.0.0</span>
    </el-footer>
  </el-container>
</template>

<script setup lang="ts">
import GlobalHeader from '@/components/GlobalHeader.vue'
</script>
