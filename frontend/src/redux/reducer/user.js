import {handleActions} from 'redux-actions'
import {UPDATE_USER_INFO} from '../actionType'

export default handleActions({
	[UPDATE_USER_INFO] (state, action) {
		let data = {...state, ...action.payload}
		return data
	},
}, null)
