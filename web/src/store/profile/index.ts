import {
  VuexModule,
  Module,
  Mutation,
  Action
} from 'vuex-module-decorators'
import { Account } from 'msal'
import {
  Options,
  url
} from 'gravatar'
import { AuthenticationService } from '../../services/AuthenticationService'

@Module({ namespaced: true })
class Profile extends VuexModule {
  user?: Account
  error = false
  isAuthenticated = false

  get fullName (): string {
    const user = this.user
    const firstName = (user && user.idToken && user.idToken.given_name) || ''
    const lastName = (user && user.idToken && user.idToken.family_name) || ''
    return `${firstName} ${lastName}`
  }

  get gravatarImage (): string {
    const user = this.user
    const email = (user && user.idToken && user.idToken.emails[0]) || ''
    return url(email)
  }

  @Action
  async login () {
    console.log('DEBUG: Store login called')
    const service = AuthenticationService.getInstance()
    const loginResult = await service.login()
    if (loginResult) {
      const user = service.getUser()
      this.context.commit('updateUser', user)
      this.context.commit('updateIsAuthenticated', true)
      this.context.commit('updateError', false)
      console.log('DEBUG: Store login complete.') // Should send a message to the UI indicating that login failed
    } else {
      this.context.commit('updateUser', {})
      this.context.commit('updateIsAuthenticated', false)
      this.context.commit('updateError', true)
    }
  }

  @Action
  logout () {
    const service = AuthenticationService.getInstance()
    service.logout() // What if error
    this.context.commit('updateUser', {})
    this.context.commit('updateIsAuthenticated', false)
  }

  @Mutation
  updateError (value: boolean) {
    this.error = value
  }

  @Mutation
  updateIsAuthenticated (value: boolean) {
    this.isAuthenticated = value
  }

  @Mutation
  updateUser (value: Account) {
    this.user = value
  }
}

export default Profile
