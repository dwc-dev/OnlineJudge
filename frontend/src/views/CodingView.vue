<template>
  <div class="flex h-full w-full gap-1 p-2">
    <el-card class="h-full flex-1" :body-style="{ height: '100%' }">
      <el-skeleton :rows="10" animated v-if="loading" />
      <div v-else class="flex h-full w-full flex-col">
        <MdPreview :modelValue="questionTitle" previewTheme="vuepress" />
        <el-scrollbar class="h-full flex-1 pr-2">
          <MdPreview :modelValue="questionContent" previewTheme="vuepress" />
        </el-scrollbar>
      </div>
    </el-card>

    <el-card class="h-full flex-1" :body-style="{ height: '100%' }">
      <div class="flex h-full w-full flex-col">
        <div class="mb-4 flex w-full justify-between gap-2">
          <el-select v-model="selectedLanguage" class="!w-1/4" title="选择语言">
            <el-option v-for="lang in languages" :key="lang" :label="lang" :value="lang" />
          </el-select>
          <div class="flex items-center justify-end gap-4">
            <el-icon
              @click="submitCode"
              class="cursor-pointer rounded-md hover:bg-gray-200"
              title="提交代码"
              :size="25"
            >
              <Loading v-if="submitting" class="is-loading" />
              <CaretRight v-else />
            </el-icon>
            <el-popconfirm title="确定重置代码吗？" placement="bottom" @confirm="resetCode">
              <template #reference>
                <el-icon
                  class="cursor-pointer rounded-md !p-0.5 hover:bg-gray-200"
                  title="重置代码"
                  :size="25"
                >
                  <RefreshRight /> </el-icon
              ></template>
            </el-popconfirm>
            <el-icon
              @click="formatCode"
              class="cursor-pointer rounded-md hover:bg-gray-200"
              title="格式化代码"
              :size="25"
            >
              <MagicStick />
            </el-icon>
          </div>
        </div>
        <div ref="editorContainer" class="h-full w-full flex-1 border-1 border-gray-300"></div>
      </div>
    </el-card>

    <div
      class="relative flex h-full flex-col"
      :class="{ 'w-12': !aiChatExpanded, 'flex-1': aiChatExpanded }"
    >
      <el-button
        type="primary"
        plain
        class="absolute top-1/2 -left-4 z-10 -translate-y-1/2 !rounded-md !px-1 !py-3"
        size="small"
        @click="aiChatExpanded = !aiChatExpanded"
      >
        <el-icon>
          <arrow-right v-if="aiChatExpanded" />
          <arrow-left v-else />
        </el-icon>
      </el-button>
      <el-card v-if="aiChatExpanded" class="h-full w-full">
        <AIChat />
      </el-card>
      <div v-else class="flex h-full w-full items-center justify-center rounded-md bg-gray-100">
        <div class="flex flex-col items-center py-4 text-gray-500">
          <span>AI</span>
          <span>助</span>
          <span>手</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import api from '@/api'
import { useRoute } from 'vue-router'
import { ref, onMounted, onBeforeUnmount, watch } from 'vue'
import { MdPreview } from 'md-editor-v3'
import { ElMessage } from 'element-plus'
import { ArrowRight, ArrowLeft, CaretRight, Loading } from '@element-plus/icons-vue'
import AIChat from '@/components/AIChat.vue'
import 'md-editor-v3/lib/preview.css'
import * as monaco from 'monaco-editor'

const route = useRoute()
const id = route.params.id
const questionTitle = ref('')
const questionContent = ref('')
const loading = ref(true)
const submitting = ref(false)

const code = ref('')
const selectedLanguage = ref('C')
const languages = ['C', 'C++', 'Java', 'Python', 'Go', 'Rust']
const judgeResults = ref([])
const editorContainer = ref<HTMLElement | null>(null)
let editor: monaco.editor.IStandaloneCodeEditor | null = null

const aiChatExpanded = ref(false)

// 语言映射配置
const editorLanguageMap: Record<string, string> = {
  C: 'c',
  'C++': 'cpp',
  Java: 'java',
  Python: 'python',
  Go: 'go',
  Rust: 'rust',
}
const judgeLanguageMap: Record<string, string> = {
  C: 'c',
  'C++': 'cpp',
  Java: 'java',
  Python: 'python',
  Go: 'golang',
  Rust: 'rust',
}

// 语言默认代码模板
const defaultCodes: Record<string, string> = {
  C: '#include <stdio.h>\n\nint main() {\n    // 在此编写代码\n    return 0;\n}',
  'C++': '#include <iostream>\n\nint main() {\n    // 在此编写代码\n    return 0;\n}',
  Java: 'public class Main {\n    public static void main(String[] args) {\n        // 在此编写代码\n    }\n}',
  Python: '# 在此编写代码',
  Go: 'package main\n\nfunc main() {\n    // 在此编写代码\n}',
  Rust: 'fn main() {\n    // 在此编写代码\n}',
}

onMounted(() => {
  api.question
    .getQuestionTitleAndContent(Number(id))
    .then((res) => {
      questionTitle.value = '## ' + res.data.title
      questionContent.value = res.data.content
      document.title = id + ' - ' + res.data.title
      loading.value = false

      // 初始化Monaco编辑器
      initEditor()
    })
    .catch(() => {
      ElMessage.error('获取题目失败')
      loading.value = false
    })
})

onBeforeUnmount(() => {
  if (editor) {
    editor.dispose()
    editor = null
  }
})

watch(selectedLanguage, (newLang) => {
  if (editor) {
    // 保存当前代码
    code.value = editor.getValue()

    // 设置新的语言模式
    monaco.editor.setModelLanguage(editor.getModel()!, editorLanguageMap[newLang])

    // 如果当前代码为空或是上一个语言的默认代码，则设置新语言的默认代码
    const currentCode = editor.getValue()
    if (currentCode === '' || Object.values(defaultCodes).includes(currentCode)) {
      editor.setValue(defaultCodes[newLang])
    }
  }
})

const initEditor = () => {
  if (editorContainer.value) {
    editor = monaco.editor.create(editorContainer.value, {
      value: defaultCodes[selectedLanguage.value],
      language: editorLanguageMap[selectedLanguage.value],
      theme: 'vs-light',
      minimap: { enabled: false },
      fontSize: 16,
    })

    // 监听编辑器内容变化
    editor.onDidChangeModelContent(() => {
      if (editor) {
        code.value = editor.getValue()
      }
    })
  }
}

const submitCode = () => {
  if (!editor) return
  if (submitting.value) {
    ElMessage.warning('代码已在评测中~')
    return
  }

  code.value = editor.getValue()
  submitting.value = true

  api.judge
    .judgeCode(code.value, judgeLanguageMap[selectedLanguage.value], Number(id))
    .then((res) => {
      judgeResults.value = res.data.results
      ElMessage.success('代码已提交')
    })
    .catch((error) => {
      ElMessage.error(error.message)
    })
    .finally(() => {
      submitting.value = false
    })
}

const resetCode = () => {
  if (!editor) return
  editor?.setValue(defaultCodes[selectedLanguage.value])
}

const formatCode = () => {
  if (!editor) return
  editor?.getAction('editor.action.formatDocument')?.run()
}
</script>

<style scoped>
.is-loading {
  animation: loading-rotate 2s linear infinite;
}

@keyframes loading-rotate {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}
</style>
