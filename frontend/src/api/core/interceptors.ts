import type { ApiResponse } from './types'
import { ElMessage } from 'element-plus'
import jwt from './jwt'

// 请求拦截器
export const requestInterceptor = (config: RequestInit): RequestInit => {
  // 添加 JWT 到请求头
  const token = jwt.getToken()
  if (token) {
    const headers = new Headers(config.headers)
    headers.set('Authorization', `Bearer ${token}`)
    return { ...config, headers }
  }
  return config
}

// 响应拦截器
export const responseInterceptor = async <T>(response: Response): Promise<ApiResponse<T>> => {
  const jsonData = (await response.json()) as ApiResponse<T>
  if (!response.ok) {
    ElMessage.error(jsonData.msg)
    throw new Error(jsonData.msg) // throw 会自动触发 reject
  } else {
    return jsonData
  }
}

export default {
  requestInterceptor,
  responseInterceptor,
}
