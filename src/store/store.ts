import { createStore, applyMiddleware } from 'redux';
import promiseMiddleware from 'redux-promise-middleware';
import * as createLogger from "redux-logger";
import AuthenticationReducer  from '../reducers/authentication'

const store = createStore(
  AuthenticationReducer,
  applyMiddleware(promiseMiddleware(), createLogger())
);

export default store;
