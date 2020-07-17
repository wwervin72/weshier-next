import React, {Component} from 'react'
import {Form, Input, Button} from 'antd'

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
	render () {
		return (
			<Form {...layout} ref={this.formRef} name="login-form">
				<Form.Item name="username" label="账号" rules={[{required: true, message: '请输入登录账号'}]}>
					<Input type="text"></Input>
				</Form.Item>
				<Form.Item name="password" label="密码" rules={[{required: true, message: '请输入登录密码'}]}>
					<Input type="password"></Input>
				</Form.Item>
				<Form.Item {...tailLayout}>
					<Button htmlType="submit" type="primary">登录</Button>
				</Form.Item>
			</Form>
		)
	}
}
