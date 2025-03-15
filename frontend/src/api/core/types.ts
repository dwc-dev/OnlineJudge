// 后端统一响应格式
export interface ApiResponse<T = unknown> {
  code: number
  msg: string
  data: T
}
