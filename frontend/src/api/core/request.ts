import type { ApiResponse } from './types'
import { requestInterceptor, responseInterceptor } from './interceptors'
import { ElMessage } from 'element-plus'

const BASE_URL = 'http://localhost:8001/api/v1'

// 基础请求方法
const baseRequest = async <T>(
  url: string,
  config: RequestInit,
  params: Record<string, string | number> = {},
): Promise<ApiResponse<T>> => {
  const interceptedConfig = requestInterceptor(config)
  const fullUrl = `${BASE_URL}${url}${formatGetParams(params)}`

  let response: Response

  try {
    response = await fetch(fullUrl, interceptedConfig)
  } catch {
    ElMessage.error('网络请求失败')
    throw new Error('网络请求失败')
  }

  return responseInterceptor(response)
}

// GET 请求
export const get = async <T>(
  url: string,
  params: Record<string, string | number> = {},
  config: RequestInit = {},
): Promise<ApiResponse<T>> => {
  return baseRequest<T>(url, { ...config, method: 'GET' }, params)
}

// 格式化 GET 参数
const formatGetParams = (params: Record<string, string | number>): string => {
  const queryString = Object.keys(params)
    .map((key) => `${encodeURIComponent(key)}=${encodeURIComponent(params[key])}`)
    .join('&')
  return queryString ? `?${queryString}` : ''
}

// POST 请求
export const post = async <T>(
  url: string,
  body: unknown,
  config: RequestInit = {},
): Promise<ApiResponse<T>> => {
  const headers = new Headers(config.headers)
  headers.set('Content-Type', 'application/json')

  return baseRequest<T>(url, {
    ...config,
    method: 'POST',
    headers,
    body: JSON.stringify(body),
  })
}

export default {
  get,
  post,
}
