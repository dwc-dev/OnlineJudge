<template>
  <div class="flex h-full w-full">
    <el-menu class="!w-50" :router="true" :default-active="activeIndex">
      <el-menu-item index="/manage/user">用户管理</el-menu-item>
      <el-menu-item index="/manage/question">题目管理</el-menu-item>
      <el-menu-item index="/manage/competition">比赛管理</el-menu-item>
    </el-menu>
    <RouterView />
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const activeIndex = computed(() => route.path)
</script>
