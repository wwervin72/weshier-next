import React, {Component, Fragment} from 'react'
import {connect} from 'react-redux'
import {Switch, Route, Redirect} from 'react-router-dom'
import AppHeader from '../../components/header/index.js'
import AppFooter from '../../components/footer'
import Blog from "../blog"
import Home from "../home"
import Admin from "../admin"
import Editor from "../editor"
import User from "../user"

@connect((state) => ({
	userInfo: state.userInfo
}))
class Layout extends Component {
	render () {
		return (
			<Fragment>
				<AppHeader></AppHeader>
				<Switch>
					<Route exact path="/" component={Home}></Route>
					<Route path="/editor" component={Editor}></Route>
					<Route path='/admin' component={Admin}></Route>
					<Route path='/u/:userId' component={User}></Route>
					<Route path="/blog/:blogId" component={Blog}></Route>
					<Redirect from="*" to="/404"></Redirect>
				</Switch>
				<AppFooter></AppFooter>
			</Fragment>
		)
	}
}

export default Layout
