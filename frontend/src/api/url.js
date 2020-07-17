import {baseURL} from './baseUrl'

export const prefixURL = 'api/'

export const userURL = 'user/'

export const base = baseURL + prefixURL

// 登录
export const loginURL = base + userURL + '/login'
// 登出
export const logoutURL = base + userURL + '/logout'
