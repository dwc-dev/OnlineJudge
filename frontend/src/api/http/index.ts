import axios from 'axios'
import jwt from './jwt'
import type { InternalRequestConfig, RequestConfig } from './types'

const instance = axios.create({
  baseURL: import.meta.env.VITE_APP_API_BASE_URL || 'http://localhost:8888/api/v1',
  timeout: 10000,
})

instance.interceptors.request.use((config: InternalRequestConfig) => {
  if (config.needToken) {
    const token = jwt.getToken()
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
  }
  return config
})

instance.interceptors.response.use(
  (response) => {
    const data = response.data?.data
    // data.token != null 会同时过滤掉 null 和 undefined 的值
    // null == undefined  true
    // null === undefined false
    if (data && data.token != null) {
      jwt.setToken(data.token)
    }
    return response.data
  },
  (error) => {
    if (!error.response) {
      error.message = '网络异常，请检查你的网络连接'
      return Promise.reject(error)
    }
    const { data } = error.response
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
  data?: Record<string, unknown> | FormData | Blob | File,
  config?: RequestConfig,
) {
  return instance.post(url, data, config)
}
