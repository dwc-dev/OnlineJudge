import { post } from '../http'

export const question = {
  pagination(page: number, pageSize: number, filter: unknown) {
    return post('/question/pagination', { page, page_size: pageSize, filter })
  },
  getQuestionTitleAndContent(id: number) {
    return post('/question/info', { id })
  },
}
