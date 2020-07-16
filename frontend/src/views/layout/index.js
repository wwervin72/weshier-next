import React, {Component, Fragment} from 'react'
import {Switch, Route, Redirect} from 'react-router-dom'
import AppHeader from '../../components/header'
import AppFooter from '../../components/footer'
import Blog from "../blog";
import Home from "../home";
import Admin from "../admin";

export default class Layout extends Component {
	render () {
		return (
			<Fragment>
				<AppHeader></AppHeader>
				<Switch>
					<Route exact path="/" component={Home}></Route>
					<Route path='/admin' component={Admin}></Route>
					<Route path="/blog/:blogId" component={Blog}></Route>
					<Redirect from="*" to="/404"></Redirect>
				</Switch>
				<AppFooter></AppFooter>
			</Fragment>
		)
	}
}
