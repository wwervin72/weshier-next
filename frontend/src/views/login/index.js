import React, {Component} from 'react'
import {Form, Input, Button} from 'antd'
// import {PoweroffOutlined} from '@ant-design/icons'
import {login} from '../../api/fetch/user'

const layout = {
	labelCol: { span: 8 },
	wrapperCol: { span: 16 },
  };
  const tailLayout = {
	wrapperCol: { offset: 8, span: 16 },
  };

// @Form.create({
// 	onFieldsChange(props, items) {},
// })
export default class Login extends Component {
	formRef = React.createRef()
	constructor (props) {
		super(props)
		this.loginHandler = this.loginHandler.bind(this)
		this.state = {
			logining: false
		}
	}
	loginHandler (data) {
		if (this.state.logining) return
		this.setState({
			logining: true
		})
		login(data).then(res => {
			console.log(res);
		}).catch(err => {
			console.log(err);
		}).finally(() => {
			this.setState({
				logining: false
			})
		})
	}
	render () {
		const {logining} = this.state
		return (
			<Form {...layout} ref={this.formRef} name="login-form" onFinish={this.loginHandler}>
				<Form.Item name="username" hasFeedback label="账号" rules={[{required: true, message: '请输入登录账号'}]}>
					<Input type="text"></Input>
				</Form.Item>
				<Form.Item name="password" hasFeedback label="密码" rules={[{required: true, message: '请输入登录密码'}]}>
					<Input type="password"></Input>
				</Form.Item>
				<Form.Item {...tailLayout}>
					<Button loading={logining} htmlType="submit" type="primary">登录</Button>
				</Form.Item>
			</Form>
		)
	}
}
