import * as axios from 'axios';

export const ACTION_FETCH_STATE = 'FETCH_STATE';
export const ACTION_AUTHENTICATE = 'AUTHENTICATE';
export const ACTION_LOGOUT = 'LOGOUT';

export function fetchState() {
  return {
    type: ACTION_FETCH_STATE,
    payload: axios.get('/v1/validation')
  };
}

export function authenticate(username: string, password: string) {
  return {
    type: ACTION_AUTHENTICATE,
    payload: axios.post('/v1/authentication', {
      username: username,
      password: password
    })
  };
}

export function logout() {
  return {
    type: ACTION_LOGOUT,
    payload: axios.post('/v1/logout')
  };
}
