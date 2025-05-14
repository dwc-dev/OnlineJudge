import { post } from '../http'

export const judge = {
  judgeCode(code: string, language: string, question_id: number) {
    return post(
      '/judge/judgecode',
      { code, language, question_id },
      { timeout: 60000, needToken: true },
    )
  },
  getJudgeList(page: number, pageSize: number) {
    return post('/judge/list', { page, page_size: pageSize }, { needToken: true })
  },
}
