<template>
  <div style="display: flex; flex-direction: column; height: 230px; justify-content: space-between">
    <BubbleList :list="list" />
    <Sender />
  </div>
</template>

<script setup lang="ts">
import { BubbleList, Sender } from 'vue-element-plus-x'
const list = [
  {
    content: 'Hello, Element Plus X',
    role: 'user',
  },
]
</script>

<style scoped></style>
