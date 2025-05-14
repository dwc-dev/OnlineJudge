<template>
  <div class="px-8 py-4">
    <div class="mb-6 flex gap-4">
      <el-input
        v-model="searchName"
        placeholder="比赛名称"
        class="w-64"
        clearable
        @clear="handleSearch"
      >
        <template #prefix>
          <el-icon><Search /></el-icon>
        </template>
      </el-input>

      <el-select
        v-model="searchType"
        placeholder="比赛类型"
        clearable
        class="w-48"
        @clear="handleSearch"
        @change="handleSearch"
      >
        <el-option label="OI赛制" value="oi" />
        <el-option label="ACM赛制" value="acm" />
      </el-select>

      <el-button type="primary" @click="handleSearch">搜索</el-button>
    </div>
    <el-table :data="competitions" table-layout="fixed" style="width: 100%" stripe class="mb-4">
      <el-table-column prop="name" label="比赛名称">
        <template #default="scope">
          <div class="flex items-center">
            <el-link :underline="false" target="_blank" :href="`/competition/${scope.row.id}`">
              {{ scope.row.name }}
            </el-link>
            <el-icon :size="15" class="mr-2" v-if="scope.row.password_required">
              <Lock />
            </el-icon>
          </div>
        </template>
      </el-table-column>
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
      <el-table-column prop="status" label="操作">
        <template #default="scope">
          <el-button size="small" plain @click="handleJoin(scope.row)">报名</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
  <el-dialog v-model="dialogVisible" title="密码" width="30%">
    <el-input v-model="password" placeholder="请输入比赛密码" />
    <template #footer>
      <el-button type="primary" @click="handleDialogJoin">报名</el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { Lock } from '@element-plus/icons-vue'
import api from '@/api'
import { useUserStore } from '@/stores/user'
import { ElMessage } from 'element-plus'
const userStore = useUserStore()
const competitions = ref([])
const total = ref(0)
const currentPage = ref(1)
const pageSize = ref(10)
const searchName = ref('')
const searchType = ref('')
const dialogVisible = ref(false)
const password = ref('')
const currentCompetition = ref(0)
const handleSearch = () => {
  getCompetitions()
}

const handleTypeFilter = (type: string) => {
  searchType.value = type
  handleSearch()
}

const getCompetitions = () => {
  api.competition
    .getCompetitionList(currentPage.value, pageSize.value, {
      name: searchName.value,
      type: searchType.value,
    })
    .then((res) => {
      competitions.value = res.data.competition_list
      total.value = res.data.total
    })
}

onMounted(() => {
  document.title = 'OnlineJudge - 比赛'
  getCompetitions()
})

const handleJoin = async (row: { password_required: boolean; id: number }) => {
  if (row.password_required) {
    currentCompetition.value = row.id
    dialogVisible.value = true
  } else {
    try {
      await api.competition.joinCompetition(row.id, userStore.uid, '')
      ElMessage.success('报名成功')
    } catch (error: unknown) {
      ElMessage.error('报名失败，错误信息：' + (error as Error).message)
    }
  }
}

const handleDialogJoin = async () => {
  try {
    await api.competition.joinCompetition(currentCompetition.value, userStore.uid, password.value)
    dialogVisible.value = false
    password.value = ''
    ElMessage.success('报名成功')
  } catch (error: unknown) {
    dialogVisible.value = false
    password.value = ''
    ElMessage.error('报名失败，错误信息：' + (error as Error).message)
  }
}
</script>
