import * as React from 'react';
import * as ReactDOM from 'react-dom';

import { Provider } from 'react-redux';
import store from './store/store';
import { App } from './app';

import { fetchState } from './actions/authentication';

store.dispatch(fetchState());

ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('app-container')
);
