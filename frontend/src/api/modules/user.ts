import { post } from '../http'

export const user = {
  getUserInfo() {
    return post('/user/info', undefined, { needToken: true })
  },

  login(email: string, password: string) {
    return post('/user/login', { email, password })
  },

  register(username: string, email: string, password: string) {
    return post('/user/register', { username, email, password })
  },

  getUserPublicInfoById(id: number) {
    return post('/user/public_info', { user_id: id })
  },
}
