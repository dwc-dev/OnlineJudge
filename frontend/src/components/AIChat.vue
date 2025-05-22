<template>
  <div class="flex h-full w-full flex-col" v-loading="pageLoading">
    <div class="mb-2 flex items-center">
      <el-select v-model="currentSessionId" placeholder="新会话" @change="handleSessionChange">
        <el-option
          v-for="(session, idx) in sessions"
          :key="idx"
          :label="session.title"
          :value="session.session_id"
        />
      </el-select>
      <el-icon
        class="ml-2 cursor-pointer rounded-md p-0.5 hover:bg-gray-200"
        @click="handleCreateSession"
        :size="25"
        title="新建会话"
      >
        <Plus />
      </el-icon>
    </div>
    <BubbleList :list="list" class="bubble-list flex-1" maxHeight="99999px" ref="bubbleListRef" />
    <Sender
      placeholder="👋🤖 你好，我是AI助手！"
      clearable
      @submit="handleSubmit"
      @cancel="handleCancel"
      v-model="senderValue"
      :disabled="senderDisabled"
      class="mt-2"
      :loading="isLoading"
    />
  </div>
</template>

<script setup lang="ts">
import { BubbleList, Sender, useXStream } from 'vue-element-plus-x'
import type { BubbleListItemProps, BubbleListProps } from 'vue-element-plus-x/types/BubbleList'
import { onMounted, computed, ref, watch } from 'vue'
import { useUserStore } from '@/stores/user'
import { useRoute } from 'vue-router'
import { Plus } from '@element-plus/icons-vue'
import token from '@/api/http/token'
import api from '@/api'

// import 'vue-element-plus-x/styles/prism.min.css'
// import 'vue-element-plus-x/styles/prism-coy.min.css'

type Session = {
  session_id: string
  question_id: number
  title: string
}

const bubbleListRef = ref()
const aiAvatarURL = ref('/deepseek-color.svg')
const list = ref<BubbleListProps<BubbleListItemProps>['list']>([])
const userStore = useUserStore()
const pageLoading = ref(true)
const senderDisabled = ref(false)
const route = useRoute()
const userId = computed(() => userStore.uid)
const questionId = computed(() => Number(route.params.id))
const currentSessionId = ref<string>('')
const sessions = ref<Session[]>([])
const senderValue = ref<string>('')

onMounted(() => {
  api.ai.getQuestionSessions(userId.value, questionId.value).then((res) => {
    sessions.value = res.data
    if (sessions.value.length == 0) {
      pageLoading.value = false
      return
    }
    // 默认会话为第一个会话
    currentSessionId.value = sessions.value[0].session_id
    api.ai.getSessionChatHistory(userId.value, currentSessionId.value).then((res) => {
      for (const item of res.data) {
        list.value.push({
          content: item.content,
          placement: item.role === 'user' ? 'end' : 'start',
          isMarkdown: item.role === 'user' ? false : true,
          avatar: item.role === 'user' ? userStore.avatar : aiAvatarURL.value,
          avatarSize: '24px', // 头像占位大小
          avatarGap: '12px', // 头像与气泡之间的距离
        })
      }
      pageLoading.value = false
    })
  })
})

const handleSessionChange = () => {
  list.value = []
  pageLoading.value = true
  api.ai.getSessionChatHistory(userId.value, currentSessionId.value).then((res) => {
    for (const item of res.data) {
      list.value.push({
        content: item.content,
        placement: item.role === 'user' ? 'end' : 'start',
        isMarkdown: item.role === 'user' ? false : true,
        avatar: item.role === 'user' ? userStore.avatar : aiAvatarURL.value,
        avatarSize: '24px', // 头像占位大小
        avatarGap: '12px', // 头像与气泡之间的距离
      })
    }
    pageLoading.value = false
  })
}

const handleCreateSession = () => {
  currentSessionId.value = ''
  list.value = []
}

const { startStream, cancel, data, isLoading } = useXStream()
const handleSubmit = async () => {
  senderDisabled.value = true
  const aiContent = computed(() => {
    let text = ''
    for (let index = 0; index < data.value.length; index++) {
      const chunk = data.value[index].data
      try {
        if (!chunk) {
          break
        }
        const parsed = JSON.parse(chunk)
        if (parsed.msg === 'DONE') {
          text = aiContent.value
          break
        }
        text += parsed.content || ''
      } catch (err) {
        console.error('解析失败:', err)
      }
    }
    return text
  })
  const sessionInfo = computed(() => {
    const sessionInfo: Session = {
      session_id: '',
      question_id: questionId.value,
      title: '',
    }
    for (let index = data.value.length - 2; index >= 0; index--) {
      const chunk = data.value[index].data
      const parsed = JSON.parse(chunk)
      if (parsed.msg === 'DONE') {
        sessionInfo.session_id = parsed.session_id
        sessionInfo.title = parsed.title
        break
      }
    }
    return sessionInfo
  })

  const message = senderValue.value.trim()
  if (!message) return

  // 添加用户消息
  list.value.push({
    content: message,
    placement: 'end',
    isMarkdown: false,
    avatar: userStore.avatar,
    avatarSize: '24px', // 头像占位大小
    avatarGap: '12px', // 头像与气泡之间的距离
  })
  senderValue.value = ''

  // 添加 AI 占位气泡
  const aiMsg = ref<BubbleListItemProps>({
    content: '',
    placement: 'start',
    isMarkdown: true,
    typing: true,
    loading: true,
    avatar: aiAvatarURL.value,
    avatarSize: '24px', // 头像占位大小
    avatarGap: '12px', // 头像与气泡之间的距离
  })
  list.value.push(aiMsg.value)
  bubbleListRef.value.scrollToBottom()
  try {
    let res = await fetch(`${api.BASE_URL}/ai/chat`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token.getAccessToken()}`,
      },
      body: JSON.stringify({
        user_id: userId.value,
        question_id: questionId.value,
        session_id: currentSessionId.value,
        message: message,
      }),
    })
    if (!res.ok && res.status === 401) {
      await token.refreshToken()
      res = await fetch(`${api.BASE_URL}/ai/chat`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token.getAccessToken()}`,
        },
        body: JSON.stringify({
          user_id: userId.value,
          question_id: questionId.value,
          session_id: currentSessionId.value,
          message: message,
        }),
      })
    }
    const readableStream = res.body!
    const stopWatchAiContent = watch(
      aiContent,
      (val) => {
        if (val.length > 0) {
          aiMsg.value.loading = false
          aiMsg.value.content = val
        } else {
          console.log('val的值为空')
        }
      },
      { immediate: false },
    )
    const stopWatchSessionInfo = watch(
      sessionInfo,
      (val) => {
        if (val.session_id !== '' && val.title !== '' && currentSessionId.value === '') {
          currentSessionId.value = val.session_id
          sessions.value = [
            {
              session_id: val.session_id,
              question_id: questionId.value,
              title: val.title,
            },
            ...sessions.value,
          ]
        }
        if (val.session_id !== '') {
          console.log('sessionInfo:', val)
        }
      },
      { immediate: false },
    )
    await startStream({ readableStream })
    stopWatchAiContent()
    stopWatchSessionInfo()
  } catch (err) {
    console.error('请求出错:', err)
    aiMsg.value.content = '❌ 出错了，请稍后重试'
    aiMsg.value.loading = false
    cancel()
  } finally {
    senderDisabled.value = false
  }
}
const handleCancel = () => {
  list.value.pop()
  cancel()
}
</script>

<style scoped></style>
