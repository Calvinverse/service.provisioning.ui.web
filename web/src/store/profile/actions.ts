import { ActionTree } from 'vuex'
import { ProfileState } from './ProfileState'
import { RootState } from './RootState'
import { AuthenticationService } from '../../services/AuthenticationService'
import { MsalConfig } from '../../config/msalConfig'

const config = new MsalConfig()
const authenticationService = new AuthenticationService(config.clientID, config.authority, config.scopes)

export const actions: ActionTree<ProfileState, RootState> = {
  async login ({ commit }) {
    console.log('DEBUG: Store login called')
    const user = await authenticationService.login() // What if error
    commit('updateUser', user)
    commit('updateIsAuthenticated', user !== undefined)
    console.log('DEBUG: Store login complete.') // Should send a message to the UI indicating that login failed
  },
  logout ({ commit }) {
    authenticationService.logout() // What if error
    commit('updateUser', {})
    commit('updateIsAuthenticated', false)
  }
}
