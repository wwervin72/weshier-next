import React, {Component} from 'react'
import {connect} from 'react-redux'
import {withRouter} from 'react-router-dom'
import {Button} from 'antd'

@connect((state, props) => ({
	userInfo: state.userInfo
}))
@withRouter
class Header extends Component {
	render () {
		const {history, userInfo} = this.props
		return (
			<header>
				<Button type='primary' onClick={(e) => {
					history.push('/login')
				}}>登录</Button>
				{
					userInfo ? <p>您好: {userInfo.nickname}</p> : ''
				}
				<Button type='primary' onClick={(e) => {
					history.push('/admin')
				}}>admin</Button>
			</header>
		)
	}
}
export default Header
