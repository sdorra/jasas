import { Action } from 'redux';

export interface Action extends Action {
  type: string;
  payload?: any;
}

export const EmptyAction: Action = {
  type: ''
};
