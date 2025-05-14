import { post } from './index'
const ACCESS_TOKEN_KEY = 'access_token'
const REFRESH_TOKEN_KEY = 'refresh_token'

const getAccessToken = (): string | null => localStorage.getItem(ACCESS_TOKEN_KEY)
const setAccessToken = (tokenValue: string) => localStorage.setItem(ACCESS_TOKEN_KEY, tokenValue)
const clearAccessToken = () => localStorage.removeItem(ACCESS_TOKEN_KEY)

const getRefreshToken = (): string | null => localStorage.getItem(REFRESH_TOKEN_KEY)
const setRefreshToken = (tokenValue: string) => localStorage.setItem(REFRESH_TOKEN_KEY, tokenValue)
const clearRefreshToken = () => localStorage.removeItem(REFRESH_TOKEN_KEY)

const refreshToken = async () => {
  await post('/user/refresh_token', { refresh_token: getRefreshToken() })
}

export default {
  getAccessToken,
  setAccessToken,
  clearAccessToken,
  getRefreshToken,
  setRefreshToken,
  clearRefreshToken,
  refreshToken,
}
