<template>
  <div class="p-4">
    <!-- 搜索区域 -->
    <div class="mb-6 flex gap-4">
      <el-input
        v-model="searchTitle"
        placeholder="题目标题"
        class="w-64"
        clearable
        @clear="handleSearch"
      >
        <template #prefix>
          <el-icon><search /></el-icon>
        </template>
      </el-input>

      <el-input
        v-model="searchTag"
        placeholder="题目标签"
        clearable
        class="w-48"
        @clear="handleSearch"
      >
        <template #prefix>
          <el-icon><search /></el-icon>
        </template>
      </el-input>

      <el-input
        v-model="searchDifficulty"
        placeholder="题目难度"
        clearable
        class="w-48"
        @clear="handleSearch"
      >
        <template #prefix>
          <el-icon><search /></el-icon>
        </template>
      </el-input>

      <el-button type="primary" @click="handleSearch">搜索</el-button>
    </div>

    <!-- 题目列表 -->
    <el-table :data="questions" table-layout="fixed" style="width: 100%" stripe class="mb-4">
      <el-table-column prop="id" label="ID" />
      <el-table-column prop="title" label="标题" />
      <el-table-column label="标签">
        <template #default="scope">
          <div class="flex flex-wrap gap-1">
            <el-tag
              v-for="tag in scope.row.tags"
              :key="tag"
              class="cursor-pointer"
              @click="quickFilterByTag(tag)"
              type="info"
              effect="plain"
            >
              {{ tag }}
            </el-tag>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="通过率">
        <template #default="scope">
          {{ scope.row.AcRate }}
        </template>
      </el-table-column>
      <el-table-column label="难度">
        <template #default="scope">
          <el-tag :type="DifficultyToType(scope.row.Difficulty)" effect="light">
            {{ getDifficulty(scope.row.Difficulty) }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template #default="scope">
          <el-button type="primary" plain @click="handleClickQuestion(scope.row.id)"
            >答题</el-button
          >
        </template>
      </el-table-column>
    </el-table>

    <!-- 分页 -->
    <div class="flex justify-end">
      <el-pagination
        v-model:current-page="currentPage"
        v-model:page-size="pageSize"
        :page-sizes="[10, 20, 30, 50]"
        layout="total, sizes, prev, pager, next, jumper"
        :total="total"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { Search } from '@element-plus/icons-vue'
import { question } from '../api/modules/question'
import { ElMessage } from 'element-plus'
// 状态
const questions = ref([])
const total = ref(0)
const currentPage = ref(1)
const pageSize = ref(10)
const searchTitle = ref('')
const searchTag = ref('')
const searchDifficulty = ref('')

// 根据难度返回对应的标签类型
const DifficultyToType = (difficulty: string) => {
  switch (difficulty) {
    case 'easy':
      return 'success'
    case 'medium':
      return 'warning'
    case 'hard':
      return 'danger'
    default:
      return 'info'
  }
}

const getDifficulty = (difficulty: string) => {
  switch (difficulty) {
    case 'easy':
      return '简单'
    case 'medium':
      return '中等'
    case 'hard':
      return '困难'
    default:
      return '未知'
  }
}

// 获取题目数据
const getQuestions = async () => {
  try {
    const filter = {
      title: searchTitle.value,
      tag: searchTag.value,
      difficulty: searchDifficulty.value,
    }
    const response = await question.pagination(currentPage.value, pageSize.value, filter)
    const data = response.data
    questions.value = data.questions
    total.value = data.total
  } catch {
    ElMessage.error('获取题目列表失败')
  }
}

// 搜索操作
const handleSearch = () => {
  currentPage.value = 1
  getQuestions()
}

// 快速按标签筛选
const quickFilterByTag = (tag: string) => {
  searchTag.value = tag
  handleSearch()
}

// 页码变化
const handleCurrentChange = (page: number) => {
  currentPage.value = page
  getQuestions()
}

// 每页数量变化
const handleSizeChange = (size: number) => {
  pageSize.value = size
  currentPage.value = 1
  getQuestions()
}

// 点击题目
const handleClickQuestion = (id: number) => {
  window.open(`/question/${id}`, '_blank')
}

// 初始化
onMounted(() => {
  // 设置浏览器标签页标题
  document.title = 'OnlineJudge - 题库'
  getQuestions()
})
</script>
