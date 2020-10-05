import { MutationTree } from 'vuex'
import { ProfileState } from './ProfileState'
import { Account } from 'msal'

export const mutations: MutationTree<ProfileState> = {
  updateIsAuthenticated (state, value: boolean) {
    state.isAuthenticated = value
  },
  updateUser (state, value: Account) {
    state.user = value
  }
}
