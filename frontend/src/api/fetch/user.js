import {fetch} from './index'
import {loginURL, logoutURL, userInfoURL} from '../url'

/**
 * 登录
 * @param {*} data
 */
export function login(data) {
	return fetch({
		method: 'post',
		url: loginURL,
		data
	})
}

/**
 * 根据 token 获取用户信息
 * @param {*} data
 */
export function fetchUserInfo() {
	return fetch({
		url: userInfoURL,
	})
}

/**
 * 登出
 */
export function logout() {
	return fetch({
		url: logoutURL,
	})
}
