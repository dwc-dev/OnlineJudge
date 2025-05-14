import { post } from '../http'

export const question = {
  getPublicQuestionList(page: number, pageSize: number, filter: unknown) {
    return post('/question/public_list', { page, page_size: pageSize, filter })
  },

  getQuestionBasicInfo(id: number) {
    return post('/question/basic_info', { id })
  },

  addQuestion(question: Record<string, unknown>) {
    return post('/admin/question/add', question, { needToken: true })
  },

  deleteQuestion(id: number) {
    return post('/admin/question/delete', { id }, { needToken: true })
  },

  updateQuestion(question: Record<string, unknown>) {
    return post('/admin/question/update', question, { needToken: true })
  },

  getQuestionList(page: number, pageSize: number, filter: unknown) {
    return post('/admin/question/list', { page, page_size: pageSize, filter }, { needToken: true })
  },
}
