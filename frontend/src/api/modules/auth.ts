import request from '../core/request'

interface RegisterReq {
  username: string
  email: string
  password: string
}

// 注册接口
export const register = (body: RegisterReq) => {
  return request.post<unknown>('/user/register', body)
}

export default {
  register,
}
