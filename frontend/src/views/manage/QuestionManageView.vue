<template>
  <div class="w-full p-4">
    <div class="mb-4 flex gap-4">
      <el-button type="primary" @click="addQuestionDrawer = true">创建题目</el-button>
    </div>
    <el-card>
      <!-- 搜索区域 -->
      <div class="mb-6 flex gap-4">
        <el-input
          v-model="filter.title"
          placeholder="题目标题"
          class="w-64"
          clearable
          @clear="search"
        >
          <template #prefix>
            <el-icon><Search /></el-icon>
          </template>
        </el-input>

        <el-input
          v-model="filter.tag"
          placeholder="题目标签"
          clearable
          class="w-48"
          @clear="search"
        >
          <template #prefix>
            <el-icon><Search /></el-icon>
          </template>
        </el-input>

        <el-select
          v-model="filter.difficulty"
          placeholder="题目难度"
          clearable
          class="w-48"
          @clear="search"
          @change="search"
        >
          <el-option label="简单" value="easy" />
          <el-option label="中等" value="medium" />
          <el-option label="困难" value="hard" />
        </el-select>

        <el-button type="primary" @click="search">搜索</el-button>
      </div>

      <el-table :data="tableData" stripe border class="w-full" v-loading="loading">
        <el-table-column prop="id" label="ID" />
        <el-table-column prop="title" label="标题" />
        <el-table-column prop="tags" label="标签">
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
            {{ calculatePassRate(scope.row.accepted_num, scope.row.submit_num) }}
          </template>
        </el-table-column>
        <el-table-column prop="difficulty" label="难度">
          <template #default="scope">
            <el-tag
              :type="DifficultyToType(scope.row.difficulty)"
              class="cursor-pointer"
              effect="light"
              @click="quickFilterByDifficulty(scope.row.difficulty)"
            >
              {{ getDifficulty(scope.row.difficulty) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template #default="scope">
            <div class="flex justify-start gap-3">
              <el-link
                type="primary"
                plain
                @click="handleQuestionDetail(scope.row)"
                :underline="false"
                >详情</el-link
              >
              <el-link
                type="danger"
                plain
                @click="handleDeleteQuestion(scope.row.id)"
                :underline="false"
                >删除</el-link
              >
            </div>
          </template>
        </el-table-column>
      </el-table>

      <el-pagination
        v-model:current-page="page"
        v-model:page-size="pageSize"
        :total="total"
        class="mt-4"
        @size-change="search"
        @current-change="search"
      />
    </el-card>

    <el-drawer v-model="addQuestionDrawer" title="创建题目" size="100%">
      <el-form label-width="100px" label-position="top" @submit.prevent>
        <el-form-item label="题目标题">
          <el-input
            v-model="newQuestionInfo.title"
            class="max-w-1/2"
            placeholder="请输入题目标题"
          />
        </el-form-item>
        <el-form-item label="题目内容">
          <MdEditor v-model="newQuestionInfo.content" previewTheme="github" codeTheme="github" />
        </el-form-item>
        <el-form-item label="测试用例">
          <div class="flex flex-col gap-2">
            <div v-for="(caseItem, index) in newQuestionInfo.judge_case" :key="index">
              <span class="text-gray-500">测试用例{{ index + 1 }}</span>
              <el-input v-model="caseItem.input" placeholder="输入" type="textarea" class="mb-2" />
              <el-input v-model="caseItem.output" placeholder="输出" type="textarea" class="mb-2" />
              <el-button type="danger" @click="removeTestCase(index)" class="mb-2">删除</el-button>
            </div>
            <el-button type="primary" @click="addTestCase">添加测试用例</el-button>
          </div>
        </el-form-item>
        <el-form-item label="题目标签">
          <el-input-tag
            v-model="newQuestionInfo.tags"
            placeholder="按Enter回车键添加输入内容为标签"
            class="max-w-1/2"
          />
        </el-form-item>
        <el-form-item label="题目难度">
          <el-select
            v-model="newQuestionInfo.difficulty"
            placeholder="请选择难度"
            class="max-w-1/8"
          >
            <el-option label="简单" value="easy" />
            <el-option label="中等" value="medium" />
            <el-option label="困难" value="hard" />
          </el-select>
        </el-form-item>
        <el-form-item label="题目可见范围">
          <el-select
            v-model="newQuestionInfo.visible_scope"
            placeholder="请选择可见范围"
            class="max-w-1/8"
          >
            <el-option label="公开" value="public" />
            <el-option label="仅比赛时可见" value="competition_only" />
          </el-select>
        </el-form-item>
        <el-form-item label="超时时间（ms）">
          <el-input
            v-model.number="newQuestionInfo.judge_config.timeout_millisecond"
            type="number"
            class="max-w-1/8"
            placeholder="请输入超时时间（ms）"
          />
        </el-form-item>
        <el-form-item label="内存限制（MB）">
          <el-input
            v-model.number="newQuestionInfo.judge_config.memory_limit_mib"
            type="number"
            class="max-w-1/8"
            placeholder="请输入内存限制（MB）"
          />
        </el-form-item>
        <el-form-item label="栈限制（MB）">
          <el-input
            v-model.number="newQuestionInfo.judge_config.stack_limit_mib"
            type="number"
            class="max-w-1/8"
            placeholder="请输入栈限制（MB）"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleCreateQuestion">创建题目</el-button>
        </el-form-item>
      </el-form>
    </el-drawer>

    <el-drawer v-model="questionDetailDrawer" title="题目详情" size="100%">
      <el-form label-width="100px" label-position="top" @submit.prevent>
        <el-form-item label="题目ID">
          <span>{{ currentQuestionInfo.id }}</span>
        </el-form-item>
        <el-form-item label="提交次数">
          <span>{{ currentQuestionInfo.submit_num }}</span>
        </el-form-item>
        <el-form-item label="通过次数">
          <span>{{ currentQuestionInfo.accepted_num }}</span>
        </el-form-item>
        <el-form-item label="题目标题">
          <el-input
            v-model="currentQuestionInfo.title"
            class="max-w-1/2"
            placeholder="请输入题目标题"
          />
        </el-form-item>
        <el-form-item label="题目内容">
          <MdEditor
            v-model="currentQuestionInfo.content"
            previewTheme="github"
            codeTheme="github"
          />
        </el-form-item>
        <el-form-item label="测试用例">
          <div class="flex flex-col gap-2">
            <div v-for="(caseItem, index) in currentQuestionInfo.judge_case" :key="index">
              <span class="text-gray-500">测试用例{{ index + 1 }}</span>
              <el-input v-model="caseItem.input" placeholder="输入" type="textarea" class="mb-2" />
              <el-input v-model="caseItem.output" placeholder="输出" type="textarea" class="mb-2" />
              <el-button type="danger" @click="removeTestCase(index)" class="mb-2">删除</el-button>
            </div>
            <el-button type="primary" @click="addTestCase">添加测试用例</el-button>
          </div>
        </el-form-item>
        <el-form-item label="题目标签">
          <el-input-tag
            v-model="currentQuestionInfo.tags"
            placeholder="按Enter回车键添加输入内容为标签"
            class="max-w-1/2"
          />
        </el-form-item>
        <el-form-item label="题目难度">
          <el-select
            v-model="currentQuestionInfo.difficulty"
            placeholder="请选择难度"
            class="max-w-1/8"
          >
            <el-option label="简单" value="easy" />
            <el-option label="中等" value="medium" />
            <el-option label="困难" value="hard" />
          </el-select>
        </el-form-item>
        <el-form-item label="题目可见范围">
          <el-select
            v-model="currentQuestionInfo.visible_scope"
            placeholder="请选择可见范围"
            class="max-w-1/8"
          >
            <el-option label="公开" value="public" />
            <el-option label="仅比赛时可见" value="competition_only" />
          </el-select>
        </el-form-item>
        <el-form-item label="超时时间（ms）">
          <el-input
            v-model.number="currentQuestionInfo.judge_config.timeout_millisecond"
            type="number"
            class="max-w-1/8"
            placeholder="请输入超时时间（ms）"
          />
        </el-form-item>
        <el-form-item label="内存限制（MB）">
          <el-input
            v-model.number="currentQuestionInfo.judge_config.memory_limit_mib"
            type="number"
            class="max-w-1/8"
            placeholder="请输入内存限制（MB）"
          />
        </el-form-item>
        <el-form-item label="栈限制（MB）">
          <el-input
            v-model.number="currentQuestionInfo.judge_config.stack_limit_mib"
            type="number"
            class="max-w-1/8"
            placeholder="请输入栈限制（MB）"
          />
        </el-form-item>
        <el-form-item label="创建时间">
          <span>{{ currentQuestionInfo.create_at }}</span>
        </el-form-item>
        <el-form-item label="更新时间">
          <span>{{ currentQuestionInfo.update_at }}</span>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleUpdateQuestion">更新题目信息</el-button>
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup lang="ts">
import api from '@/api'
import { ref, onMounted } from 'vue'
import { Search } from '@element-plus/icons-vue'
import { MdEditor } from 'md-editor-v3'
import { questionTemplate } from '@/config/questionTemplateConfig'
import 'md-editor-v3/lib/style.css'
import { ElMessage, ElMessageBox } from 'element-plus'

const tableData = ref<Question[]>([])
const total = ref(0)
const page = ref(1)
const pageSize = ref(10)
const addQuestionDrawer = ref(false)
const questionDetailDrawer = ref(false)
const loading = ref(false)

interface JudgeCase {
  input: string
  output: string
}

interface JudgeConfig {
  memory_limit_mib: number
  stack_limit_mib: number
  timeout_millisecond: number
}

interface Question {
  id: number
  title: string
  content: string
  tags: string[]
  difficulty: string
  submit_num: number
  accepted_num: number
  judge_case: JudgeCase[]
  judge_config: JudgeConfig
  visible_scope: string
  create_at: string
  update_at: string
}

interface NewQuestionInfo {
  title: string
  content: string
  tags: string[]
  difficulty: string
  judge_case: JudgeCase[]
  judge_config: JudgeConfig
  visible_scope: string
}

const newQuestionInfo = ref<NewQuestionInfo>({
  title: '',
  content: questionTemplate,
  tags: [],
  difficulty: '',
  judge_case: [],
  judge_config: {
    memory_limit_mib: NaN,
    stack_limit_mib: NaN,
    timeout_millisecond: NaN,
  },
  visible_scope: '',
})

const currentQuestionInfo = ref<Question>({
  id: NaN,
  title: '',
  content: '',
  tags: [],
  difficulty: '',
  submit_num: NaN,
  accepted_num: NaN,
  judge_case: [],
  judge_config: {
    memory_limit_mib: NaN,
    stack_limit_mib: NaN,
    timeout_millisecond: NaN,
  },
  visible_scope: '',
  create_at: '',
  update_at: '',
})

const filter = ref({
  title: '',
  tag: '',
  difficulty: '',
})

onMounted(() => {
  document.title = 'OnlineJudge - 题目管理'
  search()
})

const search = () => {
  loading.value = true
  api.question
    .getQuestionList(page.value, pageSize.value, filter.value)
    .then((res) => {
      tableData.value = res.data.questions
      total.value = res.data.total
    })
    .catch(() => {
      ElMessage.error('获取题目列表失败')
    })
    .finally(() => {
      loading.value = false
    })
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

const quickFilterByTag = (tag: string) => {
  filter.value.tag = tag
  search()
}

const quickFilterByDifficulty = (difficulty: string) => {
  filter.value.difficulty = difficulty
  search()
}

const addTestCase = () => {
  newQuestionInfo.value.judge_case.push({ input: '', output: '' })
}

const removeTestCase = (index: number) => {
  newQuestionInfo.value.judge_case.splice(index, 1)
}

const handleCreateQuestion = async () => {
  try {
    await api.question.addQuestion(newQuestionInfo.value)
    newQuestionInfo.value = {
      title: '',
      content: questionTemplate,
      tags: [],
      difficulty: '',
      judge_case: [],
      judge_config: {
        memory_limit_mib: NaN,
        stack_limit_mib: NaN,
        timeout_millisecond: NaN,
      },
      visible_scope: '',
    }
    addQuestionDrawer.value = false
    ElMessage.success('创建成功')
    search()
  } catch {
    ElMessage.error('创建失败')
  }
}

const handleQuestionDetail = (row: Question) => {
  currentQuestionInfo.value = JSON.parse(
    JSON.stringify({
      id: row.id,
      title: row.title,
      content: row.content,
      tags: row.tags,
      difficulty: row.difficulty,
      submit_num: row.submit_num,
      accepted_num: row.accepted_num,
      judge_case: row.judge_case,
      judge_config: row.judge_config,
      visible_scope: row.visible_scope,
      create_at: row.create_at,
      update_at: row.update_at,
    }),
  )
  questionDetailDrawer.value = true
}

const handleUpdateQuestion = async () => {
  try {
    await api.question.updateQuestion({
      id: currentQuestionInfo.value.id,
      title: currentQuestionInfo.value.title,
      content: currentQuestionInfo.value.content,
      tags: currentQuestionInfo.value.tags,
      difficulty: currentQuestionInfo.value.difficulty,
      judge_case: currentQuestionInfo.value.judge_case,
      judge_config: currentQuestionInfo.value.judge_config,
      visible_scope: currentQuestionInfo.value.visible_scope,
    })
    questionDetailDrawer.value = false
    ElMessage.success('更新成功')
    search()
  } catch {
    ElMessage.error('更新失败')
  }
}

const handleDeleteQuestion = async (id: number) => {
  await ElMessageBox.confirm('确定删除该题目吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  })
  await api.question.deleteQuestion(id)
  ElMessage.success('删除成功')
  search()
}

// 计算通过率（保留两位小数）
const calculatePassRate = (acceptedNum: number, submitNum: number) => {
  if (submitNum === 0) {
    return '0.00%'
  }
  return ((acceptedNum / submitNum) * 100).toFixed(2) + '%'
}
</script>
