import React, {Component, Fragment} from 'react'
import {connect} from 'react-redux'
import {Switch, Route, Redirect} from 'react-router-dom'
import store from '../../redux/store'
import AppHeader from '../../components/header'
import AppFooter from '../../components/footer'
import Blog from "../blog"
import Home from "../home"
import Admin from "../admin"
import { ADMIN_ROLE } from '../../utils/variables';

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
					<Route path='/admin' render={() => {
						const {userInfo} = store.getState()
						if (userInfo && userInfo.role === ADMIN_ROLE) {
							return <Admin></Admin>
						} else {
							return <Redirect to={
								userInfo ? '/' : '/login'
							}></Redirect>
						}
					}}></Route>
					<Route path="/blog/:blogId" component={Blog}></Route>
					<Redirect from="*" to="/404"></Redirect>
				</Switch>
				<AppFooter></AppFooter>
			</Fragment>
		)
	}
}

export default Layout
