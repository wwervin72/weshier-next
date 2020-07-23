import React, {Component} from 'react'
import {Form, Input, Button} from 'antd'
import {connect} from 'react-redux'
import {withRouter} from 'react-router-dom'
import {login} from '../../api/fetch/user'
import {setLoginToken} from '../../utils/cookie'
import {updateUserInfo} from '../../redux/action/user'
import {GITHUB_OAUTH_URI, GITHUB_REDIRECT_URI, GITHUB_CLIENT_ID} from '../../utils/variables'

const layout = {
	labelCol: { span: 8 },
	wrapperCol: { span: 16 },
  };
  const tailLayout = {
	wrapperCol: { offset: 8, span: 16 },
  };

@connect((state) => ({
	userInfo: state.userInfo
}), {
	updateUserInfo
})
@withRouter
class Login extends Component {
	formRef = React.createRef()
	constructor (props) {
		super(props)
		this.loginHandler = this.loginHandler.bind(this)
		this.githubLogin = this.githubLogin.bind(this)
		this.state = {
			logining: false
		}
	}
	componentWillReceiveProps (nextProps) {
		const {userInfo, history} = nextProps
		// 已登录
		if (userInfo) {
			history.replace('/')
		}
	}
	loginHandler (data) {
		if (this.state.logining) return
		this.setState({
			logining: true
		})
		login(data).then(res => {
			this.setState({
				logining: false
			})
			setLoginToken(res.token)
			this.props.updateUserInfo(res)
			this.props.history.replace('/')
		}).catch(err => {
			this.setState({
				logining: false
			})
		})
	}
	githubLogin () {
		const state = Date.now()
		const url = `${GITHUB_OAUTH_URI}?client_id=${GITHUB_CLIENT_ID}&redirect_uri=${GITHUB_REDIRECT_URI}&scope=user&state=${state}`
		const myWindow = window.open(
			url,
			'weshier-github-login',
			'modal=yes,toolbar=no,titlebar=no,menuba=no,location=no,top=200,left=500,width=600,height=400'
		  )
		myWindow.focus()
	}
	render () {
		const {logining} = this.state
		return (
			<Form {...layout} ref={this.formRef} name="login-form" onFinish={this.loginHandler}>
				<Form.Item name="userName" hasFeedback label="账号" rules={[{required: true, message: '请输入登录账号'}]}>
					<Input type="text"></Input>
				</Form.Item>
				<Form.Item name="passWord" hasFeedback label="密码" rules={[{required: true, message: '请输入登录密码'}]}>
					<Input type="password"></Input>
				</Form.Item>
				<Form.Item {...tailLayout}>
					<Button loading={logining} htmlType="submit" type="primary">登录</Button>
				</Form.Item>
				<Form.Item {...tailLayout}>
					<Button onClick={this.githubLogin} htmlType="button" type="primary">github</Button>
				</Form.Item>
			</Form>
		)
	}
}

export default Login
