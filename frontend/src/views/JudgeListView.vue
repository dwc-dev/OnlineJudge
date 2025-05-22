<template>
  <div class="p-4">
    <h1 class="text-2xl font-bold">评测记录</h1>

    <div class="mt-6">
      <el-table :data="judgeList" stripe class="w-full">
        <el-table-column prop="judge_id" label="评测ID" width="100" />
        <el-table-column prop="user_id" label="用户ID" width="100" />
        <el-table-column prop="question_id" label="题目ID" width="100" />
        <el-table-column prop="language" label="语言" width="100">
          <template #default="scope">
            {{ judgeLanguageMap[scope.row.language] }}
          </template>
        </el-table-column>
        <el-table-column prop="accepted" label="评测结果" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.accepted ? 'success' : 'danger'">
              {{ scope.row.accepted ? '通过' : '未通过' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="create_at" label="提交时间" />
        <el-table-column label="操作">
          <template #default="scope">
            <el-button size="small" @click="showCode(scope.row)">查看代码</el-button>
            <el-button size="small" @click="showExecDetail(scope.row)">评测详情</el-button>
            <el-button v-if="!scope.row.accepted" size="small" @click="handleCodeCheck(scope.row)"
              >AI分析</el-button
            >
          </template>
        </el-table-column>
      </el-table>

      <div class="mt-4 flex justify-end">
        <el-pagination
          v-model:current-page="currentPage"
          v-model:page-size="pageSize"
          :total="total"
          @current-change="handleCurrentChange"
        />
      </div>
    </div>

    <el-dialog v-model="codeDialogVisible" title="代码" width="50%">
      <pre class="rounded bg-[#f0f2f5] p-4 text-[#303133]"><code>{{ currentCode }}</code></pre>
      <div class="mt-4 flex justify-end">
        <el-button @click="copyCode">复制代码</el-button>
      </div>
    </el-dialog>
    <el-dialog v-model="codeCheckDialogVisible" title="代码分析" width="50%">
      <pre class="rounded bg-[#f0f2f5] p-4 text-[#303133]"><code>{{ currentCode }}</code></pre>
      <div class="mt-4">
        <Bubble
          :content="codeCheckRef"
          :loading="codeCheckLoading"
          :typing="true"
          :is-markdown="true"
          maxWidth="full"
        />
      </div>
    </el-dialog>
    <JudgeResult ref="judgeResultRef" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import JudgeResult from '@/components/JudgeResult.vue'
import api from '@/api'
import { Bubble } from 'vue-element-plus-x'
const judgeResultRef = ref()
const codeCheckRef = ref('')
const codeCheckLoading = ref(false)
const judgeLanguageMap: Record<string, string> = {
  c: 'C',
  cpp: 'C++',
  java: 'Java',
  python: 'Python',
  golang: 'Go',
  rust: 'Rust',
}

interface JudgeItem {
  judge_id: number
  user_id: number
  question_id: number
  language: string
  code: string
  exec_result: string
  accepted: boolean
  create_at: string
  update_at: string
}

interface JudgeListResponse {
  judge_list: JudgeItem[]
  total: number
  page: number
  page_size: number
}

const judgeList = ref<JudgeItem[]>([])
const total = ref(0)

const currentPage = ref(1)
const pageSize = ref(10)

// 对话框控制
const codeDialogVisible = ref(false)
const currentCode = ref('')
const questionId = ref(0)
const codeCheckDialogVisible = ref(false)
// 加载评测记录
const loadJudgeList = () => {
  api.judge
    .getJudgeList(currentPage.value, pageSize.value)
    .then((res) => {
      const data = res.data as JudgeListResponse
      judgeList.value = data.judge_list
      total.value = data.total
    })
    .catch((error) => {
      ElMessage.error('获取评测记录失败！')
      console.error(error)
    })
}

const handleCurrentChange = (val: number) => {
  currentPage.value = val
  loadJudgeList()
}

// 查看代码
const showCode = (row: JudgeItem) => {
  currentCode.value = row.code
  codeDialogVisible.value = true
}

// 查看评测详情
const showExecDetail = (row: JudgeItem) => {
  try {
    judgeResultRef.value.open(JSON.parse(row.exec_result))
  } catch (error) {
    ElMessage.error('查看评测详情失败！')
    console.error(error)
  }
}

// 复制代码
const copyCode = () => {
  navigator.clipboard.writeText(currentCode.value)
  ElMessage.success('代码已复制到剪贴板！')
}

onMounted(() => {
  document.title = 'OnlineJudge - 评测记录'
  loadJudgeList()
})

const handleCodeCheck = (row: JudgeItem) => {
  questionId.value = row.question_id
  currentCode.value = row.code
  codeCheckDialogVisible.value = true
  codeCheckLoading.value = true
  api.ai
    .codeCheck(questionId.value, currentCode.value)
    .then((res) => {
      codeCheckRef.value = res.data
    })
    .finally(() => {
      codeCheckLoading.value = false
    })
}
</script>
