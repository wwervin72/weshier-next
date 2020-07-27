import React, {Component} from 'react'
import {connect} from 'react-redux'
import {withRouter} from 'react-router-dom'
import { CAN_EDITOR_ROLE } from '../../utils/variables';

@withRouter
@connect((state) => ({
	userInfo: state.userInfo
}))
class Editor extends Component {
	UNSAFE_componentWillReceiveProps (nextProps) {
		const {userInfo, history} = nextProps
		if (!userInfo) {
			history.replace('/login')
		} else if (!CAN_EDITOR_ROLE.includes(userInfo.role)) {
			history.replace('/')
		}
	}
	render () {
		return (
			<div>
				editor
			</div>
		)
	}
}

export default Editor
