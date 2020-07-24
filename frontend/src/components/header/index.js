import React, {Component, Fragment} from 'react'
import {connect} from 'react-redux'
import {withRouter, Link} from 'react-router-dom'
import {Menu, Dropdown} from 'antd'
import {HomeOutlined, LogoutOutlined} from '@ant-design/icons'
import Logo from '../logo'
import Avatar from '../avatar/index'
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
	userDropdown (userInfo) {
		return (
			<Menu>
				<Menu.Item>
					<Link className="nav" to={"/u/" + userInfo.id}>
						<HomeOutlined />
						<label>主页</label>
					</Link>
				</Menu.Item>
				<Menu.Item>
					<Link className="nav" to="/logout">
						<LogoutOutlined />
						<label>登出</label>
					</Link>
				</Menu.Item>
			</Menu>
		)
	}
	logined () {
		const {userInfo} = this.props
		return (
			<Fragment>
				<Link className="nav" to="/editor">写点啥</Link>
				<Link className="nav" to="/admin">后台</Link>
				<Dropdown overlayClassName="avatar_op" overlay={this.userDropdown(userInfo)} placement="bottomRight">
					<span className="avatar_trigger">
						<Avatar user={userInfo}></Avatar>
					</span>
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
