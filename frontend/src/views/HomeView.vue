<template>
  <div class="p-4">
    <h1 class="text-2xl font-bold">主页</h1>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
onMounted(() => {
  document.title = 'OnlineJudge - 主页'
})
</script>
