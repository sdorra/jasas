import { } from '../actions/authentication';

export interface AuthenticationState {
  error: boolean
  loading: boolean
  authenticated: boolean
  username?: string
  message?: string
}

const INITIAL_STATE: AuthenticationState = {
  error: false,
  loading: true,
  authenticated: false
};

function authenticationReducer(state = INITIAL_STATE, action: any = { type: '' }) {
  switch (action.type) {
    case 'AUTHENTICATE_PENDING':
    case 'FETCH_STATE_PENDING':
      return Object.assign({}, state, {loading: true, authenticated: false, error: false, message: null});
    case 'AUTHENTICATE_FULFILLED':
    case 'FETCH_STATE_FULFILLED':
      return Object.assign({}, state, {loading: false, authenticated: true, error: false, message: null}, action.payload.data);
    case 'AUTHENTICATE_REJECTED':
      return Object.assign({}, state, {loading: false, authenticated: false, error: true, message: "authentication failed"});
    case 'FETCH_STATE_REJECTED':
      return Object.assign({}, state, {loading: false});

    case 'LOGOUT_PENDING':
      return Object.assign({}, state, {loading: true, error: false, message: null});

    case 'LOGOUT_FULFILLED':
      return Object.assign({}, state, {loading: false, authenticated: false, error: false, message: null, username: null});

    case 'LOGOUT_REJECTED':
      return Object.assign({}, state, {loading: false, error: true, message: "logout failed"});

    default:
      return state;
  }
}

export default authenticationReducer;
