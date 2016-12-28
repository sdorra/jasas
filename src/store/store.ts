import { createStore, combineReducers, applyMiddleware } from 'redux';
import promiseMiddleware from 'redux-promise-middleware';
import * as createLogger from "redux-logger";
import AuthenticationReducer  from '../reducers/authentication'

const store = createStore(
  combineReducers({
    auth: AuthenticationReducer
  }),
  applyMiddleware(promiseMiddleware(), createLogger())
);

export default store;
