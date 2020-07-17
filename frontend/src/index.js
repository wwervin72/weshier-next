import React from 'react';
import ReactDOM from 'react-dom';
import {Provider} from 'react-redux'
import Route from './routes'
import store from './redux/store'
import './styles/index.css';

ReactDOM.render(
	<Provider store={store}>
		<Route></Route>
	</Provider>,
  document.getElementById('root')
);
