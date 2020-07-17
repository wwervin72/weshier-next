import {fetch} from './index'
import {loginURL, logoutURL} from '../url'

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
 * 登出
 */
export function logout() {
	return fetch({
		url: logoutURL,
	})
}
