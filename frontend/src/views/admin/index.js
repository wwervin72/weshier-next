import React, {Component} from 'react'
import {connect} from 'react-redux'
import {withRouter} from 'react-router-dom'
import {ADMIN_ROLE} from '../../utils/variables'

@connect((state) => ({
	userInfo: state.userInfo
}))
@withRouter
class Admin extends Component {
	UNSAFE_componentWillReceiveProps (nextProps) {
		const {userInfo, history} = nextProps
		if (!userInfo) {
			history.replace('/login')
		} else if (userInfo.role !== ADMIN_ROLE) {
			history.replace('/')
		}
	}
	render () {
		return (
			<div>Admin</div>
		)
	}
}

export default Admin
