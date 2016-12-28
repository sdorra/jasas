import { Action, EmptyAction } from '../actions/action';
import { Reducer } from './reducer';

import {
  AuthenticationForm,

  FETCH_STATE_PENDING,
  FETCH_STATE_FULFILLED,
  FETCH_STATE_REJECTED,

  CHANGE_FORM,

  AUTHENTICATE_PENDING,
  AUTHENTICATE_FULFILLED,
  AUTHENTICATE_REJECTED,

  LOGOUT_PENDING,
  LOGOUT_FULFILLED,
  LOGOUT_REJECTED
} from '../actions/authentication';

export interface AuthenticationState {
  error: boolean
  loading: boolean
  authenticated: boolean
  username?: string
  message?: string,
  form: AuthenticationForm
}

const EMTPY_FORM: AuthenticationForm = {
  username: '',
  password: ''
};

const INITIAL_STATE: AuthenticationState = {
  error: false,
  loading: true,
  authenticated: false,
  form: EMTPY_FORM
};

let authenticationReducer: Reducer<AuthenticationState> = (state = INITIAL_STATE, action: Action = EmptyAction) => {
  switch (action.type) {
    case AUTHENTICATE_PENDING:
    case FETCH_STATE_PENDING:
      return Object.assign({}, state, {
        loading: true, 
        authenticated: false, 
        error: false, 
        message: null
      });

    case AUTHENTICATE_FULFILLED:
      return Object.assign({}, state, {
        loading: false, 
        authenticated: true, 
        error: false, 
        message: null,
        form: EMTPY_FORM
      }, action.payload.data);
      
    case FETCH_STATE_FULFILLED:
      return Object.assign({}, state, {
        loading: false, 
        authenticated: true, 
        error: false, 
        message: null
      }, action.payload.data);

    case AUTHENTICATE_REJECTED:
      return Object.assign({}, state, {
        loading: false, 
        authenticated: false, 
        error: true, 
        message: "authentication failed"
      });

    case FETCH_STATE_REJECTED:
      return Object.assign({}, state, {
        loading: false
      });

    // logout

    case LOGOUT_PENDING:
      return Object.assign({}, state, {
        loading: true, 
        error: false, 
        message: null
      });

    case LOGOUT_FULFILLED:
      return Object.assign({}, state, {
        loading: false, 
        authenticated: false, 
        error: false, 
        message: null, 
        username: null
      });

    case LOGOUT_REJECTED:
      return Object.assign({}, state, {
        loading: false, 
        error: true, 
        message: "logout failed"
      });

    case CHANGE_FORM:
      return Object.assign({}, state, {
        form: action.payload
      });

    default:
      return state;
  }
}

export default authenticationReducer;
