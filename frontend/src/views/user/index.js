import React, {Component} from 'react'
import {withRouter, matchPath} from 'react-router-dom'

@withRouter
@matchPath
class User extends Component {
	render () {
		const {userId} = this.props
		console.log(userId);
		return (
			<div>user nihao {userId}</div>
		)
	}
}

export default User
