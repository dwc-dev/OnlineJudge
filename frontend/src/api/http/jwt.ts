// JWT 操作
const TOKEN_KEY = 'jwt'
const getToken = (): string | null => localStorage.getItem(TOKEN_KEY)
const setToken = (tokenValue: string) => localStorage.setItem(TOKEN_KEY, tokenValue)
const clearToken = () => localStorage.removeItem(TOKEN_KEY)

export default {
  getToken,
  setToken,
  clearToken,
}
