// JWT 操作
const TOKEN_KEY = 'jwt'
export const getToken = (): string | null => localStorage.getItem(TOKEN_KEY)
export const setToken = (tokenValue: string) => localStorage.setItem(TOKEN_KEY, tokenValue)
export const clearToken = () => localStorage.removeItem(TOKEN_KEY)

export default {
  getToken,
  setToken,
  clearToken,
}
