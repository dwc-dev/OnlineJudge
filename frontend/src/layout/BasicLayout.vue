<template>
  <el-container class="h-full">
    <el-header class="!box-content !border-b-1 !border-b-gray-200 !p-0">
      <GlobalHeader />
    </el-header>
    <el-main class="!p-0">
      <RouterView />
    </el-main>
    <el-footer
      v-if="!isCodingPage && !isCompetitionQuestionPage"
      class="flex !h-10 items-center justify-center bg-gray-100"
    >
      <span class="text-base text-gray-500">OnlineJudge</span>
    </el-footer>
  </el-container>
</template>

<script setup lang="ts">
import GlobalHeader from '@/components/GlobalHeader.vue'
import { computed } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const isCodingPage = computed(() => route.name === 'questionDetail')
const isCompetitionQuestionPage = computed(() => route.name === 'competitionQuestion')
</script>
