import React from 'react';
import ReactDOM from 'react-dom';
import {Provider} from 'react-redux'
import Route from './routes'
import store from './redux/store'
import {initUserInfo} from './redux/action/user'
import './styles/index.css';

store.dispatch(initUserInfo())

ReactDOM.render(
	<Provider store={store}>
		<Route></Route>
	</Provider>,
  document.getElementById('root')
);
