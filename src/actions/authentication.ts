import { Action } from './action';

import * as axios from 'axios';

const PENDING_SUFFIX = '_PENDING';
const FULFILLED_SUFFIX = '_FULFILLED';
const REJECTED_SUFFIX = '_REJECTED';

const FETCH_STATE = 'FETCH_STATE';
export const FETCH_STATE_PENDING = FETCH_STATE + PENDING_SUFFIX;
export const FETCH_STATE_FULFILLED = FETCH_STATE + FULFILLED_SUFFIX;
export const FETCH_STATE_REJECTED = FETCH_STATE + REJECTED_SUFFIX;

export function fetchState(): Action {
  return {
    type: FETCH_STATE,
    payload: axios.get('/v1/validation')
  };
}

export const CHANGE_FORM = 'CHANGE_FORM';

export interface AuthenticationForm {
  username: string
  password: string
}

export function changeForm(form: AuthenticationForm): Action {
  return {
    type: CHANGE_FORM,
    payload: form
  };
}

const AUTHENTICATE = 'AUTHENTICATE';
export const AUTHENTICATE_PENDING = AUTHENTICATE + PENDING_SUFFIX;
export const AUTHENTICATE_FULFILLED = AUTHENTICATE + FULFILLED_SUFFIX;
export const AUTHENTICATE_REJECTED = AUTHENTICATE + REJECTED_SUFFIX;

export function authenticate(username: string, password: string): Action {
  return {
    type: AUTHENTICATE,
    payload: axios.post('/v1/authentication', {
      username: username,
      password: password
    })
  };
}

const LOGOUT = 'LOGOUT';
export const LOGOUT_PENDING = LOGOUT + PENDING_SUFFIX;
export const LOGOUT_FULFILLED = LOGOUT + FULFILLED_SUFFIX;
export const LOGOUT_REJECTED = LOGOUT + REJECTED_SUFFIX;


export function logout(): Action {
  return {
    type: LOGOUT,
    payload: axios.post('/v1/logout')
  };
}
