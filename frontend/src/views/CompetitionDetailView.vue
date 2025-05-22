<template>
  <div class="h-full" v-loading="loading" v-if="loading"></div>
  <div
    v-if="access_no"
    class="flex h-1/3 flex-col items-center justify-center rounded-lg text-[#F56C6C]"
  >
    <el-icon size="100">
      <Warning />
    </el-icon>
    <span class="mt-4 text-2xl">无访问权限，请先报名</span>
  </div>
  <div v-else-if="access_yes" class="h-full w-full px-4 py-4">
    <el-card class="mb-4" shadow="never">
      <template #header>
        <div>
          <span class="font-bold">{{ compInfo.name }}</span>
        </div>
      </template>
      <span>{{ compInfo.description }}</span>
    </el-card>
    <div class="grid grid-cols-3 gap-4">
      <div class="col-span-2">
        <el-tabs @tab-change="handleTabChange">
          <el-tab-pane label="题目集">
            <el-card class="mb-4" shadow="never">
              <template #header>
                <div>
                  <span class="font-bold">题目集</span>
                </div>
              </template>
              <div v-if="!questionOpen">
                <span>题目集将在比赛开始后开放</span>
              </div>
              <div v-else>
                <el-table :data="questionList" style="width: 100%">
                  <el-table-column prop="qid" label="题号" />
                  <el-table-column prop="title" label="题目名称">
                    <template #default="{ row }">
                      <el-link
                        :href="`/competition/${id}/question/${row.qid}`"
                        target="_blank"
                        :underline="false"
                        >{{ row.title }}</el-link
                      >
                    </template>
                  </el-table-column>
                </el-table>
              </div>
            </el-card>
          </el-tab-pane>
          <el-tab-pane label="排行榜" name="rank">
            <el-card shadow="never">
              <template #header>
                <div>
                  <span class="font-bold">排行榜</span>
                </div>
              </template>
              <el-table :data="rankTableData" style="width: 100%" :row-class-name="rowClassName">
                <el-table-column prop="username" label="用户名" align="center">
                  <template #default="{ row }">
                    <el-link :href="`/space/${row.user_id}`" target="_blank" :underline="false">{{
                      row.username
                    }}</el-link>
                  </template>
                </el-table-column>
                <el-table-column label="各题得分" align="center">
                  <el-table-column
                    v-for="qid in allQids"
                    :key="qid"
                    :label="qid"
                    :prop="qid"
                    align="center"
                  />
                </el-table-column>
                <el-table-column prop="total_score" label="总得分" align="center" />
              </el-table>
            </el-card>
          </el-tab-pane>
        </el-tabs>
      </div>
      <div class="col-span-1">
        <el-card class="mb-4" shadow="never">
          <template #header>
            <div>
              <span class="font-bold">比赛倒计时</span>
            </div>
          </template>
          <span>{{ time_str }}</span>
        </el-card>
        <el-card class="mb-4" shadow="never">
          <template #header>
            <div>
              <span class="font-bold">时间线</span>
            </div>
          </template>
          <el-timeline>
            <el-timeline-item
              v-for="(activity, index) in activities"
              :key="index"
              :timestamp="activity.timestamp"
              :color="activity.color"
            >
              {{ activity.content }}
            </el-timeline-item>
          </el-timeline>
        </el-card>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref, computed, onUnmounted } from 'vue'
import api from '@/api'
import { useRoute } from 'vue-router'
import { useUserStore } from '@/stores/user'
interface ScoreDetail {
  qid: string
  score: number
}
interface RankInfo {
  user_id: string
  username: string
  total_score: number
  score_details: ScoreDetail[]
}
const userStore = useUserStore()
const now = ref(new Date())
const rankList = ref<RankInfo[]>([])
const time_str = computed(() => {
  const start_time = new Date(compInfo.value.start_time)
  const end_time = new Date(compInfo.value.end_time)
  if (now.value < start_time) {
    const diffTime = start_time.getTime() - now.value.getTime()
    const days = Math.floor(diffTime / (1000 * 60 * 60 * 24))
    const hours = Math.floor((diffTime % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60))
    const minutes = Math.floor((diffTime % (1000 * 60 * 60)) / (1000 * 60))
    const seconds = Math.floor((diffTime % (1000 * 60)) / 1000)
    return `${days}天${hours}小时${minutes}分钟${seconds}秒`
  } else if (now.value > end_time) {
    return '比赛已结束'
  } else {
    return '比赛进行中'
  }
})

const activities = [
  {
    content: '比赛开始',
    timestamp: '',
    color: '#67C23A',
  },
  {
    content: '比赛结束',
    timestamp: '',
    color: '#F56C6C',
  },
]

interface CompInfo {
  name: string
  description: string
  start_time: string
  end_time: string
  competition_type: string
  password_required: boolean
}
const compInfo = ref<CompInfo>({
  name: '',
  description: '',
  start_time: '',
  end_time: '',
  competition_type: '',
  password_required: false,
})
interface QuestionBasicInfo {
  qid: string
  title: string
}

const route = useRoute()
const id = route.params.id
const access_no = ref(false)
const access_yes = ref(false)
const loading = ref(true)
const questionOpen = ref(false)
const questionList = ref<QuestionBasicInfo[]>([])

onMounted(() => {
  document.title = 'OnlineJudge - 比赛详情'
  api.competition
    .getCompetitionInfo(Number(id))
    .then((res) => {
      access_yes.value = true
      compInfo.value = res.data
      document.title = 'OnlineJudge - ' + compInfo.value.name
      activities[0].timestamp = compInfo.value.start_time
      activities[1].timestamp = compInfo.value.end_time
      loading.value = false
      api.competition.getCompetitionQuestionList(Number(id)).then((res) => {
        questionList.value = res.data.question_list
        questionOpen.value = true
      })
    })
    .catch(() => {
      access_no.value = true
      loading.value = false
    })
  const timer = setInterval(() => {
    now.value = new Date()
  }, 1000)

  onUnmounted(() => {
    clearInterval(timer)
  })
})

const handleTabChange = (tab: string) => {
  if (tab === 'rank') {
    api.competition.getCompetitionRankList(Number(id)).then((res) => {
      rankList.value = res.data.rank_list
    })
  }
}
// 提取所有题号（去重）
const allQids = computed(() => {
  const qidSet = new Set<string>()
  rankList.value.forEach((item) => {
    item.score_details.forEach((sd) => {
      qidSet.add(sd.qid)
    })
  })
  return Array.from(qidSet)
})
// 将原始数据格式化为表格可识别的形式
const rankTableData = computed(() => {
  return rankList.value.map((item) => {
    const row: Record<string, unknown> = {
      user_id: item.user_id,
      username: item.username,
      total_score: item.total_score,
    }
    item.score_details.forEach((sd) => {
      row[sd.qid] = sd.score
    })
    return row
  })
})
const rowClassName = ({ row }: { row: { user_id: number } }) => {
  // console.log(row)
  if (row.user_id === userStore.uid) {
    return 'my-row'
  }
  return ''
}
</script>
<style>
.el-table .my-row {
  --el-table-tr-bg-color: var(--el-color-primary-light-9);
}
</style>
