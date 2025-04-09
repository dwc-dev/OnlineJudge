<template>
  <el-dialog v-model="visible" title="评测结果">
    <el-table :data="results" style="width: 100%">
      <el-table-column prop="output" label="输出" />
      <el-table-column prop="time_milliseconds" label="时间 (ms)" />
      <el-table-column prop="memory_usage" label="内存 (KB)" />
      <el-table-column prop="accepted" label="通过" />
    </el-table>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, defineProps } from 'vue'

const props = defineProps({
  visible: Boolean,
  results: Array,
})
</script>
