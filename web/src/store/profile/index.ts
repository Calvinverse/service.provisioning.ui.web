import { Module } from 'vuex'
import { getters } from './getters'
import { actions } from './actions'
import { mutations } from './mutations'
import { ProfileState } from './ProfileState'
import { RootState } from './RootState'

export const state: ProfileState = {
  user: undefined,
  error: false,
  isAuthenticated: false
}

const namespaced = true

export const profile: Module<ProfileState, RootState> = {
  namespaced,
  state,
  getters,
  actions,
  mutations
}
