import axios from 'axios'
import token from './token'
import type { InternalRequestConfig, RequestConfig } from './types'

export const BASE_URL = import.meta.env.VITE_APP_API_BASE_URL || 'http://localhost:8888/api/v1'

const instance = axios.create({baseURL: BASE_URL})

instance.interceptors.request.use((config: InternalRequestConfig) => {
  if (config.needToken) {
    const accessToken = token.getAccessToken()
    if (accessToken) {
      config.headers.Authorization = `Bearer ${accessToken}`
    }
  }
  return config
})

instance.interceptors.response.use(
  (response) => {
    const data = response.data.data
    // data.access_token != null 会同时过滤掉 null 和 undefined 的值
    // null == undefined  true
    // null === undefined false
    if (data && data.access_token != null) {
      token.setAccessToken(data.access_token)
    }
    if (data && data.refresh_token != null) {
      token.setRefreshToken(data.refresh_token)
    }
    return response.data
  },
  async (error) => {
    if (!error.response) {
      error.message = '网络异常，请检查你的网络连接'
      return Promise.reject(error)
    }
    const { data } = error.response
    // AccessToken 无效
    if (data && data.code === 40101) {
      console.log('AccessToken 无效', error.config.url)
      await token.refreshToken()
      error.config.headers.Authorization = `Bearer ${token.getAccessToken()}`
      return instance.request(error.config)
    }
    // RefreshToken 无效
    if (data && data.code === 40102) {
      console.log('RefreshToken 无效', error.config.url)
      token.clearAccessToken()
      token.clearRefreshToken()
      window.location.href = '/auth/login'
      return Promise.reject(error)
    }
    const msg = data ? data.msg : undefined
    //a || b 表达式的返回值是“第一个 truthy 值”或最后一个值
    error.message = msg || '未知错误'
    return Promise.reject(error)
  },
)

export function get(url: string, params?: Record<string, unknown>, config?: RequestConfig) {
  return instance.get(url, { params, ...config })
}

export function post(
  url: string,
  data?: Record<string, unknown> | FormData | Blob | File | unknown,
  config?: RequestConfig,
) {
  return instance.post(url, data, config)
}
