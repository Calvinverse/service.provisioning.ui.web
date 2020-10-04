import { Account } from 'msal'

export interface ProfileState {
  user?: Account
  error: boolean
  isAuthenticated: boolean
}