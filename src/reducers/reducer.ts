import { Action } from '../actions/action';

export interface Reducer<T> {
  (state: T, action: Action): T;
}
