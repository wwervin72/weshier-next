import React, {Component, Fragment} from 'react'
import {connect} from 'react-redux'
import {withRouter, Link} from 'react-router-dom'
import {Menu, Dropdown} from 'antd'
import {HomeOutlined, LogoutOutlined, EditOutlined, LoginOutlined} from '@ant-design/icons'
import Logo from '../logo'
import Avatar from '../avatar/index'
import {ADMIN_ROLE} from '../../utils/variables'
import './index.scss'

@connect((state, props) => ({
	userInfo: state.userInfo
}))
@withRouter
class Header extends Component {
	noLogin () {
		return (
			<Link to="/login"><LoginOutlined /></Link>
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
				<Link className="nav" to="/editor"><EditOutlined /></Link>
				{
					userInfo.role === ADMIN_ROLE ?
						<Link className="nav" to="/admin">
							<img className="admin_back" alt="weshier" src={require('../../assets/img/logo/shier.png')} />
						</Link> :
						''
				}
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
