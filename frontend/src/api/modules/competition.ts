import { post } from '../http'

export const competition = {
  getCompetitionList(page: number, pageSize: number, filter: Record<string, string>) {
    return post('/competition/list', { page, page_size: pageSize, filter })
  },
  getCompetitionInfo(id: number) {
    return post(`/competition/info`, { competition_id: id }, { needToken: true })
  },
  getCompetitionQuestionList(id: number) {
    return post(`/competition/question/list`, { competition_id: id }, { needToken: true })
  },
  getCompetitionQuestionInfo(competitionId: number, qid: string) {
    return post(
      `/competition/question/detail`,
      { competition_id: competitionId, qid },
      { needToken: true },
    )
  },
  getCompetitionRankList(id: number) {
    return post(`/competition/rank/list`, { competition_id: id }, { needToken: true })
  },
  joinCompetition(competition_id: number, user_id: number, password: string) {
    return post(`/competition/attend`, { competition_id, user_id, password }, { needToken: true })
  },
  submitCompetitionQuestion(competitionId: number, qid: string, code: string, language: string) {
    return post(
      `/competition/question/submit`,
      { competition_id: competitionId, qid, code, language },
      { needToken: true },
    )
  },
  adminGetCompetitionList(page: number, pageSize: number, filter: Record<string, string>) {
    return post(
      '/admin/competition/list',
      { page, page_size: pageSize, filter },
      { needToken: true },
    )
  },
  addCompetition(competition: Record<string, unknown>) {
    return post('/admin/competition/add', competition, { needToken: true })
  },
  deleteCompetition(id: number) {
    return post('/admin/competition/delete', { id }, { needToken: true })
  },
  updateCompetition(competition: Record<string, unknown>) {
    return post('/admin/competition/update', competition, { needToken: true })
  },
}
