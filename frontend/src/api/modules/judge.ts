import { post } from '../http'

export const judge = {
  judgeCode(code: string, language: string, question_id: number) {
    return post('/judge/judgecode', { code, language, question_id }, { timeout: 60000 })
  },
}
