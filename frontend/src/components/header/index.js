import React, {Component, Fragment} from 'react'
import {connect} from 'react-redux'
import {withRouter, Link} from 'react-router-dom'
import {Menu, Dropdown} from 'antd'
import Logo from '../logo'
import Avatar from '../avatar'
import './index.scss'

@connect((state, props) => ({
	userInfo: state.userInfo
}))
@withRouter
class Header extends Component {
	noLogin () {
		return (
			<Link to="/login">登录</Link>
		)
	}
	userDropdown () {
		return (
			<Menu>
				<Menu.Item>
					<Link className="nav" to="/logout">登出</Link>
				</Menu.Item>
			</Menu>
		)
	}
	logined () {
		return (
			<Fragment>
				<Link className="nav" to="/editor">写点啥</Link>
				<Link className="nav" to="/admin">后台</Link>
				<Dropdown overlay={this.userDropdown()} placement="bottomLeft">
					<Avatar></Avatar>
				</Dropdown>
			</Fragment>
		)
	}
	render () {
		const {userInfo} = this.props
		return (
			<header>
				<Logo></Logo>
				<nav>
					{
						userInfo ? this.logined() : this.noLogin()
					}
				</nav>
			</header>
		)
	}
}
export default Header
