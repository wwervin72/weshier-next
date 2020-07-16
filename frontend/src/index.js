import React from 'react';
import ReactDOM from 'react-dom';
import {Provider} from 'react-redux'
import './index.css';
import Route from './routes'
import store from './redux/store'

ReactDOM.render(
  <React.StrictMode>
	  <Provider store={store}>
		  <Route></Route>
	  </Provider>
  </React.StrictMode>,
  document.getElementById('root')
);
