import {baseURL} from './baseUrl'

export const prefixURL = 'api/'

export const userURL = 'user'

export const base = baseURL + prefixURL

// 登录
export const loginURL = base + userURL + '/login'
// 根据 token 获取用户信息
export const userInfoURL = base + userURL
// 登出
export const logoutURL = base + userURL + '/logout'
