<template>
  <el-dialog
    v-model="visible"
    title="评测结果"
    width="80%"
    :close-on-click-modal="false"
    :close-on-press-escape="true"
  >
    <div class="w-full">
      <div v-if="judgeResults.length === 0" class="py-4 text-center text-gray-500">
        暂无评测结果
      </div>
      <div v-else>
        <div class="mb-4 flex items-center justify-between">
          <div class="text-lg font-bold">共 {{ judgeResults.length }} 个测试用例</div>
          <div class="flex items-center gap-4">
            <div class="flex items-center gap-1">
              <div class="h-3 w-3 rounded-full bg-[--el-color-success-light-9]"></div>
              <span>通过: {{ passedCount }}</span>
            </div>
            <div class="flex items-center gap-1">
              <div class="h-3 w-3 rounded-full bg-[--el-color-danger-light-9]"></div>
              <span>未通过: {{ judgeResults.length - passedCount }}</span>
            </div>
          </div>
        </div>

        <el-scrollbar height="500px" class="w-full">
          <el-row :gutter="12" class="w-full">
            <el-col
              :xs="24"
              :sm="12"
              :md="8"
              :lg="6"
              v-for="(result, index) in judgeResults"
              :key="index"
              class="mb-3"
            >
              <el-card
                shadow="never"
                :body-style="{ padding: '12px' }"
                :style="{
                  backgroundColor: result.accepted
                    ? 'var(--el-color-success-light-9)'
                    : 'var(--el-color-danger-light-9)',
                }"
              >
                <div class="flex flex-col gap-2">
                  <div class="flex items-center justify-between">
                    <div class="text-sm font-bold">测试用例 #{{ index + 1 }}</div>
                    <el-tag size="small" :type="result.accepted ? 'success' : 'danger'">
                      {{ result.accepted ? '通过' : '未通过' }}
                    </el-tag>
                  </div>
                  <div class="flex justify-between text-xs">
                    <div class="flex items-center gap-1" title="运行时间">
                      <el-icon><Clock /></el-icon>
                      <span>{{ result.time_milliseconds }} ms</span>
                    </div>
                    <div class="flex items-center gap-1" title="内存使用">
                      <el-icon><Cpu /></el-icon>
                      <span>{{ result.memory_usage }} KB</span>
                    </div>
                  </div>
                  <div class="mt-1 cursor-pointer" @click="showOutputDetail(result.output)">
                    <div class="mb-1 text-xs font-medium">输出：</div>
                    <div
                      class="overflow-hidden rounded bg-gray-100 p-1 text-xs text-ellipsis whitespace-nowrap"
                    >
                      {{ result.output }}
                    </div>
                  </div>
                </div>
              </el-card>
            </el-col>
          </el-row>
        </el-scrollbar>
      </div>
    </div>

    <!-- 输出详情对话框 -->
    <el-dialog v-model="outputDialogVisible" title="输出详情" width="50%" append-to-body>
      <pre class="max-h-96 overflow-auto rounded bg-gray-100 p-3">{{ selectedOutput }}</pre>
    </el-dialog>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, defineExpose, computed } from 'vue'
import { Clock, Cpu } from '@element-plus/icons-vue'

interface JudgeResult {
  output: string
  time_milliseconds: number
  memory_usage: number
  accepted: boolean
}

const visible = ref(false)
const judgeResults = ref<JudgeResult[]>([])
const outputDialogVisible = ref(false)
const selectedOutput = ref('')

// 计算通过的用例数
const passedCount = computed(() => {
  return judgeResults.value.filter((result) => result.accepted).length
})

const open = (results: JudgeResult[]) => {
  judgeResults.value = results
  visible.value = true
}

const showOutputDetail = (output: string) => {
  selectedOutput.value = output
  outputDialogVisible.value = true
}

defineExpose({
  open,
})
</script>

<style scoped></style>
