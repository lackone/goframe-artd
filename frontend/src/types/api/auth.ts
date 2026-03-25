declare namespace Api.Auth {
  interface LoginParams {
    account: string
    password: string
  }

  interface LoginResponse {
    token: string
    refreshToken: string
  }

  interface UserInfo {
    buttons: string[]
    roles: string[]
    userId: number
    userName: string
    email: string
    avatar?: string
    phone?: string
    address?: string
    real_name?: string
    nick_name?: string
    remark?: string
    weixin?: string
    account?: string
    sex?: number
  }
}
