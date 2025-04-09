<template>
  <el-container class="flex h-screen flex-col">
    <el-header class="!p-0">
      <GlobalHeader />
    </el-header>
    <el-main class="flex-1 !p-0">
      <RouterView />
    </el-main>
    <el-footer v-if="!isCodingPage" class="flex !h-10 items-center justify-center bg-gray-100">
      <span class="text-base text-gray-500">OnlineJudge Version 1.0.0</span>
    </el-footer>
  </el-container>
</template>

<script setup lang="ts">
import GlobalHeader from '@/components/GlobalHeader.vue'
import { computed } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const isCodingPage = computed(() => route.name === 'questionDetail')
</script>
