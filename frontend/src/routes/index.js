import React from "react";
import { Switch, Route, BrowserRouter as Router } from "react-router-dom";

import Layout from "../views/layout";
import Login from "../views/login";
import NotMatch from "../views/notFound";

export default () => (
  <Router>
    <Switch>
      <Route path="/login" component={Login}></Route>
      <Route path="/404" component={NotMatch}></Route>
      <Route path="/" component={Layout}></Route>
      <Route path="*" component={NotMatch} />
    </Switch>
  </Router>
);
