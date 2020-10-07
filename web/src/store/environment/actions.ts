import { ActionTree } from 'vuex'
import {
  EnvironmentState
} from './types'
import { RootState } from '../RootState'
import { EnvironmentService } from '../../services/EnvironmentService'

const environmentService = new EnvironmentService(config.clientID, config.authority, config.scopes)

export const actions: ActionTree<EnvironmentState, RootState> = {
  async get({ commit }) {
    console.log('DEBUG: Store login called')
    const user = await environmentService.get() // What if error
    commit('updateUser', user)
    commit('updateIsAuthenticated', user !== undefined)
    console.log('DEBUG: Store login complete.') // Should send a message to the UI indicating that login failed
  },

  async getAll({ commit }) {
    console.log('DEBUG: Store login called')
    const user = await environmentService.getAll() // What if error
    commit('updateUser', user)
    commit('updateIsAuthenticated', user !== undefined)
    console.log('DEBUG: Store login complete.') // Should send a message to the UI indicating that login failed
  },

  async delete({ commit }) {
    await environmentService.delete() // What if error
    commit('updateUser', {})
    commit('updateIsAuthenticated', false)
  }
}
