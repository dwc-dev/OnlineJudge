import { post } from '../http'

export const ai = {
  getQuestionSessions(userId: number, questionId: number) {
    return post(
      '/ai/question/sessions',
      { user_id: userId, question_id: questionId },
      { needToken: true, timeout: 600000 },
    )
  },
  getSessionChatHistory(userId: number, sessionId: string) {
    return post(
      '/ai/session/chat/history',
      { user_id: userId, session_id: sessionId },
      { needToken: true, timeout: 600000 },
    )
  },
  codeCheck(questionId: number, code: string) {
    return post(
      '/ai/code/check',
      { question_id: questionId, code },
      { needToken: true, timeout: 600000 },
    )
  },
}
