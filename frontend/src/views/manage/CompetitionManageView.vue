<template>
  <div class="w-full p-4">
    <div class="mb-4 flex gap-4">
      <el-button type="primary" @click="addCompetitionDrawer = true">创建比赛</el-button>
    </div>
    <el-card>
      <!-- 搜索区域 -->
      <div class="mb-6 flex gap-4">
        <el-input
          v-model="filter.name"
          placeholder="比赛名称"
          class="w-64"
          clearable
          @clear="search"
        >
          <template #prefix>
            <el-icon><Search /></el-icon>
          </template>
        </el-input>

        <el-select
          v-model="filter.type"
          placeholder="比赛类型"
          clearable
          class="w-48"
          @clear="search"
          @change="search"
        >
          <el-option label="OI赛制" value="oi" />
          <el-option label="ACM赛制" value="acm" />
        </el-select>

        <el-button type="primary" @click="search">搜索</el-button>
      </div>

      <el-table :data="tableData" stripe border class="w-full" v-loading="loading">
        <el-table-column prop="id" label="ID" />
        <el-table-column prop="name" label="比赛名称" />
        <el-table-column prop="start_time" label="开始时间" />
        <el-table-column prop="end_time" label="结束时间" />
        <el-table-column prop="competition_type" label="比赛类型">
          <template #default="scope">
            <el-tag
              v-if="scope.row.competition_type === 'oi'"
              type="info"
              @click="handleTypeFilter('oi')"
              class="cursor-pointer"
              >OI赛制</el-tag
            >
            <el-tag v-else type="info" @click="handleTypeFilter('acm')" class="cursor-pointer"
              >ACM赛制</el-tag
            >
          </template>
        </el-table-column>
        <el-table-column prop="password_required" label="是否需要密码">
          <template #default="scope">
            <span>{{ scope.row.password_required ? '是' : '否' }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template #default="scope">
            <div class="flex justify-start gap-3">
              <el-link
                type="primary"
                plain
                :underline="false"
                @click="handleCompetitionDetail(scope.row)"
                >详情</el-link
              >
              <el-link
                type="danger"
                plain
                :underline="false"
                @click="handleDeleteCompetition(scope.row.id)"
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

    <el-drawer v-model="addCompetitionDrawer" title="创建比赛" size="50%">
      <el-form label-width="100px">
        <el-form-item label="比赛名称" prop="name">
          <el-input v-model="newCompetitionInfo.name" placeholder="请输入比赛名称" />
        </el-form-item>
        <el-form-item label="比赛描述" prop="description">
          <el-input
            v-model="newCompetitionInfo.description"
            type="textarea"
            placeholder="请输入比赛描述"
          />
        </el-form-item>
        <el-form-item label="开始时间" prop="start_time">
          <el-date-picker
            v-model="newCompetitionInfo.start_time"
            type="datetime"
            placeholder="请选择开始时间"
            value-format="YYYY-MM-DD HH:mm:ss"
            format="YYYY-MM-DD HH:mm:ss"
          />
        </el-form-item>
        <el-form-item label="结束时间" prop="end_time">
          <el-date-picker
            v-model="newCompetitionInfo.end_time"
            type="datetime"
            placeholder="请选择结束时间"
            value-format="YYYY-MM-DD HH:mm:ss"
            format="YYYY-MM-DD HH:mm:ss"
          />
        </el-form-item>
        <el-form-item label="比赛类型" prop="competition_type">
          <el-select v-model="newCompetitionInfo.competition_type" placeholder="请选择比赛类型">
            <el-option label="OI赛制" value="oi" />
            <el-option label="ACM赛制" value="acm" />
          </el-select>
        </el-form-item>
        <el-form-item label="比赛题目" prop="questions">
          <div class="flex flex-col gap-2">
            <div v-for="(question, index) in newCompetitionInfo.questions" :key="index">
              <span class="text-gray-500">题目{{ index + 1 }}</span>
              <el-input v-model="question.qid" placeholder="请输入题号" class="mb-2" />
              <el-input
                v-model.number="question.id"
                placeholder="请输入题目ID（题库中题目的ID）"
                class="mb-2"
                type="number"
              />
              <el-button type="danger" @click="removeCompetitionQuestion(index)" class="mb-2"
                >删除</el-button
              >
            </div>
            <el-button type="primary" @click="addCompetitionQuestion">添加比赛题目</el-button>
          </div>
        </el-form-item>
        <el-form-item label="密码开关" prop="password_required">
          <el-switch
            v-model="newCompetitionInfo.password_required"
            @change="handlePasswordRequiredChange"
          />
        </el-form-item>
        <el-form-item label="密码" prop="password" v-show="newCompetitionInfo.password_required">
          <el-input
            v-model="newCompetitionInfo.password"
            type="password"
            placeholder="请输入密码"
            readonly
            @focus="(e: any) => e.target.removeAttribute('readonly')"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleAddCompetition">创建比赛</el-button>
        </el-form-item>
      </el-form>
    </el-drawer>

    <el-drawer v-model="competitionDetailDrawer" title="比赛详情" size="50%">
      <el-form label-width="100px">
        <el-form-item label="比赛ID" prop="id">
          <span>{{ currentCompetition.id }}</span>
        </el-form-item>
        <el-form-item label="比赛名称" prop="name">
          <el-input v-model="currentCompetition.name" placeholder="请输入比赛名称" />
        </el-form-item>
        <el-form-item label="比赛描述" prop="description">
          <el-input
            v-model="currentCompetition.description"
            type="textarea"
            placeholder="请输入比赛描述"
          />
        </el-form-item>
        <el-form-item label="开始时间" prop="start_time">
          <el-date-picker
            v-model="currentCompetition.start_time"
            type="datetime"
            placeholder="请选择开始时间"
            value-format="YYYY-MM-DD HH:mm:ss"
            format="YYYY-MM-DD HH:mm:ss"
          />
        </el-form-item>
        <el-form-item label="结束时间" prop="end_time">
          <el-date-picker
            v-model="currentCompetition.end_time"
            type="datetime"
            placeholder="请选择结束时间"
            value-format="YYYY-MM-DD HH:mm:ss"
            format="YYYY-MM-DD HH:mm:ss"
          />
        </el-form-item>
        <el-form-item label="比赛类型" prop="competition_type">
          <el-select v-model="currentCompetition.competition_type" placeholder="请选择比赛类型">
            <el-option label="OI赛制" value="oi" />
            <el-option label="ACM赛制" value="acm" />
          </el-select>
        </el-form-item>
        <el-form-item label="比赛题目" prop="questions">
          <div class="flex flex-col gap-2">
            <div v-for="(question, index) in currentCompetition.questions" :key="index">
              <span class="text-gray-500">题目{{ index + 1 }}</span>
              <el-input v-model="question.qid" placeholder="请输入题号" class="mb-2" />
              <el-input
                v-model.number="question.id"
                placeholder="请输入题目ID（题库中题目的ID）"
                class="mb-2"
                type="number"
              />
              <el-button type="danger" @click="removeCompetitionQuestion(index)" class="mb-2"
                >删除</el-button
              >
            </div>
            <el-button type="primary" @click="addCompetitionQuestion">添加比赛题目</el-button>
          </div>
        </el-form-item>
        <el-form-item label="密码开关" prop="password_required">
          <el-switch
            v-model="currentCompetition.password_required"
            @change="handlePasswordRequiredChange"
          />
        </el-form-item>
        <el-form-item label="密码" prop="password" v-show="currentCompetition.password_required">
          <el-input
            v-model="currentCompetition.password"
            type="password"
            placeholder="请输入密码（若之前设置过密码，将覆盖之前的密码）"
            readonly
            @focus="(e: any) => e.target.removeAttribute('readonly')"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleUpdateCompetition">更新比赛信息</el-button>
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>
<script setup lang="ts">
import { ref, onMounted } from 'vue'
import api from '@/api'
import { ElMessage, ElMessageBox } from 'element-plus'

const loading = ref(false)
const addCompetitionDrawer = ref(false)
const competitionDetailDrawer = ref(false)
interface CompetitionQuestion {
  qid: string
  id: number
}

interface Competition {
  id: number
  name: string
  description: string
  start_time: string
  end_time: string
  questions: CompetitionQuestion[]
  competition_type: string
  password_required: boolean
  password: string
}

interface NewCompetitionInfo {
  name: string
  description: string
  start_time: string
  end_time: string
  questions: CompetitionQuestion[]
  competition_type: string
  password_required: boolean
  password: string
}

const newCompetitionInfo = ref<NewCompetitionInfo>({
  name: '',
  description: '',
  start_time: '',
  end_time: '',
  questions: [],
  competition_type: '',
  password_required: false,
  password: '',
})

const currentCompetition = ref<Competition>({
  id: NaN,
  name: '',
  description: '',
  start_time: '',
  end_time: '',
  questions: [],
  competition_type: '',
  password_required: false,
  password: '',
})

const filter = ref({
  name: '',
  type: '',
})

const tableData = ref<Competition[]>([])
const total = ref(0)
const page = ref(1)
const pageSize = ref(10)

const search = () => {
  loading.value = true
  api.competition
    .adminGetCompetitionList(page.value, pageSize.value, filter.value)
    .then((res) => {
      tableData.value = res.data.competition_list
      total.value = res.data.total
    })
    .catch(() => {
      ElMessage.error('获取比赛列表失败')
    })
    .finally(() => {
      loading.value = false
    })
}

onMounted(() => {
  document.title = 'OnlineJudge - 比赛管理'
  search()
})

const handleTypeFilter = (type: string) => {
  filter.value.type = type
  search()
}

const addCompetitionQuestion = () => {
  if (addCompetitionDrawer.value) {
    newCompetitionInfo.value.questions.push({
      qid: '',
      id: NaN,
    })
  } else if (competitionDetailDrawer.value) {
    currentCompetition.value.questions.push({
      qid: '',
      id: NaN,
    })
  }
}

const removeCompetitionQuestion = (index: number) => {
  if (addCompetitionDrawer.value) {
    newCompetitionInfo.value.questions.splice(index, 1)
  } else if (competitionDetailDrawer.value) {
    currentCompetition.value.questions.splice(index, 1)
  }
}

const handleAddCompetition = async () => {
  try {
    await api.competition.addCompetition(newCompetitionInfo.value)
    newCompetitionInfo.value = {
      name: '',
      description: '',
      start_time: '',
      end_time: '',
      questions: [],
      competition_type: '',
      password_required: false,
      password: '',
    }
    addCompetitionDrawer.value = false
    ElMessage.success('创建成功')
    search()
  } catch {
    ElMessage.error('创建失败')
  }
}

const handleCompetitionDetail = (row: Competition) => {
  currentCompetition.value = JSON.parse(
    JSON.stringify({
      id: row.id,
      name: row.name,
      description: row.description,
      start_time: row.start_time,
      end_time: row.end_time,
      questions: row.questions,
      competition_type: row.competition_type,
      password_required: row.password_required,
      password: '',
    }),
  )
  competitionDetailDrawer.value = true
}

const handleDeleteCompetition = async (id: number) => {
  try {
    await ElMessageBox.confirm('确定删除该比赛吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    })
    await api.competition.deleteCompetition(id)
    ElMessage.success('删除成功')
    search()
  } catch {
    ElMessage.error('删除失败')
  }
}

const handlePasswordRequiredChange = () => {
  if (!newCompetitionInfo.value.password_required) {
    newCompetitionInfo.value.password = ''
  }
  if (!currentCompetition.value.password_required) {
    currentCompetition.value.password = ''
  }
}

const handleUpdateCompetition = async () => {
  try {
    await api.competition.updateCompetition({
      id: currentCompetition.value.id,
      name: currentCompetition.value.name,
      description: currentCompetition.value.description,
      start_time: currentCompetition.value.start_time,
      end_time: currentCompetition.value.end_time,
      questions: currentCompetition.value.questions,
      competition_type: currentCompetition.value.competition_type,
      password_required: currentCompetition.value.password_required,
      password: currentCompetition.value.password,
    })
    competitionDetailDrawer.value = false
    ElMessage.success('更新成功')
    search()
  } catch {
    ElMessage.error('更新失败')
  }
}
</script>
