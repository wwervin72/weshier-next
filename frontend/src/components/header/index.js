import React, {Component} from 'react'
import {connect} from 'react-redux'
import {withRouter} from 'react-router-dom'
import {Button} from 'antd'

@connect((state, props) => ({
}))
@withRouter
class Header extends Component {
	render () {
		const {history} = this.props
		return (
			<header>
				<Button type='primary' onClick={(e) => {
					history.push('/login')
				}}>登录</Button>
			</header>
		)
	}
}
export default Header
