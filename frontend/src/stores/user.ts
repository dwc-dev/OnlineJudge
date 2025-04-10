import { defineStore } from 'pinia'

interface UserInfo {
  uid: number
  username: string
  avatar: string
  email: string
  role: string
  profile: string
}

export const useUserStore = defineStore('user', {
  state: () => ({
    uid: 0,
    username: '',
    avatar: '',
    email: '',
    role: '',
    isLogin: false,
    profile: '',
  }),
  getters: {},
  actions: {
    setUserInfo(userInfo: UserInfo) {
      this.uid = userInfo.uid
      this.username = userInfo.username
      this.avatar = userInfo.avatar
      this.email = userInfo.email
      this.role = userInfo.role
      this.profile = userInfo.profile
      this.isLogin = true
    },
    logout() {
      this.uid = 0
      this.username = ''
      this.avatar = ''
      this.email = ''
      this.role = ''
      this.profile = ''
      this.isLogin = false
    },
  },
})
