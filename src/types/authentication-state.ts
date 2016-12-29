import { AuthenticationForm } from './authentication-form';

export interface AuthenticationState {
  error: boolean
  loading: boolean
  authenticated: boolean
  username?: string
  message?: string
  redirectUri?: string
  form: AuthenticationForm
}
