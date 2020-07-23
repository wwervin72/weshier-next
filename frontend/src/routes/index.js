import React from "react";
import { Switch, Route, BrowserRouter as Router } from "react-router-dom";
import Layout from "../views/layout";
import Login from "../views/login";
import Github from "../views/oauth/github";
import NotMatch from "../views/notFound";

export default () => {
	return (
	<Router>
		<Switch>
			<Route path="/login" component={Login}></Route>
			<Route path="/login-github" component={Github}></Route>
			<Route path="/404" component={NotMatch}></Route>
			<Route path="/" component={Layout}></Route>
			<Route path="*" component={NotMatch} />
		</Switch>
	</Router>)
};
