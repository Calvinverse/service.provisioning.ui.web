import { MutationTree } from 'vuex'
import { EnvironmentState } from './types'
import { Account } from 'msal'
import { Environment } from '@/types'

export const mutations: MutationTree<EnvironmentState> = {
  update(state, value: Environment) {
    state. = value
  },
  updateUser(state, value: Account) {
    state.user = value
  }
}
