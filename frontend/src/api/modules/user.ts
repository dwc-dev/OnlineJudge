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

  getUserList(page: number, pageSize: number, filter: unknown) {
    return post('/admin/user/list', { page, page_size: pageSize, filter }, { needToken: true })
  },

  addUser(user: Record<string, unknown>) {
    return post('/admin/user/add', user, { needToken: true })
  },

  updateUser(user: Record<string, unknown>) {
    return post('/admin/user/update', user, { needToken: true })
  },

  deleteUser(id: number) {
    return post('/admin/user/delete', { user_id: id }, { needToken: true })
  },

  logout() {
    return post('/user/logout', undefined, { needToken: true })
  },
}
