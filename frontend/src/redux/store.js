import {createStore, combineReducers, applyMiddleware} from 'redux'
import thunk from 'redux-thunk'
import {createLogger} from 'redux-logger'

import userInfo from './reducer/user'

const middleWare = [thunk]

if (process.env.NODE_ENV !== 'production') {
	middleWare.push(createLogger())
}

export default createStore(combineReducers({
	userInfo
}), {
	userInfo: null
}, applyMiddleware(...middleWare))
