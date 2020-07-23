import {UPDATE_USER_INFO} from '../actionType'
import {getLoginToken, removeLoginToken} from '../../utils/cookie'
import {fetchUserInfo} from '../../api/fetch/user'

export function updateUserInfo (userInfo) {
	return (dispatch, getState) => {
		dispatch({
			type: UPDATE_USER_INFO,
			payload: userInfo
		})
	}
}

/**
 * 初始化用户信息
 */
export function initUserInfo() {
	return dispatch => {
		let token = getLoginToken()
		if (token) {
			fetchUserInfo().then(res => {
				dispatch(updateUserInfo(res))
			}).catch(err => {
				dispatch(updateUserInfo(null))
				// 加载失败，清除 token
				removeLoginToken()
				if (process.env.NODE_ENV === 'development') {
					console.log(err);
				}
			})
		}
	}
}
