import * as React from 'react';
import * as ReactDOM from 'react-dom';
import * as urls from './urls';

import { Provider } from 'react-redux';
import store from './store/store';

import { App } from './app';

import '../style/jasas.scss';

import { registerRedirectUri, fetchState } from './actions/authentication';

// register redirect uri
let redirectUri = urls.getQueryParameter('redirect_uri');
if (redirectUri) {
  store.dispatch(registerRedirectUri(redirectUri));
}

// fetch intitial state
store.dispatch(fetchState());

// render application
ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('app')
);
